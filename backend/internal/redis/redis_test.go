package redis

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redismock/v8"
	"os"
	"tcms/m/internal/dry"
	"testing"
)

func TestClient_Publish(t *testing.T) {
	const message = "test_data"
	const channel = "test_channel"

	marshalData, _ := json.Marshal(message)

	db, mock := redismock.NewClientMock()

	mock.ExpectPublish(channel, marshalData)

	c := client{client: db}
	_, err := c.Publish(context.Background(), channel, message)
	dry.TestHandleError(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestClient_Publish_badData(t *testing.T) {
	const channel = "test_channel"

	message := make(chan int)

	db, _ := redismock.NewClientMock()

	c := client{client: db}
	_, err := c.Publish(context.Background(), channel, message)
	dry.TestCheckEqual(t, "json: unsupported type: chan int", err.Error())
}

func TestClient_Subscribe(t *testing.T) {
	const channel = "test_channel"

	db, _ := redismock.NewClientMock()

	c := client{client: db}

	c.Subscribe(context.Background(), channel)
}

func TestGetClient(t *testing.T) {
	err := os.Setenv("REDIS_HOST", "test")
	dry.TestHandleError(t, err)
	err = os.Setenv("REDIS_PASSWORD", "password")
	dry.TestHandleError(t, err)
	err = os.Setenv("REDIS_DATABASE", "0")
	dry.TestHandleError(t, err)

	GetClient()
}
