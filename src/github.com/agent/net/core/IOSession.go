package core

import (
	"errors"
	"net"
	"strconv"
	"sync/atomic"
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
}

//GetID 获取sessionId
func (s *IOSession) GetID() string {
	return s.sessionID
}

//ToString 格式化对象的方法
func (s *IOSession) String() string {
	return s.sessionID
}
func (s *IOSession)setProtocolOutputFactory(protocolOutputFactory *ProtocolOutputFactory)   {
	s.protocolOutputFactory=protocolOutputFactory;
}

func (s *IOSession) GetRemoteAddr() net.Addr {
	return s.conn.RemoteAddr()
}

//Close 关闭当前链接
func (s *IOSession) Close() {
	s.isConnection = false
	//s.conn.SetDeadline
	s.conn.Close()
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
		s.protocolCodecFactory.GetProtobufEncoder().Encode(s, mesg, s.protocolOutputFactory.protocolEncoderOutput)
		return nil
	}
	return errors.New("连接已经断开了")
}

//NewIOSession 得到一个IOSession 的是实例
func NewIOSession(conn net.Conn, protocolCodecFactory ProtocolCodecFactory) *IOSession {
	ioBuffer := NewIoBuffer(conn)
	sessionID := strconv.FormatInt(atomic.AddInt64(&in64, 1), 10)
	return &IOSession{sessionID: sessionID, conn: conn, isConnection: true, ioBuffer: ioBuffer, protocolCodecFactory: protocolCodecFactory}
}
