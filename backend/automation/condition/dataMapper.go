package condition

import (
	"fmt"
	"strconv"
	"tcms/m/automation/core"
	"tcms/m/db/model"
)

type DataMapper struct {
	Condition *model.Condition
}

func (a DataMapper) getFromMapInt(trigger core.Trigger, key string) (int, error) {
	s, err := a.getFromMap(trigger, key)
	if err != nil {
		return 0, err
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return i, nil
}

func (a DataMapper) getFromMap(trigger core.Trigger, key string) (string, error) {
	mappingData, ok := a.Condition.Mapping[key]
	if ok {
		if mappingData.Simple {
			return mappingData.Value, nil
		} else {
			triggerData := trigger.GetData()
			value, ok := triggerData[mappingData.Value]
			if ok {
				return value, nil
			} else {
				return "", fmt.Errorf("key %s not found in trigger data", key)
			}
		}

	}
	return "", fmt.Errorf("key %s not found", key)
}
