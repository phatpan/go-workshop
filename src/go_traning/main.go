package main

import (
  "fmt"
  "encoding/json"
  "net/http"
  "io/ioutil"

  // "math/rand"
  // "time"

)

func main(){
  // var msg string = "hello world"
  // var f float64= 10.1234
  // var n int = 10
  // var gopher string= "Gopher"
  // var t bool=true
  // var a bool= false
  //
  // fmt.Println(msg)
  // fmt.Printf("%.2f\n",f)
  // fmt.Printf("%d\n", n)
  // fmt.Printf("%s\n", gopher)
  // fmt.Println(t)
  // fmt.Println(a)

  // msg := "hello world"
  // f := 10.1234
  // n := 10
  // gopher := "Gopher"
  // t := true
  // a := false
  //
  // fmt.Println(msg)
  // fmt.Printf("%.2f\n",f)
  // fmt.Printf("%d\n", n)
  // fmt.Printf("%s\n", gopher)
  // fmt.Println(t)
  // fmt.Println(a)
  //
  // f1, f2 := 10.0, 20.0
  // fmt.Printf("%f", f1 + f2)
  // fmt.Printf("%f\n", f1 - f2)
  // fmt.Println(f1)
  // fmt.Println(fmt.Sprintf("http://%s:%d", "localhost"))
  // var name string
  // fmt.Scanln(&name)
  // fmt.Println("Hello " + name)

  //
  // res, _ := http.Get("http://www.google.com/finance/info?q=INDEXBKK:S")
  // content, _ := ioutil.ReadAll(res.Body)
  // res.Body.Close()
  // fmt.Printf("%s\n", content)


 // num := 3
 // name := "phatpan"
 // if num := 2; num == 2 {
 //   fmt.Println("one")
 // }



 // if name == "phatpan" && num ==1 {
 //   fmt.Println("Equal")
 // }else {
 //   fmt.Println("Not Equal")
 // }


  // var num int
  // fmt.Scanln(&num)
  // if num % 15 == 0{
  //   fmt.Println("FizzBuzz")
  // }else if num % 5 == 0{
  //   fmt.Println("Buzz")
  // }else if num % 3 == 0{
  //   fmt.Println("Fizz")
  // }else{
  //   fmt.Println(num)
  // }

  // num := 10
  // switch num {
  // case 5:
  //   fmt.Println("Five")
  // case 10:
  //   fmt.Println("Ten")
  // default:
  //   fmt.Println(num)
  // }

  // num := 10
  // switch  {
  // case num % 15 ==0:
  //   fmt.Println("FizzBuzz")
  // case num % 3 == 0:
  //   fmt.Println("Fizz")
  // case num % 5 ==0:
  //   fmt.Println("Buzz")
  //   default:
  //     fmt.Println(num)
  // }


  // for i:= 0; i <= 10; i++ {
  //
  // }


  // rand.Seed(time.Now().UnixNano())
  // guess := rand.Intn(100)+1
  // GameLoop(guess)


  // num1, _ := GetTwoNumber()
  // a, b := 10,20
  // a, b =Swap(a, b
  // )
  // fmt.Println(num1)
  // fmt.Println(a, b)



  // var data interface{}
  // reqJSON := "{\"name\": \"Phatpan\"}"
  // err := json.Unmarshal([]byte(reqJSON), &data)
  // if err != nil {
  //   fmt.Println("Error : ", err)
  //   return
  // }
  // fmt.Println(data)


  // var nums [5]int = [5]int {1,2,3,4,5}
  // var snums []int
  // fmt.Println(len(nums))
  // fmt.Println(nums[0])
  //
  // for i, n:=  range nums {
  //   fmt.Println(i, ":", n)
  // }
  // var nums2 [3]int = nums
  // snums = nums[0:2]
  // fmt.Println(snums)


  // var m map[string]string = map[string]string{
  //   "name": "phatpan Azumo",
  //   "nickname": "kea",
  // }
  // fmt.Println(m["name"], ":", m["nickname"]);
  //




  // reqJSON := `["Phatpan", "Azumo"]`
  // var names []string
  // json.Unmarshal([]byte(reqJSON), &names)
  // fmt.Println(names)


  // var m map[string]string
  // reqJSON := `{"name": "phatpancharaphan ananpreechakun"}`
  // json.Unmarshal([]byte(reqJSON), &m)
  // fmt.Println(m["name"])
  //
  // for k, v := range m {
  //   fmt.Println(k, v)
  // }


  // var data []map[string]string
  // reqJSON := `[{"name": "phatpan", "name":"Azumo"}]`
  // json.Unmarshal([]byte(reqJSON), &data)
  // for _, d := range data{
  //   fmt.Println(d["name"])
  // }


  var data []map[string]string
    res, _ := http.Get("http://www.google.com/finance/info?q=INDEXBKK:SET")
    reqJSON, _ := ioutil.ReadAll(res.Body)
    res.Body.Close()
    json.Unmarshal(reqJSON[3:], &data)
    fmt.Printf("Index of", data[0]["l_fix"])
}

func Swap(a,b int)(int, int){
  return b,a
}

func GetNumber() int{
  var num int
  fmt.Scanln(&num)
  return num
}

func GetTwoNumber()(int, int){
  return 2,3
}

func GameLoop(guess int){
   var num int
   END:
   for {
     fmt.Print("Enter number: ")
     num = GetNumber()
     switch  {
     case num == guess:
       break END
     case num > guess:
       fmt.Println("Greater than")
     case num < guess:
        fmt.Println("less than")
     }
   }
}
