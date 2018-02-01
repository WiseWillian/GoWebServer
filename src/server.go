package main

import (
	"log"
	"time"
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
	var retorno = new (ApiString)
	str := make([]byte, size)

	for i := range str {
		str[i] = letters[generator.Intn(len(letters))]
	}

	retorno.Data = string(str)
	return retorno
}

//Função que retorna um numero inteiro randomico
func randomizeInt(generator *rand.Rand, max int) *ApiInt {
	var retorno = new (ApiInt)
	retorno.Data = generator.Intn(max)
	return retorno
}

func checkDecade(generator *rand.Rand) {
	if decadeCounter == decadeMax {
		decadeCounter = 0
		stringPosition = generator.Intn(decadeMax)
	} else {
		decadeCounter++
	}
}

func getValue(w http.ResponseWriter, r *http.Request) {
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)

	if decadeCounter == stringPosition {
		json.NewEncoder(w).Encode(randomizeString(generator, 14))
	} else {
		json.NewEncoder(w).Encode(randomizeInt(generator, 100))
	}

	checkDecade(generator)
}

func resetServer() {
	decadeCounter = 0
	stringPosition = 0
}

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)
	stringPosition = generator.Intn(decadeMax)

	router := mux.NewRouter()
	router.HandleFunc("/reset", getValue).Methods("GET")
	router.HandleFunc("/value", getValue).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}