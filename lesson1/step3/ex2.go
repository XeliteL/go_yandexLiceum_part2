package step3

func Receive(ch chan int) int {
	return <-ch
}
