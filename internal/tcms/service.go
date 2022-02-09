package tcms

import (
	"context"
	"google.golang.org/grpc"
	"net"
	"tcms/m/internal/connections/db/model"
	"tcms/m/internal/connections/db/repository"
	"tcms/m/pkg/tcms"
)

type gRPCServer struct {
	tcms.UnimplementedTcmsServer
	repo repository.AutomationRepository
}

func (s gRPCServer) AddAutomation(ctx context.Context, automation *tcms.Automation) (*tcms.Result, error) {

	actions := make([]model.Action, len(automation.GetActions()))
	for key, act := range automation.GetActions() {
		mapping := act.GetMapping()
		newMapping := make(map[string]model.Mapping, len(act.GetMapping()))
		for _, m := range mapping {
			newMapping[m.GetName()] = model.Mapping{
				Simple: m.GetSimple(),
				Name:   m.GetName(),
				Value:  m.GetValue(),
			}
		}

		action := model.Action{
			Name:    act.GetName(),
			Mapping: newMapping,
		}

		actions[key] = action
	}

	record := model.Automation{
		Triggers:  automation.GetTriggers(),
		Condition: nil,
		Actions:   actions,
	}

	err := s.repo.Save(ctx, record)
	if err != nil {
		return nil, err
	}

	return &tcms.Result{}, nil
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
