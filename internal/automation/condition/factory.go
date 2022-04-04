package condition

import (
	"fmt"

	condition2 "github.com/BlenderistDev/automation/condition"
	"github.com/BlenderistDev/automation/datamapper"
	"github.com/BlenderistDev/automation/interfaces"
	"tcms/m/internal/model"
	"tcms/m/pkg/tcms"
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
		condition = condition2.CreateEqualCondition(dm)
	case "not":
		if len(subConditions) != 1 {
			return nil, fmt.Errorf("not condition can have only one subcondition")
		}
		condition = condition2.CreateNotCondition(subConditions[0])
	case "or":
		condition, err = condition2.CreateOrCondition(subConditions)
		if err != nil {
			return nil, err
		}
	case "and":
		condition, err = condition2.CreateAndCondition(subConditions)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown condition %s", conditionData.Name)
	}
	return condition, nil
}

func GetList() *tcms.ConditionList {
	return &tcms.ConditionList{Conditions: []*tcms.ConditionDescription{
		{
			Name: "equal",
			Fields: []*tcms.Field{
				{
					Name:     "value1",
					Type:     "string",
					Required: true,
				},
				{
					Name:     "value2",
					Type:     "string",
					Required: true,
				},
			},
		},
		{
			Name:                 "not",
			MinSubConditionCount: 1,
			MaxSubConditionCount: 1,
		},
		{
			Name:                 "or",
			MinSubConditionCount: 2,
		},
		{
			Name:                 "and",
			MinSubConditionCount: 2,
		},
	}}
}
