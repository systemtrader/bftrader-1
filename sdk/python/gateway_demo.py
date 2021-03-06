# coding=utf-8

import time

import bftrader_pb2
import bfgateway_pb2
import bfrobot_pb2

from grpc.beta import implementations

_BF_VOID = bftrader_pb2.BfVoid()
_ONE_DAY_IN_SECONDS = 60 * 60 * 24
_TIMEOUT_SECONDS = 1

class Gateway(bfgateway_pb2.BetaBfGatewayServiceServicer):
    def __init__(self):
        print "init gateway"
        self.robot_channel = ""
        self.robot = ""
        
    def SetKv(self, request, context):
        print "SetKv"
        return _BF_VOID
        
    def GetKv(self, request, context):
        print "GetKv"
        return _BF_VOID
        
    def GetContract(self, request, context):
        print "GetContract"
        
    def GetContractList(self, request, context):
        print "GetContractList"

    def Connect(self, request, context):
        print "Connect"
        print request.robotId,request.robotIp,request.robotPort
        self.robot_channel = implementations.insecure_channel(request.robotIp, request.robotPort)
        self.robot = bfrobot_pb2.beta_create_BfRobotService_stub(self.robot_channel)
        return bftrader_pb2.BfConnectResp(exchangeOpened = True)

    def Subscribe(self, request, context):
        print "Subscribe"
        print request.symbol,request.exchange
        mt =  context.invocation_metadata()
        for it in mt:
            print "metadata: %s:%s" % (it[0],it[1])
        bfvoid = self.robot.OnTick(bftrader_pb2.BfTickData(),_TIMEOUT_SECONDS)        
        return _BF_VOID

    def SendOrder(self, request, context):
        print "SendOrder"

    def CancelOrder(self, request, context):
        print "CancelOrder"

    def QueryAccount(self, request, context):
        print "QueryAccount"

    def QueryPosition(self, request, context):
        print "QueryPosition"

    def Close(self, request, context):
        print "Close"
    
def run():
    gateway = bfgateway_pb2.beta_create_BfGatewayService_server(Gateway())
    gateway.add_insecure_port('[::]:50051')
    gateway.start()
    print "start gateway"
    try:
        while True:
            time.sleep(_ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        print "stop gateway"
        gateway.stop(0)

if __name__ == '__main__':
    run()
