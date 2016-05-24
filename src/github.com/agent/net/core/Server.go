package core

import (
	"net"
	"strconv"

	"github.com/agent/util"
)

//Server 服务的开启都需要这个类
type Server struct {
	ip                   string               // 如果不指定这个IP  将默认为127.0.0.1
	port                 int                  // 端口号
	ioHandler            IoHandler            // 用户自己定义的处理对象
	protocolCodecFactory ProtocolCodecFactory // 编解码工厂
}

//NewServer2 创建一个server
func NewServer2(ip string, port int, ioHandler IoHandler, protocolCodecFactory ProtocolCodecFactory) *Server {

	return &Server{ip, port, ioHandler, protocolCodecFactory}
}

//NewServer 创建一个server
func NewServer(port int, ioHandler IoHandler, protocolCodecFactory ProtocolCodecFactory) *Server {
	return NewServer2("127.0.0.1", port, ioHandler, protocolCodecFactory)
}

func (s *Server) handler(conn net.Conn) {
	session := NewIOSession(conn, s.protocolCodecFactory, s.ioHandler)
	s.ioHandler.SessionCreated(session)
	protocolDecoderOutput := NewProtocolDecoderOutputImpl()
	// 启动解码数据携程
	go protocolDecoderOutput.flush(s.ioHandler, session)
	protocolEncoderOutput := NewProtocolEncoderOutputImpl()
	// 启动编码数据的携程
	go protocolEncoderOutput.flush(s.ioHandler, session)
	session.setProtocolOutputFactory(newProtocolOutputFactory(protocolDecoderOutput, protocolEncoderOutput))

	s.ioHandler.SessionOpened(session)

	for {
		result, err := s.protocolCodecFactory.GetProtobuDecoder().Decoder(session, session.GetIoBuffer(), protocolDecoderOutput)
		session.SetlastReaderIdleTime()
		if checkError(err, "Connection") == false {
			session.Close()
			break
		}
		if result == false {
			session.Close()
			break
		}
	}

}

//StartServer 启动一个服务器
func (s *Server) StartServer() error {
	logger := util.NewLogger()
	service := s.ip + ":" + strconv.Itoa(s.port)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if !checkError(err, "ResolveTCPAddr") {
		return err
	}

	l, err := net.ListenTCP("tcp", tcpAddr)
	if !checkError(err, "ListenTCP") {
		return err
	}

	logger.Infof("startServer:[%s:%d]", s.ip, s.port)

	for {
		logger.Infoln("Listening...")
		conn, err := l.Accept()
		checkError(err, "Accept")
		logger.Infoln("Accepting...")

		go s.handler(conn)
	}

}

func checkError(err error, info string) bool {
	logger := util.NewLogger()
	if err != nil {
		logger.Errorf("%s:%s", info, err.Error())
		return false
	}
	return true
}
