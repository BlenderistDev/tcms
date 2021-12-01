package condition

import (
	"fmt"
	"tcms/m/internal/automation/core"
	"tcms/m/internal/db/model"
)

func CreateCondition(conditionData *model.Condition) (core.Condition, error) {
	var condition core.Condition
	switch conditionData.Name {
	case "equal":
		condition = createEqualCondition(conditionData)
	default:
		return nil, fmt.Errorf("unknown action %s", conditionData.Name)
	}
	return condition, nil
}
