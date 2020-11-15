FROM grpc/python
ADD . /app/src/Proyecto2/
WORKDIR /app/src/Proyecto2/servidor_Py
RUN pip install --upgrade pip \
 && pip install grpcio grpcio-tools
CMD ["python", "servidor.py", "11443"]