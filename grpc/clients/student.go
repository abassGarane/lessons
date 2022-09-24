package main

import (
	"github.com/abassGarane/lessons/grpc/protos"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

type StudentClientServer struct {
	l hclog.Logger
}

func (s *StudentClientServer) Init() protos.StudentClient {
	conn, err := grpc.Dial("localhost:8080")
	if err != nil {
		s.l.Info("Unable to initialize client grpc server")
	}
	cc := protos.NewStudentClient(conn)
	return cc
}
