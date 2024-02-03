package noa_test

import (
	"testing"

	"github.com/malikilamalik/noa"
)

func TestNewLog(t *testing.T) {
	noa.New()
	t.Error()
}
