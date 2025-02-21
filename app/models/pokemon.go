package models

type Pokemon struct {
	Id int `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Type1 string `json:"type1" bson:"type1"`
	Type2 string `json:"type2" bson:"type2"`
}
