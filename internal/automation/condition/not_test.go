package condition

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"tcms/m/internal/dry"
	mock_interfaces "tcms/m/internal/testing/automation/interfaces"
	"testing"
)

func TestNotCondition_createNotCondition(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	subCondition := mock_interfaces.NewMockCondition(ctrl)

	createdCondition := CreateNotCondition(subCondition)

	switch condition := createdCondition.(type) {
	case notCondition:
		dry.TestCheckEqual(t, subCondition, condition.subCondition)
	default:
		t.Errorf("condition type is not notCondition")
	}
}

func TestNotCondition_SubConditionError(t *testing.T) {
	const errText = "some error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	trigger := mock_interfaces.NewMockTrigger(ctrl)
	subCondition := mock_interfaces.NewMockCondition(ctrl)

	subCondition.
		EXPECT().
		Check(gomock.Eq(trigger)).
		Return(false, fmt.Errorf(errText))

	createdCondition := CreateNotCondition(subCondition)

	res, err := createdCondition.Check(trigger)
	dry.TestCheckEqual(t, false, res)
	dry.TestCheckEqual(t, errText, err.Error())
}

func TestNotCondition_CheckWithTrueSubCondition(t *testing.T) {
	testNotConditionCheckWithSubCondition(t, true)
}

func TestNotCondition_CheckWithFalseSubCondition(t *testing.T) {
	testNotConditionCheckWithSubCondition(t, false)
}

func testNotConditionCheckWithSubCondition(t *testing.T, subConditionRes bool) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	trigger := mock_interfaces.NewMockTrigger(ctrl)
	subCondition := mock_interfaces.NewMockCondition(ctrl)

	subCondition.
		EXPECT().
		Check(gomock.Eq(trigger)).
		Return(subConditionRes, nil)

	createdCondition := CreateNotCondition(subCondition)

	res, err := createdCondition.Check(trigger)
	dry.TestHandleError(t, err)
	dry.TestCheckEqual(t, !subConditionRes, res)
}
