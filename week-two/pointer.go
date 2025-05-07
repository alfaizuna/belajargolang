package main

import "fmt"

func tambahNilai(x *int) {
    *x += 10
}

func main() {
    nilai := 20
    tambahNilai(&nilai)
    fmt.Println(nilai)
}

/*
Pada file pointer.go ini kita mempelajari:

1. Pointer
   - Pointer adalah variabel yang menyimpan alamat memori dari variabel lain
   - Menggunakan operator & untuk mendapatkan alamat memori suatu variabel
   - Menggunakan operator * untuk mengakses/mengubah nilai dari alamat yang ditunjuk

2. Fungsi dengan Parameter Pointer
   - Fungsi tambahNilai() menerima parameter bertipe pointer (*int)
   - Memungkinkan fungsi mengubah nilai asli variabel yang dipass sebagai argument
   - Perubahan nilai di dalam fungsi akan mempengaruhi variabel di luar fungsi

3. Penggunaan Pointer
   - nilai := 20 membuat variabel biasa
   - &nilai mengambil alamat memori dari variabel nilai
   - *x += 10 mengubah nilai yang ditunjuk pointer x
   - Menunjukkan cara kerja pass by reference di Go

4. Manfaat Pointer
   - Menghemat memori karena tidak membuat copy dari data
   - Memungkinkan fungsi mengubah nilai asli parameter
   - Berguna untuk data besar atau ketika perlu mengubah state
*/
