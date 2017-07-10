#!/usr/bin/python
#coding=utf-8
# import random, string
import requests
import json

import urllib2
from urllib2 import urlopen
import numpy as np
import json
import time
import cv2
import argparse

# post方式访问serving,post传送的数据未post形式
def http_post(url,data):
	json_data = json.dumps(data)
	req = urllib2.Request(url, json_data)
	response = urllib2.urlopen(req)
	return response.read()

def create_opencv_image_from_url(url, cv2_img_flag=3):
	print("create_opencv_image_from_url")
	request = urlopen(url)
	img_array = np.asarray(bytearray(request.read()), dtype=np.uint8)
	return cv2.imdecode(img_array, cv2_img_flag)

#https://ofp0w1txg.qnssl.com/img/texture/格纹/48b42c872afd821c327e88e7e1a4b74f.jpg
#https://ofp0w1txg.qnssl.com/img/texture/格纹/d6055a3a727c2a6325d7dfc22e79194e.jpg

# config
login_url = "http://taas.caicloud.io:30122/api/v1/login"
cloth_recognition_url = "http://taas.caicloud.io:30122/api/v1/servings/cloth_recognition"

if __name__=="__main__":

	# parse arg
	ap = argparse.ArgumentParser()
	ap.add_argument("-i", "--image", help="url to the image file")
	args = vars(ap.parse_args())

	# login to get token
	loginPayload = {}
	loginPayload['username'] = 'admin'
	loginPayload['password'] = '*****'

	headers = {'content-type': 'application/json'}


	loginResponse = requests.post(login_url, data=json.dumps(loginPayload), headers=headers)
	print 'loginResponse.status_code :', loginResponse.status_code
	print 'loginResponse.text :', loginResponse.text
	accessToken =  (json.loads(loginResponse.text))['accessToken']
	print 'token: ', accessToken



	url='http://192.168.48.14:8000'
	imageUrl=args["image"]
	img=create_opencv_image_from_url(imageUrl)
	img_w=img.shape[1]
	img_h=img.shape[0]
	data = {"imageUrl": imageUrl,"sizeFilterThres":"0.005"}
	start_time=time.time()

	headers = {'content-type': 'application/json',
           'Authorization': accessToken,
           'sizeFilterThres': '0.005'}

	payload = {}
	payload['imageUrl'] = imageUrl

	response = requests.post(cloth_recognition_url, data=json.dumps(payload), headers=headers)
	print 'response.code :', response.status_code
	print 'response.text :', response.text
    # if response.status_code == 200:
 #        print 'Create new user: ' + payload['username'] + ', password: ' + payload['password']

	# #解析serving返回的数据并打印
	resp = list(json.loads(response.text))

	for i in range(len(resp)):
		print resp[i]['tag'],resp[i]['position']['xmin'],resp[i]['position']['ymin'],resp[i]['position']['xmax'],resp[i]['position']['ymax']
		xmin=int(img_w*float(resp[i]['position']['xmin']))
		ymin=int(img_h*float(resp[i]['position']['ymin']))
		xmax=int(img_w*float(resp[i]['position']['xmax']))
		ymax=int(img_h*float(resp[i]['position']['ymax']))
		print xmin,ymin,xmax,ymax
		cv2.rectangle(img,(abs(xmin-img_w),ymin),(abs(xmax-img_w),ymax),(0,0,255),2)

	end_time=time.time()
	print "count time is :",end_time-start_time
	cv2.imshow("img",img)
	k=cv2.waitKey(0)
	if k==ord("q"):
		cv2.destroyAllWindows()

