package core

//ProtobufDecoder 这是一个解析类的接口
type ProtobufDecoder interface {
	Decoder(iOSession *IOSession, ioBuffer *IoBuffer, protocolDecoderOutput ProtocolDecoderOutput) (result bool, err error)
}
