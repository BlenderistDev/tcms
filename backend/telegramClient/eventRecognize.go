package telegramClient

import (
	"fmt"
	"github.com/k0kubun/pp"
	"github.com/xelaj/mtproto/telegram"
	"reflect"
	"strconv"
	"strings"
	"tcms/m/automation"
)

func recognizeTrigger(i interface{}) []automation.TelegramUpdateTrigger {
	var triggerType string
	var triggerList []automation.TelegramUpdateTrigger
	switch message := i.(type) {
	case *telegram.UpdateShort:
		triggerType = getTriggerType(message.Update)
		triggerData := parsePtr(message.Update)
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
		triggerType = "unknown"
	}

	pp.Println(triggerList)
	return triggerList
}

func getTriggerType(i interface{}) string {
	return strings.Replace(reflect.TypeOf(i).String(), "*telegram.", "", 1)
}

func parsePtr(i interface{}, prefixArr ...string) map[string]interface{} {
	if i == nil {
		return nil
	}
	if reflect.ValueOf(i).IsNil() || reflect.ValueOf(i).IsZero() {
		return nil
	}

	prefix := getPrefix(prefixArr...)

	v := reflect.Indirect(reflect.ValueOf(i))

	values := make(map[string]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		fieldData := v.Field(i).Interface()
		filedName := v.Type().Field(i).Name
		data := parseUnknown(fieldData, filedName)
		for key, value := range data {
			values[prefix+key] = value
		}
	}
	return values
}

func parseUnknown(i interface{}, prefixArr ...string) map[string]interface{} {
	if i == nil {
		return nil
	}

	prefix := getPrefix(prefixArr...)

	values := make(map[string]interface{})

	valueType := reflect.ValueOf(i).Kind()

	switch valueType {
	case reflect.Ptr:
		data := parsePtr(i)
		for key, value := range data {
			values[prefix+"."+key] = value
		}
	case reflect.Slice:
		fmt.Println("slice")
		data := parseSlice(i)
		fmt.Println(data)
		if len(data) > 0 {
			for key, value := range data {
				values[prefix+"."+key] = value
			}
		} else {
			values[prefix] = nil
		}

	case reflect.Map:
	case reflect.Array:
	case reflect.Struct:

	default:
		values[prefix] = i
	}

	return values
}

func parseSlice(i interface{}, prefixArr ...string) map[string]interface{} {

	if i == nil {
		return nil
	}

	prefix := ""

	if len(prefixArr) > 0 {
		prefix = prefixArr[0]
	}

	if prefix != "" {
		prefix += "."
	}

	listVal := reflect.ValueOf(i)

	values := make(map[string]interface{}, listVal.Len())

	for key := 0; key < listVal.Len(); key++ {
		values[prefix+strconv.Itoa(key)] = parseUnknown(listVal.Interface())
	}

	return values
}

func getPrefix(prefixArr ...string) string {
	prefix := ""

	if len(prefixArr) > 0 {
		prefix = prefixArr[0]
	}

	return prefix
}
