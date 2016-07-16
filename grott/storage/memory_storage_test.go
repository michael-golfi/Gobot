package storage_test

import (
	"testing"
	"github.com/michael-golfi/Grott/grott/types"
	"github.com/stretchr/testify/assert"
	"github.com/michael-golfi/Grott/grott/storage"
)

func TestInMemoryStorage(t *testing.T) {
	ctx := &types.DialogContext{}

	storage := storage.NewInMemoryStorage()
	err := storage.Save("Example", ctx)
	assert.NoError(t, err)

	ctx, err = storage.Get("Example")
	assert.NoError(t, err)

	storage.Delete("Example")

	ctx, err = storage.Get("Example")
	assert.Error(t, err)
	assert.Nil(t, ctx)
}