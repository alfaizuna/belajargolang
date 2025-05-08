package main

import (
	"fmt"
	"time"
)

func getRateFromSlowSource(ch chan float64) {
	time.Sleep(2 * time.Second)
	ch <- 15700.0
}

func main() {
	rateCh := make(chan float64)

	go getRateFromSlowSource(rateCh)

	select {
	case rate := <-rateCh:
		fmt.Printf("Kurs diterima: Rp%.2f\n", rate)
	case <-time.After(1500 * time.Millisecond):
		fmt.Println("⏰ Timeout: gagal ambil kurs dari sumber")
	}
}