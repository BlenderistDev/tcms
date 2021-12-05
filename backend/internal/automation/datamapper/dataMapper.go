package datamapper

import (
	"fmt"
	"strconv"
	"tcms/m/internal/automation/core"
	"tcms/m/internal/db/model"
)

type DataMapper struct {
	Mapping map[string]model.Mapping
}

func (a DataMapper) GetFromMapInt64(trigger core.Trigger, key string) (int64, error) {
	s, err := a.GetFromMap(trigger, key)
	if err != nil {
		return 0, err
	}

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}

	return i, nil
}

func (a DataMapper) GetFromMapInt32(trigger core.Trigger, key string) (int32, error) {
	i, err := a.GetFromMapInt64(trigger, key)
	return int32(i), err
}

func (a DataMapper) GetFromMap(trigger core.Trigger, key string) (string, error) {
	mappingData, ok := a.Mapping[key]
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
