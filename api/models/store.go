package models

import (
	"database/sql"
)

type Store struct {
	Db *sql.DB
}

type Cart struct {
	CartId    int64  `json:"cart_id"`
	RobotName string `json:"robot_name"`
	Cost      string `json:"cost"`
	CreatedOn string `json:"created_on"` // should be time.Time, will look into later
}
