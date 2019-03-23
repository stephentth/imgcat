package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())

}

type resultAPI struct {
	// We only interest in field "name"
	Name string `json:"name"`
}

func GetRandomPokemon(id int) (string, error) {
	if id == 0 {
		id = rand.Intn(TotalNumberPokemon-1) + 1
	}

	url := PokemonAPIEndpoint + strconv.Itoa(id)
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	jsn, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var data resultAPI
	err = json.Unmarshal(jsn, &data)
	if err != nil {
		return "", err
	}
	return data.Name, nil
}

func GetImageFromHTTP(url string) (io.Reader, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}
