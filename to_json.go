package entity

import (
	"encoding/json"
)

func (o *Entity) ToJSON() ([]byte, error) {
	return json.Marshal(o)
}

func (a *Array) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}
