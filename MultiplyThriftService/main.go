package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"com.example.multiplyservice/gen-go/multiply"

	thrift "github.com/apache/thrift/lib/go/thrift"
)

type MultiplyHandler struct{}

func (h *MultiplyHandler) Multiply(ctx context.Context, N1 multiply.Int, N2 multiply.Int) (multiply.Int, error) {
	log.Print("Received Call")
	return N1 * N2, nil
}

func main() {
	handler := &MultiplyHandler{} //the handler is an object with all the necessary methods defined in the service implemented
	processor := multiply.NewMultiplicationServiceProcessor(handler)
	//if using the TBinaryProtocol
	transport, err := thrift.NewTServerSocket(":9090")
	if err != nil {
		fmt.Println("Error opening socket:", err)
		os.Exit(1)
	}
	//if using TCompactProtocol
	if err != nil {
		fmt.Println("Error opening socket:", err)
		os.Exit(1)
	}
	server := thrift.NewTSimpleServer2(processor, transport)
	fmt.Println("Starting the simple server...")
	err = server.Serve()

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
