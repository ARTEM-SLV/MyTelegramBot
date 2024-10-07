package telegram

import (
	"MyTelegramBot/pkg/clients/telgram"
	"MyTelegramBot/storage"
)

type Processor struct {
	tg      *telgram.Client
	offset  int
	storage storage.Storage
}

func New(client *telgram.Client) {}
