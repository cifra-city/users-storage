package events

import (
	"context"

	"github.com/cifra-city/cifractx"
	"github.com/cifra-city/users-storage/internal/config"
	"github.com/cifra-city/users-storage/internal/service/events/callbacks"
	"github.com/sirupsen/logrus"
)

const (
	accountCreate = "account.create"
)

func Listener(ctx context.Context) {
	server, err := cifractx.GetValue[*config.Service](ctx, config.SERVER)
	if err != nil {
		logrus.Fatalf("failed to get server from context: %v", err)
	}

	err = server.Broker.Listen(ctx, server.Logger, "account.create", "sso.events", callbacks.CreateAccount)
	if err != nil {
		logrus.Fatalf("Listener encountered an error: %v", err)
	}
}