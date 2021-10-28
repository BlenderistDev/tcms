package telegramClient

import "github.com/xelaj/mtproto/telegram"

func getTriggerType(i interface{}) string {
	var triggerType string
	switch message := i.(type) {
	case *telegram.UpdateShort:
		switch message.Update.(type) {
		case *telegram.UpdateUserStatus:
			triggerType = "UpdateUserStatus"
		default:
			triggerType = "unknown"
		}
	}
	return triggerType
}
