package main

import (
	f "github.com/uplatform-ai/foundation"
)

var (
	svc = f.InitOutboxCourier("chats-outbox")
)

func main() {
	svc.Start(&f.OutboxCourierOptions{
		BatchSize: 100,
	})
}
