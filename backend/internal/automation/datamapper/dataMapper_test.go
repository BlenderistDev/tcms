package datamapper

import (
	"github.com/golang/mock/gomock"
	mock_interfaces "tcms/m/internal/automation/interfaces/mock"
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
	trigger := mock_interfaces.NewMockTrigger(ctrl)

	datamapper := DataMapper{Mapping: mapping}

	mapValue, err := datamapper.GetFromMap(trigger, "name")
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

	trigger := mock_interfaces.NewMockTrigger(ctrl)
	trigger.
		EXPECT().
		GetData().
		Return(triggerData)

	datamapper := DataMapper{Mapping: mapping}

	mapValue, err := datamapper.GetFromMap(trigger, name)
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

	trigger := mock_interfaces.NewMockTrigger(ctrl)

	datamapper := DataMapper{Mapping: mapping}

	mapValue, err := datamapper.GetFromMapInt32(trigger, "name")
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

	trigger := mock_interfaces.NewMockTrigger(ctrl)
	trigger.
		EXPECT().
		GetData().
		Return(triggerData)

	datamapper := DataMapper{Mapping: mapping}

	mapValue, err := datamapper.GetFromMapInt32(trigger, name)
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

	trigger := mock_interfaces.NewMockTrigger(ctrl)

	datamapper := DataMapper{Mapping: mapping}

	mapValue, err := datamapper.GetFromMapInt64(trigger, "name")
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

	trigger := mock_interfaces.NewMockTrigger(ctrl)
	trigger.
		EXPECT().
		GetData().
		Return(triggerData)

	datamapper := DataMapper{Mapping: mapping}

	mapValue, err := datamapper.GetFromMapInt64(trigger, name)
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

	trigger := mock_interfaces.NewMockTrigger(ctrl)

	datamapper := DataMapper{Mapping: mapping}

	_, err := datamapper.GetFromMapInt64(trigger, key)
	dry.TestCheckEqual(t, "key "+key+" not found", err.Error())
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

	trigger := mock_interfaces.NewMockTrigger(ctrl)

	datamapper := DataMapper{Mapping: mapping}

	_, err := datamapper.GetFromMapInt64(trigger, key)
	dry.TestCheckEqual(t, "strconv.ParseInt: parsing \""+value+"\": invalid syntax", err.Error())
}

func TestGetFromMap_valueNotExist(t *testing.T) {
	const key = "name"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mapping := map[string]model.Mapping{}

	trigger := mock_interfaces.NewMockTrigger(ctrl)

	datamapper := DataMapper{Mapping: mapping}

	_, err := datamapper.GetFromMap(trigger, key)

	dry.TestCheckEqual(t, "key "+key+" not found", err.Error())
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

	trigger := mock_interfaces.NewMockTrigger(ctrl)
	trigger.
		EXPECT().
		GetData().
		Return(triggerData)

	datamapper := DataMapper{Mapping: mapping}

	_, err := datamapper.GetFromMap(trigger, key)

	dry.TestCheckEqual(t, "key "+key+" not found in trigger data", err.Error())
}
