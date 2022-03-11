package model

type Tests struct {
	ID   int64 `gorm:"primaryKey"`
	Name string
}

func (Tests) TableName() string {
	return "tests"
}
