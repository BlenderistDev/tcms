package redis

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redismock/v8"
	"tcms/m/internal/dry"
	"testing"
)

func TestClient_Publish(t *testing.T) {
	const message = "test_data"
	const channel = "test_channel"

	marshalData, err := json.Marshal(message)

	db, mock := redismock.NewClientMock()

	mock.ExpectPublish(channel, marshalData)

	c := client{client: db}
	_, err = c.Publish(context.Background(), channel, message)
	dry.TestHandleError(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}
