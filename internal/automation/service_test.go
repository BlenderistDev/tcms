package automation

import (
	"github.com/golang/mock/gomock"
	"tcms/m/internal/automation/interfaces"
	"tcms/m/internal/dry"
	mock_interfaces "tcms/m/internal/testing/automation/interfaces"
	"testing"
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
