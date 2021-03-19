package main
import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "./Paquete"
)
func main(){
  fmt.Println("Ingrese Operación")
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  op := scanner.Text()
  if strings.Contains(op,"+") {
    resultado := Paquete.Sumar(op)
    fmt.Println(resultado)
  } else if strings.Contains(op,"-") {
    resultado := Paquete.Restar(op)
    fmt.Println(resultado)
  } else if strings.Contains(op,"*") {
    resultado := Paquete.Multiplicar(op)
    fmt.Println(resultado)
  }  else if strings.Contains(op,"/") {
    resultado := Paquete.Dividir(op)
    fmt.Println(resultado)
  } else {

  }
}
