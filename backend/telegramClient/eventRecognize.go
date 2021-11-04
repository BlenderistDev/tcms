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

func getTriggerData(i interface{}) map[string]interface{} {

	v := reflect.ValueOf(i)

	values := make([]interface{}, reflect.Indirect(v).NumField())

	for i := 0; i < reflect.Indirect(v).NumField(); i++ {
		values[i] = reflect.Indirect(v).Field(i).Interface()
	}

	fmt.Printf("%v", values)

	var inInterface map[string]interface{}

	return inInterface
}
