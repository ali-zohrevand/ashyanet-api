package models

import "gopkg.in/mgo.v2/bson"

type PermissionModelInDB struct {
}
type PermissionModel struct {
}
type PermiisonPolicy struct {
	PolicyList []Policy
}
type Policy struct {
	Sub string
	Obj string
	Act string
}
type CasbinPermision struct {
	Id     bson.ObjectId `json:"id" bson:"_id"`
	Model  string
	Policy string
}
