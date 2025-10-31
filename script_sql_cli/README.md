# Consulta ao Banco de Dados em Go
## Descrição
Este projeto é um aplicativo em Go que executa uma consulta ao banco de dados MySQL. O aplicativo recebe como argumento de linha de comando o nome de um nó (node) e realiza uma consulta ao banco de dados para obter informações relacionadas a esse nó.

## Principais Conceitos
* **Consulta SQL**: O aplicativo utiliza uma consulta SQL para buscar informações no banco de dados.
* **Conexão ao Banco de Dados**: O aplicativo se conecta ao banco de dados utilizando o driver MySQL para Go.
* **Tratamento de Erros**: O aplicativo trata os erros que ocorrem durante a execução da consulta e imprime uma resposta de erro padronizada em formato JSON.
* **Resposta em JSON**: O aplicativo imprime a resposta da consulta em formato JSON.

## Como Usar
1. Clone o repositório e instale as dependências necessárias: `go get -u github.com/go-sql-driver/mysql` e `go get -u github.com/joho/godotenv`.
2. Crie um arquivo `.env` com as variáveis de ambiente necessárias:
	* `DB_USERG`: nome do usuário do banco de dados
	* `DB_PASSG`: senha do usuário do banco de dados
	* `DB_HOSTG`: endereço do host do banco de dados
	* `DB_NAMEG`: nome do banco de dados
3. Execute o aplicativo passando o nome do nó como argumento de linha de comando: `go run main.go NOME_DO_NÓ`

## Contato
Para mais informações ou para relatar problemas, por favor entre em contato com [Valdemir Bezerra de Souza Júnior](mailto:badmoon25@gmail.com).