package model

type ControlHistory struct {
	Id      int64  `gorm:"column:id;primary_key"`
	Command string `json:"command" gorm:"column:command"`
	Serial  string `json:"serial" gorm:"column:serial"`
	PoolId  int64  `gorm:"column:pool_id"`
}
