package telegramClient

import (
	"fmt"
	"github.com/xelaj/mtproto/telegram"
	"reflect"
	"strings"
	"tcms/m/automation"
)

func recognizeTrigger(i interface{}) []automation.TelegramUpdateTrigger {
	var triggerType string
	var triggerList []automation.TelegramUpdateTrigger
	switch message := i.(type) {
	case *telegram.UpdateShort:
		triggerType = getTriggerType(message.Update)
		triggerData := getTriggerData(message.Update)
		trigger := automation.TelegramUpdateTrigger{
			Name: triggerType,
			Data: triggerData,
		}
		triggerList = append(triggerList, trigger)
	case *telegram.UpdatesObj:
		for _, event := range message.Updates {
			triggerType = getTriggerType(event)
			triggerData := getTriggerData(event)
			trigger := automation.TelegramUpdateTrigger{
				Name: triggerType,
				Data: triggerData,
			}
			triggerList = append(triggerList, trigger)
		}

	default:
		triggerType = "unknown"
	}

	return triggerList
}

func getTriggerType(i interface{}) string {
	return strings.Replace(reflect.TypeOf(i).String(), "*telegram.", "", 1)
}

func getTriggerData(i interface{}, prefixArr ...string) map[string]interface{} {

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

	v := reflect.Indirect(reflect.ValueOf(i))

	values := make(map[string]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {

		valueType := v.Field(i).Kind().String()

		switch valueType {
		case "interface":
			data := getTriggerData(v.Field(i).Interface(), v.Type().Field(i).Name)
			for key, value := range data {
				values[prefix+key] = value
			}
		case "slice":
			if v.Field(i).IsZero() || v.Field(i).IsNil() {
				values[prefix+v.Type().Field(i).Name] = nil
			} else {
				fmt.Println("slice")
			}
		case "ptr":
			fmt.Println("ptr")
		default:
			values[prefix+v.Type().Field(i).Name] = v.Field(i).Interface()
		}
	}
	return values
}
