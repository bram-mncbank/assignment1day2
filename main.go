package main

import (
	"fmt"
	"os"
	"strconv"
)

type Siswa struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	var kelas = []*Siswa{}

	var a, _ = strconv.Atoi(os.Args[1])

	var addSiswa = func(s *Siswa) {
		kelas = append(kelas, s)
	}

	var getSiswa = func(s int, kls []*Siswa) {
		if s >= len(kls) {
			fmt.Println("Sorry data not available")
		} else {
			fmt.Printf("Nama     : %v\n", kls[s].nama)
			fmt.Printf("Pekerjaan: %v\n", kls[s].pekerjaan)
			fmt.Printf("Alamat   : %v\n", kls[s].alamat)
			fmt.Printf("Alasan   : %v\n", kls[s].alasan)
		}
	}

	addSiswa(&Siswa{nama: "bram", alamat: "bandung", pekerjaan: "programmer", alasan: "cool"})
	addSiswa(&Siswa{nama: "sigit", alamat: "jakarta", pekerjaan: "hacker", alasan: "nice"})
	addSiswa(&Siswa{nama: "frisky", alamat: "bogor", pekerjaan: "programmer", alasan: "gabut"})
	getSiswa(a, kelas)

}
