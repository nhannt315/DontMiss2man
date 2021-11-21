package test

import (
	"context"
	"testing"

	"github.com/nhannt315/real_estate_api/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestHelper_Test_Main(t *testing.T) {

	h := NewHelper(t)

	// DB
	reg := h.Registry()

	ctx := context.Background()
	newUser := &model.User{
		ID:    1,
		Email: "test@gmail.com",
	}
	err := reg.UserRepository().WithContext(ctx).Create(newUser)
	if err != nil {
		t.Fatal(err)
	}

	got, err := reg.UserRepository().WithContext(ctx).FindByEmail(newUser.Email)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, newUser.Email, got.Email)
}
