package mathutil

import "testing"

func TestAdd(t *testing.T) {
    got := Add(2, 5)
    want := 7

    if got != want {
        t.Errorf("Expected %d, got %d", want, got)
    }
}

/*
Pada file mathutil_test.go dan mathutil.go ini kita mempelajari:

1. Unit Testing di Go
   - Menggunakan package "testing" untuk membuat unit test
   - File test harus berakhiran _test.go
   - Fungsi test harus diawali dengan "Test"
   - Menerima parameter t *testing.T untuk testing utilities

2. Struktur Test
   - TestAdd() menguji fungsi Add() dari mathutil.go
   - Menggunakan pattern got vs want untuk membandingkan hasil
   - t.Errorf() untuk melaporkan jika test gagal
   - Menunjukkan cara menulis test yang baik dan jelas

3. Fungsi yang Diuji (mathutil.go)
   - Package mathutil berisi fungsi-fungsi matematika
   - Fungsi Add() melakukan penjumlahan dua bilangan
   - Implementasi sederhana untuk mendemonstrasikan testing
   - Menunjukkan separation of concerns antara kode dan test

4. Best Practices
   - Test terpisah dalam file sendiri
   - Nama test menjelaskan apa yang diuji
   - Error message informatif dan jelas
   - Test fokus pada satu fungsionalitas
*/
