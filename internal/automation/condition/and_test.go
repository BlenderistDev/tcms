package condition

import (
	"github.com/golang/mock/gomock"
	"tcms/m/internal/automation/datamapper"
	"tcms/m/internal/automation/interfaces"
	"tcms/m/internal/dry"
	mock_interfaces "tcms/m/internal/testing/automation/interfaces"
	"testing"
)

func TestAndCondition_createAndCondition(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	subCondition1 := mock_interfaces.NewMockCondition(ctrl)
	subCondition2 := mock_interfaces.NewMockCondition(ctrl)

	subConditions := []interfaces.Condition{subCondition1, subCondition2}
	createdCondition, err := createAndCondition(datamapper.DataMapper{}, subConditions)
	dry.TestHandleError(t, err)

	switch condition := createdCondition.(type) {
	case andCondition:
		dry.TestCheckEqual(t, subConditions, condition.subConditions)
	default:
		t.Errorf("condition type is not andCondition")
	}
}
