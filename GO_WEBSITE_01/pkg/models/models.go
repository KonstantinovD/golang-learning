package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: подходящей записи не найдено")

type Snippet struct {
	SnippetId int
	Title     string
	Content   string
	Created   time.Time
	Expires   time.Time
}
