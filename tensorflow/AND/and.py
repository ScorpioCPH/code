# coding=utf-8

import tensorflow as tf

# 训练数据
X = [[0, 0],[0, 1],[1, 0],[1, 1]]
Y = [[0], [0], [0], [1]]

# 定义网络结构
N_INPUT_NODES = 2  # 2 个输入节点
N_OUTPUT_NODES  = 1  # 1 个输出节点

# 定义训练迭代次数
N_STEPS = 20000
N_EPOCH = 1000

# 定义学习率，即每次递减下降的大小
LEARNING_RATE = 0.02

# 定义接收训练数据的占位符
x_ = tf.placeholder(tf.float32, shape=[len(X), N_INPUT_NODES], name="x-input")
y_ = tf.placeholder(tf.float32, shape=[len(Y), N_OUTPUT_NODES], name="y-input")

# 定义权重
weight = tf.Variable(tf.random_uniform([N_INPUT_NODES, N_OUTPUT_NODES], -1, 1), name="weight")

# 定义偏差
bias = tf.Variable(tf.zeros([N_OUTPUT_NODES]), name="bias")

# 定义前向传播函数
output = tf.sigmoid(tf.matmul(x_, weight) + bias)

# 定义损失函数（最小均方差），来描述预测值和真实值之间的差距
cost = tf.reduce_mean(tf.square(Y - output))

# 定义反向传播函数，即使用梯度下降的方法，求解损失函数的最小值
train = tf.train.GradientDescentOptimizer(LEARNING_RATE).minimize(cost)

# 初始化变量
init = tf.initialize_all_variables()
sess = tf.Session()
sess.run(init)

saver = tf.train.Saver()

# 开始训练过程
for i in range(N_STEPS):
    # 执行训练函数，将训练数据 feed 到模型中
    sess.run(train, feed_dict={x_: X, y_: Y})
    if i % N_EPOCH == 0:
        # 每隔 N_EPOCH 轮，输出一次训练结果
        print('STEPS: ', i, 'cost: ', sess.run(cost, feed_dict={x_: X, y_: Y}))
        print('output: ', sess.run(output, feed_dict={x_: X}))
        print('cost: ', sess.run(cost, feed_dict={x_: X, y_: Y}))

# 训练结束，执行一次预测过程
print('output: ', sess.run(output, feed_dict={x_: X}))

# 保存参数到本地
save_path = saver.save(sess, "./model/model.ckpt")
print("Model saved in file: %s" % save_path)

# 保存 graph 到本地
writer = tf.summary.FileWriter("./model/", sess.graph)
