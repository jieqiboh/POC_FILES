package main

import (
	"example/constants"
	api "example/kitex_gen/api/helloservice"
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.ETCD_ADDRESS}) // r should not be reused.
	if err != nil {
		log.Fatal(err)
	}

	ip, _ := net.ResolveTCPAddr("tcp", constants.HelloService.UPSTREAM_URL)

	svr := api.NewServer(
		new(HelloServiceImpl),
		server.WithServiceAddr(ip),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.HelloService.SERVICE_NAME}),
		server.WithRegistry(r),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
