package main

import (
	"bufio"  // Package untuk membaca input lebih dari satu kata dari pengguna
	"fmt"    // Package untuk mencetak output ke layar
	"os"     // Package untuk interaksi dengan sistem operasi
)

const maxUsers = 100
const maxPosts = 100
const maxFriends = 100
const maxComments = 100
// Mendefinisikan batas maksimum untuk jumlah pengguna, postingan, teman, dan komentar

type Comment struct {
	Username string
	Comment  string
}
// Struktur untuk menyimpan komentar, terdiri dari username dan isi komentar

type Post struct {
	Friend       string
	Status       string
	Comments     [maxComments]Comment
	CommentCount int
}
// Struktur untuk menyimpan postingan, terdiri dari username teman, status, daftar komentar, dan jumlah komentar

type Profile struct {
	Name  string
	Email string
	Age   int
}
// Struktur untuk menyimpan profil pengguna, terdiri dari nama, email, dan umur

type User struct {
	Username    string
	Password    string
	Friends     [maxFriends]string
	FriendCount int
	Posts       [maxPosts]Post
	PostCount   int
	Profile     Profile
}
// Struktur untuk menyimpan informasi pengguna, terdiri dari username, password, daftar teman, jumlah teman, daftar postingan, jumlah postingan, dan profil

var users [maxUsers]User
var userCount int
// Mendefinisikan array untuk menyimpan pengguna dan variabel untuk menghitung jumlah pengguna

func main() {
	for {
		var pilihan int
		fmt.Println("\nSelamat datang di GoThread!")
		fmt.Println("===========================")
		fmt.Println("1 = Register")
		fmt.Println("2 = Login")
		fmt.Println("3 = Keluar dari GoThread")
		fmt.Print("Masukkan pilihan (1-3) : ")
		fmt.Scanln(&pilihan)
		// Menampilkan menu utama dan meminta pengguna untuk memilih opsi

		if pilihan == 1 { // Mengarahkan pengguna berdasarkan pilihan yang dimasukkan
			buatAkun()
		} else if pilihan == 2 {
			var username, password string
			fmt.Print("Masukkan Username: ")
			fmt.Scanln(&username)
			fmt.Print("Masukkan Password: ")
			fmt.Scanln(&password)

			if login(username, password) {
				homepage(username)
			} else {
				fmt.Println("Username/Password yang anda masukkan salah")
			}
		} else if pilihan == 3 {
			fmt.Println("\nSelamat menimbun stres kembali!\n")
			return
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
		
	}
}

func buatAkun() {
	var username, password, name, email string
	var age int

	fmt.Println("==============")
	fmt.Println("\nBuat Akun baru")
	fmt.Print("Masukan username: ")
	fmt.Scanln(&username)

	fmt.Print("Masukan password: ")
	fmt.Scanln(&password)

	fmt.Print("Masukan nama: ")
	scanner := bufio.NewScanner(os.Stdin) // Membuat scanner untuk membaca input dari terminal
	scanner.Scan() // Membaca satu baris input dari pengguna
	name = scanner.Text() // Menyimpan input nama dari pengguna ke variabel name

	fmt.Print("Masukan email: ")
	fmt.Scanln(&email)

	fmt.Print("Masukan umur: ")
	fmt.Scanln(&age)
	// Meminta pengguna untuk memasukkan informasi akun baru

	if userCount < maxUsers { // Jika jumlah pengguna saat ini kurang dari batas maksimum pengguna
		users[userCount] = User{ // Tambahkan pengguna baru ke dalam array users di indeks userCount
			Username: username,
			Password: password,
			Profile: Profile{ // Mengisi profil pengguna dengan nama, email, dan umur
				Name:  name,
				Email: email,
				Age:   age,
			},
		}
		userCount++ //Baris ini digunakan untuk menambah jumlah pengguna yang terdaftar sebanyak 1, yang berarti ada satu pengguna baru yang ditambahkan ke sistem.
		fmt.Println("Akun Berhasil Dibuat!")
	} else {
		fmt.Println("User limit reached, cannot create new account.")
	}
	// Menambahkan pengguna baru ke dalam array `users` jika belum mencapai batas maksimum
}

func login(username, password string) bool {
	for i := 0; i < userCount; i++ { // Loop melalui semua pengguna yang ada
		if users[i].Username == username && users[i].Password == password { // Jika username dan password cocok dengan pengguna ke-i
			return true // Kembalikan nilai true, menandakan login berhasil
		}
	}
	return false
	// Memeriksa apakah username dan password yang dimasukkan cocok dengan salah satu pengguna yang terdaftar
}

func homepage(username string) {
	var input string

	fmt.Println("\nSelamat datang di aplikasi GoThread,", username)
	fmt.Println("\n===================================\n")
	fmt.Println("Berikut merupakan status teman-teman anda yang terbaru hari ini :")

	lihatSemuaStatus(username)
	// Menampilkan status terbaru dari teman-teman pengguna

	fmt.Println("\nPilih opsi berikut:")
	fmt.Println("1 = Unggah postingan status baru")
	fmt.Println("2 = Tambah Komentar")
	fmt.Println("3 = Tambah teman")
	fmt.Println("4 = Hapus teman")
	fmt.Println("5 = Lihat daftar teman")
	fmt.Println("6 = Cari pengguna")
	fmt.Println("7 = Edit profil")
	fmt.Println("8 = Log out\n")

	fmt.Print("Masukkan pilihan: ")
	fmt.Scanln(&input)
	// Menampilkan menu opsi setelah login

	switch input { // Menjalankan fungsi sesuai dengan opsi yang dipilih pengguna
	case "1":
		posting(username)
		homepage(username)
	case "2":
		var statusUsername string
		fmt.Print("Masukkan username status yang ingin dikomentari : ")
		fmt.Scanln(&statusUsername)

		var statusContent string
		fmt.Print("Masukkan konten status yang ingin dikomentari : ")
		fmt.Scanln(&statusContent)

		var comment string
		fmt.Print("Masukkan komentar : ")
		scanner := bufio.NewScanner(os.Stdin) // Membuat scanner untuk membaca input dari pengguna melalui keyboard
		scanner.Scan() // Membaca input dari pengguna
		comment = scanner.Text() // Menyimpan input yang dibaca ke dalam variabel `comment`
		 
		tambahKomentar(username, statusUsername, statusContent, comment)
		homepage(username)
	case "3":
		tambahTeman(username)
		homepage(username)
	case "4":
		hapusTeman(username)
		homepage(username)
	case "5":
		lihatDaftarTeman(username)
		homepage(username)
	case "6":
		cariPengguna()
		homepage(username)
	case "7":
		editProfil(username)
		homepage(username)
	case "8":
		return
	default: //digunakan untuk menangani kasus di mana input pengguna tidak sesuai dengan salah satu opsi yang telah ditentukan.
		fmt.Println("Input salah, silahkan coba lagi")
	}
	
}

func posting(username string) {
	var isistatus string // Variabel untuk menyimpan status yang dimasukkan pengguna

	fmt.Println("Curahkan isi hati anda : ") // Menampilkan pesan untuk meminta input status dari pengguna

	scanner := bufio.NewScanner(os.Stdin) // Membuat scanner untuk membaca input dari terminal
	scanner.Scan() // Membaca satu baris input dari pengguna
	isistatus = scanner.Text() // Menyimpan input status dari pengguna ke variabel isistatus

	// Meminta pengguna untuk memasukkan status baru

	for i := 0; i < userCount; i++ { // Loop untuk mencari pengguna berdasarkan username
		if users[i].Username == username { // Jika username ditemukan
			if users[i].PostCount < maxPosts { // Jika jumlah postingan belum mencapai batas maksimum
				users[i].Posts[users[i].PostCount] = Post{
					Friend: username, // Menambahkan username ke field Friend di Post
					Status: isistatus, // Menambahkan status ke field Status di Post
				}
				users[i].PostCount++ // Meningkatkan jumlah postingan pengguna
				fmt.Println("\nStatus baru berhasil diunggah!\n") // Menampilkan pesan sukses
			} else {
				fmt.Println("Post limit reached, cannot create new post.") // Menampilkan pesan jika batas postingan tercapai
			}
			return // Keluar dari fungsi setelah menambahkan status
		}
	}
	// Menambahkan status baru ke dalam daftar postingan pengguna
}

func tambahKomentar(currentUser, statusUsername, statusContent, comment string) {	
	for i := 0; i < userCount; i++ { // Loop untuk mencari pengguna berdasarkan username status
		if users[i].Username == statusUsername { // Jika pengguna dengan username status ditemukan
			for j := 0; j < users[i].PostCount; j++ { // Loop untuk mencari status yang cocok
				if users[i].Posts[j].Friend == statusUsername && users[i].Posts[j].Status == statusContent {
					if users[i].Posts[j].CommentCount < maxComments { // Jika jumlah komentar belum mencapai batas maksimum
						users[i].Posts[j].Comments[users[i].Posts[j].CommentCount] = Comment{
							Username: currentUser, // Menambahkan username komentar ke field Username di Comment
							Comment:  comment, // Menambahkan komentar ke field Comment di Comment
						}
						users[i].Posts[j].CommentCount++ // Meningkatkan jumlah komentar pada status
						fmt.Println("Komentar berhasil ditambahkan!") // Menampilkan pesan sukses
						return // Keluar dari fungsi setelah menambahkan komentar
					} else {
						fmt.Println("Comment limit reached, cannot add new comment.") // Menampilkan pesan jika batas komentar tercapai
					}
					return // Keluar dari fungsi setelah mencoba menambahkan komentar
				}
			}
			fmt.Println("Status tidak ditemukan.") // Menampilkan pesan jika status tidak ditemukan
			return // Keluar dari fungsi setelah mencari status
		}
	}
	fmt.Println("Pengguna tidak ditemukan.") // Menampilkan pesan jika pengguna tidak ditemukan
	// Menambahkan komentar baru ke status yang dipilih
}

func tambahTeman(username string) {
	var teman string // Variabel untuk menyimpan username teman baru
	fmt.Print("Masukkan username teman yang ingin ditambahkan: ") // Menampilkan pesan untuk meminta input username teman
	fmt.Scanln(&teman) // Membaca input username teman dari pengguna
	// Meminta pengguna untuk memasukkan username teman baru

	for i := 0; i < userCount; i++ { // Loop untuk mencari pengguna berdasarkan username
		if users[i].Username == username { // Jika pengguna dengan username ditemukan
			if users[i].FriendCount < maxFriends { // Jika jumlah teman belum mencapai batas maksimum
				for j := 0; j < userCount; j++ { // Loop untuk mencari teman yang akan ditambahkan
					if users[j].Username == teman { // Jika username teman ditemukan
						users[i].Friends[users[i].FriendCount] = teman // Menambahkan teman ke daftar teman pengguna
						users[i].FriendCount++ // Meningkatkan jumlah teman pengguna
						insertionSort(users[i].Friends[:users[i].FriendCount]) // Mengurutkan daftar teman menggunakan insertion sort
						fmt.Println("Teman berhasil ditambahkan!") // Menampilkan pesan sukses
						return // Keluar dari fungsi setelah menambahkan teman
					}
				}
				fmt.Println("Pengguna tidak ditemukan.") // Menampilkan pesan jika pengguna (teman) tidak ditemukan
			} else {
				fmt.Println("Friend limit reached, cannot add new friend.") // Menampilkan pesan jika batas teman tercapai
			}
			return // Keluar dari fungsi setelah mencoba menambahkan teman
		}
	}
	// Menambahkan teman baru ke dalam daftar teman pengguna
}

func insertionSort(arr []string) {
	n := len(arr) // Mendapatkan panjang array
	for i := 1; i < n; i++ { // Loop untuk setiap elemen dalam array mulai dari indeks 1
		key := arr[i] // Menyimpan elemen yang akan dibandingkan
		j := i - 1 // Inisialisasi indeks untuk perbandingan
		for j >= 0 && arr[j] > key { // Loop untuk memindahkan elemen yang lebih besar dari key ke posisi berikutnya
			arr[j+1] = arr[j] // Memindahkan elemen ke posisi berikutnya
			j = j - 1 // Mengurangi indeks untuk perbandingan selanjutnya
		}
		arr[j+1] = key // Menempatkan key pada posisi yang benar
	}
	// Mengurutkan array teman menggunakan metode insertion sort
}

func hapusTeman(username string) {
	var teman string // Variabel untuk menyimpan username teman yang akan dihapus
	fmt.Print("Masukkan username teman yang ingin dihapus: ") // Menampilkan pesan untuk meminta input username teman
	fmt.Scanln(&teman) // Membaca input username teman dari pengguna
	// Meminta pengguna untuk memasukkan username teman yang ingin dihapus

	for i := 0; i < userCount; i++ { // Loop untuk mencari pengguna berdasarkan username
		if users[i].Username == username { // Jika pengguna dengan username ditemukan
			for j := 0; j < users[i].FriendCount; j++ { // Loop untuk mencari teman yang akan dihapus
				if users[i].Friends[j] == teman { // Jika username teman ditemukan
					for k := j; k < users[i].FriendCount-1; k++ { // Loop untuk menggeser elemen ke kiri setelah menghapus teman
						users[i].Friends[k] = users[i].Friends[k+1] // Menggeser elemen ke kiri
					}
					users[i].Friends[users[i].FriendCount-1] = "" // Mengosongkan elemen terakhir setelah penggeseran
					users[i].FriendCount-- // Mengurangi jumlah teman pengguna
					selectionSort(users[i].Friends[:users[i].FriendCount]) // Mengurutkan daftar teman menggunakan selection sort
					fmt.Println("Teman berhasil dihapus!") // Menampilkan pesan sukses
					return // Keluar dari fungsi setelah menghapus teman
				}
			}
			fmt.Println("Teman tidak ditemukan.") // Menampilkan pesan jika teman tidak ditemukan
			return // Keluar dari fungsi setelah mencari teman
		}
	}
	// Menghapus teman dari daftar teman pengguna
}

func selectionSort(arr []string) {
	n := len(arr) // Mendapatkan panjang array
	for i := 0; i < n-1; i++ { // Loop untuk setiap elemen dalam array kecuali elemen terakhir
		minIdx := i // Menyimpan indeks elemen terkecil
		for j := i + 1; j < n; j++ { // Loop untuk mencari elemen terkecil di sisa array
			if arr[j] < arr[minIdx] { // Jika elemen lebih kecil ditemukan
				minIdx = j // Mengupdate indeks elemen terkecil
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i] // Menukar elemen terkecil dengan elemen saat ini
	}
	// Mengurutkan array teman menggunakan metode selection sort
}

func lihatDaftarTeman(username string) {
	fmt.Println("Daftar Teman:") // Menampilkan pesan header
	for i := 0; i < userCount; i++ { // Loop untuk mencari pengguna berdasarkan username
		if users[i].Username == username { // Jika pengguna dengan username ditemukan
			for j := 0; j < users[i].FriendCount; j++ { // Loop untuk menampilkan setiap teman dalam daftar teman pengguna
				fmt.Println(users[i].Friends[j]) // Menampilkan username teman
			}
		}
	}
	// Menampilkan daftar teman pengguna
}

func cariPengguna() {
	var pencarian string // Variabel untuk menyimpan username yang akan dicari
	fmt.Print("Masukkan username yang ingin dicari: ") // Menampilkan pesan untuk meminta input username
	fmt.Scanln(&pencarian) // Membaca input username dari pengguna
	// Meminta pengguna untuk memasukkan username yang ingin dicari

	for i := 0; i < userCount; i++ { // Loop untuk mencari pengguna berdasarkan username
		if users[i].Username == pencarian { // Jika pengguna dengan username ditemukan
			fmt.Printf("Pengguna ditemukan: %s\n", users[i].Username) // Menampilkan pesan pengguna ditemukan
			return // Keluar dari fungsi setelah menemukan pengguna
		}
	}

	fmt.Println("Pengguna tidak ditemukan.") // Menampilkan pesan jika pengguna tidak ditemukan
	// Mencari pengguna berdasarkan username
}

func editProfil(username string) {
	for i := 0; i < userCount; i++ { // Loop untuk mencari pengguna berdasarkan username
		if users[i].Username == username { // Jika pengguna dengan username ditemukan
			fmt.Print("Masukkan nama baru: ") // Menampilkan pesan untuk meminta input nama baru
			fmt.Scanln(&users[i].Profile.Name) // Membaca input nama baru dari pengguna
			fmt.Print("Masukkan email baru: ") // Menampilkan pesan untuk meminta input email baru
			fmt.Scanln(&users[i].Profile.Email) // Membaca input email baru dari pengguna
			fmt.Print("Masukkan umur baru: ") // Menampilkan pesan untuk meminta input umur baru
			fmt.Scanln(&users[i].Profile.Age) // Membaca input umur baru dari pengguna
			fmt.Println("Profil berhasil diperbarui!") // Menampilkan pesan sukses
			return // Keluar dari fungsi setelah mengupdate profil
		}
	}
	// Mengedit profil pengguna
}

func lihatSemuaStatus(username string) {
	for i := 0; i < userCount; i++ { // Loop untuk setiap pengguna
		for j := users[i].PostCount - 1; j >= 0; j-- { // Loop untuk setiap postingan pengguna dari yang terbaru
			fmt.Printf("\nUsername : %s\nStatus : %s\n", users[i].Posts[j].Friend, users[i].Posts[j].Status) // Menampilkan username dan status
			for k := 0; k < users[i].Posts[j].CommentCount; k++ { // Loop untuk setiap komentar pada status
				fmt.Printf("\tKomentar dari %s: %s\n", users[i].Posts[j].Comments[k].Username, users[i].Posts[j].Comments[k].Comment) // Menampilkan username dan komentar
			}
		}
	}
	// Menampilkan semua status dan komentar
}	

// Git test