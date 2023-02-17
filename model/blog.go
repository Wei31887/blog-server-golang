package model

import "time"


type Blog struct {
	Id         int    `gorm:"id" json:"id"`
	Title      string `gorm:"title" json:"title"`
	TypeId     int    `gorm:"column:type_id" json:"type_id"`
	Content    string `gorm:"content" json:"content"`
	Summary    string `gorm:"summary" json:"summary"`
	ClickHit   int    `gorm:"click_hit" json:"click_hit"`
	ReplayHit  int    `gorm:"replay_hit" json:"replay_hit"` 
	AddTime    time.Time `gorm:"add_time" json:"add_time"`
	UpdateTime time.Time `gorm:"update_time" json:"update_time"`
	TypeName   string `gorm:"->" json:"type_name"`
}