package automation

import (
	"context"
	"encoding/json"
	redis2 "github.com/go-redis/redis/v8"
	"tcms/m/internal/dry"
	"tcms/m/internal/redis"
)

type TelegramUpdateTrigger struct {
	Name string
	Data map[string]string
}

func (t TelegramUpdateTrigger) GetName() string {
	return t.Name
}

func (t TelegramUpdateTrigger) GetData() map[string]string {
	return t.Data
}

func UpdateTriggerFactory() {
	var ctx = context.Background()

	redisClient := redis.GetClient()
	subscribe := redisClient.Subscribe(ctx, "update")

	defer func(subscribe *redis2.PubSub) {
		err := subscribe.Close()
		dry.HandleError(err)
	}(subscribe)

	automationService := Service{}
	automationService.Start()

	for {
		msg, err := subscribe.ReceiveMessage(ctx)
		dry.HandleError(err)
		bytes := []byte(msg.Payload)

		var trigger TelegramUpdateTrigger
		err = json.Unmarshal(bytes, &trigger)
		dry.HandleError(err)
		automationService.HandleTrigger(trigger)
	}
}
