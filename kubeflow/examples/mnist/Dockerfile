# FROM gcr.io/tensorflow/tensorflow:latest-gpu

FROM cargo.caicloud.io/tensorflow/tensorflow:1.2.1-gpu

RUN mkdir -p /opt/mnist

ADD https://raw.githubusercontent.com/tensorflow/tensorflow/master/tensorflow/examples/tutorials/mnist/mnist_softmax.py /opt/mnist/mnist_softmax.py

ENTRYPOINT ["python", "/opt/mnist/mnist_softmax.py"]
