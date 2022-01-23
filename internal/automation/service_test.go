package automation

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"tcms/m/internal/automation/interfaces"
	"tcms/m/internal/dry"
	mock_interfaces "tcms/m/internal/testing/automation/interfaces"
	"testing"
	"time"
)

func TestService_AddAutomation(t *testing.T) {
	const (
		t1 = "trigger1"
		t2 = "trigger2"
	)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	automation := mock_interfaces.NewMockAutomation(ctrl)
	automation.EXPECT().GetTriggers().Return([]string{t1, t2})

	expected := map[string][]interfaces.Automation{
		t1: {automation},
		t2: {automation},
	}
	s := Service{}
	s.AddAutomation(automation)
	dry.TestCheckEqual(t, expected, s.list)
}

func TestService_Start(t *testing.T) {
	const t1 = "trigger1"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trigger := mock_interfaces.NewMockTrigger(ctrl)
	trigger.EXPECT().GetName().Return(t1)

	automation := mock_interfaces.NewMockAutomation(ctrl)
	automation.EXPECT().GetTriggers().Return([]string{t1})
	automation.EXPECT().Execute(gomock.Eq(trigger))

	service := Service{}
	service.AddAutomation(automation)

	triggerChan := make(chan interfaces.Trigger)
	errChan := make(chan error)

	go service.Start(triggerChan, errChan)

	triggerChan <- trigger
	time.Sleep(100)
}

func TestService_Start_automationExecuteError(t *testing.T) {
	const (
		t1      = "trigger1"
		errText = "some error"
	)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trigger := mock_interfaces.NewMockTrigger(ctrl)
	trigger.EXPECT().GetName().Return(t1)

	automation := mock_interfaces.NewMockAutomation(ctrl)
	automation.EXPECT().GetTriggers().Return([]string{t1})
	automation.EXPECT().Execute(gomock.Eq(trigger)).Return(fmt.Errorf(errText))

	service := Service{}
	service.AddAutomation(automation)

	triggerChan := make(chan interfaces.Trigger)
	errChan := make(chan error)

	go service.Start(triggerChan, errChan)

	triggerChan <- trigger
	err := <-errChan
	dry.TestCheckEqual(t, errText, err.Error())
}
