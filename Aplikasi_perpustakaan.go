package main

import (
	"fmt"
	"os"
	"os/exec"
)	

const NMax int = 100

type buku struct {
	namaBuku string
	lokasi string
	jumlah float64
}

type Array struct {
	bookList[NMax] buku 
	total int
}

var n , idx int

func main(){
	var isi Array
	var cari string

	var menu int
	cmd:= exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	for menu != 7 {
		fmt.Println("------------------------------------------------------")
		fmt.Println("|====================================================|")
		fmt.Println("|                   PERPUSTAKAAN                     |")
		fmt.Println("|====================================================|")
		fmt.Println("|              1. Lihat Daftar Buku                  |")
		fmt.Println("|              2. Tambah Buku                        |")
		fmt.Println("|              3. Edit data buku                     |")
		fmt.Println("|              4. cari buku                          |")
		fmt.Println("|              5. Delete                             |")
		fmt.Println("|              6. Urutan Buku Sesuai Huruf Abjad     |")
		fmt.Println("|              7. Keluar                             |")
		fmt.Println("------------------------------------------------------")
		fmt.Println(" ")
		fmt.Print("Pilih Menu : ")
		fmt.Scan(&menu)
		
		for menu != 1 && menu != 2 && menu != 3 && menu != 4 && menu != 5 && menu != 0 && menu != 6 && menu != 7 {
			fmt.Print("Input Salah, Pilih Menu : ")
			fmt.Scan(&menu)
		}	

		if menu == 1 {

			tampil(&isi)

		} else if menu == 2 {

			insert(&isi,n)
			
		} else if menu == 3 {

			edit(&isi)

		} else if menu == 4 {

			fmt.Print("(1. Sequensial) (2. Binary) : ")
			fmt.Scan(&n)
			if n == 1 {
				fmt.Print("Masukan buku yang ingin dicari : ")
				fmt.Scan(&cari)
				findSeq(isi,cari)
			} else if n == 2 {
				fmt.Print("Masukan buku yang ingin dicari : ")
				fmt.Scan(&cari)
				findBiner(isi,cari)
			}

		} else if menu == 5 {
			delete(&isi)
			
		} else if menu == 6 {
			fmt.Print("(1. Insertion) (2. Selection) : ")
			fmt.Scan(&n)
			fmt.Println(" ")
			if n == 1{
				fmt.Print("(1. Ascending) (2. Descending) : ")
				fmt.Scan(&n)
				if n == 1 {
					sortingInsertAs(&isi)
				} else if n == 2 {
					sortingInsertDes(&isi)
				}
			} else if n == 2 {
				fmt.Print("(1. Ascending) (2. Descending) : ")
				fmt.Scan(&n)
				if n == 1 {
					sortingSelectAs(&isi)
				} else if n == 2 {
					sortingSelectDec(&isi)
				}
			}

		} else if menu == 7 {
			fmt.Println("---------------------- KELUAR ------------------------")
		}	
	}
}

func tampil(data *Array){
	var i int

	i = 0
	idx = 1
	for i < data.total && data.bookList[i].jumlah != 0 {
	fmt.Println("-------------------------")
	fmt.Println(idx,".Nama Buku   : ",(*data).bookList[i].namaBuku)
	fmt.Println("   Lokasi Buku : ",(*data).bookList[i].lokasi)
	fmt.Println("   Jumlah Buku : ",(*data).bookList[i].jumlah)
	fmt.Println(" ")

	i++
	idx++
	}
}

func insert(data *Array, n int){
	var i int
	fmt.Print("Masukan Jumlah Buku Yang Ingin Di Input : ")
	fmt.Scan(&n)
	i = 0
	idx = 1
	for i < n  {
		fmt.Println("-----------------------------")
		fmt.Print(idx,". Nama Buku   : ")
		fmt.Scan(&data.bookList[i].namaBuku)
		fmt.Print("   Lokasi Buku : ")
		fmt.Scan(&data.bookList[i].lokasi)
		fmt.Print("   Jumlah Buku : ")
		fmt.Scan(&data.bookList[i].jumlah)
		
		data.total++
		i++
		idx++
		
	}
}

func edit(data *Array){
	var nama , tempat string
	var sum float64
	fmt.Print("Buku Nomor Berapa Yang ingin Di Edit : ")
	fmt.Scan(&idx)

	fmt.Println("-----------------------------")
	fmt.Print(idx,". Nama Buku   : ")
	fmt.Scan(&nama)
	fmt.Print("   Lokasi Buku : ")   
	fmt.Scan(&tempat)
	fmt.Print("   Jumlah Buku : ")
	fmt.Scan(&sum)

	(*data).bookList[idx-1].namaBuku = nama
	(*data).bookList[idx-1].lokasi = tempat
	(*data).bookList[idx-1].jumlah = sum
}

func findSeq(data Array, name string) {
	i := 0
	indx := 1
	for (i < data.total)&&(data.bookList[i].namaBuku != name) {
		i++
		indx++
	}
	if (data.bookList[i].namaBuku == name) {
		fmt.Println("-------------------------")
		fmt.Println(indx,".Nama Buku   : ",data.bookList[i].namaBuku)
		fmt.Println("   Lokasi Buku : ",data.bookList[i].lokasi)
		fmt.Println("   Jumlah Buku : ",data.bookList[i].jumlah)
	} else {
		fmt.Println("Tidak Ada Buku")
	}
}

func findBiner(data Array, name string){
	awal := 0
	akhir := data.total -1 
	
	found := false
	for akhir >= awal && found != true {
		tengah := (awal+akhir)/2
		if data.bookList[tengah].namaBuku < name {
			awal = tengah + 1
		} else if data.bookList[tengah].namaBuku > name {
			akhir = tengah
		} else {
			found = true
			fmt.Println("-------------------------")
			fmt.Println(data.total,".Nama Buku   : ",data.bookList[tengah].namaBuku)
			fmt.Println("   Lokasi Buku : ",data.bookList[tengah].lokasi)
			fmt.Println("   Jumlah Buku : ",data.bookList[tengah].jumlah)
		}
	}
}

func delete(data *Array){
	var k , j , i int
	fmt.Print("Masukan pilihan anda yang ingin di hapus : ")
	fmt.Scan(&j)
	for i = 1 ; i < idx ; i++ {
		if j == (i + 1) {
			k = i
		}

	}
	if j == k + 1 {
		for i = k ; i < idx ; i++ {
			(*data).bookList[i] = (*data).bookList[i+1]
		}
		idx = idx - 1
	}
	i = 0
	indx := 1
	for i < idx-1 {
		fmt.Println("-----------------------------")
		fmt.Println(indx,".Nama Buku   : ",data.bookList[i].namaBuku)
		fmt.Println("   Lokasi Buku : ",data.bookList[i].lokasi)
		fmt.Println("   Jumlah Buku : ",data.bookList[i].jumlah)
		i = i + 1
		indx = indx + 1
	}
	
}

func sortingInsertAs(data *Array ) {
	var j , i , indx int
	var temp, tempat string
	var sum float64
	
	i = 0
	for i < idx - 1  {
		temp = data.bookList[i].namaBuku
		tempat = data.bookList[i].lokasi
		sum = data.bookList[i].jumlah
		j = i - 1
		for j >= 0 && data.bookList[j].namaBuku > temp {
			data.bookList[j+1].namaBuku, data.bookList[j+1].lokasi, data.bookList[j+1].jumlah = data.bookList[j].namaBuku, data.bookList[j].lokasi, data.bookList[j].jumlah
			j = j - 1
		}
		data.bookList[j+1].namaBuku = temp
		data.bookList[j+1].lokasi = tempat
		data.bookList[j+1].jumlah = sum

		i = i + 1
	}
	
	i = 0
	indx = 1
	for i < idx-1 {
		fmt.Println("-----------------------------")
		fmt.Println(indx,".Nama Buku   : ",data.bookList[i].namaBuku)
		fmt.Println("   Lokasi Buku : ",data.bookList[i].lokasi)
		fmt.Println("   Jumlah Buku : ",data.bookList[i].jumlah)
		i = i + 1
		indx = indx + 1
	}
}

func sortingInsertDes(data *Array ) {
	var j , i , indx int
	var temp, tempat string
	var sum float64
	
	i = 0
	for i < idx  {
		temp = data.bookList[i].namaBuku
		tempat = data.bookList[i].lokasi
		sum = data.bookList[i].jumlah
		j = i - 1
		for j >= 0 && data.bookList[j].namaBuku < temp {
			data.bookList[j+1].namaBuku, data.bookList[j+1].lokasi, data.bookList[j+1].jumlah = data.bookList[j].namaBuku, data.bookList[j].lokasi, data.bookList[j].jumlah
			j = j - 1
		}
		data.bookList[j+1].namaBuku = temp
		data.bookList[j+1].lokasi = tempat
		data.bookList[j+1].jumlah = sum

		i = i + 1
	}
	
	i = 0
	indx = 1
	for i < idx-1 {
		fmt.Println("-----------------------------")
		fmt.Println(indx,".Nama Buku   : ",data.bookList[i].namaBuku)
		fmt.Println("   Lokasi Buku : ",data.bookList[i].lokasi)
		fmt.Println("   Jumlah Buku : ",data.bookList[i].jumlah)
		i = i + 1
		indx = indx + 1
	}
}

func sortingSelectAs(data *Array){
	var pass,i,min int
	var temp buku
	
	pass = 0
	for pass <= idx {
		min = pass
		i = pass + 1
		for i < idx-1 {
			if data.bookList[pass].namaBuku > data.bookList[i].namaBuku {
				min = i
			}
			i = i + 1
		}
		temp = data.bookList[min]
		data.bookList[min] = data.bookList[pass]
		data.bookList[pass]= temp
		pass = pass + 1
	}
	
	i = 0
	indx := 1
	for i < idx-1 {
		fmt.Println("-----------------------------")
		fmt.Println(indx,".Nama Buku   : ",data.bookList[i].namaBuku)
		fmt.Println("   Lokasi Buku : ",data.bookList[i].lokasi)
		fmt.Println("   Jumlah Buku : ",data.bookList[i].jumlah)
		i = i + 1
		indx = indx + 1
	}
}

func sortingSelectDec(data *Array){
	var pass,i,min int
	var temp buku
	
	pass = 0
	for pass <= idx-1 {
		min = pass
		i = pass + 1
		for i <= idx {
			if data.bookList[min].namaBuku < data.bookList[i].namaBuku {
				min = i
			}
			i = i + 1
		}
		temp = data.bookList[min]
		data.bookList[min] = data.bookList[pass]
		data.bookList[pass]= temp
		pass = pass + 1
	}
	
	i = 0
	indx := 1
	for i < idx-1 {
		fmt.Println("-----------------------------")
		fmt.Println(indx,".Nama Buku   : ",data.bookList[i].namaBuku)
		fmt.Println("   Lokasi Buku : ",data.bookList[i].lokasi)
		fmt.Println("   Jumlah Buku : ",data.bookList[i].jumlah)
		i = i + 1
		indx = indx + 1
	}
}