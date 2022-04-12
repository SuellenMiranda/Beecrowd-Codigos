//2310 -  Voleibol

package main

import (
  "fmt"
)

func main() {
  
  var N int
  var nome string 
  var A, B, S, A1, B1, S1 int64
  var sumS, sumS1, sumB, sumB1, sumA, sumA1 int64

  fmt.Scanf("%d\n", &N)
  
  for i := 0; i < N; i++ {
    fmt.Scanf("%s\n", &nome)
    fmt.Scanf("%d %d %d\n", &S, &B, &A)
    fmt.Scanf("%d %d %d\n", &S1, &B1, &A1)

    sumS += S
    sumS1 += S1
    sumB += B
    sumB1 += B1
    sumA += A
    sumA1 += A1
  } 

  pcentS := float64(sumS1) / float64(sumS) * 100
  pcentB := float64(sumB1) / float64(sumB) * 100
  pcentA := float64(sumA1) / float64(sumA) * 100

	fmt.Printf("Pontos de Saque: %.2f %%.\n", pcentS)
	fmt.Printf("Pontos de Bloqueio: %.2f %%.\n", pcentB)
	fmt.Printf("Pontos de Ataque: %.2f %%.\n", pcentA)

}
