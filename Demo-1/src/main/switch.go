package main

import "fmt"
import "time"

func main() {
	i := 2
	fmt.Println("write ",i," as")
	switch i{
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
	}

	switch time.Now().Weekday(){
	case time.Saturday,time.Sunday:
		fmt.Println("It is weekend.")
	default:
		fmt.Println("It is weekday")
	}
	t := time.Now()
	switch{
	case t.Hour() < 12:
		fmt.Println("It's before noon.")
	default:
		fmt.Println("It's after noon.")
	}

	whatAmI := func(i interface{}){
		switch t := i.(type){
		case bool:
			fmt.Println("Iam a bool.")
		case int:
			fmt.Println("Iam an int.")
		default:
			fmt.Printf("Dont know type %T\n",t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}