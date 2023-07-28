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
var EchoService = Service{
	SERVICE_NAME:        "Echo", //take straight from the idl file
	UPSTREAM_URL:        "0.0.0.0:8000",
	LOAD_BALANCING_TYPE: "ROUND_ROBIN",
}
