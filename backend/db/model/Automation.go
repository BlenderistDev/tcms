package model

type Mapping struct {
	Simple bool   `bson:"simple"`
	Name   string `bson:"name"`
	Value  string `bson:"value"`
}

type Action struct {
	Name    string             `bson:"name"`
	Mapping map[string]Mapping `bson:"mapping"`
}

type Automation struct {
	Id       string   `bson:"_id"`
	Triggers []string `bson:"triggers"`
	Actions  []Action `bson:"actions"`
}
