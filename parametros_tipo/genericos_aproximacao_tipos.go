/*
~string significa "qualquer tipo cujo tipo subjacente (underlying type) seja string".

Isso é diferente de apenas string. Se a interface fosse string, apenas o tipo string
literal seria aceito. Com ~string, você pode criar seus próprios tipos, como type Status string,
e eles satisfarão a restrição, pois o tipo subjacente de Status é string.
*/
package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

// --- INÍCIO DO CÓDIGO GENÉRICO ---
// Esta é a interface de restrição para nosso tipo genérico.
// Ela exige que qualquer tipo usado com a função `unmarshalEnum`
// tenha `string` como seu tipo subjacente (graças ao `~`) e
// possua um método `IsValid()` que retorna um booleano.
type EnumValido interface {
	~string
	EValido() bool
}

// Esta é a função genérica que desserializa e valida um "enum".
// `[T ValidEnum]` significa que ela aceita um tipo `T` que cumpre as regras de `ValidEnum`.
func unmarshalEnum[T EnumValido](data []byte, alvo *T, errInvalido error) error {
	// Declara uma variável string temporária para receber o valor do JSON.
	var s string
	// Tenta desserializar o JSON para a string `s`. Se falhar, é um JSON inválido.
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	// Converte a string `s` para o nosso tipo genérico `T`.
	// Isso é seguro por causa da restrição `~string`.
	tmp := T(s)
	// Chama o método `EValido()` do nosso tipo `T` para verificar se o valor é permitido.
	if !tmp.EValido() {
		return errInvalido
	}
	// Se a validação passou, atribui o valor `tmp` ao ponteiro `alvo`
	*alvo = tmp
	// Retorna `nil` para indicar que a operação foi um sucesso.
	return nil

}

// --- FIM DO CÓDIGO GENÉRICO ---

// 1. Definindo nosso tipo enum personalizado chamado 'Status'.
// O tipo subjacente é 'string', o que o torna compatível com `~string`.
type Status string

const (
	StatusPendente Status = "pendente"
	StatusCompleto Status = "completo"
	StatusFalho    Status = "falhou"
)

// 3. Implementando o método EValido() para o tipo 'Status'.
// Esta implementação faz com que o tipo 'Status' satisfaça a interface 'EnumValido'.
func (s Status) EValido() bool {
	switch s {
	case StatusPendente, StatusCompleto, StatusFalho:
		return true
	default:
		return false
	}
}

func main() {
	jsonPendente := []byte(`"pendente"`)
	// Declara uma variável do tipo Status para armazenar o resultado.
	var statusOK Status
	// Chama a função genérica, passando o JSON, o ponteiro para a variável e o erro customizado.
	err := unmarshalEnum(jsonPendente, &statusOK, errors.New("status inválido fornecido"))

	if err != nil {
		fmt.Printf("Erro no cenário de sucesso: %v\n", err)
	} else {
		fmt.Printf("Sucesso! Status desserializado: %s\n", statusOK)
	}

	// Um JSON válido, mas com um valor que não pertence ao nosso enum.
	jsonInvalido := []byte(`"desconhecido"`)
	var statusErro Status
	err = unmarshalEnum(jsonInvalido, &statusErro, errors.New("status inválido fornecido"))

	if err != nil {
		fmt.Printf("Erro no cenário de falha: %v\n", err)
	} else {
		fmt.Printf("Falha! Status desserializado: %s\n", statusErro)
	}

}
