package datamapper

import (
	"fmt"
	"math"
	"strconv"
)

type Mapping interface {
	IsSimple() bool
	GetValue() string
}

type DataMapper struct {
	Mapping map[string]Mapping
}

func (a DataMapper) GetFromMapInt64(data map[string]string, key string) (int64, error) {
	s, err := a.GetFromMap(data, key)
	if err != nil {
		return 0, err
	}

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}

	return i, nil
}

func (a DataMapper) GetFromMapInt32(data map[string]string, key string) (int32, error) {
	i, err := a.GetFromMapInt64(data, key)
	if i > math.MaxInt32 {
		return 0, fmt.Errorf("number %d is greater, than MaxInt32", i)
	}
	return int32(i), err
}

func (a DataMapper) GetFromMap(data map[string]string, key string) (string, error) {
	mappingData, ok := a.Mapping[key]
	if ok {
		if mappingData.IsSimple() {
			return mappingData.GetValue(), nil
		} else {
			value, ok := data[mappingData.GetValue()]
			if ok {
				return value, nil
			} else {
				return "", fmt.Errorf("key %s not found in trigger data", key)
			}
		}

	}
	return "", fmt.Errorf("key %s not found", key)
}

func (a DataMapper) GetFromMapBool(data map[string]string, key string) (bool, error) {
	s, err := a.GetFromMap(data, key)
	if err != nil {
		return false, err
	}
	return s != "", nil
}
