package pokebuild

import (
	"fmt"

	"github.com/florentsorel/go-pokebuild/models"
)

type pokemonService service

func (p *pokemonService) GetAll() ([]*models.Pokemon, error) {
	req, err := p.client.newRequest("pokemon")
	if err != nil {
		return nil, err
	}

	var pokemon []*models.Pokemon
	_, err = p.client.do(req, &pokemon)
	if err != nil {
		return nil, err
	}

	return pokemon, nil
}

func (p *pokemonService) GetAllWithLimit(limit int16) ([]*models.Pokemon, error) {
	req, err := p.client.newRequest(fmt.Sprintf("pokemon/limit/%d", limit))
	if err != nil {
		return nil, err
	}

	var pokemon []*models.Pokemon
	_, err = p.client.do(req, &pokemon)
	if err != nil {
		return nil, err
	}

	return pokemon, nil
}

func (p *pokemonService) GetByName(name string) (*models.Pokemon, error) {
	req, err := p.client.newRequest(fmt.Sprintf("pokemon/%s", name))
	if err != nil {
		return nil, err
	}

	var pokemon models.Pokemon
	_, err = p.client.do(req, &pokemon)
	if err != nil {
		return nil, err
	}

	return &pokemon, nil
}

func (p *pokemonService) Detail(id int16) (*models.Pokemon, error) {
	req, err := p.client.newRequest(fmt.Sprintf("pokemon/%d", id))
	if err != nil {
		return nil, err
	}

	var pokemon models.Pokemon
	_, err = p.client.do(req, &pokemon)
	if err != nil {
		return nil, err
	}

	return &pokemon, nil
}
