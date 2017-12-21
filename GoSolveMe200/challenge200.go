
package main 

import (
		
		"fmt"			
)

type Integer struct{

	number int
}

func main(){

var input Integer
var store  = [40]int {82, 110, 21, 121, 78, 90, 97, 78, 21, 84, 121, 95, 102, 115, 121, 101, 74, 90, 75, 100, 121, 82, 110, 21, 121, 100, 67, 25, 25, 67, 116, 121, 114, 78, 21, 121, 80, 90, 21, 81}
var check Integer
check.number = 38
fmt.Println("Enter Input:\n")
_, err := fmt.Scanf("%d", &input.number)

if err != nil {
            fmt.Println(err)
         } else {

if input.number < 9999999 {
	badinput()

}else if input.number > 99999999 {
	badinput()

}else {

	var ret int = calc(input,check)
	if ret != 1 {
		badinput()
	}else{

		printflag(input,store)
	     }
	}
      }
}

func calc(input Integer, check Integer) int{


	var ret int = 0
	var temp Integer
	var tempInt Integer
	var finalTemp Integer
	
	tempInt.number = 1000000
	finalTemp.number = Div(input,tempInt)

	tempInt.number = 10
	finalTemp.number = Modulus(finalTemp,tempInt)
	
	tempInt.number = 10
	finalTemp.number = Mul(finalTemp,tempInt)

	temp.number = Modulus(input,tempInt)
	
	finalTemp.number = Add(finalTemp ,temp)

	if(finalTemp.number == check.number) {
		ret = 1
	}else {

		ret = 0
	}

	return ret

	
}



func printflag(input Integer, store [40]int) {

	var flag string
	var temp Integer
	var tempInt Integer
	var finalTemp Integer
	
	tempInt.number = 1000000
	finalTemp.number = Div(input,tempInt)

	tempInt.number = 10
	finalTemp.number = Modulus(finalTemp,tempInt)
	
	tempInt.number = 10
	finalTemp.number = Mul(finalTemp,tempInt)

	temp.number = Modulus(input,tempInt)
	
	finalTemp.number = Add(finalTemp ,temp)

	fmt.Print("\nCongratulations!!\n\nFlag: inctf{")
	for a := 0; a < 40; a++ {

		temp.number = store[a] ^ finalTemp.number
		character := string(temp.number)
		flag = flag + character
	}

	fmt.Print(flag)

	fmt.Print("}\n\n")

}

 func badinput() {

 	fmt.Println("\n\nWrong Input.")
 }


func Modulus(a Integer, b Integer ) int {

	
	return (a.number%b.number)
}

func Add(a Integer, b Integer) int{

	return (a.number+b.number)

}

func Div(a Integer, b Integer) int{
	return a.number/b.number
}

func Mul(a Integer, b Integer) int{

	return (a.number * b.number)
}
