package gotoc

import (
	"bytes"
	"fmt"
	"go/ast"
	"os"
	"strconv"
	"strings"
)

func notImplemented(n ast.Node) {
	fmt.Fprintf(os.Stderr, "not implemented: %s <%T>", n, n)
	os.Exit(1)
}

func upath(path string) string {
	return strings.Replace(path, "/", "_", -1)
}

func tmpname(w *bytes.Buffer) string {
	return "__" + strconv.Itoa(w.Len())
}
