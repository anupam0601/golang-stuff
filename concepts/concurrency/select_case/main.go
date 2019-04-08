package select_case

import "fmt"

func main() {
	sendCh := make(chan<- int)
	v := 5
	select {
	case sendCh <- v:
		fmt.Println("Sending---", v)
	}
}
