package core

//ProtocolEncoderOutput 这是一个编码工厂的输出类
type ProtocolEncoderOutput interface{
    Write(encodedMessage interface{}) error;
    //MergeAll();
    flush(ioHandler IoHandler,  ioSession *IOSession);
}