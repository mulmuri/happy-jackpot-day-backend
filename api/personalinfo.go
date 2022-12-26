package api





type PersonalInfo struct {
	ID              uint64 `gorm:"column:user_id"`
	PW              string `gorm:"-"                        json:"PW"`
	Key             string `gorm:"-"                        json:"ID"`
	UserName        string `gorm:"column:name"              json:"name"`
	Email           string `gorm:"column:email"             json:"email"`
	PhoneNumber     string `gorm:"column:phone_number"      json:"phoneNumber"`
	AccountBankName string `gorm:"column:account_bank_name" json:"accountBankName"`
	AccountNumber   string `gorm:"column:account_number"    json:"accountNumber"`
	Recommender     string `gorm:"-"                        json:"recommender"`
}

func (PersonalInfo) TableName() string {
	return "personal_info"
}



func (PersonalInfo) CheckIDValidity() error {
	return nil
}

func (PersonalInfo) CheckEmailValidity() error {
	return nil
}

func (PersonalInfo) CheckPhoneNumberValidity() error {
	return nil
}

func (this *PersonalInfo) CheckDataValidity() error {
	if err := this.CheckIDValidity(); err != nil {
		return err
	}

	if err := this.CheckEmailValidity(); err != nil {
		return err
	}

	if err := this.CheckPhoneNumberValidity(); err != nil {
		return err
	}

	return nil
}
