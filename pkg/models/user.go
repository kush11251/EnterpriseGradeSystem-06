package models

import (
	"encoding/json"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
