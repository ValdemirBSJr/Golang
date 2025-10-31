package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// APIResposta define a estrutura de saída JSON para sucesso ou falha
type APIResposta struct {
	Node     string      `json:"node"`
	CMTS     interface{} `json:"cmts,omitempty"`
	SNROF    interface{} `json:"snr_of,omitempty"`
	Data     interface{} `json:"data,omitempty"`
	Comando  string      `json:"comando"`
	Resposta string      `json:"resposta"`
	Mensagem string      `json:"mensagem",omitempty`
}

// QueryResult armazena os dados brutos do banco, lidando com valores NULL
type QueryResult struct {
	Hostname   sql.NullString
	SNROF      sql.NullFloat64
	DataColeta sql.NullTime
}

// Constante da consulta SQL
const querySQL = `
SELECT
    e.hostname,
    h.snr_of,
    h.data_coleta
FROM
    historico_SNR_ofdma h
JOIN
    equipamento e ON h.id_cmts = e.id
WHERE
    h.descritor = ?
    AND h.data_coleta >= CURDATE()
    AND h.data_coleta < (CURDATE() + INTERVAL 1 DAY)
ORDER BY
    h.data_coleta DESC
LIMIT 1;
`

// run contém a lógica principal da aplicação.
func run() error {
	// Obtem o 'node' (descritor) dos argumentos de linha de comando
	nodeNome, err := retornaNodeArgs()
	if err != nil {
		imprimeJsonErro(nodeNome, err) //Imprime o json do erro
		return err
	}

	// Carregar .env e construir a string DSN (Data Source Name)
	dsn, err := construirDSNEnv()
	if err != nil {
		imprimeJsonErro(nodeNome, err)
		return err

	}

	// Conectar ao banco
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		err = fmt.Errorf("falha ao preparar conexão com o banco: %w", err)
		imprimeJsonErro(nodeNome, err)
		return err
	}
	defer db.Close() // Garante que a conexaão será fechada

	// Executar a consulta
	resultado, err := consultaNodeNota(db, nodeNome)
	if err != nil {
		// Trata o erro 'sql.ErrNoRows' de forma amigável
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("Nenhum registro encontrado para o node %s. Verifique se o nome está correto, ou se está incompleto.", nodeNome)
		} else {
			err = fmt.Errorf("falha ao executar consulta: %w", err)
		}

		imprimeJsonErro(nodeNome, err)
		return err
	}

	// Se tudo der certo, imprimir o JSON de sucesso
	retornaJSONSucesso(nodeNome, resultado)
	return nil // retorna OK!

} // fim do run()

// retornaNodeArgs valida e retorna o argumento da linha de comando.
func retornaNodeArgs() (string, error) {

	if len(os.Args) != 2 {
		return "", errors.New("uso: main NOME_DO_NODE")
	}
	nodeNome := strings.ToUpper(os.Args[1])
	if nodeNome == "" {
		return "", errors.New("o nome do node não pode ser vazio")
	}

	return nodeNome, nil

} // fim do retornaNodeArgs()

// construirDSNEnv carrega o .env e cria a string de conexão.
func construirDSNEnv() (string, error) {

	//const caminhoEnv ="/ferramentas/clarobot/pasta_env/.env"

	if err := godotenv.Load(); err != nil {
		return "", errors.New("erro ao carregar o arquivo .env")
	}

	usuario := os.Getenv("DB_USERG")
	passwd := os.Getenv("DB_PASSG")
	host := os.Getenv("DB_HOSTG")
	porta := ""
	dbnome := os.Getenv("DB_NAMEG")

	if usuario == "" || passwd == "" || host == "" || dbnome == "" {
		return "", errors.New("Variáveis do banco devem ser definidas no .env!")
	}

	if porta == "" {
		porta = "3306"
	}

	// Formato DSN: "usuario:password@tcp(host:porta)/dbnome"
	// Adicionamos '?parseTime=true' ao final da DSN.
	// Isso instrui o driver a converter DATETIME do banco para time.Time do Go.
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", usuario, passwd, host, porta, dbnome)
	return dsn, nil

} //fim construirDSNEnv()

// consultaNodeNota executa a consulta no banco com segurança
func consultaNodeNota(db *sql.DB, nodeNome string) (*QueryResult, error) {
	var resultado QueryResult

	// Usamos um contexto com timeout, uma boa prática para queries
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// QueryRowContext é usado pois esperamos no máximo 1 linha (LIMIT 1)
	// Passamos nodeNome como parâmetro (o '?' na query) para evitar SQL Injection.
	// O .Scan() precisa corresponder à ordem das colunas no SELECT.
	err := db.QueryRowContext(ctx, querySQL, nodeNome).Scan(
		&resultado.Hostname,
		&resultado.SNROF,
		&resultado.DataColeta,
	)

	// Retorna o resultado e o erro (se tiver, pode ser nil ou sql.ErrNoRows)
	return &resultado, err

} // fim consultaNodeNota()

// imprimeJsonErro formata e imprime uma resposta de erro padronizada.
func imprimeJsonErro(nodeNome string, err error) {
	resposta := APIResposta{
		Node:     nodeNome,
		Comando:  "snrof",
		Resposta: "nok",
		Mensagem: err.Error(),
	}

	// Usamos NewEncoder para escrever diretamente no stdout, é mais eficiente
	if err := json.NewEncoder(os.Stdout).Encode(resposta); err != nil {
		log.Printf("falha ao codificar/formatar JSON de erro: %v", err)
	}

} // fim imprimeJsonErro()

/*
--- Funções Auxiliares para 'sql.Null*' ---

	Estas funções convertem os tipos 'sql.Null*' em 'interface{}'
	para que o JSON os serialize como o valor (ex: 34.5) ou 'null'.
*/
func sqlNullString(s sql.NullString) interface{} {
	if s.Valid {
		return s.String
	}
	return nil
}

func sqlNullFloat(f sql.NullFloat64) interface{} {
	if f.Valid {
		return f.Float64
	}
	return nil
}

func sqlNullTime(t sql.NullTime) interface{} {
	if t.Valid {
		// formatamos ele para uma string no layout "DD/MM/YYYY HH:MM"
		// O layout de referência do Go é "02/01/2006 15:04"
		return t.Time.Format("02/01/2006 15:04")
	}
	return nil
}

// retornaJSONSucesso formata e imprime a resposta de sucesso.
func retornaJSONSucesso(nodeNome string, resultado *QueryResult) {
	resposta := APIResposta{
		Node:     nodeNome,
		Comando:  "snrof",
		Resposta: "ok",
		// Verificamos se os valores são válidos antes de adicioná-los
		// Se forem nulos (inválidos), o JSON os mostrará como 'null'.
		CMTS:  sqlNullString(resultado.Hostname),
		SNROF: sqlNullFloat(resultado.SNROF),
		Data:  sqlNullTime(resultado.DataColeta),
	}

	if err := json.NewEncoder(os.Stdout).Encode(resposta); err != nil {
		log.Printf("falha ao formatar JSON de sucesso: %v", err)
	}
} // fim retornaJSONSucesso()

// main é o ponto de entrada. Usamos log.Fatal para capturar qualquer erro
// retornado pela nossa função principal 'run'.
func main() {
	if err := run(); err != nil {
		// 'run' já imprime o JSON de erro, então apenas saímos com status 1
		//os.Exit(1)
	}
}
