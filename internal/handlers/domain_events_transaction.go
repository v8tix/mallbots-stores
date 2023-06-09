package handlers

import (
	"context"

	"github.com/v8tix/eda/ddd"
	"github.com/v8tix/eda/di"
)

func RegisterDomainEventHandlersTx(container di.Container) {
	handlers := ddd.EventHandlerFunc[ddd.AggregateEvent](func(ctx context.Context, event ddd.AggregateEvent) error {
		domainHandlers := di.Get(ctx, "domainEventHandlers").(ddd.EventHandler[ddd.AggregateEvent])

		return domainHandlers.HandleEvent(ctx, event)
	})

	subscriber := container.Get("domainDispatcher").(*ddd.EventDispatcher[ddd.AggregateEvent])

	RegisterDomainEventHandlers(subscriber, handlers)
}
