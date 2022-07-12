package model

type Notification struct {
	UserId   string `gorm:"primaryKey"`
	Rejected int32
	Accepted int32
}
