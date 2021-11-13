package telegramClient

import (
	"fmt"
	"github.com/xelaj/mtproto/telegram"
	"reflect"
	"strconv"
	"strings"
)

type updateTrigger struct {
	Name string
	Data map[string]string
}

func recognizeTrigger(i interface{}) []updateTrigger {
	var triggerType string
	var triggerList []updateTrigger
	switch message := i.(type) {
	case *telegram.UpdateShort:
		triggerType = getTriggerType(message.Update)
		triggerData := parsePtr(message)
		triggerList = appendTrigger(triggerType, triggerData, triggerList)
	case *telegram.UpdatesObj:
		users := make(map[string]string, len(message.Users))
		for userIndex, user := range message.Users {
			userMap := parsePtr(user)
			for key, value := range userMap {
				users[strconv.Itoa(userIndex)+"."+key] = value
			}
		}

		for _, event := range message.Updates {
			triggerType = getTriggerType(event)
			triggerData := parsePtr(event)
			for userIndex, value := range users {
				triggerData["users."+userIndex] = value
			}
			triggerList = appendTrigger(triggerType, triggerData, triggerList)
		}
	default:
		val := reflect.ValueOf(i).Elem().FieldByName("Obj")

		if val.IsNil() || val.IsZero() {
			triggerData := parseUnknown(i)
			triggerList = appendTrigger(triggerType, triggerData, triggerList)
		} else {
			triggerList = recognizeTrigger(val.Interface())
		}
	}

	return triggerList
}

func appendTrigger(triggerType string, triggerData map[string]string, triggerList []updateTrigger) []updateTrigger {
	trigger := updateTrigger{
		Name: triggerType,
		Data: triggerData,
	}
	triggerList = append(triggerList, trigger)
	return triggerList
}

func getTriggerType(i interface{}) string {
	return strings.Replace(reflect.TypeOf(i).String(), "*telegram.", "", 1)
}

func parsePtr(i interface{}) map[string]string {
	if i == nil {
		return nil
	}
	if reflect.ValueOf(i).IsNil() || reflect.ValueOf(i).IsZero() {
		return nil
	}

	v := reflect.Indirect(reflect.ValueOf(i))

	values := make(map[string]string, v.NumField())

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

func parseUnknown(i interface{}, prefixArr ...string) map[string]string {
	if i == nil {
		return nil
	}

	prefix, originalPrefix := getPrefix(prefixArr...)

	values := make(map[string]string)

	valueType := reflect.ValueOf(i).Kind()

	switch valueType {
	case reflect.Ptr:
		data := parsePtr(i)
		for key, value := range data {
			values[prefix+key] = fmt.Sprintf("%v", value)
		}
	case reflect.Slice:
		data := parseSlice(i)
		if len(data) > 0 {
			for key, value := range data {
				values[prefix+key] = fmt.Sprintf("%v", value)
			}
		} else {
			values[prefix] = ""
		}

	case reflect.Map:
		data := parseMap(i)
		for key, value := range data {
			values[prefix+key] = fmt.Sprintf("%v", value)
		}
	case reflect.Array:
		panic("array in parse unknown!")
	case reflect.Struct:
		data := parseStruct(i)
		for key, value := range data {
			values[prefix+key] = fmt.Sprintf("%v", value)
		}

	default:
		values[originalPrefix] = fmt.Sprintf("%v", i)
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
