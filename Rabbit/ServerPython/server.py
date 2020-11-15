import pika
import json
from flask import Flask
from flask_pymongo import PyMongo
from flask_cors import CORS
import redis

connection = pika.BlockingConnection(pika.ConnectionParameters(host='104.154.53.252:5672'))
channel= connection.channel()

channel.queue_declare(queue='lab-so1')

def callback(ch, method, properties ,body):
    print('Recibido: %r'% body)
    y = json.loads(body)
    x = json.loads(body)
    # print(y)
    # the result is a Python dictionary:
    # print(y["age"])

    app = Flask(__name__)
    CORS(app)

    app.config['MONGO_URI'] = 'mongodb://3.139.95.193:27017/test'
    mongo = PyMongo(app)
    create_mesaje(y, mongo)

    # print("redis")
    rediss(x)

def rediss(info):
    r = redis.StrictRedis(host="ec2-3-139-95-193.us-east-2.compute.amazonaws.com", port=6379, db=0, charset="utf-8", decode_responses=True)
    json_images = json.dumps(info)
    r.rpush('casos', json_images)
    #unpacked_images = json.loads(r.get('images'))

    tamano = r.llen('casos')
    print(tamano)
    unpacked_images = json.loads(r.lrange('casos', tamano - 1, tamano - 1)[0])

def create_mesaje(info, mongo):
    # Receiving Data
    mongo.db.casos.insert_one(info)
    # print("insertar")
channel.basic_consume(queue='lab-so1',on_message_callback=callback, auto_ack=True)
print('Esperando mensajes :)')
channel.start_consuming()