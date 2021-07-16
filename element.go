package cacheflow

import "time"

type element struct {
	Data   []byte
	Expiry time.Time
}

func (c element) Expired() bool {
	if c.Expiry.After(time.Now()) {
		return false
	}
	return true
}
