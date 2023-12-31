// Code generated by Thrift Compiler (0.18.1). DO NOT EDIT.

package multiply

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"time"
	thrift "github.com/apache/thrift/lib/go/thrift"
	"strings"
	"regexp"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = errors.New
var _ = context.Background
var _ = time.Now
var _ = bytes.Equal
// (needed by validator.)
var _ = strings.Contains
var _ = regexp.MatchString

type Int int32

func IntPtr(v Int) *Int { return &v }

type MultiplicationService interface {
  // Parameters:
  //  - N1
  //  - N2
  Multiply(ctx context.Context, n1 Int, n2 Int) (_r Int, _err error)
}

type MultiplicationServiceClient struct {
  c thrift.TClient
  meta thrift.ResponseMeta
}

func NewMultiplicationServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *MultiplicationServiceClient {
  return &MultiplicationServiceClient{
    c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
  }
}

func NewMultiplicationServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *MultiplicationServiceClient {
  return &MultiplicationServiceClient{
    c: thrift.NewTStandardClient(iprot, oprot),
  }
}

func NewMultiplicationServiceClient(c thrift.TClient) *MultiplicationServiceClient {
  return &MultiplicationServiceClient{
    c: c,
  }
}

func (p *MultiplicationServiceClient) Client_() thrift.TClient {
  return p.c
}

func (p *MultiplicationServiceClient) LastResponseMeta_() thrift.ResponseMeta {
  return p.meta
}

func (p *MultiplicationServiceClient) SetLastResponseMeta_(meta thrift.ResponseMeta) {
  p.meta = meta
}

// Parameters:
//  - N1
//  - N2
func (p *MultiplicationServiceClient) Multiply(ctx context.Context, n1 Int, n2 Int) (_r Int, _err error) {
  var _args0 MultiplicationServiceMultiplyArgs
  _args0.N1 = n1
  _args0.N2 = n2
  var _result2 MultiplicationServiceMultiplyResult
  var _meta1 thrift.ResponseMeta
  _meta1, _err = p.Client_().Call(ctx, "multiply", &_args0, &_result2)
  p.SetLastResponseMeta_(_meta1)
  if _err != nil {
    return
  }
  return _result2.GetSuccess(), nil
}

type MultiplicationServiceProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler MultiplicationService
}

func (p *MultiplicationServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *MultiplicationServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *MultiplicationServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewMultiplicationServiceProcessor(handler MultiplicationService) *MultiplicationServiceProcessor {

  self3 := &MultiplicationServiceProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self3.processorMap["multiply"] = &multiplicationServiceProcessorMultiply{handler:handler}
return self3
}

func (p *MultiplicationServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err2 := iprot.ReadMessageBegin(ctx)
  if err2 != nil { return false, thrift.WrapTException(err2) }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(ctx, seqId, iprot, oprot)
  }
  iprot.Skip(ctx, thrift.STRUCT)
  iprot.ReadMessageEnd(ctx)
  x4 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(ctx, name, thrift.EXCEPTION, seqId)
  x4.Write(ctx, oprot)
  oprot.WriteMessageEnd(ctx)
  oprot.Flush(ctx)
  return false, x4

}

type multiplicationServiceProcessorMultiply struct {
  handler MultiplicationService
}

func (p *multiplicationServiceProcessorMultiply) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  var _write_err5 error
  args := MultiplicationServiceMultiplyArgs{}
  if err2 := args.Read(ctx, iprot); err2 != nil {
    iprot.ReadMessageEnd(ctx)
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err2.Error())
    oprot.WriteMessageBegin(ctx, "multiply", thrift.EXCEPTION, seqId)
    x.Write(ctx, oprot)
    oprot.WriteMessageEnd(ctx)
    oprot.Flush(ctx)
    return false, thrift.WrapTException(err2)
  }
  iprot.ReadMessageEnd(ctx)

  tickerCancel := func() {}
  // Start a goroutine to do server side connectivity check.
  if thrift.ServerConnectivityCheckInterval > 0 {
    var cancel context.CancelFunc
    ctx, cancel = context.WithCancel(ctx)
    defer cancel()
    var tickerCtx context.Context
    tickerCtx, tickerCancel = context.WithCancel(context.Background())
    defer tickerCancel()
    go func(ctx context.Context, cancel context.CancelFunc) {
      ticker := time.NewTicker(thrift.ServerConnectivityCheckInterval)
      defer ticker.Stop()
      for {
        select {
        case <-ctx.Done():
          return
        case <-ticker.C:
          if !iprot.Transport().IsOpen() {
            cancel()
            return
          }
        }
      }
    }(tickerCtx, cancel)
  }

  result := MultiplicationServiceMultiplyResult{}
  if retval, err2 := p.handler.Multiply(ctx, args.N1, args.N2); err2 != nil {
    tickerCancel()
    err = thrift.WrapTException(err2)
    if errors.Is(err2, thrift.ErrAbandonRequest) {
      return false, thrift.WrapTException(err2)
    }
    _exc6 := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing multiply: " + err2.Error())
    if err2 := oprot.WriteMessageBegin(ctx, "multiply", thrift.EXCEPTION, seqId); err2 != nil {
      _write_err5 = thrift.WrapTException(err2)
    }
    if err2 := _exc6.Write(ctx, oprot); _write_err5 == nil && err2 != nil {
      _write_err5 = thrift.WrapTException(err2)
    }
    if err2 := oprot.WriteMessageEnd(ctx); _write_err5 == nil && err2 != nil {
      _write_err5 = thrift.WrapTException(err2)
    }
    if err2 := oprot.Flush(ctx); _write_err5 == nil && err2 != nil {
      _write_err5 = thrift.WrapTException(err2)
    }
    if _write_err5 != nil {
      return false, thrift.WrapTException(_write_err5)
    }
    return true, err
  } else {
    result.Success = &retval
  }
  tickerCancel()
  if err2 := oprot.WriteMessageBegin(ctx, "multiply", thrift.REPLY, seqId); err2 != nil {
    _write_err5 = thrift.WrapTException(err2)
  }
  if err2 := result.Write(ctx, oprot); _write_err5 == nil && err2 != nil {
    _write_err5 = thrift.WrapTException(err2)
  }
  if err2 := oprot.WriteMessageEnd(ctx); _write_err5 == nil && err2 != nil {
    _write_err5 = thrift.WrapTException(err2)
  }
  if err2 := oprot.Flush(ctx); _write_err5 == nil && err2 != nil {
    _write_err5 = thrift.WrapTException(err2)
  }
  if _write_err5 != nil {
    return false, thrift.WrapTException(_write_err5)
  }
  return true, err
}


// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - N1
//  - N2
type MultiplicationServiceMultiplyArgs struct {
  N1 Int `thrift:"n1,1" db:"n1" json:"n1"`
  N2 Int `thrift:"n2,2" db:"n2" json:"n2"`
}

func NewMultiplicationServiceMultiplyArgs() *MultiplicationServiceMultiplyArgs {
  return &MultiplicationServiceMultiplyArgs{}
}


func (p *MultiplicationServiceMultiplyArgs) GetN1() Int {
  return p.N1
}

func (p *MultiplicationServiceMultiplyArgs) GetN2() Int {
  return p.N2
}
func (p *MultiplicationServiceMultiplyArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField2(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *MultiplicationServiceMultiplyArgs)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  temp := Int(v)
  p.N1 = temp
}
  return nil
}

func (p *MultiplicationServiceMultiplyArgs)  ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  temp := Int(v)
  p.N2 = temp
}
  return nil
}

func (p *MultiplicationServiceMultiplyArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "multiply_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
    if err := p.writeField2(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *MultiplicationServiceMultiplyArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "n1", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:n1: ", p), err) }
  if err := oprot.WriteI32(ctx, int32(p.N1)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.n1 (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:n1: ", p), err) }
  return err
}

func (p *MultiplicationServiceMultiplyArgs) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "n2", thrift.I32, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:n2: ", p), err) }
  if err := oprot.WriteI32(ctx, int32(p.N2)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.n2 (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:n2: ", p), err) }
  return err
}

func (p *MultiplicationServiceMultiplyArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("MultiplicationServiceMultiplyArgs(%+v)", *p)
}

// Attributes:
//  - Success
type MultiplicationServiceMultiplyResult struct {
  Success *Int `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewMultiplicationServiceMultiplyResult() *MultiplicationServiceMultiplyResult {
  return &MultiplicationServiceMultiplyResult{}
}

var MultiplicationServiceMultiplyResult_Success_DEFAULT Int
func (p *MultiplicationServiceMultiplyResult) GetSuccess() Int {
  if !p.IsSetSuccess() {
    return MultiplicationServiceMultiplyResult_Success_DEFAULT
  }
return *p.Success
}
func (p *MultiplicationServiceMultiplyResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *MultiplicationServiceMultiplyResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField0(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *MultiplicationServiceMultiplyResult)  ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 0: ", err)
} else {
  temp := Int(v)
  p.Success = &temp
}
  return nil
}

func (p *MultiplicationServiceMultiplyResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "multiply_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *MultiplicationServiceMultiplyResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin(ctx, "success", thrift.I32, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := oprot.WriteI32(ctx, int32(*p.Success)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(ctx); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *MultiplicationServiceMultiplyResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("MultiplicationServiceMultiplyResult(%+v)", *p)
}


