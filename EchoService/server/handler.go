package main

import (
	"context"
	echoapi "echoservice/kitex_gen/echoapi"
)

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// Echo implements the EchoImpl interface.
func (s *EchoImpl) Echo(ctx context.Context, req *echoapi.Request) (resp *echoapi.Response, err error) {
	// TODO: Your code here...
	return &echoapi.Response{Message: req.Message}, nil
}
