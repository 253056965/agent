package core

// ProtobufEncoder 这是一个封包的类
type ProtobufEncoder interface{
    Encode(iOSession *IOSession,obj interface{},protocolEncoderOutput ProtocolEncoderOutput)(result bool ,err error);
}