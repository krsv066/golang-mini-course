package dto

import "sync"

type Account struct {
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

type AccountStorage struct {
	Mtx     sync.RWMutex
	Storage map[string]*Account
}
