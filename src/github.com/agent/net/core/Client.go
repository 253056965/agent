package core

import (
	"net"
	"strconv"

	"github.com/agent/util"
)

type Client struct {
	ip                   string               // 如果不指定这个IP  将默认为127.0.0.1
	port                 int                  // 端口号
	ioHandler            IoHandler            // 用户自己定义的处理对象
	protocolCodecFactory ProtocolCodecFactory // 编解码工厂
	log                  *util.Logger
	session              *IOSession // 当前玩家的session 对象
}

// 连接服务器
func (c *Client) Connection() error {
	addr := c.ip + ":" + strconv.Itoa(c.port)
	conn, err := net.Dial("tcp4", addr)
	if err != nil {
		c.log.Errorln("连接出错了:", addr, err)
		return err
	}
	go c.handler(conn)
	return err
}

func (c *Client) handler(conn net.Conn) {

	session := NewIOSession(conn, c.protocolCodecFactory, c.ioHandler)
	c.session = session
	c.ioHandler.SessionCreated(session)
	protocolDecoderOutput := NewProtocolDecoderOutputImpl()
	// 启动解码数据携程
	go protocolDecoderOutput.flush(c.ioHandler, session)
	protocolEncoderOutput := NewProtocolEncoderOutputImpl()
	// 启动编码数据的携程
	go protocolEncoderOutput.flush(c.ioHandler, session)
	session.setProtocolOutputFactory(newProtocolOutputFactory(protocolDecoderOutput, protocolEncoderOutput))
	c.ioHandler.SessionOpened(session)
	for {
		result, err := c.protocolCodecFactory.GetProtobuDecoder().Decoder(session, session.GetIoBuffer(), protocolDecoderOutput)
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
func (c *Client) GetSession() *IOSession {
	return c.session
}

//NewServer2 创建一个server
func NewClient2(ip string, port int, ioHandler IoHandler, protocolCodecFactory ProtocolCodecFactory) *Client {
	log := util.NewLogger()
	return &Client{ip, port, ioHandler, protocolCodecFactory, log, nil}
}

//NewServer 创建一个server
func NewClent(port int, ioHandler IoHandler, protocolCodecFactory ProtocolCodecFactory) *Client {
	return NewClient2("127.0.0.1", port, ioHandler, protocolCodecFactory)
}
