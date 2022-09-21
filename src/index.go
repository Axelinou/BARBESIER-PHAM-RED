package main

import (
	"fmt"
	"time"
)

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
	skill                    []string
	deathcount               int
	islearned                int
	marchandinventairesort   []string
	persoinventairesort      []string
}

func main() {
	var p1 Personnage
	var temp int = 0
	temp = temp + 0
	p1.Nom = "Suad"
	p1.Classe = "Chômeur"
	p1.Niveau = 3
	p1.Pvmax = 100

	p1.Pv = 50

	p1.inventaire = []string{"Cv", "allocation"}
	p1.inventairepotion = []string{"potion"}
	p1.capacitéinventairepotion = 3
	p1.inventairemarchand = []string{"uzi", "mp5", "ermaemp", "ruby", "remington700", "boule de feu"}

	p1.skill = []string{"Coup de poing"}
	p1.marchandinventairesort = []string{"boule de feu"}
	p1.islearned = 0
	p1.deathcount = 0
	p1.persoinventairesort = []string{}
	//fmt.Println(p1)
	//p1.display()
	//p1.setPower()
	//p1.LessPower()
	fmt.Println("-------------------------------------------------------------------------------")
	fmt.Println("===============================================================================")
	fmt.Println("                 •BIENVENUE SUR  A CHOMAGE ADVENTURE• ")
	fmt.Println("===============================================================================")
	fmt.Println("-------------------------------------------------------------------------------")
	DisplayMenu(Personnage{"Suad", "Chômeur", 3, 100, 50, []string{"Cv", "allocation"}, []string{"potion", "potion", "potion"}, 3, []string{"uzi", "mp5", "ermaemp", "ruby", "remington700"}, []string{"Coup de poing"}, 0, 0, []string{"boule de feu"}, []string{}})
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

//func (p1 *Personnage) setPower() {
//	p1.Niveau += 100
//	fmt.Println("Power:", p1.Niveau)
//}

//func (p1 *Personnage) LessPower() {
//	p1.Niveau -= 100
//	fmt.Println("Power:", p1.Niveau)
//}

func DisplayMenu(p1 Personnage) {

	var i int8
	var h string
	fmt.Println(i)
	fmt.Println(p1.deathcount)
	potioncount := len(p1.inventairepotion)
	fmt.Println(p1.Pv)
	fmt.Println("==========================")
	fmt.Println(p1.Nom, "                    |")
	fmt.Println("Santé:", p1.Pv, "♥/", p1.Pvmax, "♥       |               ")
	fmt.Println("Potion(s) Restante(s):", potioncount, "|              ")
	fmt.Println("Niveau:", p1.Niveau, "               |                 ")
	fmt.Println("Morts:", p1.deathcount, "                |                                                    ")
	fmt.Println("==========================")
	fmt.Println("-----")
	fmt.Println("Menu|")
	fmt.Println("-----")
	fmt.Println("----------------------------")
	fmt.Println("----------------------------")

	fmt.Println("1. Afficher l'inventaire")
	fmt.Println("2. Afficher les stats du Personnage")
	fmt.Println("3. Vendre/Acheter")
	fmt.Println("4. Ameliorer le Personnage")
	fmt.Println("5. Quitter le jeu")
	fmt.Println("----------------------------")
	fmt.Println("----------------------------")
	fmt.Println("►choisissez parmis ces options◄")

	//fmt.Print("Type a number:")
	fmt.Scanln(&i)
	fmt.Println("Your number is:", i)
	switch i {
	case 1:
		fmt.Println("Votre inventaire contient les objets suivants:")
		AcessInventory(p1)
	case 2:
		display(p1)

	case 3:
		marchand(p1)
	case 4:
		SpellBook(p1)
		//Poisonpot(p1)
		//removehealth(p1)
	case 5:
		fmt.Println("____________________________________")
		fmt.Print("Voulez-vous vraiment quitter le jeu ? ")

		fmt.Scanln(&h)
		switch h {
		case "oui":
			fmt.Println("====================================")
			fmt.Println("Jeu Fermé")
			fmt.Println("Merci d'avoir joué")
			fmt.Println("====================================")

		case "non":
			DisplayMenu(p1)
		default:
			fmt.Println("Vous n'avez pas choisi une option valide : Veuillez réessayer")
			DisplayMenu(p1)
		}
	default:
		fmt.Println("Vous n'avez pas choisi une option valide : Veuillez réessayer")
		DisplayMenu(p1)

	}
}

func (p1 *Personnage) AddItem(item string) {
	p1.inventaire = append(p1.inventaire, item)
}

func AcessInventory(p1 Personnage) {
	var j int8
	fmt.Println(p1.deathcount)

	fmt.Println("e")
	fmt.Println("================================")
	fmt.Println("Inventaire de", p1.Nom, "|")
	fmt.Println(p1.Pv)
	fmt.Println("potion(s) de soin:")
	fmt.Println(len(p1.inventairepotion))
	fmt.Println("Sorts:")
	fmt.Println(p1.persoinventairesort)
	fmt.Println("Objets:")
	fmt.Println(p1.inventaire, "|")
	fmt.Println("================================")
	fmt.Println("--------------------------------")
	fmt.Println("choisissez parmis ces options")
	fmt.Println("1. Utiliser une potion de soin")
	fmt.Println("2. Apprendre un sort")
	fmt.Println("3. Quitter l'Inventaire")
	fmt.Println("----------------------------")
	fmt.Print("choississez une option: ")
	fmt.Scanln(&j)
	fmt.Println("===================================")
	switch j {
	case 1:
		removehealthpotion(p1)
	case 2:
		if len(p1.persoinventairesort) == 0 {
			fmt.Println("Vous n'avez pas de sort à apprendre")
			AcessInventory(p1)
		}
		SpellBook(p1)
	case 3:
		DisplayMenu(p1)
	default:
		fmt.Println("Vous n'avez pas choisi une option valide : Veuillez réessayer")
		AcessInventory(p1)
		fmt.Println(p1.Pv)
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
		fmt.Println("6. Objet Magiques")
		fmt.Println("7. Quitter")
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
			buyperk(p1)
		case 7:
			marchand(p1)
		default:
			fmt.Println("Vous n'avez pas choisi une option valide : Veuillez réessayer")
			acheter(p1)
		}
	}
}
func buyperk(p1 Personnage) {
	if len(p1.marchandinventairesort) == 0 {
		fmt.Println("Le marchand n'a plus rien à vendre")
		DisplayMenu(p1)
	} else {
		var j int8
		fmt.Println("Marchand")
		fmt.Println("objets magiques à vendre:", p1.marchandinventairesort)
		fmt.Println("1. Boule de feu")
		fmt.Println("2. Boule de glace")
		fmt.Println("3. Boule de poison")
		fmt.Println("4. Boule de foudre")
		fmt.Println("5. Boule de terre")
		fmt.Println("6. Boule de vent")
		fmt.Println("7. Quitter")
		fmt.Print("Choississez l'objet que vous souhaitez acheter")
		fmt.Scanln(&j)
		switch j {
		case 1:
			if len(p1.inventairemarchand) >= 1 {
				fmt.Println("Vous avez acheté un ", p1.marchandinventairesort[:1])
				p1.persoinventairesort = append(p1.persoinventairesort, p1.marchandinventairesort[:1]...)
				p1.marchandinventairesort = append(p1.marchandinventairesort[:0], p1.marchandinventairesort[1:]...)
				fmt.Println(p1.marchandinventairesort)
				fmt.Println(p1.inventaire)

				fmt.Println(p1.persoinventairesort)
				buyperk(p1)
			} else {
				fmt.Println("Transaction Impossible")
				buyperk(p1)
			}
		case 7:
			acheter(p1)
		default:
			fmt.Println("Vous n'avez pas choisi une option valide : Veuillez réessayer")
			buyperk(p1)
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
	p1.Pv = p1.Pv - 50
	fmt.Println(p1.Pv)
	iswasted(p1)
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
		fmt.Println("????", p1.Pv)
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
			p1.Pv = p1.Pv + 25
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
	var z int8

	if p1.Pv <= 0 {
		fmt.Println("----------------------------------------------------------------------------")
		fmt.Println("_________________  \\'   /   _    _  _|_   _    _|  ________________________")
		fmt.Println("___________________  V V   (_|  _>   |_  (/_  (_|   _______________________")
		fmt.Println("__________                                                     __________")
		fmt.Println("_____________________              _|_                 ________________ ")
		fmt.Println("____________________                |                 ______________________")
		fmt.Println("----------------------------------------------------------------------------")
		fmt.Println("                             Vous êtes mort, Appuyez sur 1 pour recommencer")
		fmt.Println("                               Appuyez sur 2 pour ragequit")
		fmt.Print("Votre choix : ")
		fmt.Scan(&z)
		switch z {
		case 1:
			p1.deathcount += 1
			fmt.Println("Vous avez déjà mort", p1.deathcount, " fois")
			fmt.Println("------                              -------")
			fmt.Println("                 __ ")
			fmt.Println(" _________    __|  |__   ________           ")
			fmt.Println("  _______    |__    __|  _________        ")
			fmt.Println("                |__|                ")
			fmt.Println("------                              -------")
			p1.Pv = p1.Pvmax / 2
			fmt.Println("pv", p1.Pv)

			fmt.Println("      Vous avez ressuscité avec", p1.Pv, "♥")
			DisplayMenu(p1)

		case 2:
			fmt.Println("Jeu Fermé")
			fmt.Println("Vous avez quitté le jeu, Vous feriez mieux la prochaine fois")
		}
	} else {
		DisplayMenu(p1)
	}
}

func Poisonpot(p1 Personnage) {
	fmt.Println("Vous subissez les effets du poison")
	for i := 0; i < 3; i++ {
		fmt.Println("i", i)
		time.Sleep(1 * time.Second)
		p1.Pv = p1.Pv - 10
		fmt.Println("-10 Pv")
		if p1.Pv <= 0 {
			p1.Pv = 0
		}
		fmt.Println("vous avez ", p1.Pv, "/", p1.Pvmax, "PV")
	}

	iswasted(p1)

}

func SpellBook(p1 Personnage) {
	if p1.islearned == 0 {
		p1.skill = append(p1.skill, "Boule de feu")
		p1.persoinventairesort = append(p1.persoinventairesort[:0], p1.persoinventairesort[1:]...)
		fmt.Println("Vous avez appris le sort boule de feu ")
		p1.islearned = 1
		fmt.Println(p1.skill)
		DisplayMenu(p1)
	} else {
		fmt.Println("Vous avez déjà appris ce sort")
		DisplayMenu(p1)
	}
}
