package main

import "fmt"

type Profile struct {
	Name   string
	Health int
	Power  int
	Exp    int
}

func main() {

	profile := MakeProfile("Goku")
	fmt.Println("Name : ", profile.Name)
	fmt.Println("Health : ", profile.Health)
	fmt.Println("Power : ", profile.Power)
	fmt.Println("Exp : ", profile.Exp)
	fmt.Println("=== HEROES POWER UP===")

	powerup := PowerUp(profile, 2)
	fmt.Println("Name : ", powerup.Name)
	fmt.Println("Health : ", powerup.Health)
	fmt.Println("Power : ", powerup.Power)
	fmt.Println("Exp : ", powerup.Exp)
}
func PowerUp(nilai Profile, multiplayer int) Profile {
	nilai.Health = nilai.Health + (nilai.Health * multiplayer)
	return nilai
}

func MakeProfile(nameChar string) *Profile {
	if nameChar == "Sasuke" || nameChar == "sasuke" {
		return &Profile{
			Name:   "Sasuke",
			Health: 200,
			Power:  100,
			Exp:    0,
		}

	} else if nameChar == "Goku" {
		return &Profile{
			Name:   nameChar,
			Health: 400,
			Power:  300,
			Exp:    100,
		}
	} else if nameChar == "Naruto" {
		return &Profile{
			Name:   nameChar,
			Health: 150,
			Power:  200,
			Exp:    50,
		}
	}
	return &Profile{}
}
