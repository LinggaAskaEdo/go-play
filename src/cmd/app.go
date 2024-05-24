package main

import (
	"math/rand"
	"time"
)

const (
	DefaultMaxJitter = 2000
	DefaultMinJitter = 100
)

func SleepWithJitter() {
	min := DefaultMinJitter
	max := DefaultMaxJitter

	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	rnd := rng.Intn(max-min) + min
	time.Sleep(time.Duration(rnd) * time.Millisecond)
}
