//Media ponderada

package main
 
import (
    "fmt"
)
 
func main() {
  
  var A, B float64
  peso1 := 3.5
  peso2 := 7.5
 
  fmt.Scanf("%f\n", &A) 
  fmt.Scanf("%f\n", &B)
  
  MEDIA := (A*peso1+B*peso2) / (peso1 + peso2)
  
  fmt.Printf("MEDIA = %.5f\n",MEDIA)

}
