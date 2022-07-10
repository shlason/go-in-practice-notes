package main

import (
	"testing"
)

func TestName(t *testing.T) {
	name := getName()

	if name != "Jason" {
		t.Error("Respone form getName is unexpected value")
	}
}
