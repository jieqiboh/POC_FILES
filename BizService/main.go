package main

import (
	"genhttpserv/constants"
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/server/genericserver"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {

	r, err := etcd.NewEtcdRegistry([]string{constants.ETCD_ADDRESS}) // should be changed to env
	if err != nil {
		log.Fatal(err)
	}
	p, err := generic.NewThriftFileProvider("bizrequests.thrift")
	if err != nil {
		panic(err)
	}

	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		panic(err)
	}

	ip, _ := net.ResolveTCPAddr("tcp", constants.SERVER_ADDRESS)

	svr := genericserver.NewServer(new(BizServiceImpl), g, server.WithServiceAddr(ip),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "BizService"}),
		server.WithRegistry(r)) //no registry used yet
	if err != nil {
		panic(err)
	}

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
