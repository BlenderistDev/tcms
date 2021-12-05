package action

import (
	"github.com/golang/mock/gomock"
	"tcms/m/internal/automation/core"
	"tcms/m/internal/db/model"
	"tcms/m/internal/dry"
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
	trigger := core.NewMockTrigger(ctrl)

	datamapper := DataMapper{Mapping: mapping}

	mapValue, err := datamapper.getFromMap(trigger, "name")
	dry.TestHandleError(t, err)

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

	triggerData := map[string]string{
		"value": resultValue,
	}

	trigger := core.NewMockTrigger(ctrl)
	trigger.
		EXPECT().
		GetData().
		Return(triggerData)

	datamapper := DataMapper{Mapping: mapping}

	mapValue, err := datamapper.getFromMap(trigger, name)
	dry.TestHandleError(t, err)

	if mapValue != resultValue {
		t.Errorf("expected: %s, actual: %s", resultValue, mapValue)
	}
}

func TestGetFromInt32_simpleMapping(t *testing.T) {
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

	trigger := core.NewMockTrigger(ctrl)

	datamapper := DataMapper{Mapping: mapping}

	mapValue, err := datamapper.getFromMapInt32(trigger, "name")
	dry.TestHandleError(t, err)

	if mapValue != valueInt {
		t.Errorf("expected: %d, actual: %d", valueInt, mapValue)
	}
}

func TestGetFromMapInt32_notSimpleMapping(t *testing.T) {
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

	triggerData := map[string]string{
		"value": resultValue,
	}

	trigger := core.NewMockTrigger(ctrl)
	trigger.
		EXPECT().
		GetData().
		Return(triggerData)

	datamapper := DataMapper{Mapping: mapping}

	mapValue, err := datamapper.getFromMapInt32(trigger, name)
	dry.TestHandleError(t, err)

	if mapValue != resultValueInt {
		t.Errorf("expected: %d, actual: %d", resultValueInt, mapValue)
	}
}

func TestGetFromInt64_simpleMapping(t *testing.T) {
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

	trigger := core.NewMockTrigger(ctrl)

	datamapper := DataMapper{Mapping: mapping}

	mapValue, err := datamapper.getFromMapInt64(trigger, "name")
	dry.TestHandleError(t, err)

	if mapValue != valueInt {
		t.Errorf("expected: %d, actual: %d", valueInt, mapValue)
	}
}

func TestGetFromMapInt64_notSimpleMapping(t *testing.T) {
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

	triggerData := map[string]string{
		"value": resultValue,
	}

	trigger := core.NewMockTrigger(ctrl)
	trigger.
		EXPECT().
		GetData().
		Return(triggerData)

	datamapper := DataMapper{Mapping: mapping}

	mapValue, err := datamapper.getFromMapInt64(trigger, name)
	dry.TestHandleError(t, err)

	if mapValue != resultValueInt {
		t.Errorf("expected: %d, actual: %d", resultValueInt, mapValue)
	}
}

func TestGetFromInt64_valueNotExist(t *testing.T) {
	const key = "name"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mapping := map[string]model.Mapping{}

	trigger := core.NewMockTrigger(ctrl)

	datamapper := DataMapper{Mapping: mapping}

	_, err := datamapper.getFromMapInt64(trigger, key)
	dry.TestCheckEqual(t, err.Error(), "key "+key+" not found")
}

func TestGetFromInt64_valueIncorrect(t *testing.T) {
	const key = "key"
	const value = "test"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mapping := map[string]model.Mapping{
		key: {
			Simple: true,
			Name:   key,
			Value:  value,
		},
	}

	trigger := core.NewMockTrigger(ctrl)

	datamapper := DataMapper{Mapping: mapping}

	_, err := datamapper.getFromMapInt64(trigger, key)
	dry.TestCheckEqual(t, err.Error(), "strconv.ParseInt: parsing \""+value+"\": invalid syntax")
}

func TestGetFromMap_valueNotExist(t *testing.T) {
	const key = "name"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mapping := map[string]model.Mapping{}

	trigger := core.NewMockTrigger(ctrl)

	datamapper := DataMapper{Mapping: mapping}

	_, err := datamapper.getFromMap(trigger, key)

	dry.TestCheckEqual(t, err.Error(), "key "+key+" not found")
}

func TestGetFromMap_notSimpleMapping_valueNotExist(t *testing.T) {
	const key = "name"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mapping := map[string]model.Mapping{
		key: {
			Simple: false,
			Name:   key,
			Value:  "",
		},
	}

	triggerData := map[string]string{}

	trigger := core.NewMockTrigger(ctrl)
	trigger.
		EXPECT().
		GetData().
		Return(triggerData)

	datamapper := DataMapper{Mapping: mapping}

	_, err := datamapper.getFromMap(trigger, key)

	dry.TestCheckEqual(t, err.Error(), "key "+key+" not found in trigger data")
}
