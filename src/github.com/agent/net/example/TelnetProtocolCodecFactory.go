package example
import "github.com/agent/net/core"
//TelnetProtocolCodecFactory 一个telent的编解码工程
type TelnetProtocolCodecFactory struct{
     protobufDecoder core.ProtobufDecoder;
     protobufEncoder core.ProtobufEncoder;
}
//GetProtobuDecoder 实现ProtocolCodecFactory 这个里面的方法的解码工厂
func (t *TelnetProtocolCodecFactory) GetProtobuDecoder()  core.ProtobufDecoder {  
    return t.protobufDecoder;
}
//GetProtobufEncoder 实现ProtocolCodecFactory 这个里面的方法的解码工厂
func (t *TelnetProtocolCodecFactory) GetProtobufEncoder()  core.ProtobufEncoder {
     return t.protobufEncoder;
}
//NewTelnetProtocolCodecFactory 得到一个Telnet的编解码工厂
func NewTelnetProtocolCodecFactory(protobufDecoder core.ProtobufDecoder,protobufEncoder core.ProtobufEncoder)  *TelnetProtocolCodecFactory{
     return &TelnetProtocolCodecFactory{protobufDecoder,protobufEncoder};
}
