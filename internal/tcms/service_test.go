package tcms

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"tcms/internal/dry"
	"tcms/internal/model"
	mock_repository "tcms/internal/testing/repository"
	"tcms/pkg/tcms"
)

func TestGRPCServer_AddAutomation(t *testing.T) {
	inputAutomation := getInputAutomation()
	outputAutomation := getOutputAutomation()

	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_repository.NewMockAutomationRepository(ctrl)
	repo.EXPECT().Save(gomock.Eq(ctx), gomock.Eq(outputAutomation))

	s := gRPCServer{
		UnimplementedTcmsServer: tcms.UnimplementedTcmsServer{},
		repo:                    repo,
	}

	_, err := s.AddAutomation(ctx, inputAutomation)
	dry.TestHandleError(t, err)
}

func TestGRPCServer_AddAutomation_repoReturnError(t *testing.T) {
	inputAutomation := &tcms.Automation{}

	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_repository.NewMockAutomationRepository(ctrl)

	returnError := fmt.Errorf("some error")
	repo.EXPECT().Save(gomock.Eq(ctx), gomock.Any()).Return(returnError)

	s := gRPCServer{
		UnimplementedTcmsServer: tcms.UnimplementedTcmsServer{},
		repo:                    repo,
	}

	_, err := s.AddAutomation(ctx, inputAutomation)
	dry.TestCheckEqual(t, returnError, err)
}

func TestGRPCServer_UpdateAutomation(t *testing.T) {
	const id string = "some_id"

	inputAutomation := getInputAutomation()
	request := &tcms.UpdateAutomationRequest{
		Id:         id,
		Automation: inputAutomation,
	}
	outputAutomation := getOutputAutomation()

	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_repository.NewMockAutomationRepository(ctrl)
	repo.EXPECT().Update(gomock.Eq(ctx), gomock.Eq(id), gomock.Eq(outputAutomation))

	s := gRPCServer{
		UnimplementedTcmsServer: tcms.UnimplementedTcmsServer{},
		repo:                    repo,
	}

	_, err := s.UpdateAutomation(ctx, request)
	dry.TestHandleError(t, err)
}

func TestGRPCServer_UpdateAutomation_repoReturnError(t *testing.T) {
	const id string = "some_id"

	inputAutomation := getInputAutomation()
	request := &tcms.UpdateAutomationRequest{
		Id:         id,
		Automation: inputAutomation,
	}
	outputAutomation := getOutputAutomation()

	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_repository.NewMockAutomationRepository(ctrl)

	returnError := fmt.Errorf("some error")
	repo.EXPECT().Update(gomock.Eq(ctx), gomock.Eq(id), gomock.Eq(outputAutomation)).Return(returnError)

	s := gRPCServer{
		UnimplementedTcmsServer: tcms.UnimplementedTcmsServer{},
		repo:                    repo,
	}

	_, err := s.UpdateAutomation(ctx, request)
	dry.TestCheckEqual(t, returnError, err)
}

func TestGRPCServer_RemoveAutomation(t *testing.T) {
	const id = "some_id"

	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_repository.NewMockAutomationRepository(ctrl)
	repo.EXPECT().Remove(gomock.Eq(ctx), gomock.Eq(id))

	request := &tcms.RemoveAutomationRequest{Id: id}

	s := gRPCServer{
		UnimplementedTcmsServer: tcms.UnimplementedTcmsServer{},
		repo:                    repo,
	}

	_, err := s.RemoveAutomation(ctx, request)
	dry.TestHandleError(t, err)
}

func TestGRPCServer_RemoveAutomation_repoReturnError(t *testing.T) {
	const id = "some_id"

	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_repository.NewMockAutomationRepository(ctrl)

	returnError := fmt.Errorf("some error")
	repo.EXPECT().Remove(gomock.Eq(ctx), gomock.Eq(id)).Return(returnError)

	request := &tcms.RemoveAutomationRequest{Id: id}

	s := gRPCServer{
		UnimplementedTcmsServer: tcms.UnimplementedTcmsServer{},
		repo:                    repo,
	}

	_, err := s.RemoveAutomation(ctx, request)
	dry.TestCheckEqual(t, returnError, err)
}

func getInputAutomation() *tcms.Automation {
	inputAutomation := &tcms.Automation{
		Triggers: []string{"test1", "test2"},
		Condition: &tcms.Condition{
			Name: "conditionName",
			Mapping: map[string]*tcms.Mapping{
				"map1": {
					Simple: false,
					Name:   "map1",
					Value:  "val1",
				},
				"map2": {
					Simple: false,
					Name:   "map2",
					Value:  "val2",
				},
			},
			SubConditions: []*tcms.Condition{
				{
					Name: "conditionName",
					Mapping: map[string]*tcms.Mapping{
						"map3": {
							Simple: false,
							Name:   "map3",
							Value:  "val3",
						},
						"map4": {
							Simple: false,
							Name:   "map4",
							Value:  "val4",
						},
					},
				},
				{
					Name: "conditionName",
					Mapping: map[string]*tcms.Mapping{
						"map5": {
							Simple: false,
							Name:   "map5",
							Value:  "val5",
						},
						"map6": {
							Simple: false,
							Name:   "map6",
							Value:  "val6",
						},
					},
				},
			},
		},
		Actions: []*tcms.Action{
			{
				Name: "act1",
				Mapping: map[string]*tcms.Mapping{
					"map7": {
						Simple: false,
						Name:   "map7",
						Value:  "val7",
					},
					"map8": {
						Simple: false,
						Name:   "map8",
						Value:  "val8",
					},
				},
			},
			{
				Name: "act2",
			},
		},
	}
	return inputAutomation
}

func getOutputAutomation() model.NewAutomation {
	outputAutomation := model.NewAutomation{
		Triggers: []string{"test1", "test2"},
		Condition: &model.Condition{
			Name: "conditionName",
			Mapping: map[string]model.Mapping{
				"map1": {
					Simple: false,
					Name:   "map1",
					Value:  "val1",
				},
				"map2": {
					Simple: false,
					Name:   "map2",
					Value:  "val2",
				},
			},
			SubConditions: []model.Condition{
				{
					Name: "conditionName",
					Mapping: map[string]model.Mapping{
						"map3": {
							Simple: false,
							Name:   "map3",
							Value:  "val3",
						},
						"map4": {
							Simple: false,
							Name:   "map4",
							Value:  "val4",
						},
					},
				},
				{
					Name: "conditionName",
					Mapping: map[string]model.Mapping{
						"map5": {
							Simple: false,
							Name:   "map5",
							Value:  "val5",
						},
						"map6": {
							Simple: false,
							Name:   "map6",
							Value:  "val6",
						},
					},
				},
			},
		},
		Actions: []model.Action{
			{
				Name: "act1",
				Mapping: map[string]model.Mapping{
					"map7": {
						Simple: false,
						Name:   "map7",
						Value:  "val7",
					},
					"map8": {
						Simple: false,
						Name:   "map8",
						Value:  "val8",
					},
				},
			},
			{
				Name: "act2",
			},
		},
	}
	return outputAutomation
}
