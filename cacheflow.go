package cacheflow

import (
	"encoding/json"
	"fmt"
	"time"
)

type Cacher struct {
	Cache         map[string]element
	DefaultExpiry time.Duration
}

func NewCacherDefault() *Cacher {
	return &Cacher{
		Cache:         make(map[string]element),
		DefaultExpiry: time.Hour * 24,
	}
}

func NewCacher(defaultExpiry time.Duration) *Cacher {
	return &Cacher{
		Cache:         make(map[string]element),
		DefaultExpiry: defaultExpiry,
	}
}

func (c *Cacher) NumElements() int {
	return len(c.Cache)
}

func (c *Cacher) Get(key string) ([]byte, error) {
	if data, ok := c.Cache[key]; ok {
		if !data.Expired() {
			return data.Data, nil
		}
	}
	return []byte{}, fmt.Errorf("object not found")
}

func (c Cacher) GetObject(key string, v interface{}) error {
	data, err := c.Get(key)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func (c *Cacher) Insert(key string, data []byte) {
	c.Cache[key] = element{
		Data:   data,
		Expiry: time.Now().Add(c.DefaultExpiry),
	}
}

func (c *Cacher) InsertObject(key string, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	c.Cache[key] = element{
		Data:   data,
		Expiry: time.Now().Add(c.DefaultExpiry),
	}
	return nil
}

func (c *Cacher) InsertObjectWithExpiry(key string, expiry time.Duration, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	c.Cache[key] = element{
		Data:   data,
		Expiry: time.Now().Add(expiry),
	}
	return nil
}

func (c *Cacher) InsertWithExpiry(key string, expiry time.Duration, data []byte) {
	c.Cache[key] = element{
		Data:   data,
		Expiry: time.Now().Add(expiry),
	}
}

func (c *Cacher) Delete(key string) error {
	ele, ok := c.Cache[key]
	if !ok {
		return fmt.Errorf("element not found for deletion %+v", c.Cache)
	}
	delete(c.Cache, key)
}
