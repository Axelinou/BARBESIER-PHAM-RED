package main 
import "fmt"

type Personnage struct {
	Nom string
	Classe string
	Niveau int
	Pvmax int
	Pv int 
	inventaire []string

}
func main() {
	init(perso1)
}

func init(perso1) {
	var perso1
	perso1.nom = "éclairus"
	perso1.classe =  "warrior"
	perso1.niveau = 1
	perso1.pvmax = 150
	perso1.pv = 100
	perso1.inventaire = []string{"mace","shield","armor,healthpotion"}

}

func (Personnage) display() {
	fmt.Println("Nom:",perso01.nom)
	fmt.Println("Classe:",perso01.classe)
	fmt.Println("Niveau:",perso01.niveau)
	fmt.Println("Pvmax:",perso01.pvmax)
	fmt.Println("Pv:",perso01.pv)
	fmt.Println("Inventaire:",perso01.inventaire)
}