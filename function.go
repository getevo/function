package function

import (
	"reflect"
	"runtime"
	"strings"
)

func Name(fn interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
}

func GetPKG(fn interface{}) string {
	fullName := Name(fn)
	parts := strings.Split(fullName, "/")
	lastPart := parts[len(parts)-1]
	pkgParts := strings.Split(lastPart, ".")
	return strings.Join(pkgParts[:len(pkgParts)-1], ".") // Exclude function name
}
