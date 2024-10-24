package main

import (
    "fmt"
)

type Peserta struct {
    nama       string
    soalSelesai int
    totalWaktu  int
}

func hitungSkor(nama string, waktu []int) Peserta {
    const maxWaktu = 301
    soalSelesai := 0
    totalWaktu := 0

    for _, w := range waktu {
        if w < maxWaktu {
            soalSelesai++
            totalWaktu += w
        }
    }

    return Peserta{nama, soalSelesai, totalWaktu}
}

func main() {
    var n int
    fmt.Print("Masukkan jumlah peserta: ")
    fmt.Scan(&n)

    pesertaList := make([]Peserta, n)

    for i := 0; i < n; i++ {
        var nama string
        var waktu [8]int

        fmt.Print("Masukkan nama peserta: ")
        fmt.Scan(&nama)

        fmt.Print("Masukkan waktu penyelesaian untuk 8 soal (dalam menit): ")
        for j := 0; j < 8; j++ {
            fmt.Scan(&waktu[j])
        }

        pesertaList[i] = hitungSkor(nama, waktu[:])
    }

    pemenang := pesertaList[0]
    for _, peserta := range pesertaList {
        if peserta.soalSelesai > pemenang.soalSelesai || 
           (peserta.soalSelesai == pemenang.soalSelesai && peserta.totalWaktu < pemenang.totalWaktu) {
            pemenang = peserta
        }
    }

    fmt.Printf("Pemenang: %s, Soal Selesai: %d, Total Waktu: %d menit\n", pemenang.nama, pemenang.soalSelesai, pemenang.totalWaktu)
}
