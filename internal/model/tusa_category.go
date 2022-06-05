package model

import (
	"encoding/json"
	"fmt"
)

type TusaCategory int64

const (
	Other TusaCategory = iota
	BoardGame
	Karaoke
	Clubbing
	BarEvening
	BBQ
	Sports
	Picnic
)

func (t TusaCategory) String() string {
	return [...]string{
		"Other",
		"BoardGame",
		"Karaoke",
		"Clubbing",
		"BarEvening",
		"BBQ",
		"Sports",
		"Picnic",
	}[t]
}

func FromString(category string) TusaCategory {
	return map[string]TusaCategory{
		"Other":      Other,
		"BoardGame":  BoardGame,
		"Karaoke":    Karaoke,
		"Clubbing":   Clubbing,
		"BarEvening": BarEvening,
		"BBQ":        BBQ,
		"Sports":     Sports,
		"Picnic":     Picnic,
	}[category]
}

func (t TusaCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *TusaCategory) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*t = FromString(s)
	if t.String() != s {
		return fmt.Errorf("'%s' is not a valid TusaCategory", s)
	}
	return nil
}
