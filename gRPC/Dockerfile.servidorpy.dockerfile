FROM grpc/python
ADD . /microservices-grpc-go-python
WORKDIR /Proyecto2/servidor_Py
RUN pip install --upgrade pip \
 && pip install grpcio grpcio-tools
CMD ["python", "servidor.py", "11443"]