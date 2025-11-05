# Redirecionador SSH em Go
==========================

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
Para mais informações, contribuições, ou relatos de bugs, por favor utilize a seção de issues deste repositório ou contate-nos em [seu-email@example.com](mailto:seu-email@example.com).

---

Lembre-se de substituir `seu-usuario`, `nome-do-repositorio`, `http://localhost:porta` e `seu-email@example.com` com as informações específicas do seu projeto.