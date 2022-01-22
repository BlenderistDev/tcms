package condition

import (
	"fmt"
	"tcms/m/internal/automation/datamapper"
	"tcms/m/internal/automation/interfaces"
	"tcms/m/internal/db/model"
)

func CreateCondition(conditionData *model.Condition) (interfaces.Condition, error) {
	var condition interfaces.Condition
	var err error

	dm := datamapper.DataMapper{Mapping: model.ConvertMappingToDmMapping(conditionData.Mapping)}

	var subConditions []interfaces.Condition

	for _, subCondition := range conditionData.SubConditions {
		c, err := CreateCondition(&subCondition)
		if err != nil {
			return nil, err
		}
		subConditions = append(subConditions, c)
	}

	switch conditionData.Name {
	case "equal":
		condition = createEqualCondition(dm, subConditions)
	case "not":
		if len(subConditions) != 1 {
			return nil, fmt.Errorf("not condition can have only one subcondition")
		}
		condition = createNotCondition(subConditions[0])
	case "or":
		condition, err = createOrCondition(subConditions)
		if err != nil {
			return nil, err
		}
	case "and":
		condition, err = createAndCondition(subConditions)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown action %s", conditionData.Name)
	}
	return condition, nil
}
