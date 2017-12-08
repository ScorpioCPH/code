package main

import (
	"fmt"
	"time"
	// "io/ioutil"
	// "os"

	// "github.com/golang/glog"
	pluginapi "k8s.io/kubernetes/pkg/kubelet/apis/deviceplugin/v1alpha"
	dp "k8s.io/kubernetes/pkg/kubelet/cm/deviceplugin"
)

func main() {
	// Marshal()
	// unmarshal()
	fmt.Printf("XXX\n")

	devs := []*pluginapi.Device{
		{ID: "Dev1", Health: pluginapi.Healthy},
		{ID: "Dev2", Health: pluginapi.Healthy},
	}

	p1 := dp.NewDevicePluginStub(devs, pluginapi.DevicePluginPath+"dp.1")
	err := p1.Start()
	// framework.ExpectNoError(err)
	if err != nil {
		fmt.Printf("p1.Start() error: %v", err)
	}
	fmt.Printf("Creating stub device plugin pod\n")

	resourceName := "mock1.com/resource"
	p1.Register(pluginapi.KubeletSocket, resourceName)

	for {
		time.Sleep(1 * time.Second)
	}

}
