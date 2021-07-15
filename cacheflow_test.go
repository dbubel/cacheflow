package cacheflow

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCacher(t *testing.T) {
	t.Run("test return type is correct", func(t *testing.T) {
		cache := NewCacherDefault()
		assert.IsType(t, Cacher{}, *cache)
	})

	t.Run("test insert", func(t *testing.T) {
		cache := NewCacherDefault()
		cache.Insert("insertkey", []byte("hello"))
	})

	t.Run("test get", func(t *testing.T) {
		cache := NewCacherDefault()
		cache.Insert("insertkey", []byte("hello"))
		data, err := cache.Get("insertkey")
		assert.NoError(t, err)
		assert.Equal(t, "hello", string(data))
	})

	t.Run("test get marshal", func(t *testing.T) {
		cache := NewCacherDefault()
		type teststruct struct {
			Name       string
			AddressNum int
		}
		var target teststruct
		cache.Insert("insertkey", []byte(`{"name":"name here", "addressNum":123}`))
		err := cache.GetObject("insertkey", &target)
		assert.NoError(t, err)
		assert.Equal(t, target.Name, "name here")
		assert.Equal(t, target.AddressNum, 123)
	})

	t.Run("test insert struct", func(t *testing.T) {
		cache := NewCacherDefault()
		type teststruct struct {
			Name       string
			AddressNum int
		}
		cache.InsertObject("insertkey", teststruct{
			Name:       "name here",
			AddressNum: 123,
		})
		data, err := cache.Get("insertkey")
		assert.NoError(t, err)
		assert.JSONEq(t, string(data), "{\"Name\":\"name here\",\"AddressNum\":123}")

	})
}

func BenchmarkJSON(b *testing.B) {

	type teststruct struct {
		Name       string
		AddressNum int
	}
	ts := teststruct{
		Name:       RandStringRunes(256),
		AddressNum: 0,
	}
	b.Run("bench insert and get", func(b *testing.B) {
		cache := NewCacherDefault()
		var gobTs teststruct
		for n := 0; n < b.N; n++ {
			key := RandStringRunes(32)
			cache.InsertObject(key, &ts)
			cache.GetObject(key, &gobTs)
		}
	})

	b.Run("bench insert", func(b *testing.B) {
		cache := NewCacherDefault()
		for n := 0; n < b.N; n++ {
			cache.InsertObject(RandStringRunes(32), &ts)
		}
	})

	b.Run("bench retrieve", func(b *testing.B) {
		cache := NewCacherDefault()
		for n := 0; n < b.N; n++ {
			cache.InsertObject(RandStringRunes(32), &ts)
		}
	})
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
