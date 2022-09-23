package servers

import (
	"context"
	"github.com/abassGarane/lessons/grpc/protos"
	"github.com/hashicorp/go-hclog"
)

type StudentServer struct {
	log hclog.Logger
}

var studentsStore = map[int32]*protos.StudentResponse{}

func NewStudentServer(l hclog.Logger) StudentServer {
	return StudentServer{l}
}
func (s StudentServer) RegisterStudent(ctx context.Context, srq *protos.StudentRequest) (*protos.Response, error) {
	s.log.Info("Handle RegisterStudent:: name", srq.GetName(), "Student AdmissionNUmber", srq.GetAdmNumber())
	age := uint32(2020 - srq.GetYob())
	student := &protos.StudentResponse{
		Form:      srq.GetForm(),
		Name:      srq.GetName(),
		AdmNumber: srq.GetAdmNumber(),
		Age:       age,
	}
	studentsStore[srq.AdmNumber] = student
	return &protos.Response{}, nil
}

func (s StudentServer) GetStudent(ctx context.Context, srq *protos.GetStudentRequest) (*protos.StudentResponse, error) {
	sr, _ := studentsStore[srq.GetAdmNumber()]
	return sr, nil
}
