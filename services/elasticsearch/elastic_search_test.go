package elasticsearch

import (
	"context"
	"github.com/golang/mock/gomock"
"testing"
)

func TestEsConnector_NewIndex(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockEsService(ctrl)

	ctx := context.Background()

	m.
		EXPECT().
		NewIndex(ctx, "test").
		Return(nil)
}

func TestEsConnector_GetDocumentByIndexAndID(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockEsService(ctrl)

	ctx := context.Background()

	m.
		EXPECT().
		GetDocumentByIndexAndID(ctx, "test", "hey").
		Return(nil)
}
