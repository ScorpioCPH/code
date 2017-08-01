# coding=utf-8

import tensorflow as tf

# cat.jpg 放到相同目录
image_raw_data = tf.gfile.FastGFile("./cat.jpg",'r').read()

with tf.Session() as sess:
    img_data = tf.image.decode_jpeg(image_raw_data)
    
    # 输出解码之后的三维矩阵
    print 'img_data: ', img_data.eval()
    img_data.set_shape([1797, 2673, 3])
    print 'img_shape: ', img_data.get_shape()
