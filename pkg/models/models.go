package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("modles: no matching record found!")

type Snippet struct {
	ID 		int
	Title	string
	Content string
	Created	time.Time
	Expires	time.Time
}