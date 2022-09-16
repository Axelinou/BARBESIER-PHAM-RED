package main

import "fmt"

type Personnage struct {
	Nom                      string
	Classe                   string
	Niveau                   int
	Pvmax                    int
	Pv                       int
	inventaire               []string
	inventairepotion         []string
	capacitéinventairepotion int
	inventairemarchand       []string
}

func main() {
	var p1 Personnage
	p1.Nom = "Suad"
	p1.Classe = "Chômeur"
	p1.Niveau = 3
	p1.Pvmax = 100
	p1.Pv = 50
	p1.inventaire = []string{"Cv", "allocation"}
	p1.inventairepotion = []string{"potion"}
	p1.capacitéinventairepotion = 3
	p1.inventairemarchand = []string{"uzi", "mp5", "ermaemp", "ruby", "remington700"}
	//fmt.Println(p1)
	//p1.display()
	//p1.setPower()
	//p1.LessPower()
	fmt.Println("----------------------------")
	fmt.Println("===================================")
	fmt.Println("•BIENVENUE SUR  A CHOMAGE ADVENTURE• ")
	fmt.Println("===================================")
	fmt.Println("----------------------------")
	DisplayMenu(Personnage{"Suad", "Chômeur", 3, 100, 50, []string{"Cv", "allocation"}, []string{"potion", "potion", "potion"}, 3, []string{"uzi", "mp5", "ermaemp", "ruby", "remington700"}})
}

func display(p1 Personnage) {
	fmt.Println("----------------")
	fmt.Println("Nom du joueur:", p1.Nom)
	fmt.Println("Classe:", p1.Classe)
	fmt.Println("Santé:", p1.Pv)
	fmt.Println("Niveau:", p1.Niveau)
	fmt.Println("Santé Max :", p1.Pvmax)
	fmt.Println("----------------")
	fmt.Println("♂☻♥♦♣♠◘○☺0◙♂♀♪♫↕‼¶§▬↨↑↓→←∟↔▲▼♀♫‼▲▼")
	DisplayMenu(p1)

}
func (p1 *Personnage) setPower() {
	p1.Niveau += 100
	fmt.Println("Power:", p1.Niveau)
}

func (p1 *Personnage) LessPower() {
	p1.Niveau -= 100
	fmt.Println("Power:", p1.Niveau)
}

func DisplayMenu(p1 Personnage) {
	potioncount := len(p1.inventairepotion)
	fmt.Println("==========================")
	fmt.Println(p1.Nom, "                    |")
	fmt.Println("Santé:", p1.Pv, "♥/", p1.Pvmax, "♥       |               ")
	fmt.Println("Potion(s) Restante(s):", potioncount, "|              ")
	fmt.Println("Niveau:", p1.Niveau, "               |")
	fmt.Println("==========================")
	fmt.Println("-----")
	fmt.Println("Menu|")
	fmt.Println("-----")
	fmt.Println("----------------------------")
	var i int8
	var h string
	fmt.Println("1. Afficher l'inventaire")
	fmt.Println("2. Afficher les stats du Personnage")
	fmt.Println("3. Vendre/Acheter")
	fmt.Println("4. Ameliorer le Personnage")
	fmt.Println("5. Quitter le jeu")
	fmt.Println("----------------------------")

	fmt.Println("►choisissez parmis ces options◄")

	fmt.Print("Type a number:")
	fmt.Scanln(&i)
	ans := (i)
	fmt.Println(ans)
	fmt.Println("Your number is:", i)
	switch i {
	case 1:
		AcessInventory(p1)
	case 2:
		display(p1)
	case 3:
		marchand(p1)
	case 4:
		p1.AddItem("potion")
		fmt.Println(p1.inventaire)
	case 5:
		fmt.Print("Voulez-vous vraiment quitter le jeu ?")
		fmt.Scanln(&h)
		switch h {
		case "oui":
			fmt.Println("Jeu Fermé")
		case "non":
			DisplayMenu(p1)

		default:
			fmt.Println("Vous n'avez pas choisi une option valide : Veuillez réessayer")
			DisplayMenu(p1)

		}
	}
}
func (p1 *Personnage) AddItem(item string) {
	p1.inventaire = append(p1.inventaire, item)
}

func AcessInventory(p1 Personnage) {
	var j int8
	fmt.Println("================================")
	fmt.Println("Inventaire:|")
	fmt.Println(p1.inventaire, "|")
	fmt.Println("================================")
	fmt.Println("--------------------------------")
	fmt.Println("choisissez parmis ces options")
	fmt.Println("1. Utiliser une potion de soin")
	fmt.Println("2. Quitter l'Inventaire")
	fmt.Println("----------------------------")
	fmt.Print("choississez une option: ")
	fmt.Scanln(&j)
	fmt.Println("===================================")
	switch j {
	case 1:
		removehealthpotion(p1)
	case 2:
		DisplayMenu(p1)
	case 3:
		removeinventory(p1)
	case 4:
		addhealthpotion(p1)
	default:
		fmt.Println("Vous n'avez pas choisi une option valide : Veuillez réessayer")
		AcessInventory(p1)

	}
}

func marchand(p1 Personnage) {
	//p1.inventairemarchand = []string {"uzi,mp5,ermaemp,ruby,remington700"}
	var i int8
	fmt.Println("1. Acheter")
	fmt.Println("2. Vendre")
	fmt.Println("3. Quitter")
	fmt.Println("choisissez parmis ces options")
	fmt.Print("Type a number: ")
	fmt.Scanln(&i)
	ans := (i)
	fmt.Println(ans)
	fmt.Println("Your number is:", i)
	switch i {
	case 1:
		acheter(p1)
	case 2:
		vendre(p1)
	case 3:
		DisplayMenu(p1)
	default:
		fmt.Println("Vous n'avez pas choisi une option valide : Veuillez réessayer")
		marchand(p1)
	}
}

func acheter(p1 Personnage) {
	fmt.Println("____________________________________")
	fmt.Println("|                                  |")
	fmt.Println("|=============ACHETER =============|")
	fmt.Println("|__________________________________|")
	fmt.Println("____________________________________")
	fmt.Println("Inventaire du marchand:|", p1.inventairemarchand, "|")
	if len(p1.inventairemarchand) == 0 {
		fmt.Println("Le marchand n'a plus rien à vendre")
		DisplayMenu(p1)
	} else {
		var t int8
		fmt.Println("1. Uzi")
		fmt.Println("2. Mp5")
		fmt.Println("3. Ermaemp")
		fmt.Println("4. Ruby")
		fmt.Println("5. Remington700")
		fmt.Println("6. Quitter")
		fmt.Println(len(p1.inventairemarchand))
		fmt.Println(p1.inventairemarchand)
		fmt.Print("Choississez le produit que vous souhaitez acheter")
		fmt.Scanln(&t)
		switch t {
		case 1:
			if len(p1.inventairemarchand) >= 1 {
				fmt.Println("Vous avez acheté un ", p1.inventairemarchand[:1])
				p1.inventaire = append(p1.inventaire, p1.inventairemarchand[:1]...)
				p1.inventairemarchand = append(p1.inventairemarchand[:0], p1.inventairemarchand[1:]...)
				fmt.Println(p1.inventairemarchand)

				fmt.Println(p1.inventaire)
				acheter(p1)
			} else {
				fmt.Println("Transaction Impossible")
				acheter(p1)
			}
		case 2:
			if len(p1.inventairemarchand) >= 2 {
				fmt.Println("Vous avez acheté un ", p1.inventairemarchand[1])
				p1.inventaire = append(p1.inventaire, p1.inventairemarchand[1])
				p1.inventairemarchand = append(p1.inventairemarchand[:1], p1.inventairemarchand[2:]...)
				fmt.Println(p1.inventairemarchand)
				fmt.Println(p1.inventaire)
				acheter(p1)
			} else {
				fmt.Println("Transaction Impossible")
				acheter(p1)
			}

		case 3:
			if len(p1.inventairemarchand) >= 3 {
				fmt.Println("Vous avez acheté un ", p1.inventairemarchand[2])
				p1.inventaire = append(p1.inventaire, p1.inventairemarchand[2])
				p1.inventairemarchand = append(p1.inventairemarchand[:2], p1.inventairemarchand[3:]...)
				fmt.Println(p1.inventairemarchand)
				fmt.Println(p1.inventaire)
				acheter(p1)
			} else {
				fmt.Println("Transaction Impossible")
				acheter(p1)
			}
		case 4:
			if len(p1.inventairemarchand) >= 4 {
				fmt.Println("Vous avez acheté un ", p1.inventairemarchand[3])
				p1.inventaire = append(p1.inventaire, p1.inventairemarchand[3])
				p1.inventairemarchand = append(p1.inventairemarchand[:3], p1.inventairemarchand[4:]...)
				fmt.Println(p1.inventairemarchand)
				fmt.Println(p1.inventaire)
				acheter(p1)
			} else {
				fmt.Println("Transaction Impossible")
				acheter(p1)
			}
		case 5:
			if len(p1.inventairemarchand) >= 5 {
				fmt.Println("Vous avez acheté un ", p1.inventairemarchand[4])
				p1.inventaire = append(p1.inventaire, p1.inventairemarchand[4])
				p1.inventairemarchand = append(p1.inventairemarchand[:4], p1.inventairemarchand[5:]...)
				fmt.Println(p1.inventairemarchand)
				acheter(p1)
			} else {
				fmt.Println("Transaction Impossible")
				acheter(p1)
			}
		case 6:
			marchand(p1)
		default:
			fmt.Println("Vous n'avez pas choisi une option valide : Veuillez réessayer")
			acheter(p1)
		}
	}
}
func vendre(p1 Personnage) {
	fmt.Println("____________________________________")
	fmt.Println("|                                  |")
	fmt.Println("|=============VENDRE ==============|")
	fmt.Println("|__________________________________|")
	fmt.Println("____________________________________")
	var u int8
	var y string
	if len(p1.inventaire) == 0 {
		fmt.Println("Vous n'avez plus rien à vendre")
		DisplayMenu(p1)
	} else {
		fmt.Println("Articles du Marchand:", p1.inventairemarchand)
		fmt.Println("Votre inventaire", p1.inventaire)
		fmt.Println("1. Uzi")
		fmt.Println("2. Mp5")
		fmt.Println("3. Ermaemp")
		fmt.Println("4. Ruby")
		fmt.Println("5. Remington700")
		fmt.Println("6. Quitter")
		fmt.Print("Choissiez l'article que vous souhaitez vendre")
		fmt.Scanln(&u)
		switch u {
		case 1:
			if len(p1.inventaire) >= 1 {
				if len(p1.inventaire) == 1 {
					fmt.Print("il ne vous reste que  1 article dans votre inventaire, Etes vous sur de vouloir de le  vendre ? votre article : ", p1.inventaire[0])
					fmt.Scanln(&y)
					switch y {
					case "non":
						vendre(p1)
						fmt.Println("Vente annulée")
					case "oui":
						fmt.Println("Vous avez vendu un ", p1.inventaire[0])
						p1.inventairemarchand = append(p1.inventairemarchand, p1.inventaire[:1]...)
						p1.inventaire = append(p1.inventaire[:0], p1.inventaire[1:]...)
						fmt.Println(p1.inventairemarchand)
						fmt.Println(p1.inventaire)
						vendre(p1)
					default:
						fmt.Println("Vous n'avez pas choisi une option valide : Veuillez réessayer")
						vendre(p1)
					}
				} else {
					fmt.Println("Vous avez vendu un ", p1.inventaire[:1])
					p1.inventairemarchand = append(p1.inventairemarchand, p1.inventaire[:1]...)
					p1.inventaire = append(p1.inventaire[:0], p1.inventaire[1:]...)
					fmt.Println(p1.inventairemarchand)
					fmt.Println(p1.inventaire)
					vendre(p1)
				}
			} else {
				fmt.Println("Transaction Impossible")
				vendre(p1)
			}
		case 2:
			if len(p1.inventaire) >= 2 {
				fmt.Println("Vous avez vendu un ", p1.inventaire[1])
				p1.inventairemarchand = append(p1.inventairemarchand, p1.inventaire[1])
				p1.inventaire = append(p1.inventaire[:1], p1.inventaire[2:]...)
				fmt.Println(p1.inventairemarchand)
				fmt.Println(p1.inventaire)
				vendre(p1)
			} else {
				fmt.Println("Transaction Impossible")
				vendre(p1)
			}

		case 3:
			if len(p1.inventaire) >= 3 {
				fmt.Println("Vous avez vendu un ", p1.inventaire[2])
				p1.inventairemarchand = append(p1.inventairemarchand, p1.inventaire[2])
				p1.inventaire = append(p1.inventaire[:2], p1.inventaire[3:]...)
				fmt.Println(p1.inventairemarchand)
				fmt.Println(p1.inventaire)
				vendre(p1)
			} else {
				fmt.Println("Transaction Impossible")
				vendre(p1)
			}
		case 4:
			if len(p1.inventaire) >= 4 {
				fmt.Println("Vous avez vendu un ", p1.inventaire[3])
				p1.inventairemarchand = append(p1.inventairemarchand, p1.inventaire[3])
				p1.inventaire = append(p1.inventaire[:3], p1.inventaire[4:]...)
				fmt.Println(p1.inventairemarchand)
				fmt.Println(p1.inventaire)
				vendre(p1)
			}
		case 5:
			if len(p1.inventaire) >= 5 {
				fmt.Println("Vous avez vendu un ", p1.inventaire[4])
				p1.inventairemarchand = append(p1.inventairemarchand, p1.inventaire[4])
				p1.inventaire = append(p1.inventaire[:4], p1.inventaire[5:]...)
			} else {
				fmt.Println("Transaction Impossible")
				vendre(p1)
			}
		case 6:
			marchand(p1)
		}
	}

}
func removeinventory(p1 Personnage) {
	p1.inventaire = append(p1.inventaire[:0], p1.inventaire[1:]...)
	fmt.Println(p1.inventaire)
	DisplayMenu(p1)
}

func removehealth(p1 Personnage) {
	p1.Pv -= 1
	fmt.Println(p1.Pv)
}

func addhealth(p1 Personnage) {
	p1.Pv += 1
	fmt.Println(p1.Pv)
}

func removehealthpotion(p1 Personnage) {
	potioncount := len(p1.inventairepotion)
	if len(p1.inventairepotion) == 0 {
		fmt.Println("Vous n'avez pas de potion")
		DisplayMenu(p1)
	} else {
		if p1.Pv == p1.Pvmax {
			fmt.Println("Votre vie est déja au maximum")
			fmt.Println("potion(s) restante(s)", potioncount)
			DisplayMenu(p1)
		}
		i := len(p1.inventairepotion)
		p1.inventairepotion = append(p1.inventairepotion[:0], p1.inventairepotion[:i-1]...)
		if p1.Pv+25 > p1.Pvmax {
			p1.Pv = p1.Pvmax
		} else {
			p1.Pv += 25
		}
		potioncount = potioncount - 1
		fmt.Println("vous avez consommé une potion de soin")

		fmt.Println("potion(s) restante(s)", potioncount)
		DisplayMenu(p1)

	}
}

func addhealthpotion(p1 Personnage) {
	if len(p1.inventairepotion) < p1.capacitéinventairepotion {
		p1.inventairepotion = append(p1.inventairepotion, "potion")
		fmt.Println(p1.inventairepotion)
	} else {
		fmt.Println("Vous ne pouvez pas avoir plus de 3 potions, pour cela améliorez votre inventaire")
	}
	DisplayMenu(p1)
}

func iswasted(p1 Personnage) { // fonction isdead()
	if p1.Pv <= 0 {
		fmt.Println("Vous êtes mort")
	}

}
