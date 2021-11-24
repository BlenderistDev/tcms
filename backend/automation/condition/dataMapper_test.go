package condition

import (
	"github.com/golang/mock/gomock"
	"tcms/m/automation/core"
	"tcms/m/db/model"
	"testing"
)

func TestGetFromMap_simpleMapping(t *testing.T) {
	const name = "name"
	const value = "value"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mapping := map[string]model.Mapping{
		name: {
			Simple: true,
			Name:   name,
			Value:  value,
		},
	}
	condition := model.Condition{
		Name:    "test",
		Mapping: mapping,
	}
	trigger := core.NewMockTrigger(ctrl)

	datamapper := DataMapper{Condition: condition}

	mapValue, err := datamapper.getFromMap(trigger, "name")
	if err != nil {
		t.Error(err)
	}

	if mapValue != value {
		t.Errorf("expected: %s, actual: %s", value, mapValue)
	}
}

func TestGetFromMap_notSimpleMapping(t *testing.T) {
	const value = "value"
	const name = "name"
	const resultValue = "test_value"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mapping := map[string]model.Mapping{
		name: {
			Simple: false,
			Name:   name,
			Value:  value,
		},
	}

	condition := model.Condition{
		Name:    "test",
		Mapping: mapping,
	}

	triggerData := map[string]string{
		"value": resultValue,
	}

	trigger := core.NewMockTrigger(ctrl)
	trigger.
		EXPECT().
		GetData().
		Return(triggerData)

	datamapper := DataMapper{Condition: condition}

	mapValue, err := datamapper.getFromMap(trigger, name)
	if err != nil {
		t.Error(err)
	}

	if mapValue != resultValue {
		t.Errorf("expected: %s, actual: %s", resultValue, mapValue)
	}
}

func TestGetFromInt_simpleMapping(t *testing.T) {
	const name = "name"
	const value = "123"
	const valueInt = 123

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mapping := map[string]model.Mapping{
		name: {
			Simple: true,
			Name:   name,
			Value:  value,
		},
	}
	condition := model.Condition{
		Name:    "test",
		Mapping: mapping,
	}
	trigger := core.NewMockTrigger(ctrl)

	datamapper := DataMapper{Condition: condition}

	mapValue, err := datamapper.getFromMapInt(trigger, "name")
	if err != nil {
		t.Error(err)
	}

	if mapValue != valueInt {
		t.Errorf("expected: %d, actual: %d", valueInt, mapValue)
	}
}

func TestGetFromMapInt_notSimpleMapping(t *testing.T) {
	const value = "value"
	const name = "name"
	const resultValue = "123"
	const resultValueInt = 123

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mapping := map[string]model.Mapping{
		name: {
			Simple: false,
			Name:   name,
			Value:  value,
		},
	}

	condition := model.Condition{
		Name:    "test",
		Mapping: mapping,
	}

	triggerData := map[string]string{
		"value": resultValue,
	}

	trigger := core.NewMockTrigger(ctrl)
	trigger.
		EXPECT().
		GetData().
		Return(triggerData)

	datamapper := DataMapper{Condition: condition}

	mapValue, err := datamapper.getFromMapInt(trigger, name)
	if err != nil {
		t.Error(err)
	}

	if mapValue != resultValueInt {
		t.Errorf("expected: %d, actual: %d", resultValueInt, mapValue)
	}
}
