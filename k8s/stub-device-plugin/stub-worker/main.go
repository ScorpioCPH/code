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
	"time"
	"log"
	// "os"
	// "os/signal"
	// "runtime"
	// "runtime/debug"
	// "syscall"

	// "github.com/fsnotify/fsnotify"
	// pluginapi "k8s.io/kubernetes/pkg/kubelet/apis/deviceplugin/v1alpha1"
)

// func check(err error) {
// 	if err != nil {
// 		log.Panicln("Fatal:", err)
// 	}
// }

// func exit() {
// 	if err := recover(); err != nil {
// 		if _, ok := err.(runtime.Error); ok {
// 			log.Println(err)
// 		}
// 		if os.Getenv("NV_DEBUG") != "" {
// 			log.Printf("%s", debug.Stack())
// 		}
// 		os.Exit(1)
// 	}

// 	os.Exit(0)
// }

func main() {
	// defer exit()

	// Should it be in the device plugin Serve?
	if len(getDevices()) == 0 {
		log.Println("No devices found. Looping")
		select {}
	}

	stubWorker := NewStubWorker()
	stubWorker.Serve()

	// watcher, err := fsnotify.NewWatcher()
	// check(err)
	// defer watcher.Close()
	// err = watcher.Add(pluginapi.DevicePluginPath)
	// check(err)

	// sigs := make(chan os.Signal, 1)
	// signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

// L:
	for {
		time.Sleep(10)
		// restart := false
		// select {
		// case event := <-watcher.Events:
		// 	if event.Name == pluginapi.KubeletSocket && event.Op&fsnotify.Create == fsnotify.Create {
		// 		log.Printf("inotify: %s created, restarting", pluginapi.KubeletSocket)
		// 		restart = true
		// 	}
		// case err := <-watcher.Errors:
		// 	log.Printf("inotify: %s", err)
		// case s := <-sigs:
		// 	switch s {
		// 	case syscall.SIGHUP:
		// 		log.Println("Received SIGHUP, restarting")
		// 		restart = true
		// 	default:
		// 		log.Printf("Received signal %d, shutting down", s)
		// 		stubWorker.Stop()
		// 		break L
		// 	}
		// }
		// if restart {
		// 	stubWorker.Stop()
		// 	stubWorker = NewStubWorker()
		// 	stubWorker.Serve()
		// }
	}
}
