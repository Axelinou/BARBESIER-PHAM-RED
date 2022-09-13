package main

import "fmt"

type Personnage struct {
	Nom         string
	Classe      string
	Niveau      int
	Pvmax       int
	Pv          int
	inp1entaire []string
}

func main() {
	var p1 Personnage
	p1.Nom = "Suad"
	p1.Classe = "Chômeur"
	p1.Niveau = 3
	p1.Pvmax = 10
	p1.Pv = 1
	//fmt.Println(p1)
	p1.display()
	//p1.setPower()
	//p1.LessPower()
}

func (p1 Personnage) display() {
	fmt.Println("----------------")
	fmt.Println("Nom du joueur:", p1.Nom)
	fmt.Println("Classe:", p1.Classe)
	fmt.Println("Santé:", p1.Pv)
	fmt.Println("Niveau:", p1.Niveau)
	fmt.Println("Santé Max :", p1.Pvmax)
	fmt.Println("----------------")

}
func (p1 *Personnage) setPower() {
	p1.Niveau += 100
	fmt.Println("Power:", p1.Niveau)
}

func (p1 *Personnage) LessPower() {
	p1.Niveau -= 100
	fmt.Println("Power:", p1.Niveau)
}
