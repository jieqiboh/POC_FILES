## Test RPC Microservice - EchoService
This is a test RPC microservice that connects to an etcd cluster

Constants folder:
Contains constant values for stuff like the server ports and etcd ports

## To run
Do `docker compose up`  
If the etcd cluster already exists, just boot it up  
Navigate into /server  
Do `sh build.sh`  
Do `sh output/bootstrap.sh`

## To interact with the microservice
Navigate into /client
Do `go run client.go` to make a generic json call to the server

OR

Using postman, make a POST request with a json payload to the address of the microservice.