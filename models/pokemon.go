package models

import (
	"encoding/json"
)

type Pokemon struct {
	ID        int    `json:"id"`
	PokedexID int    `json:"pokedexId"`
	Name      string `json:"name"`
	Image     string `json:"image"`
	Sprite    string `json:"sprite"`
	Slug      string `json:"slug"`
	Stats     struct {
		HP             int `json:"HP"`
		Attack         int `json:"attack"`
		Defense        int `json:"defense"`
		SpecialAttack  int `json:"special_attack"`
		SpecialDefense int `json:"special_defense"`
		Speed          int `json:"speed"`
	} `json:"stats"`
	Types []struct {
		Name  string `json:"name"`
		Image string `json:"image"`
	} `json:"apiTypes"`
	Generation  int `json:"apiGeneration"`
	Resistances []struct {
		Name             string  `json:"name"`
		DamageMultiplier float64 `json:"damage_multiplier"`
		DamageRelation   string  `json:"damage_relation"`
	} `json:"apiResistances"`
	ResistanceModifyingAbilities *ResistanceModifyingAbilities `json:"resistanceModifyingAbilitiesForApi"`
	Evolutions                   []Evolution                   `json:"apiEvolutions"`
	PreEvolution                 *PreEvolution                 `json:"apiPreEvolution"`
	ResistancesWithAbilities     []struct {
		Name             string  `json:"name"`
		DamageMultiplier float64 `json:"damage_multiplier"`
		DamageRelation   string  `json:"damage_relation"`
	} `json:"apiResistancesWithAbilities"`
}

type ResistanceModifyingAbilities struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Evolution struct {
	Name      string `json:"name"`
	PokedexId int    `json:"pokedexId"`
}

type PreEvolution struct {
	Name      string `json:"name"`
	PokedexId int    `json:"pokedexIdd"`
}

func (rma *ResistanceModifyingAbilities) UnmarshalJSON(data []byte) error {
	if string(data) == "[]" {
		return nil
	}

	var v map[string]string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	rma.Name = v["name"]
	rma.Slug = v["slug"]

	return nil
}

func (pe *PreEvolution) UnmarshalJSON(data []byte) error {
	if string(data) == "\"none\"" {
		return nil
	}

	var v map[string]any
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	pe.Name = v["name"].(string)
	pe.PokedexId = int(v["pokedexIdd"].(float64))

	return nil
}

func (p *Pokemon) UnmarshalJSON(data []byte) error {
	type pk Pokemon
	var tmp pk

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	defaultRma := ResistanceModifyingAbilities{}
	defaultPe := PreEvolution{}

	if *tmp.ResistanceModifyingAbilities == defaultRma {
		tmp.ResistanceModifyingAbilities = nil
	}

	if len(tmp.Evolutions) == 0 {
		tmp.Evolutions = nil
	}

	if *tmp.PreEvolution == defaultPe {
		tmp.PreEvolution = nil
	}

	*p = Pokemon(tmp)
	return nil
}
