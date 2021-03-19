package Operadores
import (
  "strings"
  "strconv"
  "fmt"
)

//Realiza la suma
func Sumar(operacion string) int {
  valores := strings.Split(operacion,"+")
  resultado := 0
  for i := 0; i < len(valores); i++ {
    num, error := strconv.Atoi(valores[i])
    if error != nil {
      fmt.Println("Tienes que ingresar un numero")
    } else {
      resultado += num
    }

  }
  return resultado
}

//Realiza la resta
func Restar(operacion string) int {
  valores := strings.Split(operacion,"-")
  resultado := 0
  for i := 0; i < len(valores); i++ {
    num, error := strconv.Atoi(valores[i])
    if error != nil {
      fmt.Println("Tienes que ingresar un numero")
    } else {
      if resultado == 0 {
        resultado = num
      } else {
        resultado -= num
      }

    }

  }
  return resultado
}
//Mult
func Multiplicar(operacion string) int {
  valores := strings.Split(operacion,"*")
  resultado := 0
  for i := 0; i < len(valores); i++ {
    num, error := strconv.Atoi(valores[i])
    if error != nil {
      fmt.Println("Tienes que ingresar un numero")
    } else {
      if resultado == 0 {
        resultado = num
      } else {
        resultado *= num
      }

    }

  }
  return resultado
}

//Div
func Dividir(operacion string) int {
  valores := strings.Split(operacion,"/")
  resultado := 0
  for i := 0; i < len(valores); i++ {
    num, error := strconv.Atoi(valores[i])
    if error != nil {
      fmt.Println("Tienes que ingresar un numero")
    } else {
      if resultado == 0 {
        resultado = num
      } else {
        resultado /= num
      }

    }

  }
  return resultado
}
