package tcms

import (
	"context"
	"encoding/json"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"tcms/m/internal/automation/action"
	"tcms/m/internal/model"
	"tcms/m/internal/repository"
	"tcms/m/pkg/tcms"
)

type gRPCServer struct {
	tcms.UnimplementedTcmsServer
	repo repository.AutomationRepository
}

func (s gRPCServer) AddAutomation(ctx context.Context, automation *tcms.Automation) (*tcms.Result, error) {
	str, _ := json.Marshal(automation)

	record := model.NewAutomation{}
	_ = json.Unmarshal(str, &record)

	err := s.repo.Save(ctx, record)
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

	str, err := json.Marshal(automationList)
	if err != nil {
		return nil, err
	}

	aList := make([]*tcms.Automation, len(automationList))
	err = json.Unmarshal(str, &aList)
	if err != nil {
		return nil, err
	}

	return &tcms.AutomationList{AutomationList: aList}, nil
}

func (s gRPCServer) RemoveAutomation(ctx context.Context, r *tcms.RemoveAutomationRequest) (*tcms.Result, error) {
	err := s.repo.Remove(ctx, r.Id)
	if err != nil {
		return nil, err
	}

	return &tcms.Result{}, nil
}

func (s gRPCServer) GetActionList(_ context.Context, _ *emptypb.Empty) (*tcms.ActionList, error) {
	return action.GetList(), nil
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
