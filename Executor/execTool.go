package Executor

import (
	"MQTT_Middleware/databaseConnection"
	"MQTT_Middleware/model"
	"MQTT_Middleware/util"
	"encoding/json"
	"errors"
	"strconv"
)

func SaveStatusToDatabase(jsonStr, poolId string) error {
	iPoolId, err := strconv.ParseInt(poolId, 10, 64)
	if err != nil {
		return errors.New("wrong pool id by save status")
	}
	var obj model.StatusList
	json.Unmarshal([]byte(jsonStr), &obj)
	for _, item := range obj.Status {
		id := util.IdCreator.GetId()
		item.Id = id
		item.PoolId = iPoolId
		err = databaseConnection.MysqlClient.Table("tb_device_status").Create(&item).Error
		if err != nil {
			return errors.New("create row fail,but i dont know why :)")
		}
	}
	return nil
}

func SaveHistoryToDatabase(jsonStr, poolId string) error {
	var obj model.ControlHistory
	json.Unmarshal([]byte(jsonStr), &obj)
	id := util.IdCreator.GetId()
	obj.Id = id
	iPoolId, err := strconv.ParseInt(poolId, 10, 64)
	if err != nil {
		return errors.New("wrong pool id by save history")
	}
	obj.PoolId = iPoolId
	err = databaseConnection.MysqlClient.Table("tb_device_behaviour").Create(&obj).Error
	return nil
}
