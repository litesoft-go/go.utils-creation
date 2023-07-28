package go_utils_creation

import (
	"fmt"
	"strings"
)

// NIY -- (Not Implemented Yet) panics with method information (name and optional params)
// -
//
//goland:noinspection GoUnusedExportedFunction
func NIY(method string, params ...any) {
	panic(niy(method, "", params...))
}

// NIY_withComment -- (Not Implemented Yet) panics like indicated above, but message has an
// optional comment (optional because if effectively empty/blank then it is not added).
// -
//
//goland:noinspection GoUnusedExportedFunction,GoSnakeCaseUsage
func NIY_withComment(method, withComment string, params ...any) {
	panic(niy(method, withComment, params...))
}

// niy func for testability
func niy(method, withComment string, params ...any) string {
	withComment = strings.TrimSpace(withComment)
	if withComment != "" {
		withComment = " // " + withComment
	}
	str, prefix := "", ""
	for _, param := range params {
		if s, ok := param.(string); ok {
			str += prefix + "'" + s + "'"
		} else {
			str += prefix + fmt.Sprintf("%v", param)
		}
		prefix = ", "
	}
	return method + "(" + str + ")" + withComment
}
