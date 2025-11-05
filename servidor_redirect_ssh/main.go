package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
)

func sshHandler(w http.ResponseWriter, r *http.Request) {
	// Remove a barra inicial, se existir
	ip := strings.TrimPrefix(r.URL.Path, "/")

	// Validação básica de IP
	if ip == "" {
		http.Error(w, "IP não fornecido. Use /<IP>", http.StatusBadRequest)
		return
	}

	// Valida se é um IP válido (IPv4 ou IPv6)
	if net.ParseIP(ip) == nil {
		http.Error(w, "IP inválido, verifique o ip passado.", http.StatusBadRequest)
		return
	}

	// Redireciona para ssh://<IP>
	// NOTA: Alguns navegadores bloqueiam redirecionamento para protocolos não-HTTP/HTTPS
	// Então, alternativamente, exibimos um link clicável
	url := fmt.Sprintf("ssh://%s", ip)

	// Opção 1: Tentar redirecionamento (pode não funcionar em todos os navegadores)
	// http.Redirect(w, r, url, http.StatusSeeOther)

	// Opção 2 (mais confiável): Página com link clicável + meta refresh
	html := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="utf-8">
			<title>Conectar via SSH - CLARO NE</title>
			<meta http-equiv="refresh" content="2;url=%s">
		</head>
		<body>
			<h2>Conectando a %s...</h2>
			<p>Se a conexão não iniciar automaticamente, 
			<a href="%s">clique aqui</a>.</p>
			<p>Seu sistema deve ter um handler redirector para <code>ssh://</code> (ex: PuTTY no Windows, OpenSSH no Linux/macOS ou MobaXterm ;p).</p>

			<p><i>Desenvolvido pro Alan no final do expediente. Tenham paciência ;p</i></p>
			<br>
			<p><i>Desenvolvido por valdemir Bezerra para o Datacenter NE. Todos os direitos reservados.</i></p>
		</body>
		</html>
	`, url, ip, url)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(html))

} // sshHandler()fim

func main() {
	porta := os.Getenv("PORT")
	if porta == "" {
		porta = "8080"
	}

	http.HandleFunc("/", sshHandler)
	fmt.Printf("Servidor SSH redirect rodando na porta %s...\n", porta)

	if err := http.ListenAndServe(":"+porta, nil); err != nil {
		fmt.Fprintf(os.Stderr, "Erro: %v\n", err)
		os.Exit(1)
	}

} // main() fim
