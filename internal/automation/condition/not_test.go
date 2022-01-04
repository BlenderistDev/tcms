package condition

import (
	"github.com/golang/mock/gomock"
	"tcms/m/internal/automation/datamapper"
	"tcms/m/internal/automation/interfaces"
	"tcms/m/internal/dry"
	mock_interfaces "tcms/m/internal/testing/automation/interfaces"
	"testing"
)

func TestNotCondition_createNotCondition(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	subCondition := mock_interfaces.NewMockCondition(ctrl)

	subConditions := []interfaces.Condition{subCondition}
	createdCondition, err := createNotCondition(datamapper.DataMapper{}, subConditions)
	dry.TestHandleError(t, err)

	switch condition := createdCondition.(type) {
	case notCondition:
		dry.TestCheckEqual(t, subCondition, condition.subCondition)
	default:
		t.Errorf("condition type is not notCondition")
	}
}

func TestNotCondition_CheckWithTrueSubcondition(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	trigger := mock_interfaces.NewMockTrigger(ctrl)
	subCondition := mock_interfaces.NewMockCondition(ctrl)

	subCondition.
		EXPECT().
		Check(gomock.Eq(trigger)).
		Return(true, nil)

	subConditions := []interfaces.Condition{subCondition}
	createdCondition, err := createNotCondition(datamapper.DataMapper{}, subConditions)
	dry.TestHandleError(t, err)

	res, err := createdCondition.Check(trigger)
	dry.TestHandleError(t, err)
	dry.TestCheckEqual(t, false, res)
}
