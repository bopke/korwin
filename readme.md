# Korwin

Superprosta biblioteka do generowania wypowiedzi w stylu krula.

Instalacja
--
`
go get https://github.com/bopke/korwin
`

Użycie
--
```go
package main

import(
	"fmt"
	"github.com/bopke/korwin"
)

func main(){
	fmt.Print(korwin.GenerateStatement())
}
```