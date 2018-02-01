package main

import (
	"log"
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"math/rand"
)

const stringSize = 16 //Constante que guarda o tamanho padrão da string
const decadeMax = 9 //Constante que indica até onde o contador de década vai
const letters = "abcdefghijklmnopqrstuvwxyz" //Letras para geracao da string randomica

var decadeCounter int //Conta quantas vezes o endpoint foi chamado de 10 em 10
var stringPosition int //Guarda a posicao que a string deve ser mostrada na decada

type ApiString struct {
	Data string `json:"result"`
}

type ApiInt struct {
	Data int `json:"result"`
}

//Função que retorna uma string randomica
func randomizeString(generator *rand.Rand, size int) *ApiString {

}

//Função que retorna um numero inteiro randomico
func randomizeInt(generator *rand.Rand, max int) *ApiInt {
	var retorno = new (ApiInt)
	retorno.Data = generator.Intn(max)
	return retorno
}

func getValue(w http.ResponseWriter, r *http.Request) {

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/roulette", getValue).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}