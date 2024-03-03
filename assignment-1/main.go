package main

import (
	"fmt"
	"os"
	"strconv"
)

type Biodata struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	data := []Biodata{
		{
			Nama:      "Agung",
			Alamat:    "Jakarta",
			Pekerjaan: "Mobile Developer",
			Alasan:    "Alasan Agung",
		},
		{
			Nama:      "Aditya",
			Alamat:    "Semarang",
			Pekerjaan: "Frontend Engineer",
			Alasan:    "Alasan Aditya",
		},
		{
			Nama:      "Dimas",
			Alamat:    "Sleman",
			Pekerjaan: "Backend Engineer",
			Alasan:    "Alasan Dimas",
		},
		{
			Nama:      "Kadek",
			Alamat:    "Bali",
			Pekerjaan: "DevOps",
			Alasan:    "Alasan Kadek",
		},
		{
			Nama:      "Bagus",
			Alamat:    "Klaten",
			Pekerjaan: "Data Engineer",
			Alasan:    "Alasan Bagus",
		},
	}

	if len(os.Args) > 1 {
		arg, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Format salah")
		} else if arg == 0 || arg > len(data) {
			fmt.Println("Data tidak ditemukan")
		} else {
			arg--
			fmt.Println("Nama\t\t:", data[arg].Nama)
			fmt.Println("Alamat\t\t:", data[arg].Alamat)
			fmt.Println("Pekerjaan\t:", data[arg].Pekerjaan)
			fmt.Println("Alasan\t\t:", data[arg].Alasan)
		}
	} else {
		fmt.Println("Format salah")
	}
}
