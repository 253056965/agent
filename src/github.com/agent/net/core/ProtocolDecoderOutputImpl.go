package core

import "github.com/Damnever/goqueue"

const nomalsize = 20

//ProtocolDecoderOutputImpl 这个类是对ProtocolDecoderOutput 的实现
type ProtocolDecoderOutputImpl struct {
	messageQueue *goqueue.Queue
}

//Write 把接完的数据包放到队列里面  这样可以不阻塞 解包过程 增加程序的利用
func (p *ProtocolDecoderOutputImpl) Write(message interface{}) error {
	return p.messageQueue.PutNoWait(message)
}

//flus 把数据 传送给 玩家的业务层
func (p *ProtocolDecoderOutputImpl) flush(ioHandler IoHandler, ioSession *IOSession) {
	for {
		if !ioSession.IsConection() {
			break
		}
		if !p.messageQueue.IsEmpty() {
			mes, er := p.messageQueue.Get(10)
			if er != nil {
				continue
			}
			go ioHandler.MessageReceived(ioSession, mes)
		}
	}
}

//NewProtocolDecoderOutputImplForSize 得到一个输出类  这个方法主要用来对应自己的业务来调优
func NewProtocolDecoderOutputImplForSize(queueSize int) *ProtocolDecoderOutputImpl {
	protocolDecoderOutputImpl := new(ProtocolDecoderOutputImpl)
	protocolDecoderOutputImpl.messageQueue = goqueue.New(queueSize)
	return protocolDecoderOutputImpl
}

//NewProtocolDecoderOutputImpl // 默认初始化 大小为nomalsize 的输出队列
func NewProtocolDecoderOutputImpl() *ProtocolDecoderOutputImpl {
	return NewProtocolDecoderOutputImplForSize(nomalsize)
}
