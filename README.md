[![Open in Visual Studio Code](https://classroom.github.com/assets/open-in-vscode-f059dc9a6f8d3a56e377f745f24479a46679e63a5d9fe6f495e02850cd0d8118.svg)](https://classroom.github.com/online_ide?assignment_repo_id=5989683&assignment_repo_type=AssignmentRepo)
[![GoDoc](https://godoc.org/github.com/dgraph-io/dgo?status.svg)](https://godoc.org/github.com/dgraph-io/dgo)





# API-TAREFAS
> Thomaz Flanklin de Souza Jorge

* API-REST implementada utilizando GoLang, como resultado de um trabalho da disciplina **GCC129 - Sistemas Distribuídos** 
* **Professor** Neumar Costa Malheiros

### Documentação e testes
É possível verificar a documentação e realizar testes utilizando o Swagger

[![logo swagger](https://emojis.slackmojis.com/emojis/images/1582769394/7925/swagger.png?1582769394)Documentação Swagger](https://app.swaggerhub.com/apis-docs/ThomazSIUFLA/API-TAREFAS-ThomazSIUFLA/1.0.1)

### API-Disponível
A API encontra-se disponível e livre para uso no heroku, inclusive com banco de dados postgres

## [ ![logo heroku](https://emojis.slackmojis.com/emojis/images/1465929657/511/heroku.png?1465929657) API- heroku](https://api-tarefas2.herokuapp.com/tarefas/)

## Como executar

**Necessário Instalação GoLang**
Para instalar, [CLIQUE AQUI](https://golang.org/doc/install) e instale de acordo com seu sistema operacional

![golang](https://emojis.slackmojis.com/emojis/images/1454546974/291/golang.png?1454546974)

### Postgres em Docker
*Foi utilizado um container com a imagem do postgres na versão 12.8*

[Instalar Docker](https://docs.docker.com/engine/install/)

[Instalar docker-compose](https://docs.docker.com/compose/install/)

* Na raíz do projeto, execute o comando 
```docker
docker-compose up
```
  - Para conferir se o banco está rodando, execute
  ```docker
  docker-compose ps
  ```
  Deverá ver um resultado semelhante a esse:



### API
1- Clone o repositório e entre na raíz do projeto

```git
git clone https://github.com/dcc-ufla/api-tarefas-ThomazSIUFLA.git
cd api-tarefas-ThomazSIUFLA
```

2- Execute o comando para definir o GOPATH

```go
go mod init
go mod tidy
```
> Obs. Se receber o erro `go: modules disabled by GO111MODULE=off; see 'go help modules'` execute `export GO111MODULE=on` e torne executar o comando 
```go
go mod init
```

3- Execute o comando para executar o servidor

```go
go run main.go
```

3.1- Build do projeto

```go
go build
```
