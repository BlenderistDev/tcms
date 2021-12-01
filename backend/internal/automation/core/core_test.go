package core

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestAutomation_ExecuteNoCondition(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trigger := NewMockTrigger(ctrl)

	action := NewMockAction(ctrl)
	action.
		EXPECT().
		Execute(gomock.Eq(trigger))

	actions := []Action{action}

	automation := Automation{Actions: actions}

	automation.Execute(trigger)
}

func TestAutomation_ExecuteManyActions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trigger := NewMockTrigger(ctrl)

	action1 := NewMockAction(ctrl)
	action1.
		EXPECT().
		Execute(gomock.Eq(trigger))

	action2 := NewMockAction(ctrl)
	action2.
		EXPECT().
		Execute(gomock.Eq(trigger))

	actions := []Action{action1, action2}

	automation := Automation{Actions: actions}

	automation.Execute(trigger)
}

func TestAutomation_ExecuteWithConditionTrue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trigger := NewMockTrigger(ctrl)

	action := NewMockAction(ctrl)
	action.
		EXPECT().
		Execute(gomock.Eq(trigger))

	condition := NewMockCondition(ctrl)
	condition.
		EXPECT().
		Check(gomock.Eq(trigger)).
		Return(true, nil)

	actions := []Action{action}

	automation := Automation{Actions: actions, Condition: condition}

	automation.Execute(trigger)
}

func TestAutomation_ExecuteWithConditionFalse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trigger := NewMockTrigger(ctrl)

	action := NewMockAction(ctrl)

	condition := NewMockCondition(ctrl)
	condition.
		EXPECT().
		Check(gomock.Eq(trigger)).
		Return(false, nil)

	actions := []Action{action}

	automation := Automation{Actions: actions, Condition: condition}

	automation.Execute(trigger)
}

func TestAutomation_ExecuteWithConditionError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trigger := NewMockTrigger(ctrl)

	action := NewMockAction(ctrl)

	condition := NewMockCondition(ctrl)
	condition.
		EXPECT().
		Check(gomock.Eq(trigger)).
		Return(true, fmt.Errorf("some error"))

	actions := []Action{action}

	automation := Automation{Actions: actions, Condition: condition}

	automation.Execute(trigger)
}
