package condition

import (
	"tcms/m/internal/db/model"
	"tcms/m/internal/dry"
	"testing"
)

func TestCreateSendMessageAction(t *testing.T) {
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
