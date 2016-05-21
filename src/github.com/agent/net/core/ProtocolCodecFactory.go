package core
//ProtocolCodecFactory 编解码接口类
type ProtocolCodecFactory interface{
    //得到一个解析类
    GetProtobuDecoder() ProtobufDecoder;
    //得到几个加密类
    GetProtobufEncoder() ProtobufEncoder;
}