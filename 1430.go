//1430 - Composição de Jingles
package main

import(
  "fmt"
)

const MAX int = 200
const MIN int = 3

var mapaNotas = map[byte]float64{
    'W': 1.0,
    'H': 1.0/2,
    'Q': 1.0/4,
    'E': 1.0/8,
    'S': 1.0/16,
    'T': 1.0/32,
    'X': 1.0/64,
  }

func main(){

  for{
    input,err := inputUser()
    if err != nil{
      return
    }
    
    lista := brokeString(input)
    
    qnt := 0
    start := false
    for _,uni := range lista{
      qnt = qnt + verify(uni,mapaNotas)
      if uni != "" && !start{
        start = true
      }else if uni == "" && start{
        fmt.Printf("%d\n",qnt)
        qnt = 0
      }
    }
  }
}

func brokeString(input string)[]string{
  var strings []string
  var str string
  for i,_ := range input{
    if input[i] != '/'{
      str += string(input[i])
    }else if input[i] == '/'{
      strings = append(strings, str)
      str = ""
    }
  }

  strings = append(strings, "")

  return strings
}

func inputUser() (string, error){
  var input string
  

  fmt.Scanf("%s", &input)

  if input[len(input)-1] == '*'{
    return " ", fmt.Errorf("ERRO")
  }else if len(input) > MAX || len(input) < MIN{
    return " ", fmt.Errorf("ERRO")
  }

  return input, nil
}

func verify(input string, mapa map[byte]float64)int{
  sum := 0.0
  for i,_ := range input{
    sum += mapa[input[i]]
  }
  
  if modulo(sum-1) < 0.0001{
    return 1
  }

  return 0
}

func modulo(num float64)float64{
  if num < 0{
    num = num * -1.0;
  }

  return num
}
