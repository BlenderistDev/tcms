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

func TestAndCondition_createAndCondition_withLessConditions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	subCondition := mock_interfaces.NewMockCondition(ctrl)
	subConditions := []interfaces.Condition{subCondition}
	_, err := createAndCondition(datamapper.DataMapper{}, subConditions)
	dry.TestCheckEqual(t, "and condition should have at least two subconditions", err.Error())
}

func TestAndCondition_SetConditions_checkResult(t *testing.T) {
	testAndConditionCheckWithSubCondition(t, false, false)
	testAndConditionCheckWithSubCondition(t, false, true)
	testAndConditionCheckWithSubCondition(t, true, false)
	testAndConditionCheckWithSubCondition(t, true, true)
}

func testAndConditionCheckWithSubCondition(t *testing.T, res1, res2 bool) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	trigger := mock_interfaces.NewMockTrigger(ctrl)
	subCondition1 := mock_interfaces.NewMockCondition(ctrl)
	subCondition2 := mock_interfaces.NewMockCondition(ctrl)

	subCondition1.
		EXPECT().
		Check(gomock.Eq(trigger)).
		Return(res1, nil)

	subCondition2.
		EXPECT().
		Check(gomock.Eq(trigger)).
		Return(res2, nil)

	subConditions := []interfaces.Condition{subCondition1, subCondition2}
	createdCondition, err := createAndCondition(datamapper.DataMapper{}, subConditions)
	dry.TestHandleError(t, err)

	res, err := createdCondition.Check(trigger)
	dry.TestHandleError(t, err)
	dry.TestCheckEqual(t, res1 && res2, res)
}
