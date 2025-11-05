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

## 🛠️ Pré-requisitos
* **Para Executar:** [Docker](https://www.docker.com/)
* **Para Desenvolver:** [Go](https://golang.org/) (versão 1.25 ou superior)

## 🚀 Como Executar (Recomendado)

A forma mais simples de rodar este projeto é usando o Docker.

1.  **Clone o repositório:**
    ```bash
    git clone [https://github.com/seu-usuario/nome-do-repositorio.git](https://github.com/seu-usuario/nome-do-repositorio.git)
    cd nome-do-repositorio
    ```

2.  **Construa a imagem Docker:**
    ```bash
    docker build -t ssh-redirect .
    ```

3.  **Execute o contêiner:**
    Você **deve** fornecer a variável de ambiente `SSH_TARGET_URL` com o seu destino SSH.

    ```bash
    docker run -d \
      --name ssh-redirect \
      --restart=unless-stopped \
      -p 8081:8080 \
      -e SSH_TARGET_URL="ssh://usuario@seu-servidor-remoto.com:22" \
      ssh-redirect
    ```
    * `-p 8081:8080`: Mapeia a porta `8081` da sua máquina (host) para a porta `8080` do contêiner.
    * `-e SSH_TARGET_URL=...`: **(OBRIGATÓRIO)** Define o endereço SSH para onde os usuários serão redirecionados.

4.  **Acesse no navegador:**
    Abra `http://localhost:8081` no seu navegador. Você será imediatamente solicitado a abrir seu cliente SSH.

## ⚙️ Configuração

O servidor é configurado via variáveis de ambiente:

| Variável | Obrigatório | Padrão | Descrição |
| :--- | :--- | :--- | :--- |
| `SSH_TARGET_URL` | **Sim** | `""` | O URI completo para o redirecionamento. Ex: `ssh://user@host.com` |
| `PORT` | Não | `8080` | A porta interna que o servidor Go irá escutar. |

## 💻 Como Executar (Localmente para Desenvolvimento)

1.  **Clone o repositório:**
    ```bash
    git clone [https://github.com/seu-usuario/nome-do-repositorio.git](https://github.com/seu-usuario/nome-do-repositorio.git)
    cd nome-do-repositorio
    ```

2.  **Instale as dependências:**
    ```bash
    go mod tidy
    ```

3.  **Execute o projeto:**
    (Lembre-se de definir as variáveis de ambiente)
    ```bash
    export SSH_TARGET_URL="ssh://usuario@seu-servidor-remoto.com"
    export PORT="8080"
    
    go run main.go
    ```
4.  Acesse `http://localhost:8080` no seu navegador.

<details>
  <summary>🐳 Ver Dockerfile</summary>
  
  ```dockerfile
  # Dockerfile
  FROM golang:1.25.2-alpine AS builder

  WORKDIR /app
  COPY go.mod .
  # Opcional: baixe as dependências se houver alguma
  # RUN go mod download
  COPY main.go .
  
  # Constrói o binário estático
  RUN CGO_ENABLED=0 GOOS=linux go build -a -o ssh-redirect .

  # Fase final: imagem leve
  FROM alpine:latest
  
  # Adiciona certificados (bom para qualquer chamada https futura)
  RUN apk --no-cache add ca-certificates
  
  WORKDIR /root/
  COPY --from=builder /app/ssh-redirect .
  
  # Expõe a porta padrão
  EXPOSE 8080
  
  # O servidor irá rodar na porta $PORT, ou 8080 se não definida
  CMD ["./ssh-redirect"]
5. Acesse `http://localhost:porta` no seu navegador (a porta padrão configurada no projeto, geralmente 8080, mas pode ser alterada no código).
6. O navegador redirecionará automaticamente para `ssh://IP_SERVIDOR_REMOTO`.
 ```

## Contato

* **Valdemir Bezerra de Souza Júnior**
* Analista Infraestrutura | Devops | SRE | Cloud | Oracle Cloud | Linux | Docker | Kubernets | Python | Go | Rust | Lua | N8N | No Code
* [Linkedin](https://www.linkedin.com/in/valdemirbezerra/)
