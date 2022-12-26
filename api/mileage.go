package api

import (
	"time"
)

type Mileage struct {
	ID     uint64 `gorm:"column:user_id" json:"-"`
	Key    string `gorm:"-"              json:"ID"`
	Amount int    `gorm:"column:amount"  json:"amount"`
}

type MileageRequest struct {
	ReqId  uint64    `gorm:"column:req_id"     json:"requestNo"`
	ID     uint64    `gorm:"column:user_id"    json:"-"`
	Key    string    `gorm:"-"                 json:"userID"`
	Amount int       `gorm:"column:amount"     json:"amount"`
	Time   time.Time `gorm:"column:created_at" json:"requestTime"`
	State  string    `gorm:"column:state"      json:"state"`
}

var MileageReqStatusAccepted string = "Accepted"
var MileageReqStatusRejected string = "Rejected"
var MileageReqStatusPending string = "Pending"
var MileageReqStatusFailed string = "Failed"

func (Mileage) TableName() string {
	return "mileage"
}

func (MileageRequest) TableName() string {
	return "mileage_request"
}
