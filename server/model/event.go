package model

import "gorm.io/gorm"

type Event struct {
	gorm.Model

	Name      string `json:"name"`
	Content   string `json:"content"`
	EndTimeAt int64  `json:"endTime_at"`
	State     int    `json:"state,omitempty"`
	Type      int    `json:"type,omitempty"`
}
