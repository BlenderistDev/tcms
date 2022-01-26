package action

import (
	"github.com/golang/mock/gomock"
	"tcms/m/internal/db/model"
	"tcms/m/internal/dry"
	mock_interfaces "tcms/m/internal/testing/action/interfaces"
	mock_interfaces2 "tcms/m/internal/testing/automation/interfaces"
	"testing"
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
