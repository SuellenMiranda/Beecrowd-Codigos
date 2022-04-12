//1180 - menor e posição

package main

import(
  "fmt"
)

func main(){
  var NX, menorValor, posicao int

  fmt.Scanf("%d", &NX)

  if NX <= 1 || NX >= 1000{
      return
  }
  
  vetorX := make([]int, NX)

  for i,_ := range vetorX{
    var num int
    fmt.Scanf("%d", &num)
    
    if i == 0{
      menorValor = num
    }else if menorValor > num{
      menorValor = num
      posicao = i
    }
    vetorX[i] = num
  }

  fmt.Printf("Menor valor: %d\n", menorValor)
  fmt.Printf("Posicao: %d\n", posicao)
}
