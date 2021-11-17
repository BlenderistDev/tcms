package action

import (
	"github.com/golang/mock/gomock"
	"tcms/m/automation/core"
	"tcms/m/db/model"
	"testing"
)

func TestGetFromMap_simpleMapping(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	const name = "value"

	mapping := map[string]model.Mapping{
		"name": {
			Simple: true,
			Name:   "name",
			Value:  name,
		},
	}
	action := model.Action{
		Name:    "test",
		Mapping: mapping,
	}
	trigger := core.NewMockTrigger(ctrl)

	datamapper := DataMapper{Action: action}

	value, err := datamapper.getFromMap(trigger, "name")
	if err != nil {
		t.Error(err)
	}

	if value != name {
		t.Errorf("expected: %s, actual: %s", name, value)
	}
}
