package example

import (
	"fmt"
     "reflect"
	"github.com/agent/net/core"
)

type TelnetProtobufEncoder struct {
}

func NewTelnetProtobufEncoder() *TelnetProtobufEncoder {

	return &TelnetProtobufEncoder{};
}

func (t *TelnetProtobufEncoder) Encode(ioSession *core.IOSession, obj interface{}, protocolEncoderOutput core.ProtocolEncoderOutput) (result bool, err error) {
	message := []byte(fmt.Sprintf("%v", reflect.ValueOf(obj).String()))
	er := protocolEncoderOutput.Write(message)
	if er == nil {
		return true, nil
	}
	return false, er
}

