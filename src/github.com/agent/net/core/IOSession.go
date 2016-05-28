package core

import (
	"errors"
	"net"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/agent/util"
	"gopkg.in/robfig/cron.v2"
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
	ioHandler             IoHandler              // 对方处理方法
	tickID                cron.EntryID           // 定时任务的ID
	timeOut               int64                  // 超时时间设置
	log                   *util.Logger           // 日志对象
	dataMap               map[string]interface{} // 用来保存对象值的
}

//GetID 获取sessionId
func (s *IOSession) GetID() string {
	return s.sessionID
}

// PutUserDate 保存用户的数据
func (s *IOSession) PutUserDate(key string, obj interface{}) {
	s.dataMap[key] = obj
}

// RemoveUserDate 删除用户的数据
func (s *IOSession) RemoveUserDate(key string) {
	if _, ok := s.dataMap[key]; ok {
		delete(s.dataMap, key)
	}
}

//GetUserDate 获取用户设置的数据 第二个返回值表示有值还是没值
func (s *IOSession) GetUserDate(key string) (interface{}, bool) {
	result, ok := s.dataMap[key]
	return result, ok
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
	return "当前sessionId:" + s.sessionID
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
		RemoveSessionJob(s.tickID)
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
func (s *IOSession) timeOutTesting(duration time.Duration) {
	var err error
	s.tickID, err = AddSessionJobForDuration(duration, func() {
		now := time.Now().Unix()
		rt := now - s.lastReaderIdleTime
		wt := now - s.lastWriterIdleTime
		if rt >= s.timeOut && wt >= s.timeOut {
			// 两者都超时了
			s.ioHandler.SessionIdle(s, BOTH_IDLE)

		}

		if rt >= s.timeOut {
			// 读超时
			s.ioHandler.SessionIdle(s, READER_IDLE)

		}
		if wt >= s.timeOut {
			// 写超时
			s.ioHandler.SessionIdle(s, WRITER_IDLE)

		}
	})
	if err != nil {

	}
}

//SetIdleTime 设置超时设置
func (s *IOSession) SetIdleTime(timeOut time.Duration) {
	s.timeOut = timeOut.Nanoseconds() / 1000 / 1000 / 1000
	s.timeOutTesting(timeOut)
}

//NewIOSession 得到一个IOSession 的是实例
func NewIOSession(conn net.Conn, protocolCodecFactory ProtocolCodecFactory, ioHandler IoHandler) *IOSession {
	ioBuffer := NewIoBuffer(conn)
	sessionID := strconv.FormatInt(atomic.AddInt64(&in64, 1), 10)
	log := util.NewLogger()
	dataMap := make(map[string]interface{}, 10)
	iosession := &IOSession{sessionID: sessionID, conn: conn, isConnection: true, ioBuffer: ioBuffer, lastReaderIdleTime: time.Now().Unix(), lastWriterIdleTime: time.Now().Unix(), protocolCodecFactory: protocolCodecFactory, ioHandler: ioHandler, log: log, dataMap: dataMap}
	return iosession
}
