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


