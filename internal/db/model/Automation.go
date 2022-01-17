package model

import "tcms/m/internal/automation/datamapper"

type Mapping struct {
	Simple bool   `bson:"simple"`
	Name   string `bson:"name"`
	Value  string `bson:"value"`
}

type DmMapper struct {
	Simple bool   `bson:"simple"`
	Value  string `bson:"value"`
}

func (m DmMapper) IsSimple() bool {
	return m.Simple
}

func (m DmMapper) GetValue() string {
	return m.Value
}

func ConvertMappingToDmMapping(mapping map[string]Mapping) map[string]datamapper.Mapping {
	dmMap := make(map[string]datamapper.Mapping)
	for key, val := range mapping {
		dmMap[key] = DmMapper{
			Simple: val.Simple,
			Value:  val.Value,
		}
	}
	return dmMap
}

type Action struct {
	Name    string             `bson:"name"`
	Mapping map[string]Mapping `bson:"mapping"`
}

type Condition struct {
	Name          string             `bson:"name"`
	Mapping       map[string]Mapping `bson:"mapping"`
	SubConditions []Condition        `bson:"sub_conditions"`
}

type Automation struct {
	Id        string     `bson:"_id"`
	Triggers  []string   `bson:"triggers"`
	Condition *Condition `bson:"condition"`
	Actions   []Action   `bson:"actions"`
}
