package main 
import "fmt"

var numberSteps int
func numberOfSteps(num int) int {
	if num == 0 {
		return numberSteps
	} else {
		if num%2 == 0 {
			num = num/2
			numberSteps++
			numberOfSteps(num)
		} else {
			num = num -1
			numberSteps++
			numberOfSteps(num)
		}
	}
	return numberSteps
}

func main() {
	
	fmt.Println(numberOfSteps(123))
}