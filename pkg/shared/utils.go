package shared

import "fmt"

const RunShared string = "I am Run from shared lib - v0.1.0"

/*Run ---*/
func Run() (int, error) {
	return fmt.Println(RunShared)
}
