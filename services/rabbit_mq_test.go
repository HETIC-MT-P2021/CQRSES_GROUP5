package services

import (
	"github.com/golang/mock/gomock"
	"testing"
)

func TestPublishEventToRBMQ(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	m := NewMockRabbitMQ(ctrl)

	// Asserts that the first and only call to Bar() is passed 99.
	// Anything else will fail.
	m.
		EXPECT().
		PublishEventToRBMQ(gomock.Eq(99)).
		Return(nil)
}
