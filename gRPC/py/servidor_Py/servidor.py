import sys
import time
import grpc
import decimal
import comunicacion_pb2
import comunicacion_pb2_grpc
from concurrent import futures
import json
from flask import Flask, jsonify, request, Response
from flask_pymongo import PyMongo
from flask_cors import CORS
import redis


class Comunicacion(comunicacion_pb2_grpc.ComunicandoServicer):
   def call(self, request, content):
        #print(request.body)
        y = json.loads(request.body)
        x = json.loads(request.body)
        #print(y)
        # the result is a Python dictionary:
        #print(y["age"])

        app = Flask(__name__)
        CORS(app)

        # print("Mongo")
        app.config['MONGO_URI'] = 'mongodb://3.139.95.193:27017/test'
        mongo = PyMongo(app)
        create_mesaje(y, mongo)
        get_mensajescantidad(mongo)
        print("redis")
        rediss(x)

        return comunicacion_pb2.Llamada(body="nice nice nice")

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
    #print("insertar")


def get_mensajescantidad(mongo):
    users = mongo.db.casos.count()
    message = {
        'number': users
    }
    print(message)


def get_server(host):
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=5))
    #keys_dir = os.path.abspath(os.path.join('.', os.pardir, 'keys'))
    #with open('%s/private.key' % keys_dir, 'rb') as f:
    #    private_key = f.read()
    #with open('%s/cert.pem' % keys_dir, 'rb') as f:
    #    certificate_chain = f.read()
    #server_credentials = grpc.ssl_server_credentials(((private_key, certificate_chain),))
    server.add_insecure_port(host)
    comunicacion_pb2_grpc.add_ComunicandoServicer_to_server(Comunicacion(), server)
    return server


if __name__ == '__main__':
    #port = sys.argv[1] if len(sys.argv) > 1 else 443
    #host = '[::]:%s' % port
    host = '[::]:9000'
    server = get_server(host)
    try:
        server.start()
        print('Running Comunicando service on %s' % host)
        while True:
            time.sleep(1)
    except Exception as e:
        print('[error] %s' % e)
        server.stop(0)

