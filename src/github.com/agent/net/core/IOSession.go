package core

import (
	"errors"
	"net"
	"strconv"
	"sync/atomic"
	"time"
)

var in64 int64

/*
IOSession 这个session 对象
*/
type IOSession struct {
	sessionID             string               // 当前sessionID
	lastReaderIdleTime    int64                // 最后一次ioSession 的读取时间
	lastWriterIdleTime    int64                // 最后一次ioSession 的访问时间
	conn                  net.Conn             // 当前session 的连接
	isConnection          bool                 // 是不是链接
	ioBuffer              *IoBuffer            // 用来读写数据的
	protocolCodecFactory  ProtocolCodecFactory // 编码输出流
	protocolOutputFactory *ProtocolOutputFactory
	ioHandler             IoHandler    // 对方处理方法
	ticker                *time.Ticker // 接受对象
	timeOut               int64        // 超时时间设置
}

//GetID 获取sessionId
func (s *IOSession) GetID() string {
	return s.sessionID
}

//SetlastReaderIdleTime 设置读的超时时间
func (s *IOSession) SetlastReaderIdleTime() {
	s.lastReaderIdleTime = time.Now().Unix()
}

//SetlastWriterIdleTime 设置写的超时时间
func (s *IOSession) SetlastWriterIdleTime() {
	s.lastReaderIdleTime = time.Now().Unix()
}

//ToString 格式化对象的方法
func (s *IOSession) String() string {
	return s.sessionID
}
func (s *IOSession) setProtocolOutputFactory(protocolOutputFactory *ProtocolOutputFactory) {
	s.protocolOutputFactory = protocolOutputFactory
}

//GetRemoteAddr 获取远程的地址
func (s *IOSession) GetRemoteAddr() net.Addr {
	return s.conn.RemoteAddr()
}

//Close 关闭当前链接
func (s *IOSession) Close() {
	if s.isConnection {
		s.isConnection = false
		s.ticker.Stop()
		s.conn.Close()
		go s.ioHandler.SessionClosed(s)
	}

}

//IsConection 看看当前链接是否关闭
func (s *IOSession) IsConection() bool {
	return s.isConnection
}

//GetIoBuffer 得到当前的IoBuffer
func (s *IOSession) GetIoBuffer() *IoBuffer {
	return s.ioBuffer
}

func (s *IOSession) Write(mesg interface{}) error {
	if s.isConnection {
		s.lastWriterIdleTime = time.Now().Unix()
		s.protocolCodecFactory.GetProtobufEncoder().Encode(s, mesg, s.protocolOutputFactory.protocolEncoderOutput)
		return nil
	}
	return errors.New("连接已经断开了")
}

//此方法是超时检测
func (s *IOSession) timeOutTesting() {
	go func() {
		for _ = range s.ticker.C {
			now := time.Now().Unix()
			rt := now - s.lastReaderIdleTime
			wt := now - s.lastWriterIdleTime

			if rt >= s.timeOut && wt >= s.timeOut {
				// 两者都超时了
				s.ioHandler.SessionIdle(s, BOTH_IDLE)
				continue
			}

			if rt >= s.timeOut {
				// 读超时
				s.ioHandler.SessionIdle(s, READER_IDLE)
				continue
			}

			if wt >= s.timeOut {
				// 写超时
				s.ioHandler.SessionIdle(s, WRITER_IDLE)
				continue
			}
			//fmt.Printf("ticked at %v", time.Now())
		}
	}()
}

//SetIdleTime 设置超时设置
func (s *IOSession) SetIdleTime(idleTime time.Duration) {
	if s.ticker != nil {
		s.ticker.Stop()
	}
	if idleTime <= 0 {
		return
	}
	s.ticker = time.NewTicker(idleTime)
	s.timeOutTesting()
}

//NewIOSession 得到一个IOSession 的是实例
func NewIOSession(conn net.Conn, protocolCodecFactory ProtocolCodecFactory, ioHandler IoHandler) *IOSession {
	ioBuffer := NewIoBuffer(conn)
	sessionID := strconv.FormatInt(atomic.AddInt64(&in64, 1), 10)
	iosession := &IOSession{sessionID: sessionID, conn: conn, isConnection: true, ioBuffer: ioBuffer, lastReaderIdleTime: time.Now().Unix(), lastWriterIdleTime: time.Now().Unix(), protocolCodecFactory: protocolCodecFactory, ioHandler: ioHandler}
	return iosession
}
