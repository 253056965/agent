package core

//ProtocolDecoderOutput 解码完毕之后 把用户自己的数据输出
type ProtocolDecoderOutput interface{
    Write(message interface{}) error ;
    flush(ioHandler IoHandler,  ioSession *IOSession);
} 