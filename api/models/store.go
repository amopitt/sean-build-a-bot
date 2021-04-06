package models

import "database/sql"

type Store struct {
	Db *sql.DB
}

type Cart struct {
	OrderId   int64   `json:"order_id"`
	RobotName string  `json:"robot_name"`
	Cost      float32 `json:"cost"`
}
