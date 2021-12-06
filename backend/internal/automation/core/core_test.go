package core

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"tcms/m/internal/automation/interfaces"
	mock_interfaces "tcms/m/internal/testing/automation/interfaces"
	"testing"
)

func TestAutomation_ExecuteNoCondition(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trigger := mock_interfaces.NewMockTrigger(ctrl)

	action := mock_interfaces.NewMockAction(ctrl)
	action.
		EXPECT().
		Execute(gomock.Eq(trigger))

	actions := []interfaces.Action{action}

	automation := Automation{Actions: actions}

	automation.Execute(trigger)
}

func TestAutomation_ExecuteManyActions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trigger := mock_interfaces.NewMockTrigger(ctrl)

	action1 := mock_interfaces.NewMockAction(ctrl)
	action1.
		EXPECT().
		Execute(gomock.Eq(trigger))

	action2 := mock_interfaces.NewMockAction(ctrl)
	action2.
		EXPECT().
		Execute(gomock.Eq(trigger))

	actions := []interfaces.Action{action1, action2}

	automation := Automation{Actions: actions}

	automation.Execute(trigger)
}

func TestAutomation_ExecuteWithConditionTrue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trigger := mock_interfaces.NewMockTrigger(ctrl)

	action := mock_interfaces.NewMockAction(ctrl)
	action.
		EXPECT().
		Execute(gomock.Eq(trigger))

	condition := mock_interfaces.NewMockCondition(ctrl)
	condition.
		EXPECT().
		Check(gomock.Eq(trigger)).
		Return(true, nil)

	actions := []interfaces.Action{action}

	automation := Automation{Actions: actions, Condition: condition}

	automation.Execute(trigger)
}

func TestAutomation_ExecuteWithConditionFalse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trigger := mock_interfaces.NewMockTrigger(ctrl)

	action := mock_interfaces.NewMockAction(ctrl)

	condition := mock_interfaces.NewMockCondition(ctrl)
	condition.
		EXPECT().
		Check(gomock.Eq(trigger)).
		Return(false, nil)

	actions := []interfaces.Action{action}

	automation := Automation{Actions: actions, Condition: condition}

	automation.Execute(trigger)
}

func TestAutomation_ExecuteWithConditionError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trigger := mock_interfaces.NewMockTrigger(ctrl)

	action := mock_interfaces.NewMockAction(ctrl)

	condition := mock_interfaces.NewMockCondition(ctrl)
	condition.
		EXPECT().
		Check(gomock.Eq(trigger)).
		Return(true, fmt.Errorf("some error"))

	actions := []interfaces.Action{action}

	automation := Automation{Actions: actions, Condition: condition}

	automation.Execute(trigger)
}
