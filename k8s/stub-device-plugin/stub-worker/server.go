/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"path"
	// "strings"
	"time"
	"sync"
	"io/ioutil"
	"regexp"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pluginapi "k8s.io/kubernetes/pkg/kubelet/apis/deviceplugin/v1alpha"
)

const (
	// stub device name
	resourceName = "stub-device.com/dummy"
	serverSocket = pluginapi.DevicePluginPath + "stub.sock"

	// dummy device file saved in "/tmp" directory,
	// e.g. "/tmp/stub-device-1", "/tmp/stub-device-2"
	dummyDeviceDirectory string = "/tmp"
	dummyDeviceRE string = "^stub-device-[0-9]*$"
)

type StubWorker struct {
	mutex sync.Mutex
	socket string
	grpcServer *grpc.Server

	devices map[string]pluginapi.Device

	stop   chan interface{}
	update chan map[string]pluginapi.Device
}

func getDevices() map[string]pluginapi.Device {
	reg := regexp.MustCompile(dummyDeviceRE)
	
	files, err := ioutil.ReadDir(dummyDeviceDirectory)
	if err != nil {
		return nil
	}

	var devices = make(map[string]pluginapi.Device)

	for _, f := range files {
		if f.IsDir() {
			continue
		}

		if reg.MatchString(f.Name()) {
			fmt.Printf("found dummy device %q\n", f.Name())
			devices[f.Name()] = pluginapi.Device{
				ID: f.Name(),
				Health: pluginapi.Healthy,
			}
		}
	}

	fmt.Printf("getDevices: %+v\n", devices)
	return devices
}

// NewStubWorker returns an initialized StubWorker.
func NewStubWorker() *StubWorker {
	return &StubWorker{
		devices: getDevices(),
		socket: serverSocket,

		stop:   make(chan interface{}),
		update: make(chan map[string]pluginapi.Device),
	}
}

// Start starts the gRPC server of the device plugin
func (w *StubWorker) Start() error {
	fmt.Printf("StubWorker.Start()\n")
	err := w.cleanup()
	if err != nil {
		return err
	}

	socket, err := net.Listen("unix", w.socket)
	if err != nil {
		return err
	}

	w.grpcServer = grpc.NewServer()
	pluginapi.RegisterDevicePluginServer(w.grpcServer, w)

	go w.grpcServer.Serve(socket)
	// Wait till grpc server is ready.
	for i := 0; i < 10; i++ {
		services := w.grpcServer.GetServiceInfo()
		if len(services) > 1 {
			break
		}
		time.Sleep(1 * time.Second)
	}
	// go m.HealthCheck()

	return nil
}

// Stop stops the gRPC server
func (w *StubWorker) Stop() error {
	fmt.Printf("StubWorker.Stop()\n")
	if w.grpcServer == nil {
		return nil
	}

	w.grpcServer.Stop()
	w.grpcServer = nil
	// close(m.stop)

	// return m.cleanup()
	return nil
}

// Register registers the device plugin for the given resourceName with Kubelet.
func (w *StubWorker) Register(kubeletEndpoint, resourceName string) error {
	conn, err := grpc.Dial(kubeletEndpoint, grpc.WithInsecure(),
		grpc.WithDialer(func(addr string, timeout time.Duration) (net.Conn, error) {
			return net.DialTimeout("unix", addr, timeout)
		}))
	defer conn.Close()
	if err != nil {
		return err
	}
	client := pluginapi.NewRegistrationClient(conn)
	reqt := &pluginapi.RegisterRequest{
		Version:      pluginapi.Version,
		Endpoint:     path.Base(w.socket),
		ResourceName: resourceName,
	}

	_, err = client.Register(context.Background(), reqt)
	if err != nil {
		return fmt.Errorf("Cannot register stub worker to kubelet service: %v", err)
	}
	return nil
}

// ListAndWatch lists devices and update that list according to the Update call
func (w *StubWorker) ListAndWatch(emtpy *pluginapi.Empty, server pluginapi.DevicePlugin_ListAndWatchServer) error {
	for {
		resp := new(pluginapi.ListAndWatchResponse)
		for _, dev := range w.devices {
			resp.Devices = append(resp.Devices, &pluginapi.Device{dev.ID, dev.Health})
		}
		fmt.Printf("ListAndWatch: send devices %v", resp)

		if err := server.Send(resp); err != nil {
			fmt.Printf("device-plugin: cannot update device states: %v\n", err)
			return err
		}

		time.Sleep(5 * time.Second)
	}
}

// // Update allows the device plugin to send new devices through ListAndWatch
// func (w *StubWorker) Update(devs []*pluginapi.Device) {
// 	m.update <- devs
// }

// Allocate which return list of devices.
func (w *StubWorker) Allocate(ctx context.Context, r *pluginapi.AllocateRequest) (*pluginapi.AllocateResponse, error) {
	// devs := w.devices
	var response pluginapi.AllocateResponse

	// for i, id := range r.DevicesIDs {
	// 	if !deviceExists(devs, id) {
	// 		return nil, fmt.Errorf("Invalid allocation request: unknown device: %s", id)
	// 	}

	// 	devRuntime := new(pluginapi.DeviceRuntimeSpec)
	// 	devRuntime.ID = id
	// 	// Only add env vars to the first device.
	// 	// Will be fixed by: https://github.com/kubernetes/kubernetes/pull/53031
	// 	if i == 0 {
	// 		devRuntime.Envs = make(map[string]string)
	// 		devRuntime.Envs["NVIDIA_VISIBLE_DEVICES"] = strings.Join(r.DevicesIDs, ",")
	// 	}

	// 	response.Spec = append(response.Spec, devRuntime)
	// }

	return &response, nil
}

func (w *StubWorker) cleanup() error {
	if err := os.Remove(w.socket); err != nil && !os.IsNotExist(err) {
		return err
	}

	return nil
}

// func (w *StubWorker) HealthCheck() {
// 	eventSet := nvml.NewEventSet()
// 	defer nvml.DeleteEventSet(eventSet)

// 	err := nvml.RegisterEvent(eventSet, nvml.XidCriticalError)
// 	check(err)

// 	for {
// 		select {
// 		case <-m.stop:
// 			return
// 		default:
// 			// FIXME: there is a race condition if another goroutine calls m.Update concurrently.
// 			devs := m.devs
// 			healthy := checkXIDs(eventSet, devs)
// 			if !healthy {
// 				m.Update(devs)
// 			}
// 		}
// 	}
// }

func (w *StubWorker) Serve() error {
	fmt.Printf("StubWorker.Serve()\n")
	err := w.Start()
	if err != nil {
		log.Printf("Could not start stub worker: %s", err)
		return err
	}
	log.Println("Starting to serve on:", w.socket)
	log.Println("Register to:", pluginapi.KubeletSocket)

	err = w.Register(pluginapi.KubeletSocket, resourceName)
	if err != nil {
		log.Printf("Could not register device plugin: %s", err)
		w.Stop()
		return err
	}
	log.Println("Registered device plugin with Kubelet")

	return nil
}