package db

import (
	"backend/api"
	"time"
)



func (db *DBORM) AddMileageRequest(request *api.MileageRequest) error {
	request.Time = time.Now()
	request.State = api.MileageReqStatusPending
	return db.Create(&request).Error
}



func (db *DBORM) GetAllMileageRequests() (requests []api.MileageRequest, err error) {
	if err := db.Model(api.MileageRequest{State: api.MileageReqStatusPending}).Find(&requests).Error; err != nil {
		return nil, err
	}

	for idx, request := range requests {
		if requests[idx].Key, err = db.GetUserKey(request.ID); err != nil {
			return nil, err
		}
	}

	return requests, nil
}

func (db *DBORM) GetAllMileageRequestsByID(userid uint64) (requests []api.MileageRequest, err error) {
	if err := db.Model(api.MileageRequest{
		State: api.MileageReqStatusPending,
		ID:    userid,
	}).Find(&requests).Error; err != nil {
		return nil, err
	}

	userkey, err := db.GetUserKey(userid)
	if err != nil {
		return nil, err
	}

	for idx := range requests {
		requests[idx].Key = userkey
	}

	return requests, nil
}



func (db *DBORM) AcceptMileageRequest(requestId uint64) error {
	request := api.MileageRequest{ReqId: requestId}
	return db.Model(&api.MileageRequest{}).Where(request).Update("state", api.MileageReqStatusAccepted).Error
}

func (db *DBORM) RejectMileageRequest(requestId uint64) error {
	request := api.MileageRequest{ReqId: requestId}
	return db.Model(&api.MileageRequest{}).Where(request).Update("state", api.MileageReqStatusRejected).Error
}

