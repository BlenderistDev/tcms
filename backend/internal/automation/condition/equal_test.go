package condition

import (
	"github.com/golang/mock/gomock"
	"tcms/m/internal/automation/core"
	"tcms/m/internal/db/model"
	"tcms/m/internal/dry"
	"testing"
)

func TestEqualCondition(t *testing.T) {
	conditionModel := &model.Condition{
		Name: "name",
		Mapping: map[string]model.Mapping{
			"test": {
				Simple: true,
				Name:   "name",
				Value:  "value",
			}},
	}
	createdCondition := createEqualCondition(conditionModel)

	switch condition := createdCondition.(type) {
	case equalCondition:
		dry.TestCheckEqual(t, conditionModel.Mapping, condition.DataMapper.Mapping)
	default:
		t.Errorf("condition type is not sendMessageAction")
	}
}

func TestEqualCondition_Check(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	trigger := core.NewMockTrigger(ctrl)

	conditionModel := &model.Condition{
		Name: "name",
		Mapping: map[string]model.Mapping{
			"value1": {
				Simple: true,
				Name:   "value1",
				Value:  "value",
			},
			"value2": {
				Simple: true,
				Name:   "value2",
				Value:  "value",
			},
		},
	}
	createdCondition := createEqualCondition(conditionModel)
	res, err := createdCondition.Check(trigger)
	dry.TestHandleError(t, err)
	dry.TestCheckEqual(t, true, res)
}
