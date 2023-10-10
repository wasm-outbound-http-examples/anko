package main

import (
	"fmt"

	"github.com/mattn/anko/core"
	"github.com/mattn/anko/env"
	_ "github.com/mattn/anko/packages"
	"github.com/mattn/anko/vm"
)

func main() {
	envrmt := env.NewEnv()
	core.Import(envrmt)

	code := `
	var http, ioutil = import("net/http"), import("io/ioutil")
	res, _ = http.DefaultClient.Get("https://httpbin.org/anything")
	b, _ = ioutil.ReadAll(res.Body)
	println(toString(b))
	res.Body.Close()
`

	_, err := vm.Execute(envrmt, nil, code)
	if err != nil {
		fmt.Println(err)
	}
}
