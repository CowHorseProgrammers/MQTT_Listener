package model

type StatusList struct {
	Status []Status `json:"status"`
}

type Status struct {
	Id int64 `gorm:"column:id;primary_key"`
	// 是否故障
	IsFaulty bool `json:"isFaulty" gorm:"column:is_faulty"`
	// 在线状态
	IsOnline bool `json:"isOnline" gorm:"column:is_online"`
	// 工作状态
	IsWorking bool `json:"isWorking" gorm:"column:is_working"`
	// 设备序列号
	Serial string `json:"serial" gorm:"column:serial"`
	PoolId int64  `gorm:"column:pool_id"`
}
