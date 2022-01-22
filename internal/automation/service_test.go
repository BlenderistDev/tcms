package automation

import (
	"tcms/m/internal/automation/core"
	"tcms/m/internal/dry"
	"testing"
)

func TestService_AddAutomation(t *testing.T) {
	const (
		t1 = "trigger1"
		t2 = "trigger2"
	)
	automation := core.Automation{}
	automation.AddTrigger(t1)
	automation.AddTrigger(t2)
	expected := map[string][]core.Automation{
		t1: {automation},
		t2: {automation},
	}
	s := Service{}
	s.AddAutomation(automation)
	dry.TestCheckEqual(t, expected, s.list)
}
