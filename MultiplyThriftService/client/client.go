package main

import (
	"context"
	"fmt"

	"com.example.multiplyservice/gen-go/multiply"
	thrift "github.com/apache/thrift/lib/go/thrift"
)

func main() {
	transport, err := thrift.NewTSocket("localhost:9090")
	if err != nil {
		fmt.Println("Error opening socket:", err)
		return
	}

	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	// protocolFactory := thrift.NewTCompactProtocolFactory()

	client := multiply.NewMultiplicationServiceClientFactory(transport, protocolFactory)

	if err := transport.Open(); err != nil {
		fmt.Println("Error opening transport:", err)
		return
	}
	defer transport.Close()

	result, err := client.Multiply(context.Background(), 3, 2)
	if err != nil {
		fmt.Println("Error calling Multiply:", err)
		return
	}

	fmt.Println(result)
}
