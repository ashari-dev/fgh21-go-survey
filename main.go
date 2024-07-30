package main

import "fmt"

type Respondent struct{
	name string
	age int
	gender string
	isSmoker bool
	cigarVarian []string
}
func main() {


	var name string
	fmt.Print("Siapa nama anda ? ")
	fmt.Scan(&name)

	var age int
	fmt.Print("Berapa usia anda ? ")
	fmt.Scan(&age)

	var gender string
	var scanGender string
	fmt.Print("Apa jenis kelami anda ?\n 1.Laki-laki 2.Perempuan ")
	fmt.Scan(&scanGender)
	if scanGender == "1" {
		gender = "Laki-laki"
	}else{
		gender = "Petempuan"
	}

	var isSmoker bool
	var scanSmoke string
	fmt.Print("Apakan anda merokok ?\n 1.ya 2.Tidak ")
	fmt.Scan(&scanSmoke)
	
	if scanSmoke == "1" {
		isSmoker = true
	}else{
		isSmoker = false
	}
	
  var	cigarVarian []string =[]string{"surya","sampoerna"}
	fmt.Println("Jenis rokok yang pernah anda coba ?")
	
	
	survey := []Respondent{
		{name: name ,age: age, gender: gender ,isSmoker: isSmoker, cigarVarian: cigarVarian},
	}

	fmt.Print(survey)
	
}