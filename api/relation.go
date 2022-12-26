package api



var MileageRewardingDefaultLevel int = 5
var MileageReqardingDefaultPercent int = 3



type Relation struct {
	ID            uint64 `gorm:"column:user_id"`
	RecommenderID uint64 `gorm:"column:recommender_id"`
}

func (Relation) TableName() string {
	return "relation"
}



type FriendsInfo struct {
	ID       uint64
	Amount   int    `json:"amount"`
	Key      string `json:"ID"`
	UserName string `json:"name"`
}



type FriendsCount struct {
	Count  int `json:"count"`
	Amount int `json:"amount"`
}

func (u *FriendsCount) Add(amount int) {
	u.Amount += amount
	u.Count  += 1
}
