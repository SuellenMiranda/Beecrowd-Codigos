//2764 - Entrada e Saida de Data-1
package main

import(
  "fmt"
  "strings"
)

func main(){
  var input string
  inputData(&input)
  list := divideData(input)
  printa(list)
}
func inputData(entr *string){
  fmt.Scanf("%s", entr)
}
func divideData(entr string)[]string{
  dividido := strings.Split(entr,"/")
  return dividido
}
func printa(entr []string){
  fmt.Printf("%s/%s/%s\n", entr[1], entr[0], entr[2])
  fmt.Printf("%s/%s/%s\n", entr[2], entr[1], entr[0])
  fmt.Printf("%s-%s-%s\n", entr[0], entr[1], entr[2])
}
