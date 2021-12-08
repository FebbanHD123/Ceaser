package main

import "fmt"

func main() {
	fmt.Println(ceaserVer("abcdefghijklmnopqrstuvwxyz", 13))
}

func ceaserVer(s string, schlüssel byte) string {
	//Vor.: String parameter & schlüssel parameter wird übergeben
	//Eff.: String wird verschlüsselt zurückgegeben
	var ergebnis string
	var i int
	for i < len(s) {
		
		var ascii byte = s[i]
		if ascii >= 97 && ascii <=122 {
			ascii += schlüssel
			if ascii > 122 {
				ascii -= 26
			}
		}
		
		if ascii >= 65 && ascii <= 90 {
			ascii += schlüssel
			if ascii > 90 {
				ascii -= 26
			}
		}
		
		ergebnis += string(ascii)
		i++
	}
	
	return ergebnis
}

func ceaserEnt(s string, schlüssel byte) string {
	//Vor.: String parameter & schlüssel parameter wird übergeben
	//Eff.: String wird entschlüsselt zurückgegeben
	var ergebnis string
	var i int
	for i < len(s) {
		
		var ascii byte = s[i]
		if ascii >= 97 && ascii <=122 {
			ascii -= schlüssel
			if ascii < 97 {
				ascii += 26
			}
		}
		
		if ascii >= 65 && ascii <= 90 {
			ascii -= schlüssel
			if ascii < 41 {
				ascii += 26
			}
		}
		
		ergebnis += string(ascii)
		i++
	}
	
	return ergebnis
}
