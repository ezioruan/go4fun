package main

import (
	"fmt"
	"go/token"
	"net/http"

	"github.com/sbinet/go-eval/pkg/eval"
)

var fset = token.NewFileSet()
var world = eval.NewWorld()

func evalHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		codeSource := r.FormValue("codeSource")
		code, err := world.Compile(fset, codeSource)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		result, err := code.Run()
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		if result != nil {
			w.Write([]byte(result.String()))
			return
		}

		fmt.Println("codeSource:", codeSource)
	} else if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		html := `<form class="form-horizontal" action="/" method="POST" accept-charset="utf-8">
            <textarea name="codeSource" style="height:400px;width:400px"></textarea>
            <br/>
            <input type="submit" value="提交"/>
        </form>
                `
		w.Write([]byte(html))

	}
}
func main() {
	http.HandleFunc("/", evalHandler)
	http.ListenAndServe(":8080", nil)
}
