package main

import (
	"context"
	"echoservice/constants"
	"fmt"
	"log"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	// Parse IDL with Local Files
	// YOUR_IDL_PATH thrift file path, eg:./idl/example.thrift

	p, err := generic.NewThriftFileProvider("../server/echo.thrift")
	if err != nil {
		panic(err)
	}

	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		panic(err)
	}

	r, err := etcd.NewEtcdResolver([]string{constants.ETCD_ADDRESS}) //depends on port of etcd
	if err != nil {
		log.Fatal(err)
	}

	cli, err := genericclient.NewClient(constants.EchoService.SERVICE_NAME, g, client.WithResolver(r))
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	// 'echo' method name must be passed as param
	resp, err := cli.GenericCall(ctx, "echo", "{\"message\": \"hello\"}")
	// resp is a JSON string
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
