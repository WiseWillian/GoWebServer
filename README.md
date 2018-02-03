# GoWebServer

Este repositório contém código e informações sobre a criação de um Webserver simples na linguagem Go que retorna em 90% das vezes que for chamado, um número inteiro randômico, e 10% uma string também randômica

## Design da aplicação

Para criação do código, e maior escalabilidade, foi instalado o pacote **gorilla/mux** que implementa várias funções para criação de servidores e Api's. No caso desta solução, foi utilizado o router (Roteador de requisições do pacote) que facilitou a implementação.

Primeiramente foi criada uma função que retorna um número inteiro aleatório, utilizando normalmente o pacote **math/rand**. Por outro lado, para criar uma string aleatória, foi declarada uma constante com as letras do alfabeto, na função que cria a string randômica, para cada posição da string tomou-se um inteiro aleatório para selecionar um letra da constante e aderir à string final.

Outra manobra também utilizada, foi uma variável global que conta de 0 até 9, sendo assim, quando está em 0 deve criar uma posição para a string (também entre 0 e 9). Quando esta variável chega ao 9, ela volta ao inicio (0). Essa manobra é conhecida como contador de década.

## Instalando repositório

### Instale a GoLang no seu computador

Siga os passos na documentação: https://golang.org/doc/install

### Instale o repositório no local desejado

Clone o repositório no local desejado

```
git clone https://github.com/WiseWillian/GoWebServer/
```

## Executando o Webserver

Para executar o código naveque até o diretório onde está contido

```
cd LocalDeClonagem/GoWebServer/src/webserver/
```

e rodar a aplicação: 

```
go run server.go
```

## Tecnologias utilizadas

* [gorilla/mux](https://github.com/gorilla/mux) - Pacote para criação de webservers
* [Go](https://golang.org/doc/) - Linguagem de programação para execução dos testes

## Autores

* **Rafael Willian** - *Trabalho inicial* - [WiseWillian](https://github.com/WiseWillian)
