package main

import (
	"fmt"
	"time"
)

func cetak(text string){
	for i:=0; i<3; i++ {
		fmt.Println(text)
	}
}

func main(){
	go cetak("Halo")
	cetak("Dunia")

	time.Sleep(500 * time.Millisecond)
}

/*
Pada file goroutine.go ini kita mempelajari:

1. Goroutine
   - Goroutine adalah fitur Go untuk menjalankan fungsi secara concurrent/bersamaan
   - Menggunakan keyword 'go' sebelum memanggil fungsi untuk membuatnya berjalan sebagai goroutine
   - Contoh: go cetak("Halo") akan berjalan bersamaan dengan cetak("Dunia")

2. Fungsi cetak()
   - Fungsi sederhana yang melakukan perulangan untuk mencetak text 3 kali
   - Digunakan untuk mendemonstrasikan concurrent execution

3. time.Sleep()
   - Digunakan untuk menunda eksekusi program selama waktu tertentu
   - Dalam contoh ini digunakan agar program tidak langsung berakhir
   - Memberikan waktu untuk goroutine menyelesaikan eksekusinya
   - Menggunakan durasi 500 millisecond

4. Concurrent vs Sequential
   - "Halo" dijalankan secara concurrent (dengan goroutine)
   - "Dunia" dijalankan secara sequential (tanpa goroutine)
   - Menunjukkan perbedaan antara eksekusi concurrent dan sequential
*/

