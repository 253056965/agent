package example

import "github.com/agent/net/core"

//TelnetProtobufDecoder 这是一个 telent 的一个 解码工厂
type TelnetProtobufDecoder struct {
}
func NewTelnetProtobufDecoder() *TelnetProtobufDecoder {
	return &TelnetProtobufDecoder{}
}
func (t *TelnetProtobufDecoder) Decoder(ioSession *core.IOSession, ioBuffer *core.IoBuffer, protocolDecoderOutput core.ProtocolDecoderOutput) (result bool, err error) {
	str, err := ioBuffer.ReadLine()
	if err == nil {
		protocolDecoderOutput.Write(str)
		return true, nil
	}
	return false, err
}
