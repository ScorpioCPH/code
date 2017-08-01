#!/Users/cenph/anaconda/bin/python
# coding=utf-8

import os
os.environ['TF_CPP_MIN_VLOG_LEVEL'] = '0'

import tensorflow as tf
import numpy as np

print "tf.version: ", tf.__version__

# M = np.array([
#         [[1],[-1],[0]],
#         [[-1],[2],[1]],
#         [[0],[2],[-2]]
#     ])

# # M = np.array([
# #         [[5],[-7],[3]],
# #         [[-6],[8],[10]],
# #         [[20],[12],[-24]]
# #     ])

# print "Matrix shape is: ", M.shape
# print "Matrix.1: ", M

# filter_weight = tf.get_variable('weights', [2, 2, 1, 1], initializer = tf.constant_initializer([
#                                                                         [1, -1],
#                                                                         [0, 2]]))
# biases = tf.get_variable('biases', [1], initializer = tf.constant_initializer(1))

# # M = np.asarray(M, dtype='float32')
# M = M.reshape(1, 3, 3, 1)

# # print "Matrix.2: ", M

# x = tf.placeholder('int32', [1, None, None, 1])
# # conv = tf.nn.conv2d(x, filter_weight, strides = [1, 2, 2, 1], padding = 'SAME')
# # bias = tf.nn.bias_add(conv, biases)
# pool = tf.nn.max_pool(x, ksize=[1, 2, 2, 1], strides=[1, 2, 2, 1], padding='SAME')
# print "pool: ", pool

# with tf.Session() as sess:
#     tf.global_variables_initializer().run()
#     # convoluted_M = sess.run(bias,feed_dict={x:M})
#     pooled_M = sess.run(pool,feed_dict={x:M})
    
#     # print "convoluted_M: \n", convoluted_M
#     print "pooled_M: \n", pooled_M
#     print "pool.get_shape(): ", pool.get_shape()

# input = tf.constant([[1., 1., 2., 4., -5.], [5., 6., 7., 8., -9.], [3., 2., 1., 0., -1.], [1., 2., 3., 4., 5.], [-19., -12., -13., -14., -15.]])
 
# input = tf.reshape(input, [1, 5, 5, 1])
 
# op = tf.nn.max_pool(input, ksize=[1, 2, 2, 1], strides=[1, 2, 2, 1], padding='SAME')
 
# with tf.Session() as sess:
#     print("input:")
#     print(input.eval())
#     print("op:")
#     print(op) 
#     result = sess.run(op)
#     print("result:")
#     print(result)

# import tensorflow as tf

import tensorflow as tf
import numpy as np

M = np.array([
        [[1],[-1],[0]],
        [[-1],[2],[1]],
        [[0],[2],[-2]]
    ])

M = np.asarray(M, dtype='float32')
M = M.reshape(1, 3, 3, 1)

data = tf.placeholder("float32", shape=[1, 3, 3, 1])
pool = tf.nn.avg_pool(data, [1, 2, 2, 1], [1, 2, 2, 1], padding='SAME')

print(pool.get_shape())

with tf.Session() as sess:
    pooled_M = sess.run(pool,feed_dict={data:M})
    print "pooled_M: \n", pooled_M
