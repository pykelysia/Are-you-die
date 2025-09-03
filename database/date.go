package database

import (
	"time"
)

type Date struct {
	ID          uint      `gorm:"primaryKey"` // 用户ID
	LastCheckin time.Time `gorm:"not null"`   // 最后一次签到时间
}

func NewDate() *Date {
	return &Date{}
}

func (*Date) TableName() string {
	return "dates"
}

func (*Date) Create(date *Date) error {
	return mysqlDB.Create(date).Error
}

func (*Date) Delete(id uint) error {
	return mysqlDB.Delete(&Date{}, id).Error
}

func (*Date) Update(date *Date) error {
	return mysqlDB.Where("id = ?", date.ID).Updates(date).Error
}

func (*Date) Get(id uint) (*Date, error) {
	var date Date
	err := mysqlDB.First(&date, id).Error
	return &date, err
}

// 更新用户签到时间
func (*Date) UpdateCheckin(id uint) error {
	return mysqlDB.Model(&Date{}).Where("id = ?", id).Update("last_checkin", time.Now()).Error
}

// 获取所有超过三天未签到的用户ID
func (*Date) GetInactiveUserIDs() ([]uint, error) {
	var dates []Date
	threeDaysAgo := time.Now().AddDate(0, 0, -3)
	err := mysqlDB.Where("last_checkin < ?", threeDaysAgo).Find(&dates).Error
	if err != nil {
		return nil, err
	}
	ids := make([]uint, 0, len(dates))
	for _, d := range dates {
		ids = append(ids, d.ID)
	}
	return ids, nil
}
