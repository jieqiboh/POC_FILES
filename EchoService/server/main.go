package main

import (
	"echoservice/constants"
	echoapi "echoservice/kitex_gen/echoapi/echo"
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

	ip, _ := net.ResolveTCPAddr("tcp", constants.EchoService.UPSTREAM_URL)

	svr := echoapi.NewServer(
		new(EchoImpl),
		server.WithServiceAddr(ip),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.EchoService.SERVICE_NAME}),
		server.WithRegistry(r),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
