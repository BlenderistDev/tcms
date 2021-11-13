package action

import (
	"fmt"
	"tcms/m/automation/core"
	"tcms/m/db/model"
)

func CreateAction(actionData model.Action) (core.Action, error) {
	var action core.Action
	switch actionData.Name {
	case "sendMessage":
		action = CreateSendMessageAction(actionData)
	default:
		return nil, fmt.Errorf("unknown action %s", actionData.Name)
	}
	return action, nil
}
