package main

import (
	"testing"
)

func TestGetClientConfiguration(t *testing.T) {

	if got := getClientInstance(&trelloClientConfig{
		appKey: "",
		token:  "",
		logger: nil,
	}); got.Logger != nil {
		t.Errorf("Logger should be nil!")
	}
}
