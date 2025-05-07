package main

import (
    "fmt"
    "time"
)

// Fungsi simulasi API 1
func getRateFromSource1(ch chan float64) {
    time.Sleep(2 * time.Second) // Simulasi delay
    ch <- 15700.0
}

// Fungsi simulasi API 2
func getRateFromSource2(ch chan float64) {
    time.Sleep(1 * time.Second)
    ch <- 15800.0
}

func main() {
    ch := make(chan float64)

    go getRateFromSource1(ch)
    go getRateFromSource2(ch)

    // Ambil salah satu rate tercepat
    rate := <-ch

    fmt.Printf("Menggunakan kurs USD → IDR: Rp%.2f\n", rate)
}

/*
Pada file stepByStep.go ini kita mempelajari:

1. Channel
   - Channel adalah cara komunikasi antar goroutine di Go
   - Dibuat dengan make(chan tipe_data)
   - Menggunakan operator <- untuk mengirim/menerima data
   - Channel memungkinkan koordinasi antar goroutine

2. Simulasi API Call
   - Fungsi getRateFromSource1() dan getRateFromSource2() 
   - Mensimulasikan pemanggilan API dengan time.Sleep()
   - Mengirim data rate ke channel setelah delay
   - Menunjukkan penggunaan channel untuk async operations

3. Concurrent API Calls
   - Kedua fungsi dijalankan secara concurrent dengan goroutine
   - Program mengambil hasil tercepat dengan <-ch
   - Mendemonstrasikan pattern untuk parallel API calls
   - Efisien karena tidak perlu menunggu semua API selesai

4. Manfaat Pattern Ini
   - Mempercepat response time aplikasi
   - Fault tolerance jika salah satu API gagal
   - Mudah ditambahkan sumber data baru
   - Pattern umum untuk distributed systems
*/
