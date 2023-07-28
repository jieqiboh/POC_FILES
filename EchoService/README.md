## Test RPC Microservice  
This is a test RPC microservice that connects to an etcd cluster  

Constants folder:
Contains constant values for stuff like the server ports and etcd ports  

## To run
Do `docker compose up`  
If the etcd cluster already exists, just boot it up  
Navigate into /server  
Do `sh build.sh`  
Do `sh output/bootstrap.sh`  
Navigate into /client
Do `go run client.go` to make a generic json call to the server

## Contents of Constants file
// information about etcd instance
const (
	ETCD_ADDRESS = "0.0.0.0:20000"
)

// info about API Gateway
const (
	GATEWAY_URL              = "0.0.0.0:8888"
	FILEPATH_TO_HELLOSERVICE = "./idl/hello.thrift" //relative to root directory of module!!
	FILEPATH_TO_BIZSERVICE   = "./idl/bizrequests.thrift"
	FILEPATH_TO_ECHO         = "./idl/echo.thrift"
)

// var name is [ServiceName]
var EchoService = Service{
	SERVICE_NAME:        "EchoService",
	UPSTREAM_URL:        "0.0.0.0:8000",
	LOAD_BALANCING_TYPE: "ROUND_ROBIN",
}
