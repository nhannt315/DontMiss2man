package controllers_test

import (
	"testing"

	"github.com/nhannt315/real_estate_api/internal/test"
)

func TestMain(m *testing.M) {
	shutdown, err := test.InitializeTestStorages()
	if err != nil {
		panic(err)
	}
	defer shutdown()
	m.Run()
}
