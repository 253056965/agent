package core

import (
	"bufio"
	"io"
	"net"
)

//BUFFERSIZE 默认缓冲区大小
const BUFFERSIZE = 1024

//IoBuffer 数据底层处理
type IoBuffer struct {
	conn net.Conn      // 当前session 的连接
	rb   *bufio.Reader //buffer read
	wb   io.Writer
}

//ReadMessage  读取缓冲区的数据
func (iob *IoBuffer) ReadMessage(length int) ([]byte, error) {
	buf := make([]byte, length)
	if _, err := io.ReadFull(iob.rb, buf); err != nil {
		return nil, err
	}
	return buf, nil
}

//ReadLine 读取一行数据
func (iob *IoBuffer) ReadLine() (string, error) {
	buf, _, err := iob.rb.ReadLine()
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

// SendMessage 向客户端发送数据
func (iob *IoBuffer) SendMessage(data []byte) error {
	if n, err := iob.wb.Write(data); err != nil {
		return err
	} else if n != len(data) {
		return err
	} else {
		return nil
	}
}

//NewIoBuffer 创建一个IoBuffer
func NewIoBuffer(conn net.Conn) *IoBuffer {
	ioBuffer := new(IoBuffer)
	ioBuffer.conn = conn
	ioBuffer.rb = bufio.NewReaderSize(conn, BUFFERSIZE)
	ioBuffer.wb = conn
	return ioBuffer
}
