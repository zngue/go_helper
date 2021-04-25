package grpc_run

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"net"
	"time"
)

var (
	Conn *grpc.ClientConn
)

type ServiceFn func(server *grpc.Server)
type ClientFn func(conn *grpc.ClientConn, ctx context.Context) (err error)
type GrpcHelperInterface interface {
	ServiceRegister(host, port string, fn ServiceFn) (err error)
	ClientRegister(host, port string, fn ClientFn) error
}

func (GrpcHelper) NewGrpcHelper() GrpcHelperInterface {
	return new(GrpcHelper)
}

type GrpcOprion struct {
}
type GrpcHelper struct {
	Conn *grpc.ClientConn
	//context.WithTimeout(context.Background(), time.Second/10)
	contextFn func() (context.Context, context.CancelFunc)
}

func (h *GrpcHelper) ClientRegister(host, port string, fn ClientFn) error {
	if h.Conn == nil {
		conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			return err
		}
		Conn = conn
		h.Conn = conn
	}
	defer Conn.Close()
	var ctx context.Context
	var cancel context.CancelFunc
	if h.contextFn != nil {
		ctx, cancel = h.contextFn()
	} else {
		ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	}
	defer cancel()
	if err := fn(Conn, ctx); err != nil {
		return err
	}
	return nil
}

func (h *GrpcHelper) ServiceRegister(host, port string, fn ServiceFn) (err error) {
	listener, err2 := net.Listen("tcp", host+":"+port)
	if err2 != nil {
		return err2
	}
	server := grpc.NewServer()
	fn(server)
	if err := server.Serve(listener); err != nil {
		return err
	}
	return
}

/*
*@Author Administrator
*@Date 23/4/2021 14:00
*@desc
 */
func getServiceRegisterHostBySerName(serviceName string) (string, error) {

	if ServiceHostList == nil {
		return "", errors.New(" Please register the service first ")
	}
	host, ok := ServiceHostList[serviceName]
	if !ok {
		return "", errors.New(serviceName + " service does not exist. Please register the service first")
	}
	return host, nil
}

func ServiceRegister(serviceName string, fn ServiceFn) (err error) {
	host, herr := getServiceRegisterHostBySerName(serviceName)
	if herr != nil {
		return herr
	}
	listener, err2 := net.Listen("tcp", host)
	if err2 != nil {
		return err2
	}
	server := grpc.NewServer()
	fn(server)
	if errs := server.Serve(listener); errs != nil {
		return errs
	}
	return
}
func ClientRegister(serviceName string, fn ClientFn) error {
	host, herr := getServiceRegisterHostBySerName(serviceName)
	if herr != nil {
		return herr
	}
	conn, err := grpc.Dial(host, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return err
	}
	defer Conn.Close()
	outTime := time.Second * 200
	if OutTime != 0 {
		outTime = OutTime
	}
	ctx, _ := context.WithTimeout(context.Background(), outTime)

	//defer cancel()
	if errs := fn(conn, ctx); errs != nil {
		return nil
	}
	return nil
}
