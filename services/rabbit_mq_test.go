package services

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/core/eventsourcing"
	"github.com/golang/mock/gomock"
	"testing"
	"time"
)

func TestPublishEventToRBMQ(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockRabbitMQ(ctrl)

	testEvent := eventsourcing.Event{
		Type:           eventsourcing.AddOrder,
		Payload:        "4",
		CreatedAt:      time.Time{},
		AggregateIndex: 1, // Order aggregation Index
	}

	m.
		EXPECT().
		PublishEventToRBMQ(testEvent).
		Return(nil)
}
