#!/Users/cenph/anaconda/bin/python
# coding=utf-8

import tensorflow as tf
import numpy as np
import cv2

print "tf.version: ", tf.__version__

sess=tf.Session()
img=cv2.imread("./image.jpg")

img_test=img.copy()

for i in range(img.shape[2]):
    M = img[:,:,i]
    M = M.reshape(1, M.shape[0], M.shape[1], 1)

    x = tf.placeholder('int32', [1, None, None, 1])

    pool = tf.nn.max_pool(x, ksize=[1, 10, 10, 1], strides=[1, 1, 1, 1], padding='SAME')

    pooled_M = sess.run(pool,feed_dict={x:M})
    cv2.imwrite(str(i)+"_pool.jpg",pooled_M[0])
    print "pooled_M: ", pooled_M
