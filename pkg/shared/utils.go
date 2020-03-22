package shared

import "fmt"

/*RunShared --*/
const RunShared string = "I am Run from shared lib - v0.1.1"

/*Run ---*/
func Run() (int, error) {
	return fmt.Println(RunShared)
}
