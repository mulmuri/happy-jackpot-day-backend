package api



type MileageByDegree struct {
	ID       uint64 `gorm:"column:user_id" json:"-"`
	Key      string `gorm:"-"              json:"ID"`
	Degree   int    `gorm:"column:degree"  json:"degree"`
	Amount   int    `gorm:"column:amount"  json:"amount"`
}

func (MileageByDegree) TableName() string {
	return "mileage_degree"
}