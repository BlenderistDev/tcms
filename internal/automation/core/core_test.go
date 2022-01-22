package core

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"tcms/m/internal/automation/interfaces"
	"tcms/m/internal/dry"
	mock_interfaces "tcms/m/internal/testing/automation/interfaces"
	"testing"
)

func TestAutomation_ExecuteNoCondition(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trigger := mock_interfaces.NewMockTrigger(ctrl)

	action := mock_interfaces.NewMockActionWithModel(ctrl)
	action.
		EXPECT().
		Execute(gomock.Eq(trigger))

	actions := []interfaces.Action{action}

	automation := Automation{actions: actions}

	err := automation.Execute(trigger)
	dry.TestHandleError(t, err)
}

func TestAutomation_ExecuteManyActions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trigger := mock_interfaces.NewMockTrigger(ctrl)

	action1 := mock_interfaces.NewMockActionWithModel(ctrl)
	action1.
		EXPECT().
		Execute(gomock.Eq(trigger))

	action2 := mock_interfaces.NewMockActionWithModel(ctrl)
	action2.
		EXPECT().
		Execute(gomock.Eq(trigger))

	actions := []interfaces.Action{action1, action2}

	automation := Automation{actions: actions}

	err := automation.Execute(trigger)
	dry.TestHandleError(t, err)
}

func TestAutomation_ExecuteWithConditionTrue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trigger := mock_interfaces.NewMockTrigger(ctrl)

	action := mock_interfaces.NewMockActionWithModel(ctrl)
	action.
		EXPECT().
		Execute(gomock.Eq(trigger))

	condition := mock_interfaces.NewMockCondition(ctrl)
	condition.
		EXPECT().
		Check(gomock.Eq(trigger)).
		Return(true, nil)

	actions := []interfaces.Action{action}

	automation := Automation{actions: actions, condition: condition}

	err := automation.Execute(trigger)
	dry.TestHandleError(t, err)
}

func TestAutomation_ExecuteWithConditionFalse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trigger := mock_interfaces.NewMockTrigger(ctrl)

	action := mock_interfaces.NewMockActionWithModel(ctrl)

	condition := mock_interfaces.NewMockCondition(ctrl)
	condition.
		EXPECT().
		Check(gomock.Eq(trigger)).
		Return(false, nil)

	actions := []interfaces.Action{action}

	automation := Automation{actions: actions, condition: condition}

	err := automation.Execute(trigger)
	dry.TestHandleError(t, err)
}

func TestAutomation_ExecuteWithConditionError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trigger := mock_interfaces.NewMockTrigger(ctrl)

	action := mock_interfaces.NewMockActionWithModel(ctrl)

	condition := mock_interfaces.NewMockCondition(ctrl)
	condition.
		EXPECT().
		Check(gomock.Eq(trigger)).
		Return(true, fmt.Errorf("some error"))

	actions := []interfaces.Action{action}

	automation := Automation{actions: actions, condition: condition}

	err := automation.Execute(trigger)
	dry.TestCheckEqual(t, "some error", err.Error())
}

func TestAutomation_ExecuteWithActionError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trigger := mock_interfaces.NewMockTrigger(ctrl)

	actionError := "some error"
	action := mock_interfaces.NewMockActionWithModel(ctrl)
	action.
		EXPECT().
		Execute(gomock.Eq(trigger)).
		Return(fmt.Errorf(actionError))

	actions := []interfaces.Action{action}

	automation := Automation{actions: actions}

	err := automation.Execute(trigger)
	dry.TestCheckEqual(t, "error while executing action: "+actionError, err.Error())
	if err != nil {
		return
	}
}
