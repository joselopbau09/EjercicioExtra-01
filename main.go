package main 
import "fmt"

var numberSteps int
func numberOfSteps(num int) int {
	if num == 0 {
		return numberSteps
	} else {
		if num%2 == 0 {
			numberSteps++
			numberOfSteps(num/2)
		} else {
			numberSteps++
			numberOfSteps(num - 1)
		}
	}
	return numberSteps
}

func main() {
	
	fmt.Println(numberOfSteps(8))
}