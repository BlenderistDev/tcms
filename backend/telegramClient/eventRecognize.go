package telegramClient

import (
	"github.com/xelaj/mtproto/telegram"
	"tcms/m/automation"
)

func recognizeTrigger(i interface{}) []automation.TelegramUpdateTrigger {
	var triggerType string
	var triggerList []automation.TelegramUpdateTrigger
	switch message := i.(type) {
	case *telegram.UpdateShort:
		triggerType = getTriggerType(message.Update)
		trigger := automation.TelegramUpdateTrigger{
			Name:    triggerType,
			KeyList: nil,
			Data:    nil,
		}
		triggerList = append(triggerList, trigger)
	case *telegram.UpdatesObj:
		for _, event := range message.Updates {
			triggerType = getTriggerType(event)
			trigger := automation.TelegramUpdateTrigger{
				Name:    triggerType,
				KeyList: nil,
				Data:    event,
			}
			triggerList = append(triggerList, trigger)
		}

	default:
		triggerType = "unknown"
	}

	return triggerList
}

func getTriggerType(i interface{}) string {
	var triggerType string
	switch i.(type) {
	case *telegram.UpdateUserStatus:
		triggerType = "UpdateUserStatus"
	case *telegram.UpdateNewMessage:
		triggerType = "UpdateNewMessage"
	case *telegram.UpdateNewChannelMessage:
		triggerType = "UpdateNewChannelMessage"
	case *telegram.UpdateReadChannelInbox:
		triggerType = "UpdateReadChannelInbox"
	case *telegram.UpdateEditChannelMessage:
		triggerType = "UpdateEditChannelMessage"
	default:
		triggerType = "unknown"
	}

	return triggerType
}
