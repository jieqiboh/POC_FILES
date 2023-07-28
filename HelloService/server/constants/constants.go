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

// var name is [ServiceName]
var HelloService = Service{
	SERVICE_NAME:        "HelloService", //take straight from the idl file
	UPSTREAM_URL:        "0.0.0.0:9000",
	LOAD_BALANCING_TYPE: "ROUND_ROBIN",
}
