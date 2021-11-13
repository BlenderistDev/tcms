package action

import (
	"fmt"
	"tcms/m/automation/core"
)

func CreateAction(name string) (core.Action, error) {
	var action core.Action
	switch name {
	case "sendMessage":
		action = CreateSendMessageAction()
	default:
		return nil, fmt.Errorf("unknown action %s", name)
	}
	return action, nil
}
