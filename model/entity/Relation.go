package entity

import "gorm.io/gorm"

// Relation UserID关注了FollowID
type Relation struct {
	ID       uint `gorm:"primaryKey; not null; auto_increment" json:"id"`
	UserID   uint `gorm:"not null" json:"user_id"`
	User     User `gorm:"foreignKey:UserID"`
	FollowID uint `gorm:"not null" json:"follow_id"`
	Follow   User `gorm:"foreignKey:FollowID"`
	DeleteAt gorm.DeletedAt
}
