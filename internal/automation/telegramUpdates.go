package automation

import (
	"context"
	"encoding/json"
	"tcms/m/internal/db"
	"tcms/m/internal/db/repository"
	"tcms/m/internal/dry"
)

type TelegramUpdateTrigger struct {
	Name string            `json:"name"`
	Data map[string]string `json:"data"`
}

func (t TelegramUpdateTrigger) GetName() string {
	return t.Name
}

func (t TelegramUpdateTrigger) GetData() map[string]string {
	return t.Data
}

func UpdateTriggerFactory(addConsumer chan chan []uint8) {
	ctx := context.Background()
	connection, err := db.GetConnection(ctx)
	dry.HandleErrorPanic(err)

	automationRepo := repository.CreateAutomationRepository(connection)

	automationService := Service{}
	automationService.Start(automationRepo)

	ch := make(chan []uint8)
	addConsumer <- ch
	for {
		data := <-ch
		var trigger TelegramUpdateTrigger
		err := json.Unmarshal(data, &trigger)
		dry.HandleError(err)
		automationService.HandleTrigger(trigger)
	}
}
