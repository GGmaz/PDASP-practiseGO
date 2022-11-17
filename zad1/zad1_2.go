package main

import (
	"fmt"
)

type Student struct {
	ime     string
	prezime string
	prosek  float64
}

func main() {
	//readFromFile("studenti.txt")

	studenti := make(map[string]Student)
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
			for key, value := range studenti {
				fmt.Println("Key : ", key, "Value : ", value)
			}
			if len(studenti) == 0 {
				fmt.Println("Nema ih")
			}
		} else if n == 1 {
			var student Student
			var indeks string
			fmt.Println("unesi ime, prezime, prosek i broj indeksa")
			_, err := fmt.Scanln(&student.ime, &student.prezime, &student.prosek, &indeks)
			if err != nil {
				if err.Error() != "unexpected newline" {
					return
				}
			}
			studenti[indeks] = student
		} else if n == 2 {
			var indeks string
			fmt.Println("unesi indeks onog kog zelis da obrises")
			_, err := fmt.Scanln(&indeks)
			if err != nil {
				if err.Error() != "unexpected newline" {
					return
				}
			}
			delete(studenti, indeks)
		} else if n == 4 {
			break
		}
	}
}
