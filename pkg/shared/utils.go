package shared

import "fmt"

const RunShared string = "I am Run from shared lib"

/*Run ---*/
func Run() (int, error) {
	return fmt.Println(RunShared)
}
