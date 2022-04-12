// balanço de parenteses i
package main
import(
  "fmt"
  "strings"
)
func main() {
  var input string
  var saida string
  for{
    err := inputDados(&input)
    if err != nil{
      return
    }
    if input == ""{
      return
    }
    inputsDivididas := strings.Split(input, "\n")
    corrige(inputsDivididas, &saida)
    printa(saida)
    saida = ""
    input = ""
  }
}
func inputDados(input *string) (error){
  var charSmol [1000]byte
  for i,_ := range charSmol{
    _, err := fmt.Scanf("%c", &charSmol[i])
    if err != nil{
      return fmt.Errorf("ERRO")
    }
    if charSmol[0] == '\n'{
      return nil
    }
    if charSmol[i] == '\n'{
      break
    }
  }
  for l,_ := range charSmol{
    *input += string(charSmol[l])
  }
  return nil
}
func printa(saida string){
  fmt.Printf("%s",saida)
}
func corrige(sliceinputs []string, saida *string){
  var temp string
  for i,_ := range sliceinputs{
    if(len(sliceinputs[i]) >= 1){
      temp = verifica(sliceinputs[i])
      if temp != "0"{
        *saida += temp
      }
    }
  }
}
func verifica(input string) string{
  var aberturas []byte
  top := 0
  achou := false
  temParenteses := false
  for i := 0; i < len(input); i++{
    if input[i] == '(' {
      temParenteses = true
      aberturas = append(aberturas, input[i])
    }else if input[i] == ')'{
      temParenteses = true
      achou = false
      for j,_ := range aberturas{
        if aberturas[j] == '('{
          achou = true
          break
        }
      }
      if !achou{
        return "incorrect\n"
      }else{
        temParenteses = true
        if aberturas[top] == '('{
          aberturas[top] = ' '
          top++
        }else if (top+1 >= len(aberturas)){
          aberturas[top] = ' '
        }else if aberturas[top] != '('{
          if aberturas[top+1] == '('{
            top++
            break
          }
          return "incorrect\n"
        }
      }
    }
  }
  for i,_ := range aberturas{
    if aberturas[i] == '('{
      return "incorrect\n"
    }
  }
  if temParenteses{
    return "correct\n"
  }
  return "0"
}
