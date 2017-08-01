#!/Users/cenph/anaconda/bin/python
# coding=utf-8

import os
os.environ['TF_CPP_MIN_VLOG_LEVEL'] = '8'

import tensorflow as tf
# import time

print "tf.version: ", tf.__version__

# time.sleep(5000)

# Config to turn on JIT compilation
config = tf.ConfigProto()
config.graph_options.optimizer_options.global_jit_level = tf.OptimizerOptions.ON_1
config.log_device_placement = True

print "new Session().1"
sess = tf.Session(config=config)
print "new Session().2"

x = tf.random_normal([1000, 1000])
y = tf.random_normal([1000, 1000])
res = tf.matmul(x, y)

sess.run(res)
