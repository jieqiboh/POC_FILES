package main

import (
	"context"
	api "example/kitex_gen/api"
)

// HelloServiceImpl implements the last service interface defined in the IDL.
type HelloServiceImpl struct{}

// HelloMethod implements the HelloServiceImpl interface.
func (s *HelloServiceImpl) HelloMethod(ctx context.Context, request *api.HelloReq) (resp *api.HelloResp, err error) {
	// TODO: Your code here...
	return &api.HelloResp{RespBody: "Name is " + request.GetName()}, nil
}
