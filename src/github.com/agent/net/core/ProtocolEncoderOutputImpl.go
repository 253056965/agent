package core

import "github.com/Damnever/goqueue"
import "reflect"

const SIZE = 20

type ProtocolEncoderOutputImpl struct {
	messageQueue *goqueue.Queue
}

func (p *ProtocolEncoderOutputImpl) Write(encodedMessage interface{}) error {
	return p.messageQueue.PutNoWait(encodedMessage)
}

func (p *ProtocolEncoderOutputImpl) flush(ioHandler IoHandler, ioSession *IOSession) {
	for {
		if !ioSession.IsConection() {
			break
		}
		if !p.messageQueue.IsEmpty() {
			mes, er := p.messageQueue.Get(1)
			if er != nil {
				continue
			}
           go  p.send(mes,ioHandler,ioSession);
		}
	}
}

func (p* ProtocolEncoderOutputImpl) send(message interface{},ioHandler IoHandler, ioSession *IOSession)  {
            bytes:= reflect.ValueOf(message).Bytes();
            ioSession.GetIoBuffer().SendMessage(bytes);
			ioHandler.MessageSent(ioSession,message);
}

func NewProtocolEncoderOutputImplForSize(queueSize int) *ProtocolEncoderOutputImpl {
	protocolEncoderOutputImpl := new(ProtocolEncoderOutputImpl)
	protocolEncoderOutputImpl.messageQueue = goqueue.New(queueSize)
	return protocolEncoderOutputImpl
}
func NewProtocolEncoderOutputImpl() *ProtocolEncoderOutputImpl {
	return NewProtocolEncoderOutputImplForSize(SIZE)
}
