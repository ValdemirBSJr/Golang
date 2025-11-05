# Redirecionador SSH em Go
==========================

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![HTTP](https://img.shields.io/badge/HTTP-1E90FF?style=for-the-badge&logo=data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9IjAgMCAyNCAyNCIgZmlsbD0id2hpdGUiPjxwYXRoIGQ9Ik0xMiAyQzYuNDggMiAyIDYuNDggMiAxMnM0LjQ4IDEwIDEwIDEwIDEwLTQuNDggMTAtMTBTMTcuNTIgMiAxMiAyem0wIDE4Yy00LjQxIDAtOC0zLjU5LTgtOHMzLjU5LTggOC04IDggMy41OSA4IDgtMy41OSA4LTggOHptLTQgLTRoOHYtMmgzVjloLTNWN2gtOHYYyaC0zVjEyaDN2MnpNMTAgOWg0VjdoLTR2MnpNMTQgMTdoLTR2LTJoNHYyIi8+PC9zdmc+)
![SSH](https://img.shields.io/badge/SSH-111111?style=for-the-badge&logo=data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9IjAgMCAyNCAyNCIgZmlsbD0id2hpdGUiPjxwYXRoIGQ9Ik0xOCAzSDZDRTYuNjcgMyAyIDcgMiAxMnMyLjY3IDkgNiA5aDEyYzMuMzEgMCA2LTQgNi05cy0yLjY5LTktNi05em0wIDE2SDZjLTIuMjEgMC00LTIuMjQtNC03czEuNzktNyA0LTdoMTJjMi4yMSAwIDQgMi4yNCA0IDdzLTEuNzkgNy00IDd6TTggMTBoOHYySDh6Ii8+PHBhdGggZD0iTTkgOGgydjZIOXptNCAwaDJ2NmgtMnoiLz48L3N2Zz4=)

## Descrição Curta
Este projeto é uma aplicação web simples escrita em Go que redireciona os acessos para um endereço SSH remotamente. É ideal para situações onde é necessário acessar um servidor SSH através de uma aba do navegador.

## Principais Conceitos
- **Servidor Web**: O projeto cria um servidor web que, ao ser acessado, fornece um redirecionamento automático para uma conexão SSH.
- **SSH**: Protocolo de rede seguro para operações de rede, terminal, transferência de arquivos e outros serviços de rede.
- **Go**: Linguagem de programação utilizada para desenvolver a aplicação.

## Como Usar
1. Clone o repositório:
   ```bash
   git clone https://github.com/seu-usuario/nome-do-repositorio.git
   ```
2. Navegue até o diretório do projeto:
   ```bash
   cd nome-do-repositorio
   ```
3. Construa o executável (o comando pode variar dependendo do ambiente e versão do Go):
   ```bash
   go build main.go
   ```
4. Execute o executável:
   ```bash
   ./main
   ```
   ou
   ```bash
   go run main.go
   ```
5. Acesse `http://localhost:porta` no seu navegador (a porta padrão configurada no projeto, geralmente 8080, mas pode ser alterada no código).
6. O navegador redirecionará automaticamente para `ssh://IP_SERVIDOR_REMOTO`.

## Contato

* **Valdemir Bezerra de Souza Júnior**
* Analista Infraestrutura | Devops | SRE | Cloud | Oracle Cloud | Linux | Docker | Kubernets | Python | Go | Rust | Lua | N8N | No Code
* [Linkedin](https://www.linkedin.com/in/valdemirbezerra/)
