package main

import ()

const (
	HOST = "localhost"
	PORT = "8080"
)

// thrift -r --gen go echo.thrift
func main() {

	/*tSocket, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))
	if err != nil {
		log.Fatalln("tSocket error:", err)
	}
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	transport := transportFactory.GetTransport(tSocket)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	client := echo.NewFormatDataClientFactory(transport, protocolFactory)

	if err := transport.Open(); err != nil {
		log.Fatalln("Error opening:", HOST + ":" + PORT)
	}
	defer transport.Close()


	ctx, _ := context.WithCancel(context.Background())

	data := echo.Data{Text:"hello,world!"}
	d, err := client.DoFormat(ctx,&data)
	fmt.Println(d.Text)

	*/
}
