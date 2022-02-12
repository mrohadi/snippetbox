package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecord           = errors.New("modles: no matching record found")
	ErrInvalidCredentials = errors.New("modles: invalid credentials")
	ErrDuplicateEmail     = errors.New("model: ducplicate email")
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}
