package redis


import (
	"errors"
	"strconv"
	"time"
	"fmt"
	"backend/api"
)



func (r *RedisORM) CreateAuth(userid uint64, atd api.TokenDetails, rtd api.TokenDetails) error {

	at := time.Unix(atd.Expires, 0)
	rt := time.Unix(rtd.Expires, 0)
	now := time.Now()

	if err := r.Set(atd.Uuid, strconv.Itoa(int(userid)), at.Sub(now)).Err(); err != nil {
		return err
	}
	if err := r.Set(rtd.Uuid, strconv.Itoa(int(userid)), rt.Sub(now)).Err(); err != nil {
		return err
	}

	return nil
}



func (r *RedisORM) FetchAuth(authD *api.AccessDetails) (uint64, error) {
	userid, err := r.Get(authD.AccessUuid).Result()
	if err != nil {
		return 0, err
	}

	userID, _ := strconv.ParseUint(userid, 10, 64)
	if authD.UserId != userID {
		return 0, errors.New("unauthorized")
	}

	return userID, nil
}



func (r *RedisORM) DeleteAuth(givenUuid string) (uint64, error) {
	deleted, err := r.Del(givenUuid).Result()
	if err != nil {
		return 0, err
	}

	return uint64(deleted), nil
}



func (r *RedisORM) DeleteTokens(authD *api.AccessDetails) error {

	deletedAt, err := r.Del(authD.AccessUuid).Result()
	if err != nil {
		return err
	}

	refreshUuid := fmt.Sprintf("%s++%d", authD.AccessUuid, authD.UserId)
	deletedRt, err := r.Del(refreshUuid).Result()
	if err != nil {
		return err
	}

	if deletedAt != 1 || deletedRt != 1 {
		return errors.New("something went wrong")
	}

	return nil
}
