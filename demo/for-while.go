package demo
import "fmt"

func ForWhile() {
	sum := 0
	for ;sum < 1000;  {
		sum += 100
    fmt.Printf("sum=%d\n",sum)
	}
}
