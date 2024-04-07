package main

import "fmt"

func main(){
	fmt.Println("=============Welcome to loops===============")
	days := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
	fmt.Println("-----------First way of loop----------")
	for i:=0; i < len(days); i++{
		fmt.Println(days[i])
	}
	fmt.Println("------------2nd way of for loop----------")
	for i := range days {
		fmt.Println(days[i])
	}
	fmt.Println("--------3rd way-----------------")
	for key, day := range days {
		fmt.Printf("The value is %v and the key is %v \n", key, day)
	}
	fmt.Println("-----While loop using for loop---------")
	roughValue := 1
	for roughValue < 10 {
		roughValue++
		if roughValue == 5 {
			roughValue++
			continue
		}
		if roughValue == 7 {
			break;
		}
		if roughValue == 9 {
			goto loc
		}
		fmt.Println("the value of roughValue is ", roughValue)
	}

	loc:
		fmt.Println("This is value of 9 from loc or label")
}