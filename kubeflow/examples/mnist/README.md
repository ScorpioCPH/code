
#### set nvidia-docker (2.0) as the default runtime

You can follow this doc: [set nvidia-docker as default runtime](https://github.com/NVIDIA/nvidia-docker/wiki/Advanced-topics#default-runtime) reference.

Change docker config:

```sh
$ cat /etc/docker/daemon.json 
{
    "runtimes": {
        "nvidia": {
            "path": "/usr/bin/nvidia-container-runtime",
            "runtimeArgs": []
        }
    },
    "default-runtime": "nvidia"
}
``` 

Check your setting:

```sh
$ docker info
...
Runtimes: nvidia runc
Default Runtime: nvidia
...
```


#### Deploy NVIDIA device plugin in your cluster

[Deploy as DaemonSet](https://github.com/NVIDIA/k8s-device-plugin#deploy-as-daemon-set)

#### make image with tensorflow (gpu) and mnist example

Dockerfile:

```Dockerfile
FROM gcr.io/tensorflow/tensorflow:latest-gpu

RUN mkdir -p /opt/mnist

ADD https://raw.githubusercontent.com/tensorflow/tensorflow/master/tensorflow/examples/tutorials/mnist/mnist_softmax.py /opt/mnist/mnist_softmax.py

ENTRYPOINT ["python", "/opt/mnist/mnist_softmax.py"]
```

Or you can use this built image: `cargo.caicloud.io/cenph/mnist_softmax:1.0`

#### deploy a pod to request nvidia gpu

Pod.yaml:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: mnist-with-device-plugin
  labels:
    app: mnist-with-device-plugin
spec:
  containers:
  - name: mnist-with-device-plugin
    image: cargo.caicloud.io/cenph/mnist_softmax:1.0
    resources:
      requests:
        nvidia.com/gpu: 1
      limits:
        nvidia.com/gpu: 1
    securityContext:
        privileged: true
```

You will get some logs after training is finished:

```text
I tensorflow/core/common_runtime/gpu/gpu_device.cc:940] Found device 0 with properties: 
name: GeForce GTX 750 Ti
major: 5 minor: 0 memoryClockRate (GHz) 1.15
pciBusID 0000:01:00.0
Total memory: 1.95GiB
Free memory: 1.17GiB
I tensorflow/core/common_runtime/gpu/gpu_device.cc:961] DMA: 0 
I tensorflow/core/common_runtime/gpu/gpu_device.cc:971] 0:   Y 
I tensorflow/core/common_runtime/gpu/gpu_device.cc:1030] Creating TensorFlow device (/gpu:0) -> (device: 0, name: GeForce GTX 750 Ti, pci bus id: 0000:01:00.0)
...
0.9177
```
