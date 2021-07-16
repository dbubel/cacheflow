package cacheflow

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestElement(t *testing.T) {
	t.Run("test expired true", func(t *testing.T) {
		e := element{
			Data:   []byte("hello"),
			Expiry: time.Now().Add(-1 * time.Second),
		}
		assert.True(t, e.Expired())
	})

	t.Run("test expired false", func(t *testing.T) {
		e := element{
			Data:   []byte("hello"),
			Expiry: time.Now().Add(time.Second),
		}
		assert.False(t, e.Expired())
	})
}
