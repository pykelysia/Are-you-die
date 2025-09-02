package database

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
}

func NewUser() *User {
	return &User{}
}

func (*User) Create(user *User) error {
	err := mysqlDB.Create(user).Error
	return err
}

func (*User) Delete(id uint) error {
	err := mysqlDB.Delete(&User{}, id).Error
	return err
}

func (*User) Update(user *User) error {
	err := mysqlDB.Where("id = ?", user.ID).Updates(user).Error
	return err
}

func (*User) Get(id uint) (*User, error) {
	var user User
	err := mysqlDB.First(&user, id).Error
	return &user, err
}

func (*User) TableName() string {
	return "users"
}
