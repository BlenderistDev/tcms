package action

import (
	"fmt"
	"strconv"
	"tcms/m/internal/automation/core"
	"tcms/m/internal/db/model"
)

type DataMapper struct {
	Action model.Action
}

func (a DataMapper) getFromMapInt64(trigger core.Trigger, key string) (int64, error) {
	s, err := a.getFromMap(trigger, key)
	if err != nil {
		return 0, err
	}

	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, err
	}

	return i, nil
}

func (a DataMapper) getFromMapInt32(trigger core.Trigger, key string) (int32, error) {
	i, err := a.getFromMapInt64(trigger, key)
	return int32(i), err
}

func (a DataMapper) getFromMap(trigger core.Trigger, key string) (string, error) {
	mappingData, ok := a.Action.Mapping[key]
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
