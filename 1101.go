//1101 - Sequência de Números e Soma-1

package main

import(
  "fmt"
)

func main(){
  var M, N int64
  var menor, maior int64
  var soma int64

  for{
    fmt.Scanf("%d", &M)
    fmt.Scanf("%d", &N)

    if M <= 0 || N <= 0{
      return
    }else if M <= N{
      menor = M
      maior = N
    }else{
      menor = N
      maior = M
    }

    for i := menor; i <= maior; i++{
      fmt.Printf("%d ", i)
      soma += i
    }

    fmt.Printf("Sum=%d\n",soma)

    soma = 0
    menor = 0
    maior = 0
  }
}
