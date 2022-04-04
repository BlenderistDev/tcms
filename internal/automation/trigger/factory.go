package trigger

import "tcms/pkg/tcms"

func GetList() *tcms.TriggerList {
	return &tcms.TriggerList{Triggers: []*tcms.TriggerDescription{
		{
			Name: "time",
			Fields: []*tcms.TriggerDescription_Field{
				{
					Name:        "timestamp",
					Description: "Current timestamp",
				},
			},
		},
	}}
}
