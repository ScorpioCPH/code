#!/Users/cenph/anaconda/bin/python
# coding=utf-8

import os
os.environ['TF_CPP_MIN_VLOG_LEVEL'] = '8'

import tensorflow as tf

print("tf.version: ", tf.__version__)

# Simple hello world using TensorFlow

# Create a Constant op
# The op is added as a node to the default graph.
#
# The value returned by the constructor represents the output
# of the Constant op.

hello = tf.constant('Hello, TensorFlow!')

# Start tf session
sess = tf.Session()

# Run the op
print(sess.run(hello))
