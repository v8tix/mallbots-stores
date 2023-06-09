package handlers

import (
	"context"

	"github.com/v8tix/eda/am"
	"github.com/v8tix/eda/ddd"
	"github.com/v8tix/mallbots-stores-proto/pb"
	"github.com/v8tix/mallbots-stores/internal/domain"
)

type domainHandlers[T ddd.AggregateEvent] struct {
	publisher am.MessagePublisher[ddd.Event]
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*domainHandlers[ddd.AggregateEvent])(nil)

func NewDomainEventHandlers(publisher am.MessagePublisher[ddd.Event]) ddd.EventHandler[ddd.AggregateEvent] {
	return &domainHandlers[ddd.AggregateEvent]{
		publisher: publisher,
	}
}

func RegisterDomainEventHandlers(subscriber ddd.EventSubscriber[ddd.AggregateEvent], handlers ddd.EventHandler[ddd.AggregateEvent]) {
	subscriber.Subscribe(handlers,
		domain.StoreCreatedEvent,
		domain.StoreParticipationEnabledEvent,
		domain.StoreParticipationDisabledEvent,
		domain.StoreRebrandedEvent,
		domain.ProductAddedEvent,
		domain.ProductRebrandedEvent,
		domain.ProductPriceIncreasedEvent,
		domain.ProductPriceDecreasedEvent,
		domain.ProductRemovedEvent,
	)
}
func (h domainHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case domain.StoreCreatedEvent:
		return h.onStoreCreated(ctx, event)
	case domain.StoreParticipationEnabledEvent:
		return h.onStoreParticipationEnabled(ctx, event)
	case domain.StoreParticipationDisabledEvent:
		return h.onStoreParticipationDisabled(ctx, event)
	case domain.StoreRebrandedEvent:
		return h.onStoreRebranded(ctx, event)

	case domain.ProductAddedEvent:
		return h.onProductAdded(ctx, event)
	case domain.ProductRebrandedEvent:
		return h.onProductRebranded(ctx, event)
	case domain.ProductPriceIncreasedEvent:
		return h.onProductPriceIncreased(ctx, event)
	case domain.ProductPriceDecreasedEvent:
		return h.onProductPriceDecreased(ctx, event)
	case domain.ProductRemovedEvent:
		return h.onProductRemoved(ctx, event)
	}
	return nil
}

func (h domainHandlers[T]) onStoreCreated(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domain.StoreCreated)
	return h.publisher.Publish(ctx, pb.StoreAggregateChannel,
		ddd.NewEvent(pb.StoreCreatedEvent, &pb.StoreCreated{
			Id:       event.AggregateID(),
			Name:     payload.Name,
			Location: payload.Location,
		}),
	)
}

func (h domainHandlers[T]) onStoreParticipationEnabled(ctx context.Context, event ddd.AggregateEvent) error {
	return h.publisher.Publish(ctx, pb.StoreAggregateChannel,
		ddd.NewEvent(pb.StoreParticipatingToggledEvent, &pb.StoreParticipationToggled{
			Id:            event.AggregateID(),
			Participating: true,
		}),
	)
}

func (h domainHandlers[T]) onStoreParticipationDisabled(ctx context.Context, event ddd.AggregateEvent) error {
	return h.publisher.Publish(ctx, pb.StoreAggregateChannel,
		ddd.NewEvent(pb.StoreParticipatingToggledEvent, &pb.StoreParticipationToggled{
			Id:            event.AggregateID(),
			Participating: false,
		}),
	)
}

func (h domainHandlers[T]) onStoreRebranded(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domain.StoreRebranded)
	return h.publisher.Publish(ctx, pb.StoreAggregateChannel,
		ddd.NewEvent(pb.StoreRebrandedEvent, &pb.StoreRebranded{
			Id:   event.AggregateID(),
			Name: payload.Name,
		}),
	)
}

func (h domainHandlers[T]) onProductAdded(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domain.ProductAdded)
	return h.publisher.Publish(ctx, pb.ProductAggregateChannel,
		ddd.NewEvent(pb.ProductAddedEvent, &pb.ProductAdded{
			Id:          event.AggregateID(),
			StoreId:     payload.StoreID,
			Name:        payload.Name,
			Description: payload.Description,
			Sku:         payload.SKU,
			Price:       payload.Price,
		}),
	)
}

func (h domainHandlers[T]) onProductRebranded(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domain.ProductRebranded)
	return h.publisher.Publish(ctx, pb.ProductAggregateChannel,
		ddd.NewEvent(pb.ProductRebrandedEvent, &pb.ProductRebranded{
			Id:          event.AggregateID(),
			Name:        payload.Name,
			Description: payload.Description,
		}),
	)
}

func (h domainHandlers[T]) onProductPriceIncreased(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domain.ProductPriceChanged)
	return h.publisher.Publish(ctx, pb.ProductAggregateChannel,
		ddd.NewEvent(pb.ProductPriceIncreasedEvent, &pb.ProductPriceChanged{
			Id:    event.AggregateID(),
			Delta: payload.Delta,
		}),
	)
}

func (h domainHandlers[T]) onProductPriceDecreased(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domain.ProductPriceChanged)
	return h.publisher.Publish(ctx, pb.ProductAggregateChannel,
		ddd.NewEvent(pb.ProductPriceDecreasedEvent, &pb.ProductPriceChanged{
			Id:    event.AggregateID(),
			Delta: payload.Delta,
		}),
	)
}

func (h domainHandlers[T]) onProductRemoved(ctx context.Context, event ddd.AggregateEvent) error {
	return h.publisher.Publish(ctx, pb.ProductAggregateChannel,
		ddd.NewEvent(pb.ProductRemovedEvent, &pb.ProductRemoved{
			Id: event.AggregateID(),
		}),
	)
}
