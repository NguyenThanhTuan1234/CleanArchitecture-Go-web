package entity

import "time"

type Session struct {
	Id           int
	LastActivity time.Time
}
