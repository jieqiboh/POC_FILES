package constants

type Service struct {
	SERVICE_NAME        string
	UPSTREAM_URL        string
	LOAD_BALANCING_TYPE string
}

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
