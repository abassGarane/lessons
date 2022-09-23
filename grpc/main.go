package main

import (
	"github.com/abassGarane/lessons/grpc/protos"
	"github.com/abassGarane/lessons/grpc/servers"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func main() {
	ll := hclog.Default()
	ss := servers.NewStudentServer(ll)
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		ll.Info("Unable to Create a net Listener")
	}
	gr := grpc.NewServer()
	reflection.Register(gr)
	protos.RegisterStudentServer(gr, ss)

	if err = gr.Serve(l); err != nil {
		ll.Info("Unable to start server:: ", l.Addr())
	}
}
