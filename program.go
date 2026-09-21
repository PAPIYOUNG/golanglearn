package main

import (
	"fmt"
)
func main(){
	fmt.Println("Hello PAPI");
	fmt.Println("Narak")

	//variables
	//1.Manual declaration type
	var name string="pongnapa"
	fmt.Println("My name is",name)

	var age int=32
	fmt.Println("Age :",age)

	var score float32 =98.9
	fmt.Println("Score :",score)

	var isPass bool =true
	fmt.Println("Pass Exam :",isPass)

	//2.Type interface
	name1 := "pongnapa"
	age1 := 32
	score1 :=99.99
	isPass1 := true
	// หรือ 
	// var name1 = "pongnapa"
	// var age1 = 32
	// var score1 =99.99
	//isPass1 := true
	fmt.Println(name1)
	fmt.Println(age1)
	fmt.Println(score1)
	fmt.Println(isPass1)

	var numb1,numb2 =10,3
	// หรือ  numb1,numb2 :=10,3
	fmt.Println(numb1,numb2)


	//constant value
	const name2 string ="pongnapa"
	//name2 = "pongpong"
	fmt.Println(name2)

	//type of variable
	fmt.Printf("Value: %v, Type: %T\n", name, name)
    fmt.Printf("Value: %v, Type: %T\n", age, age)
    fmt.Printf("Boolean: %t\n", isPass)


	//operation
	var number1 int =10
	var number2 int =3
	fmt.Println("Summarize =",number1+number2)
	fmt.Println("Separate =",number1-number2)
	fmt.Println("Multiple =",number1*number2)
	fmt.Println("divide =",number1/number2)
	fmt.Println("Modulus =",number1%number2)

	fmt.Println("is Equal :",number1 == number2)
	fmt.Println("is more :",number1 > number2)
	fmt.Println("is less :",number1 < number2)
	

	//Scanf
	var studentName string //ถ้าไม่ assign data จะเป็น zero value คือ empty string
	fmt.Print("inform student name =")
	//fmt.Scanf("%s",&studentName)
	fmt.Println("Hello",studentName)
	// &studentName หมายถึง ขอ Address ของตัวแปร เอาที่อยู่หน่วยความจำของตัวเเปรนี้ เพื่อให้ Scanf ต้อง แก้ค่าของตัวแปรเดิม

	var soccerScore int
	fmt.Print("Soccer score is :")
	//fmt.Scanf("%d",&soccerScore)
	fmt.Println("Answer Soccer score is",soccerScore)

	var like bool
	fmt.Print("Do you like soccer ?")
	//fmt.Scanf("%t",&like)
	fmt.Println("Soccer like :",like)

	//if/else
	var score2 int
	fmt.Print("inform student score :")
	//fmt.Scanf("%d",&score2)

	if score2 >50 {
		fmt.Println("Pass")
	}else {
		fmt.Println("Not Pass")
	}
	
	var number3 int
	fmt.Print("inform number to check :")
	//fmt.Scanf("%d",&number3)

	if number3%2 == 0 {
		fmt.Println("it's event number")
	}else {
		fmt.Println("it's odd number")
	}

	var number4 int
	fmt.Print("inform number choose service :")
	//fmt.Scanf("%d",&number4)

	// if number4 == 1 {
	// 	fmt.Println("open bank account")
	// }else if number4 == 2{
	// 	fmt.Println("withdraw cash")
	// }else {
	// 	fmt.Println("wrong number")
	// }

	//switch/case
switch number4 {
case 1:
	fmt.Println("Learn Eng")
case 2 :
	fmt.Println("Learn Go")
default:
	fmt.Println("wrong number")
}


//array
//1.Array ระบุจำนวนสมาชิก
var numberGroup [3]int
numberGroup[0] = 100
numberGroup[1] = 200
numberGroup[2] = 300
fmt.Println(numberGroup)

var numberGroup1 [3]int =[3]int{100,200,300}
fmt.Println(numberGroup1)

numberGroup2 := [3]int{100,200,300}
fmt.Println(numberGroup2)

fmt.Println(numberGroup2[0])
fmt.Println(numberGroup2[1])
fmt.Println(numberGroup2[2])


size := len(numberGroup2)
fmt.Println(size)

//2.Array เเบบไม่กำหนดจำนวนสมาชิก
var numberTeam =[...]int{100,200,300,400}
fmt.Println(numberTeam)

numberTeam1 := [...]int{100,200,300,450,5304,34}
fmt.Println(numberTeam1)

size1 := len(numberTeam1)
fmt.Println(size1)

//Slice
numberSlice :=[]int{100,200,300,400,500}
fmt.Println(numberSlice)

fmt.Println(numberSlice[4])
fmt.Println(numberSlice[3])
fmt.Println(numberSlice[:])
fmt.Println(numberSlice[1:])
fmt.Println(numberSlice[:3])
fmt.Println(numberSlice[3:5])


//เปลี่ยนค่าได้
numberSlice[0]=1000
numberSlice[2]=999 
fmt.Println(numberSlice)

//เพิ่มสมาชิก
numberSlice = append(numberSlice,888)
numberSlice = append(numberSlice,777)
fmt.Println(numberSlice)

for _, slicee := range numberSlice {
    fmt.Println(slicee)
}

servers := []string{"web", "api", "database"}
for _, server := range servers {
    fmt.Println(server)
}

//map
country :=[2]string{"thai","japan"}
fmt.Println(country[1])

countryMap :=map[string]string{"TH":"thai","JP":"japan"}
fmt.Println(countryMap["TH"])
 
//เช็คค่าว่ามี value ไหม
value , check := countryMap["JP"]

if check {
	fmt.Println("Found data",value)
}else {
	fmt.Println("Not found data")
}

//loop
for count := 1;count <=10;count++{
	fmt.Println(count)
	if count == 5{
		break
	}
}
fmt.Println("end running")

for count := 1;count <=10;count++{
	fmt.Println(count)
	if count == 5{
		continue
	}
}
fmt.Println("end running")


number7 := []int{10,20,30,40,50}
for i := 0;i<len(number7);i++{

	fmt.Println("index=",i,"value =",number7[i])
}

for index,value := range number7 {
	fmt.Println("index=",index,"value =",value)

}

for _,value := range number7 {
	fmt.Println("value =",value)

}

language := map[string]string{"TH":"Thai","EN":"English","JP":"Japan"}
for key,value := range language {
	fmt.Println("key=",key,"value=",value)

}

for _,value := range language {
	fmt.Println("value=",value)

}


//function
total(10,20)

delivery := getDelivery()
fmt.Println(delivery)

myCart := getTotalCart(500,200)
fmt.Println(myCart)

result,status := summation(100,200)
fmt.Println(result,status)


//variadic function
result2 := summation2(10,20,30,40,50)
fmt.Println("Total result =",result2)


//struct
product1 := Product{name:"Pen",price: 50,category: "tool",discount: 10}
fmt.Println(product1)
fmt.Println(product1.name)
fmt.Println(product1.price)
//change data in struct
product1.category = "device"
fmt.Println(product1)

}

//function
func total(number1 ,number2 int){
	fmt.Println(number1+number2)
}

//ใส่ Type ต่อท้าย สำหรับค่าที่ return
func getDelivery() int{
 return  50
}

func getTotalCart(num1,num2 int)int{
	total := num1+num2
	return total
}

//return multiple value
func summation(num1 , num2 int)(int,string){
	total := num1+num2
	status :=""
	if total%2 == 0{
		status = "event number"
	}else{
		status ="odd number"
	}

	return total,status
}


//variadic function
func summation2(num ...int)int{
	total := 0
	for _,value := range num {
		total +=value
	}
	return total
}


//struct
type Product struct {
	name string
	price float64
	category string
	discount int
}
