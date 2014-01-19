package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"os"
	"strconv"
	"strings"
)

func die(msg string) {
	os.Stderr.WriteString(msg + "\n")
	os.Exit(1)
}

func checkErr(err error) {
	if err == nil {
		return
	}
	die(err.Error())
}
