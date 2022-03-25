package action

import (
	"testing"

	mock_interfaces2 "github.com/BlenderistDev/automation/testing/interfaces"
	"github.com/golang/mock/gomock"
	"tcms/m/internal/dry"
	"tcms/m/internal/model"
	mock_interfaces "tcms/m/internal/testing/action/interfaces"
)

func TestGetActionWithModel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trigger := mock_interfaces2.NewMockTrigger(ctrl)

	m := model.Action{}
	action := mock_interfaces.NewMockActionWithModel(ctrl)
	action.EXPECT().Execute(gomock.Eq(m), gomock.Eq(trigger))

	a := GetActionWithModel(action, m)

	err := a.Execute(trigger)
	dry.TestHandleError(t, err)
}
