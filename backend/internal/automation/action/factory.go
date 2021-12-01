package action

import (
	"fmt"
	"tcms/m/internal/automation/core"
	"tcms/m/internal/db/model"
)

func CreateAction(actionData model.Action) (core.Action, error) {
	var action core.Action
	switch actionData.Name {
	case "sendMessage":
		action = createSendMessageAction(actionData)
	default:
		return nil, fmt.Errorf("unknown action %s", actionData.Name)
	}
	return action, nil
}
