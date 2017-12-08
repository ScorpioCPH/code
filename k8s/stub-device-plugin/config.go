package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	// "github.com/golang/glog"

	pluginapi "k8s.io/kubernetes/pkg/kubelet/apis/deviceplugin/v1alpha"
)

// DummyDevice represents the dummy devices details with resource name.
type DummyDevice struct {
	ResourceName string             `json:"resourceName,omitempty"`
	Devices      []pluginapi.Device `json:"devices,omitempty"`
}

// DummyDeviceList is a list of DummyDevice objects.
type DummyDeviceList struct {
	Items []DummyDevice `json:"items,omitempty"`
}

func Marshal() {
	dummyDeviceList := DummyDeviceList{
		Items: []DummyDevice{
			{
				Devices: []pluginapi.Device{
					{
						ID:     "dr1-1",
						Health: pluginapi.Healthy,
					},
					{
						ID:     "dr1-2",
						Health: pluginapi.Healthy,
					},
					{
						ID:     "dr1-3",
						Health: pluginapi.Unhealthy,
					},
				},
				ResourceName: "dr1",
			},
			// {
			// 	Devices: []pluginapi.Device{
			// 		{
			// 			ID:     "dr2-1",
			// 			Health: pluginapi.Healthy,
			// 		},
			// 		{
			// 			ID:     "dr2-2",
			// 			Health: pluginapi.Healthy,
			// 		},
			// 		{
			// 			ID:     "dr2-3",
			// 			Health: pluginapi.Unhealthy,
			// 		},
			// 	},
			// 	ResourceName: "dr2",
			// },
		},
	}

	d, err := json.Marshal(dummyDeviceList)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(d)
}

func Unmarshal() {
	file, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		fmt.Printf("ReadFile error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\n%s\n", string(file))

	var dummyDeviceList DummyDeviceList
	err = json.Unmarshal(file, &dummyDeviceList)
	if err != nil {
		fmt.Printf("Unmarshal error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Results: %v\n", dummyDeviceList)

}
