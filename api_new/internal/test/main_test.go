package test

import "testing"

func TestMain(m *testing.M) {
	shutdown, err := InitializeTestStorages()
	if err != nil {
		panic(err)
	}
	defer shutdown()
	m.Run()
}
