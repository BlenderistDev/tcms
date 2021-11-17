package core

import (
	"github.com/golang/mock/gomock"
	"testing"
)

func TestAutomation_Execute(t *testing.T) {
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
