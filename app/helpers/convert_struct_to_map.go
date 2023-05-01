package helpers

import "encoding/json"

type ConvertStructToMapString struct {
	Input any // struct
}

func (h ConvertStructToMapString) Do() map[string]string {
	var mapVar map[string]string

	data, _ := json.Marshal(h.Input)
	err := json.Unmarshal(data, &mapVar)
	if err != nil {
		return nil
	}

	return mapVar
}
