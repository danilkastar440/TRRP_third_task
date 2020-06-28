#!/usr/bin/env  python3
import grpc

# import the generated classes
import calculator_pb2
import calculator_pb2_grpc

# open a gRPC channel
channel = grpc.insecure_channel('localhost:50051')

# create a stub (client)
stub = calculator_pb2_grpc.CalculatorStub(channel)

while True:
    # create a valid request message
    value = float(input("Enter an integer: "))
    number = calculator_pb2.Number(value=value)

    # make the call
    response = stub.SquareRoot(number)
    print(response.value)
