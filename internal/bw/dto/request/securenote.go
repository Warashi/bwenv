package request

import "encoding/json"

type _type struct{}

func (_type) MarshalJSON() ([]byte, error) {
	return json.Marshal(2)
}

type _secureNote struct{}

func (_secureNote) MarshalJSON() ([]byte, error) {
	type s struct {
		Type int `json:"type"`
	}
	return json.Marshal(s{0})
}

type SecureNote struct {
	FolderID   string      `json:"folderId"`
	Type       _type       `json:"type"`
	Name       string      `json:"name"`
	Notes      string      `json:"notes"`
	SecureNote _secureNote `json:"secureNote"`
}
