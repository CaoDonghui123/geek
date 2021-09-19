package week09

//1.总结几种 socket 粘包的解包方式: fix length/delimiter based/length field based frame decoder。尝试举例其应用
//(1)fix length：累计读取到固定长度为LENGTH之后就认为读取到了一个完整的消息。然后将计数器复位，重新开始读下一个数据报文。
//(2)delimiter based：基于分隔符对字节分割，比如回车或者其他自定义的分隔符。 如读取文本协议时吗，一般用 \r\n 做分隔符，以实现按行读取
//(3)length field based frame decoder：通过指定长度来标识整包消息，这样就可以自动的处理黏包和半包消息，只要传入正确的参数，就可以轻松解决“读半包”的问题。
//类似 goim 、NSQ、ECTD等协议都是通过这种方式。
// 实现一个从 socket connection 中解码出 goim 协议的解码器
func main(){
	//TODO 暂时没思路
}
