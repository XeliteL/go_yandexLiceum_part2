package step3

func Send(ch chan int, num int) {
	ch <- num
}
