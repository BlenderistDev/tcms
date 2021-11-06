package telegramClient

import (
	"encoding/json"
	"github.com/xelaj/mtproto/telegram"
	"reflect"
	"strconv"
	"strings"
	"tcms/m/automation"
	"tcms/m/dry"
)

type GzipedEvent struct {
	Obj interface{}
}

func recognizeTrigger(i interface{}) []automation.TelegramUpdateTrigger {
	var triggerType string
	var triggerList []automation.TelegramUpdateTrigger
	switch message := i.(type) {
	case *telegram.UpdateShort:
		triggerType = getTriggerType(message.Update)
		triggerData := parsePtr(message)
		trigger := automation.TelegramUpdateTrigger{
			Name: triggerType,
			Data: triggerData,
		}
		triggerList = append(triggerList, trigger)
	case *telegram.UpdatesObj:
		for _, event := range message.Updates {
			triggerType = getTriggerType(event)
			triggerData := parsePtr(event)
			trigger := automation.TelegramUpdateTrigger{
				Name: triggerType,
				Data: triggerData,
			}
			triggerList = append(triggerList, trigger)
		}
	default:
		jsonStr, err := json.Marshal(i)
		dry.HandleError(err)
		var jsonData = GzipedEvent{}
		err = json.Unmarshal(jsonStr, &jsonData)
		if err == nil {
			triggerData := parseUnknown(jsonData.Obj)
			trigger := automation.TelegramUpdateTrigger{
				Name: triggerType,
				Data: triggerData,
			}
			triggerList = append(triggerList, trigger)
		} else {
			triggerData := parseUnknown(i)
			trigger := automation.TelegramUpdateTrigger{
				Name: triggerType,
				Data: triggerData,
			}
			triggerList = append(triggerList, trigger)
		}
	}

	return triggerList
}

func getTriggerType(i interface{}) string {
	return strings.Replace(reflect.TypeOf(i).String(), "*telegram.", "", 1)
}

func parsePtr(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	if reflect.ValueOf(i).IsNil() || reflect.ValueOf(i).IsZero() {
		return nil
	}

	v := reflect.Indirect(reflect.ValueOf(i))

	values := make(map[string]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		fieldData := v.Field(i).Interface()
		filedName := v.Type().Field(i).Name
		data := parseUnknown(fieldData, filedName)
		for key, value := range data {
			values[key] = value
		}
	}
	return values
}

func parseUnknown(i interface{}, prefixArr ...string) map[string]interface{} {
	if i == nil {
		return nil
	}

	prefix, originalPrefix := getPrefix(prefixArr...)

	values := make(map[string]interface{})

	valueType := reflect.ValueOf(i).Kind()

	switch valueType {
	case reflect.Ptr:
		data := parsePtr(i)
		for key, value := range data {
			values[prefix+key] = value
		}
	case reflect.Slice:
		data := parseSlice(i)
		if len(data) > 0 {
			for key, value := range data {
				values[prefix+"."+key] = value
			}
		} else {
			values[prefix] = nil
		}

	case reflect.Map:
		data := parseMap(i)
		for key, value := range data {
			values[prefix+key] = value
		}
	case reflect.Array:
		panic("array in recognize!")
	case reflect.Struct:
		data := parseStruct(i)
		for key, value := range data {
			values[prefix+key] = value
		}

	default:
		values[originalPrefix] = i
	}

	return values
}

func parseSlice(i interface{}) map[string]interface{} {

	if i == nil {
		return nil
	}

	listVal := reflect.ValueOf(i)

	values := make(map[string]interface{}, listVal.Len())

	for key := 0; key < listVal.Len(); key++ {
		data := parseUnknown(listVal.Index(key).Interface(), strconv.Itoa(key))
		for key, value := range data {
			values[key] = value
		}
	}

	return values
}

func parseMap(i interface{}) map[string]interface{} {

	if i == nil {
		return nil
	}

	listVal := reflect.ValueOf(i)

	values := make(map[string]interface{}, listVal.Len())

	original := reflect.ValueOf(i)

	for _, key := range original.MapKeys() {
		data := parseUnknown(original.MapIndex(key).Interface())
		for dataKey, value := range data {
			values[key.String()+dataKey] = value
		}
	}

	return values
}

func parseStruct(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}

	v := reflect.ValueOf(i)

	values := make(map[string]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		fieldData := v.Field(i).Interface()
		filedName := v.Type().Field(i).Name
		data := parseUnknown(fieldData, filedName)
		for key, value := range data {
			values[key] = value
		}
	}
	return values
}

func getPrefix(prefixArr ...string) (string, string) {
	originalPrefix := ""
	prefix := ""

	if len(prefixArr) > 0 {
		originalPrefix = prefixArr[0]
	}

	if originalPrefix != "" {
		prefix = originalPrefix + "."
	}

	return prefix, originalPrefix
}
