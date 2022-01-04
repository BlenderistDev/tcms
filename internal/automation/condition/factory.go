package condition

import (
	"fmt"
	"tcms/m/internal/automation/interfaces"
	"tcms/m/internal/db/model"
)

func CreateCondition(conditionData *model.Condition) (interfaces.Condition, error) {
	var condition interfaces.Condition
	var err error
	switch conditionData.Name {
	case "equal":
		condition = createEqualCondition(conditionData)
	case "not":
		condition, err = createNotCondition(conditionData)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown action %s", conditionData.Name)
	}
	return condition, nil
}
