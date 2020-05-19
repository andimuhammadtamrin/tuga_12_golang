package main

import "fmt"
import "regexp"

func main(){
  var text = "Saya adalah SeOrang KAPITEN"
  //menentukan spasi
  var regex, err = regexp.Compile(`[\t\n\v\f\r ]`)
  //memunculkan semua huruf alpabet dari huruf besar hingga kecil
  var regex1,err1 = regexp.Compile(`[A-Za-z]`)
  if err != nil{
    fmt.Println(err.Error())
  }else if err1 != nil{
    fmt.Println(err.Error())
  }
  var hasil = regex.FindStringIndex(text)
  fmt.Println(hasil)

  var hasil2 = regex1.FindAllString(text,-1)
  fmt.Println(hasil2)
}
