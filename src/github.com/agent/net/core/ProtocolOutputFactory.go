package core

type ProtocolOutputFactory struct {
	protocolDecoderOutput ProtocolDecoderOutput
	protocolEncoderOutput ProtocolEncoderOutput
}

func newProtocolOutputFactory(protocolDecoderOutput ProtocolDecoderOutput, protocolEncoderOutput ProtocolEncoderOutput) *ProtocolOutputFactory {
	return &ProtocolOutputFactory{protocolDecoderOutput, protocolEncoderOutput}
}
