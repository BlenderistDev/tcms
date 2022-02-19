package tcms

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"net"
	"tcms/m/internal/model"
	"tcms/m/internal/repository"
	"tcms/m/pkg/tcms"
)

type gRPCServer struct {
	tcms.UnimplementedTcmsServer
	repo repository.AutomationRepository
}

func (s gRPCServer) AddAutomation(ctx context.Context, automation *tcms.Automation) (*tcms.Result, error) {
	str, err := json.Marshal(automation)
	if err != nil {
		return nil, err
	}

	record := model.NewAutomation{}
	err = json.Unmarshal(str, &record)
	if err != nil {
		return nil, err
	}

	err = s.repo.Save(ctx, record)
	if err != nil {
		return nil, err
	}

	return &tcms.Result{}, nil
}

func (s gRPCServer) GetList(ctx context.Context, _ *emptypb.Empty) (*tcms.AutomationList, error) {
	automationList, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	resp := tcms.AutomationList{}

	for _, a := range automationList {
		s, err := json.Marshal(a)
		if err != nil {
			return nil, err
		}

		automation := tcms.Automation{}
		err = json.Unmarshal(s, &automation)
		if err != nil {
			return nil, err
		}

		resp.AutomationList = append(resp.AutomationList, &automation)
	}

	return &resp, nil
}

func StartTcmsGrpc(repo repository.AutomationRepository) error {
	addr, err := getTcmsHost()
	if err != nil {
		return err
	}

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	tcms.RegisterTcmsServer(s, &gRPCServer{
		repo: repo,
	})

	return s.Serve(lis)
}
