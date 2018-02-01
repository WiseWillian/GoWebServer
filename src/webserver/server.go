package main

/*  Autor: Rafael Willian                                                                *
 *  Descrição: O seguinte código trata-se de um Webserver que retorna, em 90% das        *  
 * 	chamadas, um número inteiro randômico, e nos outros 10% uma string também randômica  */

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

//Estrutura de dados que guarda uma string, para padronizar o retorno
type ApiString struct {
	Data string `json:"result"`
}

//Estrutura de dados que guarda um numero inteiro, para padronizar o retorno
type ApiInt struct {
	Data int `json:"result"`
}

//Função que retorna uma string randomica
func randomizeString(generator *rand.Rand, size int) *ApiString {
	var retorno = new (ApiString)
	str := make([]byte, size)

	//A cada iteração, uma letra da string é selecionada de 'letters'
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

//Função que reinicia a contagem de 10 ou incrementa
func checkDecade(generator *rand.Rand) {
	if decadeCounter == decadeMax {
		decadeCounter = 0
		stringPosition = generator.Intn(decadeMax)
	} else {
		decadeCounter++
	}
}

//Função que retorna o valor randomico de acordo com a contagem
func getValue(w http.ResponseWriter, r *http.Request) {
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)

	//Caso a contagem coincida com a posição da string uma string é retornada
	if decadeCounter == stringPosition {
		json.NewEncoder(w).Encode(randomizeString(generator, 14))
	} else { //Caso contrário um int é retornado
		json.NewEncoder(w).Encode(randomizeInt(generator, 100))
	}

	checkDecade(generator)
}

//Função que reinicia os dados básicos do Webserver
func resetServer() {
	decadeCounter = 0
	stringPosition = 0
}

func main() {
	source := rand.NewSource(time.Now().UnixNano()) //Cria uma seed para geração randomica
	generator := rand.New(source) //Cria um gerador randômico através da seed
	stringPosition = generator.Intn(decadeMax) //Inicia a posição da string, para a primeira chamada

	router := mux.NewRouter() //Cria um roteador de requisições
	//Configura-se dois endpoints /reset e /value para serem acessados via GET
	router.HandleFunc("/reset", getValue).Methods("GET") 
	router.HandleFunc("/value", getValue).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router)) //Inicia o servidor
}