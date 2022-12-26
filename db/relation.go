package db

import (
	"backend/api"
)

func (db *DBORM) GetRecommenderID(userid uint64) (uint64, error) {
	relation := api.Relation{ID: userid}
	return relation.RecommenderID, db.Where(&api.PersonalInfo{ID: userid}).Find(&relation).Error
}

func (db *DBORM) GetRecommender(userid uint64) (string, error) {
	user := api.Relation{}
	if err := db.Model(user).Where("user_id", userid).Find(&user).Error; err != nil {
		return "", err
	}
	return db.GetUserKey(user.RecommenderID)
}


func (db *DBORM) GetFriendsList(userid uint64) ([]api.Relation, error) {
	relations := []api.Relation{}
	return relations, db.Model(relations).Where("recommender_id = ?", userid).Find(&relations).Error;
}