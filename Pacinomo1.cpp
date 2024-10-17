#include <iostream>
#include <string>
#include <cmath> // Untuk fmod
#include <iomanip> // Untuk setprecision

// Struct Aset sebagai parent class
struct Aset {
    std::string nama;
    std::string jenis;
    int jumlah;
    double harga;
    int tahun; // Tambahkan variabel tahun di sini

    Aset(std::string nama, std::string jenis, int jumlah, double harga)
        : nama(nama), jenis(jenis), jumlah(jumlah), harga(harga), tahun(0) {}

    virtual double nextYear() = 0;  // Mengembalikan earnings
    virtual std::string toString() const = 0;
};

// Struct Saham sebagai child class dari Aset
struct Saham : public Aset {
    double dividen;
    double pertumbuhan;

    Saham(std::string nama, int jumlah, double harga, double pertumbuhan, double dividen)
        : Aset(nama, "SAHAM", jumlah, harga), pertumbuhan(pertumbuhan), dividen(dividen) {}

    double nextYear() override {
        grow();
        harga *= (1 + pertumbuhan);
        return dividen * harga * jumlah;  // Return earnings dari dividen
    }

    void grow() {
        int a = 75;
        int c = 74;
        int m = 2;
        pertumbuhan = fmod((a * pertumbuhan + c), m) - 1;
    }

    std::string toString() const override {
        return "Nama: " + nama + "\nTipe: " + jenis + "\nHarga: " + std::to_string(harga) + 
               "\nJumlah: " + std::to_string(jumlah) + "\nDividen: " + std::to_string(dividen) +
               "\nPertumbuhan: " + std::to_string(pertumbuhan);
    }
};

// Struct Obligasi sebagai child class dari Aset
struct Obligasi : public Aset {
    double bunga;
    int maturitas;
    bool jatuhTempo;

    Obligasi(std::string nama, int jumlah, double harga, double bunga, int maturitas)
        : Aset(nama, "OBLIGASI", jumlah, harga), bunga(bunga), maturitas(maturitas), jatuhTempo(false) {}

    double nextYear() override {
        tahun++; // Tambahkan ini untuk memperbarui tahun
        if (tahun >= maturitas) jatuhTempo = true;
        return jatuhTempo ? 0 : bunga * harga * jumlah;
    }

    std::string toString() const override {
        return "Nama: " + nama + "\nTipe: " + jenis + "\nHarga: " + std::to_string(harga) + 
               "\nJumlah: " + std::to_string(jumlah) + "\nBunga: " + std::to_string(bunga) +
               "\nJatuh Tempo: " + (jatuhTempo ? "true" : "false");
    }
};

// Node struct untuk menyimpan data dalam linked list
struct Node {
    Aset* data;
    Node* next;

    Node(Aset* data) : data(data), next(nullptr) {}
};

// Kelas Pacilnomo untuk mengelola linked list dari aset
class Pacilnomo {
private:
    Node* head;
    double earnings;

    void printSeparator() const {
        std::cout << std::string(64, '=') << std::endl;
    }

    void daftarAset() const {
        printSeparator();
        Node* current = head;
        int count = 0;
        while (current != nullptr) {
            std::cout << "- " << current->data->toString() << std::endl;
            current = current->next;
            count++;
        }
        std::cout << "Kamu memiliki " << count << " total aset" << std::endl;
        printSeparator();
    }

    void infoPortofolio() const {
        int jumlahSaham = 0, jumlahObligasi = 0;
        double netWorth = 0;
        double totalEarnings = 0;
        Node* current = head;
        
        while (current != nullptr) {
            netWorth += current->data->harga * current->data->jumlah;
            if (current->data->jenis == "SAHAM") {
                jumlahSaham++;
                Saham* saham = dynamic_cast<Saham*>(current->data);
                if (saham) {
                    totalEarnings += saham->dividen * saham->harga * saham->jumlah;
                }
            } else if (current->data->jenis == "OBLIGASI") {
                jumlahObligasi++;
                Obligasi* obligasi = dynamic_cast<Obligasi*>(current->data);
                if (obligasi && !obligasi->jatuhTempo) {
                    totalEarnings += obligasi->bunga * obligasi->harga * obligasi->jumlah;
                }
            }
            current = current->next;
        }
        
        netWorth += totalEarnings;
        printSeparator();
        std::cout << "Info Portofolio\n"
                  << "Jumlah Jenis Saham: " << jumlahSaham << "\n"
                  << "Jumlah Jenis Obligasi: " << jumlahObligasi << "\n"
                  << "Total Nilai Portofolio: " << std::fixed << std::setprecision(2) << netWorth << std::endl;
        printSeparator();
    }

    void nextYear() {
        Node* current = head;
        while (current != nullptr) {
            earnings += current->data->nextYear();
            current = current->next;
        }
        earnings += earnings;
    }

public:
    Pacilnomo() : head(nullptr), earnings(0) {}

    void addAset(Aset* aset) {
        Node* newNode = new Node(aset);
        newNode->next = head;
        head = newNode;
    }

    void run() {
        int banyakAset;
        std::cout << "Silakan masukkan banyak aset yang tersedia: ";
        std::cin >> banyakAset;
        std::cin.ignore();

        for (int i = 0; i < banyakAset; i++) {
            std::cout << "Aset " << i + 1 << ": ";
            std::string namaAset, jenisAset;
            int jumlah;
            double harga;
            std::cin >> namaAset >> jenisAset >> jumlah >> harga;

            if (jenisAset == "SAHAM") {
                double pertumbuhan, dividen;
                std::cin >> pertumbuhan >> dividen;
                Saham* saham = new Saham(namaAset, jumlah, harga, pertumbuhan, dividen);
                addAset(saham);
            } else if (jenisAset == "OBLIGASI") {
                double bunga;
                int maturitas;
                std::cin >> bunga >> maturitas;
                Obligasi* obligasi = new Obligasi(namaAset, jumlah, harga, bunga, maturitas);
                addAset(obligasi);
            }
        }

        std::cout << "Selamat datang di...\n"
                  << " /$$$$$$$                     /$$ /$$\n"
                  << "| $$__  $$                   |__/| $$\n"
                  << "| $$  \\ $$ /$$$$$$   /$$$$$$$ /$$| $$ /$$$$$$$   /$$$$$$  /$$$$$$/$$$$   /$$$$$$\n"
                  << "| $$$$$$$/|____  $$ /$$_____/| $$| $$| $$__  $$ /$$__  $$| $$_  $$_  $$ /$$__  $$\n"
                  << "| $$____/  /$$$$$$$| $$      | $$| $$| $$  \\ $$| $$  \\ $$| $$ \\ $$ \\ $$| $$  \\ $$\n"
                  << "| $$      /$$__  $$| $$      | $$| $$| $$  | $$| $$  | $$| $$ | $$ | $$| $$  | $$\n"
                  << "| $$     |  $$$$$$$|  $$$$$$$| $$| $$| $$  | $$|  $$$$$$/| $$ | $$ | $$|  $$$$$$/\n"
                  << "|__/      \\_______/ \\_______/|__/|__/|__/  |__/ \\______/ |__/ |__/ |__/ \\______/\n\n";

        while (true) {
            std::cout << "Silakan pilih salah satu opsi berikut:\n"
                      << "[1] Daftar aset\n"
                      << "[2] Info portofolio\n"
                      << "[3] Lanjut ke tahun berikutnya\n"
                      << "[*] Keluar\n";
            printSeparator();
            std::cout << "Input: ";
            std::string pilihan;
            std::cin >> pilihan;
            if (pilihan == "1") {
                daftarAset();
            } else if (pilihan == "2") {
                infoPortofolio();
            } else if (pilihan == "3") {
                nextYear();
                std::cout << "Setahun telah berlalu...\n";
                printSeparator();
            } else {
                std::cout << "Terima kasih telah menggunakan layanan Pacilnomo ~ !\n";
                break;
            }
        }
    }
};

int main() {
    Pacilnomo pacilnomo;
    pacilnomo.run();
    return 0;
}
