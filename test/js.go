package main

import (
	"fmt"
	"github.com/robertkrimen/otto/parser"
)

func main() {
	filename := "" // A filename is optional
	src := `
    // Sample xyzzy example
    (function(){
        if (3.14159 > 0) {
            console.log("Hello, World.");
            return;
        }

        var xyzzy = NaN;
        console.log("Nothing happens.");
        return xyzzy;
    })();
`

	// Parse some JavaScript, yielding a *ast.Program and/or an ErrorList
	program, _ := parser.ParseFile(nil, filename, src, 0)
	fmt.Println(program)
}
