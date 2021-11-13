package model

type Automation struct {
	Id       string `bson:"_id"`
	Triggers []string
	Actions  []string
}
