package tools

import (
	"fmt"
	"strings"
)

func StringToQueryLocate(s string, fieldName string) string {
	strMap := strings.Split(EmptySpaces(s), " ")
	var strLocate []string
	for _, nm := range strMap {
		strLocate = append(strLocate, fmt.Sprintf(`LOCATE('%s', %s) > 0`, nm, fieldName))
	}
	return strings.Join(strLocate[:], " OR ")
}
