package main

// import "fmt"

func Season(month int) string {
	switch month{
	case 3, 4, 5:
		return "Spring"
	case 6, 7, 8:
		return "Summer"
	case 9, 10, 11:
		return "Autumn"
	case 12, 1, 2:
		return "Winter"
	}
	return "Isn't month"
}

// func main(){
// 	fmt.Println(Season(11))
// }