package api

type AccessDetails struct {
	AccessUuid string
	UserId     uint64
	UserStatus string
}

type TokenDetails struct {
	Token string
	Uuid string
	Expires int64
}