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
	record := model.NewAutomation{
		Triggers:  automation.GetTriggers(),
		Actions:   getActions(automation),
		Condition: createCondition(automation.GetCondition()),
	}

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

func getActions(automation *tcms.Automation) []model.Action {
	actions := make([]model.Action, len(automation.GetActions()))
	for key, act := range automation.GetActions() {
		action := model.Action{
			Name:    act.GetName(),
			Mapping: convertMapping(act.GetMapping()),
		}
		actions[key] = action
	}
	return actions
}

func convertMapping(mapping map[string]*tcms.Mapping) map[string]model.Mapping {
	newMapping := make(map[string]model.Mapping, len(mapping))
	for _, m := range mapping {
		newMapping[m.GetName()] = model.Mapping{
			Simple: m.GetSimple(),
			Name:   m.GetName(),
			Value:  m.GetValue(),
		}
	}
	return newMapping
}

func convertSubConditions(list []*tcms.Condition) []model.Condition {
	if len(list) == 0 {
		return nil
	}
	subConditions := make([]model.Condition, len(list))
	for i, c := range list {
		subConditions[i] = *createCondition(c)
	}

	return subConditions
}

func createCondition(c *tcms.Condition) *model.Condition {
	return &model.Condition{
		Name:          c.GetName(),
		Mapping:       convertMapping(c.GetMapping()),
		SubConditions: convertSubConditions(c.GetSubConditions()),
	}
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
