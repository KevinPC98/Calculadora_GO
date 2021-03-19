package main
import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "./Operadores"
)
func main(){
  fmt.Println("Ingrese Operación")
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  op := scanner.Text()
  if strings.Contains(op,"+") {
    resultado := Operadores.Sumar(op)
    fmt.Println(resultado)
  } else if strings.Contains(op,"-") {
    resultado := Operadores.Restar(op)
    fmt.Println(resultado)
  } else if strings.Contains(op,"*") {
    resultado := Operadores.Multiplicar(op)
    fmt.Println(resultado)
  }  else if strings.Contains(op,"/") {
    resultado := Operadores.Dividir(op)
    fmt.Println(resultado)
  } else {

  }
}
