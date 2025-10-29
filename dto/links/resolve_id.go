package links

import (
	"eercase/pattern"
	"strings"
)

// resolveID devolve **uma string numérica** correspondente ao ID real.
//
//   - Se o valor já vier com o prefixo de banco (“DB_123”), remove-o e retorna “123”.
//   - Caso contrário procura no nodeMap (ID temporário do front) e devolve
//     o uint mapeado convertido para string.
//   - Se não encontrar, devolve "0".
func resolveID(raw string, nodeMap map[string]string) string {
	if strings.HasPrefix(raw, pattern.DatabasePrefix) {
		return raw
	}
	if v, ok := nodeMap[raw]; ok {
		return v
	}
	return "0"
}
