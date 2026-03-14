# <h1 align="center">Laporan Praktikum Modul 2 - ... </h1>
<p align="center">[Ichsan Maulana Muhammad] - [109082500093]</p>

## Unguided 

### 1. [Soal]
Telusuri program berikut dengan cara mengkompilasi dan mengeksekusi program. Silakan
masukan data yang sesuai sebanyak yang diminta program. Perhatikan keluaran yang
diperoleh. Coba terangkan apa sebenarnya yang dilakukan program tersebut?
#### soal1.go

```go
package main
import “fmt”
func main() {
var (
satu, dua, tiga string
temp string
)
fmt.Print("Masukan input string: ")
fmt.Scanln(&satu)
fmt.Print("Masukan input string: ")
fmt.Scanln(&dua)
fmt.Print("Masukan input string: ")
fmt.Scanln(&tiga)
fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
temp = satu
satu = dua
dua = tiga
tiga = temp
fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Isanmaulana/109082500093_Ichsan-Maulana-Muhammad/blob/main/output%20soal1.png)
[penjelasan]
program disuruh minginputkan string satu, dua, dan tiga, lalu program akan melakukan perubahan posisi nilai variabel menggunakan temp, yaitu nilai satu disimpan ke dalam temp, lalu nilai dua dipindahkan ke satu, nilai tiga dipindahkan ke dua, dan terakhir nilai yang ada di temp dipindahkan ke tiga


### 2. [Soal]
Siswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakan
praktikum kimia. Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana
susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta
untuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunan
warna zat cair pada gelas 1 hingga gelas 4 secara berturutan adalah ‘merah’, ‘kuning’,
‘hijau’, dan ‘ungu’ selama 5 kali percobaan berulang.
Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi
sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan
warna sesuai dengan informasi yang diberikan pada paragraf sebelumnya, dan false
untuk urutan warna lainnya.
Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah
adalah input/read):
#### soal2.go

```go
package main

import "fmt"

func main() {
	var a, b, c, d string
	hasil := true

	for i := 1; i <= 5; i++ {
		fmt.Print("Percobaan ", i, ": ")
		fmt.Scan(&a, &b, &c, &d)

		if a != "merah" || b != "kuning" || c != "hijau" || d != "ungu" {
			hasil = false
		}
	}
	fmt.Println("BERHASIL:", hasil)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Isanmaulana/109082500093_Ichsan-Maulana-Muhammad/blob/main/output%20soal2.png)
[penjelasan]
program ini menerima input warna dari 4 tabung reaksi selama 5 kali percobaan, setiap percobaan program akan mengecek apakah urutan warnanya merah, kuning, hijau, ungu. Jika semua percobaan memiliki urutan yang sama, maka program menampilkan true, tetapi jika ada percobaan yang berbeda maka hasilnya false


### 3. [Soal]
PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel. Maka,
buatlah program BiayaPos untuk menghitung biaya pengiriman tersebut dengan ketentuan
sebagai berikut!
Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam
gram). Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari 500
gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500
gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Sisa berat (yang
kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg.
#### soal3.go

```go
package main

import "fmt"

func main() {
	var kg, gram, sisa, total, biayakirim, biayasisa int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&gram)
	kg = gram / 1000
	sisa = gram % 1000
	biayakirim = kg * 10000

	if sisa >= 500 {
		biayasisa = sisa * 5
	} else {
		biayasisa = sisa * 15
	}
	total = biayakirim + biayasisa
	if kg > 10 {
		total = biayakirim
	}
	fmt.Println("Detail berat:", kg, "kg +", sisa, "gr")
	fmt.Println("Detail biaya: Rp.", biayakirim, "+ Rp.", biayasisa)
	fmt.Println("Total biaya: Rp.", total)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Isanmaulana/109082500093_Ichsan-Maulana-Muhammad/blob/main/output%20soal3.png)
[penjelasan]
program ini digunakan untuk menghitung biaya pengiriman parsel berdasarkan berat dalam gram. Berat parsel diubah menjadi kilogram dan sisa gram. Biaya kirim dihitung Rp10.000 per kg, lalu sisa gram dikenakan biaya Rp5 per gram jika ≥500 gram dan Rp15 per gram jika <500 gram. Jika berat parsel lebih dari 10 kg maka biaya sisa gram tidak dihitung / gratis