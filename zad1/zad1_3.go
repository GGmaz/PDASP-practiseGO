package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Student1 struct {
	ime     string
	prezime string
	prosek  float64
}

func main() {
	n := 0
	for {
		fmt.Println("za dodavanje studenta pritisni 1, za brisanje studenta pritisni 2, za prikaz svih pritisni 3, za izlazak pritisni 4")
		_, err := fmt.Scanln(&n)
		if err != nil {
			if err.Error() != "unexpected newline" && err.Error() != "expected integer" {
				return
			}
		}
		if n == 3 {
			readFromFile("studenti.txt")
		} else if n == 1 {
			var student Student1
			var indeks string
			fmt.Println("unesi ime, prezime, prosek i broj indeksa")
			_, err := fmt.Scanln(&student.ime, &student.prezime, &student.prosek, &indeks)
			if err != nil {
				if err.Error() != "unexpected newline" {
					return
				}
			}
			writeToFile("studenti.txt", student, indeks)
		} else if n == 2 {
			var indeks string
			fmt.Println("unesi indeks onog kog zelis da obrises")
			_, err := fmt.Scanln(&indeks)
			if err != nil {
				if err.Error() != "unexpected newline" {
					return
				}
			}

			file, err := os.OpenFile("studenti.txt", os.O_RDONLY, os.ModeExclusive)
			if err != nil {
				fmt.Println("nije pronasao file")
			}
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				parts := strings.Split(scanner.Text(), "|")
				if strings.TrimSpace(parts[len(parts)-1]) != indeks {
					writeLineToFile("pom.txt", scanner.Text())
				}
			}

			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
			file.Close()

			e := os.Remove("studenti.txt")
			if e != nil {
				log.Fatal(e)
			}
			e1 := os.Rename("pom.txt", "studenti.txt")
			if e1 != nil {
				log.Fatal(e)
			}
		} else if n == 4 {
			break
		}
	}
}

func readFromFile(filePath string) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, os.ModeExclusive)
	defer file.Close()
	if err != nil {
		fmt.Println("nije pronasao file")
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func writeToFile(filePath string, student Student1, indeks string) {
	//os.Create(filePath) // Može se otvoriti fajl u režimu koji ga kreira ako ne postoji
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("cant open file")
	}
	defer file.Close() //what is defer?
	writer := bufio.NewWriter(file)
	prosek := fmt.Sprintf("%f", student.prosek)
	_, err = writer.WriteString(student.ime + " " + student.prezime + " " + prosek + " | " + indeks + "\n")
	if err != nil {
		log.Fatal(err)
	}
	err = writer.Flush()
}

func writeLineToFile(filePath string, tekst string) {
	//os.Create(filePath) // Može se otvoriti fajl u režimu koji ga kreira ako ne postoji
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("cant open file")
	}
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(tekst + "\n")
	if err != nil {
		log.Fatal(err)
	}
	err = writer.Flush()
	defer file.Close() //what is defer?
}
