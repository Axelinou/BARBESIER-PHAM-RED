package main

import (
	"fmt"
	"time"
)

type Personnage struct {
	Nom                         string
	Classe                      string
	Niveau                      int
	Pvmax                       int
	Pv                          int
	inventaire                  []string //Inventaire du joueur
	inventairepotion            []string // Inventaire des potions
	capacitéinventaireperso     int      // Capacité de l'inventaire du joueur s'augmente grace avec la fonction upgradeslot()
	capacitéinventairepotion    int
	estsurchargé                bool     //Renvoie true si l'inventaire est plein
	inventairemarchand          []string // Inventaire du marchand
	skill                       []string // Compétences apprises du joueur
	deathcount                  int      // Compteur de mort du joueur
	islearned                   int      // Renvoie 1 si le joueur a appris le sort
	coins                       int      // Nombre de pièces du joueur
	marchandinventairesort      []string // Inventaire du marchand pour les sorts
	persoinventairesort         []string // Inventaire du joueur pour les sorts
	marchandinventairemateriaux []string // Inventaire du marchand pour les matériaux
	persoinventairemateriaux    []string // Inventaire du joueur pour les matériaux
	armor                       int      // pts d' armure du joueur
	havehelmet                  bool
	havechestplate              bool
	haveleggings                bool
	haveboots                   bool
}

type Equipement struct {
	helmet            []string
	chest             []string
	leggings          []string
	boot              []string
	helmeteffect      []string
	helmetequiped     []string
	chestequiped      []string
	leggingsequiped   []string
	bootequiped       []string
	helmetinventory   []string
	chestinventory    []string
	leggingsinventory []string
	bootinventory     []string
	resistanceeffect  bool
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
	p1.capacitéinventaireperso = 10
	p1.inventairemarchand = []string{"potion de soin", "potion de poison", "uzi", "mp5", "ak47"}
	p1.estsurchargé = false
	p1.skill = []string{"Coup de poing"}
	p1.marchandinventairesort = []string{}
	p1.islearned = 0
	p1.deathcount = 0
	p1.coins = 100
	p1.persoinventairesort = []string{"livre de sort (boule de feu)"}
	p1.marchandinventairemateriaux = []string{"fourrure de loup", "peau de troll", "cuir de sanglier", "plume de corbeaux"}
	p1.persoinventairemateriaux = []string{}
	p1.armor = 0
	p1.havehelmet = false
	p1.havechestplate = false
	p1.haveleggings = false
	p1.haveboots = false
	var equipement1 Equipement
	equipement1.helmet = []string{}
	equipement1.chest = []string{}
	equipement1.leggings = []string{}
	equipement1.boot = []string{}
	equipement1.helmetinventory = []string{"ret"}
	equipement1.chestinventory = []string{}
	equipement1.leggingsinventory = []string{"fe"}
	equipement1.bootinventory = []string{"efdr"}

	//fmt.Println(p1)
	//p1.display()
	//p1.setPower()
	//p1.LessPower()
	fmt.Println("-------------------------------------------------------------------------------")
	fmt.Println("===============================================================================")
	fmt.Println("                 •BIENVENUE SUR  A CHOMAGE ADVENTURE• ")
	fmt.Println("===============================================================================")
	fmt.Println("-------------------------------------------------------------------------------")
	DisplayMenu(Personnage{"Suad", "Chômeur", 3, 100, 50, []string{"Cv", "allocation"}, []string{"potion", "potion", "potion"}, 10, 10, false, []string{"potion de soin", "potion de poison", "uzi", "mp5", "ak47"}, []string{"Coup de poing"}, 0, 0, 100, []string{"livre de sort (boule de feu)"}, []string{}, []string{"fourrure de loup", "peau de troll", "cuir de sanglier", "plume de corbeaux", "fourrure de warg"}, []string{}, 0, false, false, false, false}, Equipement{})
	//lance la fonction menu du jeu avec en argument les valeurs de la structure Personnage
}
func display(p1 Personnage, equipement1 Equipement) { //Affiche les stats du personnage

	fmt.Println("----------------")
	fmt.Println("Nom du joueur:", p1.Nom)
	fmt.Println("Classe:", p1.Classe)
	fmt.Println("Santé:", p1.Pv)
	fmt.Println("Niveau:", p1.Niveau)
	fmt.Println("Santé Max :", p1.Pvmax)
	fmt.Println("----------------")
	fmt.Println("♂☻♥♦♣♠◘○☺0◙♂♀♪♫↕‼¶§▬↨↑↓→←∟↔▲▼♀♫‼▲▼")
	DisplayMenu(p1, equipement1)
	//DisplayMenu(p1,Equipement{"bob de kevin","débardeur","short de foot","claquette-chaussette",[]string{},[]string{},[]string{},[]string{},[]string{},[]string{},[]string{},[]string{}})

}

//func (p1 *Personnage) setPower() {
//	p1.Niveau += 100
//	fmt.Println("Power:", p1.Niveau)
//}

//func (p1 *Personnage) LessPower() {
//	p1.Niveau -= 100
//	fmt.Println("Power:", p1.Niveau)
//}

func DisplayMenu(p1 Personnage, equipement1 Equipement) { //Affiche le menu du jeu

	var i int8
	fmt.Println("dazygqqshbcxgjrfidkvczfghsdjkRGFHTRGJHTUTHRFUJU", equipement1.helmetinventory)
	fmt.Println(p1.coins)
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
	fmt.Println("5. Salles des armures")
	fmt.Println("6. Quitter le jeu")
	fmt.Println("----------------------------")
	fmt.Println("----------------------------")
	fmt.Println("►choisissez parmis ces options◄")

	//fmt.Print("Type a number:")
	fmt.Scanln(&i)
	fmt.Println("Your number is:", i)
	switch i { //Switch permettant de choisir une option dans le menu
	case 1:
		fmt.Println("Votre inventaire contient les objets suivants:")
		AcessInventory(p1, Equipement{})
	case 2:
		display(p1, equipement1) //appel de la fonction display

	case 3:
		marchand(p1, equipement1) //appel de la fonction marchand
	case 4:
		//SpellBook(p1)
		Poisonpot(p1, equipement1)
		//removehealth(p1)
	case 5:
		Armor(equipement1, p1)

	case 6:
		Quitgame(p1, equipement1) //appel de la fonction permettant de quitter le jeu
	default: //Dans le cas ou  l'utilisateur rentre une valeur autre que 1,2,3,4 ou 5
		fmt.Println("Vous n'avez pas choisi une option valide : Veuillez réessayer")
		fmt.Println("=============================================================")
		DisplayMenu(p1, equipement1) //appel de la fonction DisplayMenu

	}
}

func (p1 *Personnage) AddItem(item string) {
	p1.inventaire = append(p1.inventaire, item)
}

func AcessInventory(p1 Personnage, equipement1 Equipement) { //Affiche l'inventaire du joueur et permet d'intéragir avec  un objet à utiliser tel qu'une potion ou un sort
	var j int8
	fmt.Println(p1.deathcount)
	fmt.Println("zedsft'erfd('rthfgcvhbjnklm,;:!", equipement1.helmetequiped)
	fmt.Println("e")
	fmt.Println("=========================")
	fmt.Println("Inventaire de", p1.Nom, "|")
	fmt.Println("=========================")
	fmt.Println(p1.Pv)
	fmt.Println("potion(s) de soin:")
	fmt.Println(len(p1.inventairepotion))
	fmt.Println("================================")
	fmt.Println("Sorts:")
	fmt.Println(p1.persoinventairesort)
	fmt.Println("================================")
	fmt.Println("Objets:")
	fmt.Println(p1.inventaire, len(p1.inventaire), "|")
	fmt.Println("================================")
	fmt.Println("--------------------------------")
	fmt.Println("choisissez parmis ces options")
	fmt.Println("1. Utiliser une potion de soin")
	fmt.Println("2. Apprendre un sort")
	fmt.Println("3. Plain")
	fmt.Println("4. Quitter l'Inventaire")
	fmt.Println("----------------------------")
	fmt.Print("choississez une option: ")
	fmt.Scanln(&j)
	fmt.Println("===================================")
	switch j {
	case 1:
		removehealthpotion(p1, equipement1) //appel de la fonction permettant d'utiliser une potion et de la retirer de l'inventaire
	case 2:
		if len(p1.persoinventairesort) == 0 {
			fmt.Println("Vous n'avez pas de sort à apprendre")
			AcessInventory(p1, equipement1)
		} else {
			SpellBook(p1, equipement1) //appel de la fonction permettant d'utiliser un sort et de le retirer de l'inventaire
		}

	case 3:
		fmt.Println("C'est une fonctionnalité en cours de développement")
	case 4:
		DisplayMenu(p1, equipement1)
	case 5:
		charCreation(p1, equipement1)
	default:
		fmt.Println("=============================================================")
		fmt.Println("Vous n'avez pas choisi une option valide : Veuillez réessayer")
		fmt.Println("=============================================================")
	}

}

func marchand(p1 Personnage, equipement1 Equipement) {
	//p1.inventairemarchand = []string {"uzi,mp5,ermaemp,ruby,remington700"}
	var i int8
	fmt.Println("================================")
	fmt.Println("choisissez parmis ces options")
	fmt.Println("================================")
	fmt.Println("1. Acheter")
	fmt.Println("2. Vendre")
	fmt.Println("3. Forger un objet")
	fmt.Println("4. Quitter")
	fmt.Println("================================")
	fmt.Print("Entrez un nombre: ")
	fmt.Scanln(&i)
	ans := (i)
	fmt.Println(ans)
	fmt.Println("Your number is:", i)
	switch i {
	case 1:
		acheter(p1, equipement1) //appel de la fonction permettant d'acheter un objet
	case 2:
		vendre(p1, equipement1) //appel de la fonction permettant de vendre un objet
	case 3:
		forge(p1, equipement1) //appel de la fonction permettant de forger un objet
	case 4:
		DisplayMenu(p1, equipement1) //retour au menu principal
	default: //dans les autres cas
		fmt.Println("=============================================================")
		fmt.Println("Vous n'avez pas choisi une option valide : Veuillez réessayer")
		fmt.Println("=============================================================")
		marchand(p1, equipement1) //retour au menu du marchand
	}
}

func acheter(p1 Personnage, equipement1 Equipement) {

	inventorylimit(p1, equipement1)
	fmt.Println("BEUUUUUUUUUUUUUUUUURKKKKKKKKKKKKKKKKKKKKK", len(p1.inventaire))
	fmt.Println("____________________________________")
	fmt.Println("|                                  |")
	fmt.Println("|=============ACHETER =============|")
	fmt.Println("|__________________________________|")
	fmt.Println("____________________________________")
	fmt.Println("Inventaire du marchand:|", p1.inventairemarchand, "|")
	if len(p1.inventairemarchand) == 0 { //si l'inventaire du marchand est vide
		fmt.Println("Le marchand n'a plus rien à vendre")
		DisplayMenu(p1, equipement1) //retour au menu principal

	} else {
		var t int8
		var cout int
		var cointemp int
		var nocoin bool

		t = 0
		fmt.Println("adzsetrer", t)
		fmt.Println("choisissez parmis ces options")
		fmt.Println("____________________________________")
		fmt.Println(" ACHAT:")
		fmt.Println("1. 1er objet")
		fmt.Println("2. Uzi")
		fmt.Println("3. Mp5")
		fmt.Println("4. Ruby")
		fmt.Println("5. Remington700")
		fmt.Println("____________________________________")
		fmt.Println("AUTRES:")
		fmt.Println("6. Objet Magiques")
		fmt.Println("7. Matériaux")
		fmt.Println("8. Quitter")
		fmt.Println(len(p1.inventairemarchand))
		fmt.Println(p1.inventairemarchand)
		fmt.Print("Choississez le produit que vous souhaitez acheter")
		fmt.Scanln(&t)
		switch t {

		case 1:
			//fmt.Println(p1.inventairemarchand[0] == "potion de soin")
			cointemp = p1.coins //on stocke le nombre de pièces du joueur dans une variable temporaire
			fmt.Println("Your number is:", t)
			if len(p1.inventairemarchand) >= 1 { //si l'inventaire du marchand contient au moins 1 objet( permet de gerer l'erreur "index out of range")

				if p1.inventairemarchand[0] == "potion de soin" { //si l'objet acheté est une potion de soin
					if p1.coins >= 3 { //si le joueur a assez de pièces
						p1.coins = p1.coins - 3 //on retire 3 coins au joueur
						fmt.Println("temp", cointemp)
						cout = cointemp - p1.coins //la valeur de cout est égale à la valeur de cointemp - la valeur de coins
						fmt.Print(p1.coins)
					} else {
						nocoin = true
						fmt.Println(nocoin)

					}
				}

				if p1.inventairemarchand[0] == "potion de poison" { //meme fonctionnement
					if p1.coins >= 6 {
						p1.coins = p1.coins - 6
						fmt.Println("temp", cointemp)
						cout = cointemp - p1.coins
						fmt.Print(p1.coins)
					} else {

						nocoin = true

					}
				}
				if p1.inventairemarchand[0] == "uzi" {
					if p1.coins >= 25 {
						p1.coins = p1.coins - 25
						fmt.Println("temp", cointemp)
						cout = cointemp - p1.coins
						fmt.Print(p1.coins)
					} else {

						nocoin = true

					}
				}
				if p1.inventairemarchand[0] == "mp5" {
					if p1.coins >= 30 {
						p1.coins = p1.coins - 30
						fmt.Println("temp", cointemp)
						cout = cointemp - p1.coins
						fmt.Print(p1.coins)
					} else {

						nocoin = true

					}
				}
				if p1.inventairemarchand[0] == "Cv" {
					if p1.coins >= 30 {
						p1.coins = p1.coins - 30
						fmt.Println("temp", cointemp)
						cout = cointemp - p1.coins
						fmt.Print(p1.coins)
					} else {

						nocoin = true

					}
				}
				if p1.inventairemarchand[0] == "allocation" {
					if p1.coins >= 30 {
						p1.coins = p1.coins - 30
						fmt.Println("temp", cointemp)
						cout = cointemp - p1.coins
						fmt.Print(p1.coins)
					} else {

						nocoin = true

					}
				}
				if p1.inventairemarchand[0] == "ak47" {
					if p1.coins >= 30 {
						p1.coins = p1.coins - 30
						fmt.Println("temp", cointemp)
						cout = cointemp - p1.coins
						fmt.Print(p1.coins)
					} else {

						nocoin = true

					}
				}
				if nocoin == true {
					fmt.Println(("Vous n'avez pas assez de pièces"))
					marchand(p1, equipement1)
					break
				} else {
					fmt.Println("Vous avez acheté un ", p1.inventairemarchand[:1], "pour un cout de ,", cout, "d or, il vous reste ", p1.coins, " pièces d or")
					p1.inventaire = append(p1.inventaire, p1.inventairemarchand[:1]...)                     //on ajoute l'objet acheté dans l'inventaire du joueur
					p1.inventairemarchand = append(p1.inventairemarchand[:0], p1.inventairemarchand[1:]...) //on supprime l'objet acheté de l'inventaire du marchand le principe est le meme pour la vente d'objets
					fmt.Println(p1.inventairemarchand)
					t = 0
					fmt.Println("adzsetrer", t)
					fmt.Println(p1.inventaire)

					acheter(p1, equipement1)
				}

			} else {
				fmt.Println("Transaction Impossible") //condition gerant l'execption "index out of range"
				acheter(p1, equipement1)
			}

		case 2: //meme fonctionnement mais pour l'objet a la deuxième place dans l'inventaire du marchand
			cointemp = p1.coins
			if len(p1.inventairemarchand) >= 2 {
				fmt.Print("Vous avez acheté un", p1.inventairemarchand[1])
				if p1.inventairemarchand[1] == "potion de soin" {
					if p1.coins >= 30 {
						p1.coins = p1.coins - 3
						fmt.Println("temp", cointemp)
						cout = cointemp - p1.coins
						fmt.Print(p1.coins)
					} else {
						fmt.Println("Vous n'avez pas assez de pièces")
						nocoin = true

					}
				}
				if p1.inventairemarchand[1] == "potion de poison" {
					if p1.coins >= 30 {
						p1.coins = p1.coins - 6
						fmt.Println("temp", cointemp)
						cout = cointemp - p1.coins
						fmt.Print(p1.coins)
					} else {
						nocoin = true

					}
				}
				if p1.inventairemarchand[1] == "uzi" {
					if p1.coins >= 30 {
						p1.coins = p1.coins - 25
						fmt.Println("temp", cointemp)
						cout = cointemp - p1.coins
						fmt.Print(p1.coins)
					} else {
						nocoin = true

					}
				}
				if p1.inventairemarchand[1] == "mp5" {
					if p1.coins >= 30 {
						p1.coins = p1.coins - 30
						fmt.Println("temp", cointemp)
						cout = cointemp - p1.coins
						fmt.Print(p1.coins)
					} else {
						nocoin = true

					}
				}
				if p1.inventairemarchand[1] == "Cv" {
					if p1.coins >= 30 {
						p1.coins = p1.coins - 30
						fmt.Println("temp", cointemp)
						cout = cointemp - p1.coins
						fmt.Print(p1.coins)
					} else {
						nocoin = true

					}
				}
				if p1.inventairemarchand[1] == "allocation" {
					if p1.coins >= 30 {
						p1.coins = p1.coins - 30
						fmt.Println("temp", cointemp)
						cout = cointemp - p1.coins
						fmt.Print(p1.coins)
					} else {
						nocoin = true

					}
				}
				if p1.inventairemarchand[1] == "ak47" {
					if p1.coins >= 30 {
						p1.coins = p1.coins - 30
						fmt.Println("temp", cointemp)
						cout = cointemp - p1.coins
						fmt.Print(p1.coins)
					} else {
						nocoin = true

					}
				}
				if nocoin == true {
					fmt.Println(("Vous n'avez pas assez de pièces"))
					marchand(p1, equipement1)
					break
				} else {
					fmt.Println("Vous avez acheté un ", p1.inventairemarchand[1], "pour un cout de", cout, " pièces d or, il vous reste ", p1.coins, " pièces d or")
					fmt.Println("Vous avez acheté un ", p1.inventairemarchand[1])
					p1.inventaire = append(p1.inventaire, p1.inventairemarchand[1])
					p1.inventairemarchand = append(p1.inventairemarchand[:1], p1.inventairemarchand[2:]...)
					fmt.Println(p1.inventairemarchand)
					fmt.Println(p1.inventaire)
					acheter(p1, equipement1)
				}
			} else {
				fmt.Println("Transaction Impossible")
				acheter(p1, equipement1)
			}

			//idem mais pour le 3ème objet
		case 3:
			cointemp = p1.coins
			if len(p1.inventairemarchand) >= 3 {

				fmt.Print("Vous avez acheté un", p1.inventairemarchand[1])
				if p1.inventairemarchand[2] == "potion de soin" {
					if p1.coins >= 3 {
						p1.coins = p1.coins - 3
						fmt.Println("temp", cointemp)
						cout = cointemp - p1.coins
						fmt.Print(p1.coins)
					} else {
						nocoin = true

					}
				}

				if p1.inventairemarchand[2] == "potion de poison" {
					if p1.coins >= 6 {
						p1.coins = p1.coins - 6
						fmt.Println("temp", cointemp)
						cout = cointemp - p1.coins
						fmt.Print(p1.coins)
					} else {
						nocoin = true

					}
				}
				if p1.inventairemarchand[2] == "uzi" {
					if p1.coins >= 25 {
						p1.coins = p1.coins - 25
						fmt.Println("temp", cointemp)
						cout = cointemp - p1.coins
						fmt.Print(p1.coins)
					} else {
						nocoin = true

					}
				}
				if p1.inventairemarchand[2] == "mp5" {
					if p1.coins >= 30 {
						p1.coins = p1.coins - 30
						fmt.Println("temp", cointemp)
						cout = cointemp - p1.coins
						fmt.Print(p1.coins)
					} else {
						nocoin = true

					}
				}

				if p1.inventairemarchand[2] == "Cv" {
					if p1.coins >= 30 {
						p1.coins = p1.coins - 30
						fmt.Println("temp", cointemp)
						cout = cointemp - p1.coins
						fmt.Print(p1.coins)
					} else {
						nocoin = true

					}
				}
				if p1.inventairemarchand[2] == "allocation" {
					if p1.coins >= 30 {
						p1.coins = p1.coins - 30
						fmt.Println("temp", cointemp)
						cout = cointemp - p1.coins
						fmt.Print(p1.coins)
					} else {
						nocoin = true
					}
				}
				if p1.inventairemarchand[2] == "ak47" {
					if p1.coins >= 30 {
						p1.coins = p1.coins - 30
						fmt.Println("temp", cointemp)
						cout = cointemp - p1.coins
						fmt.Print(p1.coins)
					} else {
						nocoin = true

					}
				}
				if nocoin == true {
					fmt.Println(("Vous n'avez pas assez de pièces"))
					marchand(p1, equipement1)
					break
				} else {
					fmt.Println("Vous avez acheté un ", p1.inventairemarchand[2], "pour un cout de ", cout, "pièces d or, il vous reste ", p1.coins, " pièces d or")
					fmt.Println("Vous avez acheté un ", p1.inventairemarchand[2])
					p1.inventaire = append(p1.inventaire, p1.inventairemarchand[2])
					p1.inventairemarchand = append(p1.inventairemarchand[:2], p1.inventairemarchand[3:]...)
					fmt.Println(p1.inventairemarchand)
					fmt.Println(p1.inventaire)
					acheter(p1, equipement1)
				}
			} else {
				fmt.Println("Transaction Impossible")
				acheter(p1, equipement1)
			}
		case 4:
			cointemp = p1.coins
			if len(p1.inventairemarchand) >= 4 {
				if p1.coins >= 3 {
					fmt.Print("Vous avez acheté un", p1.inventairemarchand[1])
					if p1.inventairemarchand[3] == "potion de soin" {
						p1.coins = p1.coins - 3
						fmt.Println("temp", cointemp)
						cout = cointemp - p1.coins
						fmt.Print(p1.coins)
					} else {
						nocoin = true

					}
				}

				if p1.inventairemarchand[3] == "potion de poison" {
					if p1.coins >= 6 {
						p1.coins = p1.coins - 6
						fmt.Println("temp", cointemp)
						cout = cointemp - p1.coins
						fmt.Print(p1.coins)
					} else {
						nocoin = true
					}
				}
				if p1.inventairemarchand[3] == "uzi" {
					if p1.coins >= 25 {
						p1.coins = p1.coins - 25
						fmt.Println("temp", cointemp)
						cout = cointemp - p1.coins
						fmt.Print(p1.coins)
					} else {
						nocoin = true

					}
				}

				if p1.inventairemarchand[3] == "mp5" {
					if p1.coins >= 30 {
						p1.coins = p1.coins - 30
						fmt.Println("temp", cointemp)
						cout = cointemp - p1.coins
						fmt.Print(p1.coins)
					} else {
						nocoin = true

					}
				}

				if p1.inventairemarchand[3] == "Cv" {
					if p1.coins >= 30 {
						p1.coins = p1.coins - 30
						fmt.Println("temp", cointemp)
						cout = cointemp - p1.coins
						fmt.Print(p1.coins)
					} else {
						nocoin = true
					}
				}
				if p1.inventairemarchand[3] == "allocation" {
					if p1.coins >= 30 {
						p1.coins = p1.coins - 30
						fmt.Println("temp", cointemp)
						cout = cointemp - p1.coins
						fmt.Print(p1.coins)
					} else {
						nocoin = true

					}
				}
				if p1.inventairemarchand[3] == "ak47" {
					if p1.coins >= 30 {
						p1.coins = p1.coins - 30
						fmt.Println("temp", cointemp)
						cout = cointemp - p1.coins
						fmt.Print(p1.coins)
					} else {
						nocoin = true

					}
				}

				if nocoin == true {
					fmt.Println(("Vous n'avez pas assez de pièces"))
					marchand(p1, equipement1)
					break
				} else {
					fmt.Println("Vous avez acheté un ", p1.inventairemarchand[3], "pour un cout de ", cout, "pièces d'or, il vous reste ", p1.coins, " pièces d'or")
					fmt.Println("Vous avez acheté un ", p1.inventairemarchand[3])
					p1.inventaire = append(p1.inventaire, p1.inventairemarchand[3])
					p1.inventairemarchand = append(p1.inventairemarchand[:3], p1.inventairemarchand[4:]...)
					fmt.Println(p1.inventairemarchand)
					fmt.Println(p1.inventaire)
					acheter(p1, equipement1)
				}
			} else {
				fmt.Println("Transaction Impossible")
				acheter(p1, equipement1)
			}

		case 5:
			cointemp = p1.coins
			if len(p1.inventairemarchand) >= 5 {
				fmt.Print("Vous avez acheté un", p1.inventairemarchand[1])
				if p1.inventairemarchand[4] == "potion de soin" {
					if p1.coins >= 3 {
						p1.coins = p1.coins - 3
						fmt.Println("temp", cointemp)
						cout = cointemp - p1.coins
						fmt.Print(p1.coins)
					} else {
						nocoin = true

					}
				}

				if p1.inventairemarchand[4] == "potion de poison" {
					if p1.coins >= 6 {
						p1.coins = p1.coins - 6
						fmt.Println("temp", cointemp)
						cout = cointemp - p1.coins
						fmt.Print(p1.coins)
					} else {
						nocoin = true

					}
				}
				if p1.inventairemarchand[4] == "uzi" {
					if p1.coins >= 25 {
						p1.coins = p1.coins - 25
						fmt.Println("temp", cointemp)
						cout = cointemp - p1.coins
						fmt.Print(p1.coins)
					} else {
						nocoin = true

					}
				}
				if p1.inventairemarchand[4] == "mp5" {
					if p1.coins >= 30 {
						p1.coins = p1.coins - 30
						fmt.Println("temp", cointemp)
						cout = cointemp - p1.coins
						fmt.Print(p1.coins)
					} else {
						nocoin = true

					}
				}

				if p1.inventairemarchand[4] == "Cv" {
					if p1.coins >= 30 {
						p1.coins = p1.coins - 30
						fmt.Println("temp", cointemp)
						cout = cointemp - p1.coins
						fmt.Print(p1.coins)
					} else {
						nocoin = true

					}
				}
				if p1.inventairemarchand[4] == "allocation" {
					if p1.coins >= 30 {
						p1.coins = p1.coins - 30
						fmt.Println("temp", cointemp)
						cout = cointemp - p1.coins
						fmt.Print(p1.coins)
					} else {
						nocoin = true

					}
				}
				if p1.inventairemarchand[4] == "ak47" {
					if p1.coins >= 30 {
						p1.coins = p1.coins - 30
						fmt.Println("temp", cointemp)
						cout = cointemp - p1.coins
						fmt.Print(p1.coins)
					} else {
						nocoin = true

					}
				}

				if nocoin == true {
					fmt.Println(("Vous n'avez pas assez de pièces"))
					marchand(p1, equipement1)
					break
				} else {
					fmt.Println("Vous avez acheté un ", p1.inventairemarchand[4], "pour un cout de ", cout, "pièces d'or, il vous reste ", p1.coins, " pièces d'or")

					fmt.Println("Vous avez acheté un ", p1.inventairemarchand[4])
					p1.inventaire = append(p1.inventaire, p1.inventairemarchand[4])
					p1.inventairemarchand = append(p1.inventairemarchand[:4], p1.inventairemarchand[5:]...)
					fmt.Println(p1.inventairemarchand)
					acheter(p1, equipement1)
				}
			} else {

				fmt.Println("Transaction Impossible")
				acheter(p1, equipement1)
			}
		case 6:
			buyperk(p1, equipement1) //appel de la fonction d'achat d'objet magique
		case 7:
			materials(p1, equipement1)
		case 8:
			marchand(p1, equipement1)
		default:
			fmt.Println("Vous n'avez pas choisi une option valide : Veuillez réessayer")
			acheter(p1, equipement1)

			break
		}
	}
}

func buyperk(p1 Personnage, equipement1 Equipement) {

	if len(p1.marchandinventairesort) == 0 { //si le marchand n'a plus de sort

		fmt.Println("Le marchand n'a plus rien à vendre")
		DisplayMenu(p1, equipement1)
	} else {
		var cout int
		var cointemp int
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
		case 1: // meme principe que pour l'achat des objets normaux
			cointemp = p1.coins
			if len(p1.inventairemarchand) >= 1 {
				if p1.marchandinventairesort[0] == "livre de sort (boule de feu)" {
					p1.coins = p1.coins - 30
					fmt.Println("temp", cointemp)
					cout = cointemp - p1.coins
					fmt.Print(p1.coins)
				}
				fmt.Println("Vous avez acheté un ", p1.marchandinventairesort[0], "pour un cout de ", cout, "pièces d or, il vous reste ", p1.coins, " pièces d or")
				fmt.Println("Vous avez acheté un ", p1.marchandinventairesort[:1])
				p1.persoinventairesort = append(p1.persoinventairesort, p1.marchandinventairesort[:1]...)
				p1.marchandinventairesort = append(p1.marchandinventairesort[:0], p1.marchandinventairesort[1:]...)
				fmt.Println(p1.marchandinventairesort)
				fmt.Println(p1.inventaire)

				fmt.Println(p1.persoinventairesort)
				buyperk(p1, equipement1)
			} else {
				fmt.Println("Transaction Impossible")
				buyperk(p1, equipement1)
			}
		case 7:
			acheter(p1, equipement1)
		default:
			fmt.Println("Vous n'avez pas choisi une option valide : Veuillez réessayer")
			buyperk(p1, equipement1)
		}
	}
}
func materials(p1 Personnage, equipement1 Equipement) {
	var h int //variable pour le choix
	var cointemp int
	var cout int
	if len(p1.marchandinventairemateriaux) == 0 { //si l'inventaire du marchand est vide
		fmt.Println("Le marchand n'a plus rien à vendre")
		DisplayMenu(p1, equipement1)
	} else {
		fmt.Println("Inventaire du marchand:|", p1.marchandinventairemateriaux)
		fmt.Println("1.loup")
		fmt.Println("2.troll")
		fmt.Println("3.sanglier")
		fmt.Println("4.plume de corbeau")
		fmt.Println("5.quitter")
		fmt.Print("Choisissez votre action  : ")
		fmt.Scanln(&h)

		switch h {
		case 1:
			cointemp = p1.coins
			if len(p1.marchandinventairemateriaux) >= 1 {
				fmt.Print("Vous avez acheté un", p1.marchandinventairemateriaux[0])
				if p1.marchandinventairemateriaux[0] == "fourrure de loup" {
					p1.coins = p1.coins - 3
					fmt.Println("temp", cointemp)
					cout = cointemp - p1.coins
					fmt.Print(p1.coins)
				}

				if p1.marchandinventairemateriaux[0] == "peau de troll" {
					p1.coins = p1.coins - 45
					fmt.Println("temp", cointemp)
					cout = cointemp - p1.coins
					fmt.Print(p1.coins)
				}
				//fmt.Println(p1.marchandinventairemateriaux[0] == "cuir de sanglier")
				if p1.marchandinventairemateriaux[0] == "cuir de sanglier" {
					p1.coins = p1.coins - 15
					fmt.Println("temp", cointemp)
					cout = cointemp - p1.coins
					fmt.Print(p1.coins)
				}
				//fmt.Println(p1.marchandinventairemateriaux[0] == "cuir de sanglier")
				if p1.marchandinventairemateriaux[0] == "plume de corbeaux" {
					p1.coins = p1.coins - 9
					fmt.Println("temp", cointemp)
					cout = cointemp - p1.coins
					fmt.Print(p1.coins)
				}
				fmt.Println("Vous avez acheté un ", p1.marchandinventairemateriaux[0], "pour un cout de", cout, " pièces d or, il vous reste ", p1.coins, " pièces d or")
				p1.persoinventairemateriaux = append(p1.persoinventairemateriaux, p1.marchandinventairemateriaux[0])
				p1.marchandinventairemateriaux = append(p1.marchandinventairemateriaux[:0], p1.marchandinventairemateriaux[1:]...)
				fmt.Println(p1.marchandinventairemateriaux)
				fmt.Println(p1.persoinventairemateriaux)
				materials(p1, equipement1)
			} else {
				fmt.Println("Transaction Impossible")
				materials(p1, equipement1)
			}
		case 2:
			cointemp = p1.coins
			if len(p1.marchandinventairemateriaux) >= 2 {
				fmt.Print("Vous avez acheté un", p1.marchandinventairemateriaux[0])
				if p1.marchandinventairemateriaux[1] == "fourrure de loup" {
					p1.coins = p1.coins - 3
					fmt.Println("temp", cointemp)
					cout = cointemp - p1.coins
					fmt.Print(p1.coins)
				}

				if p1.marchandinventairemateriaux[1] == "peau de troll" {
					p1.coins = p1.coins - 45
					fmt.Println("temp", cointemp)
					cout = cointemp - p1.coins
					fmt.Print(p1.coins)
				}

				if p1.marchandinventairemateriaux[1] == "cuir de sanglier" {
					p1.coins = p1.coins - 15
					fmt.Println("temp", cointemp)
					cout = cointemp - p1.coins
					fmt.Print(p1.coins)
				}
				if p1.marchandinventairemateriaux[1] == "plume de corbeaux" {
					p1.coins = p1.coins - 9
					fmt.Println("temp", cointemp)
					cout = cointemp - p1.coins
					fmt.Print(p1.coins)
				}
				fmt.Println("Vous avez acheté un ", p1.marchandinventairemateriaux[1], "pour un cout de", cout, " pièces d or, il vous reste ", p1.coins, " pièces d or")
				p1.persoinventairemateriaux = append(p1.persoinventairemateriaux, p1.marchandinventairemateriaux[1])
				p1.marchandinventairemateriaux = append(p1.marchandinventairemateriaux[:1], p1.marchandinventairemateriaux[2:]...)
				fmt.Println(p1.marchandinventairemateriaux)
				fmt.Println(p1.persoinventairemateriaux)
				materials(p1, equipement1)
			} else {
				fmt.Println("Transaction Impossible")
				materials(p1, equipement1)
			}
		case 3:
			cointemp = p1.coins
			if len(p1.marchandinventairemateriaux) >= 3 {
				fmt.Print("Vous avez acheté un", p1.marchandinventairemateriaux[0])
				if p1.marchandinventairemateriaux[2] == "fourrure de loup" {
					p1.coins = p1.coins - 3
					fmt.Println("temp", cointemp)
					cout = cointemp - p1.coins
					fmt.Print(p1.coins)
				}

				if p1.marchandinventairemateriaux[2] == "peau de troll" {
					p1.coins = p1.coins - 45
					fmt.Println("temp", cointemp)
					cout = cointemp - p1.coins
					fmt.Print(p1.coins)
				}

				if p1.marchandinventairemateriaux[2] == "cuir de sanglier" {
					p1.coins = p1.coins - 15
					fmt.Println("temp", cointemp)
					cout = cointemp - p1.coins
					fmt.Print(p1.coins)
				}
				if p1.marchandinventairemateriaux[2] == "plume de corbeaux" {
					p1.coins = p1.coins - 9
					fmt.Println("temp", cointemp)
					cout = cointemp - p1.coins
					fmt.Print(p1.coins)
				}
				fmt.Println("Vous avez acheté un ", p1.marchandinventairemateriaux[2], "pour un cout de", cout, " pièces d or, il vous reste ", p1.coins, " pièces d or")
				p1.persoinventairemateriaux = append(p1.persoinventairemateriaux, p1.marchandinventairemateriaux[2])
				p1.marchandinventairemateriaux = append(p1.marchandinventairemateriaux[:2], p1.marchandinventairemateriaux[3:]...)
				fmt.Println(p1.marchandinventairemateriaux)
				fmt.Println(p1.persoinventairemateriaux)
				materials(p1, equipement1)
			} else {
				fmt.Println("Transaction Impossible")
				materials(p1, equipement1)
			}
		case 4:
			cointemp = p1.coins
			if len(p1.marchandinventairemateriaux) >= 4 {
				fmt.Print("Vous avez acheté un", p1.marchandinventairemateriaux[0])
				if p1.marchandinventairemateriaux[3] == "fourrure de loup" {
					p1.coins = p1.coins - 3
					fmt.Println("temp", cointemp)
					cout = cointemp - p1.coins
					fmt.Print(p1.coins)
				}

				if p1.marchandinventairemateriaux[3] == "peau de troll" {
					p1.coins = p1.coins - 45
					fmt.Println("temp", cointemp)
					cout = cointemp - p1.coins
					fmt.Print(p1.coins)
				}

				if p1.marchandinventairemateriaux[3] == "cuir de sanglier" {
					p1.coins = p1.coins - 15
					fmt.Println("temp", cointemp)
					cout = cointemp - p1.coins
					fmt.Print(p1.coins)
				}
				if p1.marchandinventairemateriaux[3] == "plume de corbeaux" {
					p1.coins = p1.coins - 9
					fmt.Println("temp", cointemp)
					cout = cointemp - p1.coins
					fmt.Print(p1.coins)
				}
				fmt.Println("Vous avez acheté un ", p1.marchandinventairemateriaux[3], "pour un cout de", cout, " pièces d or, il vous reste ", p1.coins, " pièces d or")
				p1.persoinventairemateriaux = append(p1.persoinventairemateriaux, p1.marchandinventairemateriaux[3])
				p1.marchandinventairemateriaux = append(p1.marchandinventairemateriaux[:3], p1.marchandinventairemateriaux[4:]...)
				fmt.Println(p1.marchandinventairemateriaux)
				fmt.Println(p1.persoinventairemateriaux)
				materials(p1, equipement1)
			} else {
				fmt.Println("Transaction Impossible")
				materials(p1, equipement1)
			}
		case 5:
			marchand(p1, equipement1)
		default:
			fmt.Println("Vous n'avez pas choisi une action valide")
		}

	}
}

func vendre(p1 Personnage, equipement1 Equipement) {
	var cout int
	fmt.Println("____________________________________")
	fmt.Println("|                                  |")
	fmt.Println("|=============VENDRE ==============|")
	fmt.Println("|__________________________________|")
	fmt.Println("____________________________________")
	var u int8
	var y string
	var cointemp int
	if len(p1.inventaire) == 0 {
		fmt.Println("Vous n'avez plus rien à vendre")
		DisplayMenu(p1, equipement1)
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
			cointemp = p1.coins
			if len(p1.inventaire) >= 1 {
				if len(p1.inventaire) == 1 { //avertit le joueur qu'il ne lui reste que 1 objet dans son inventaire
					fmt.Print("il ne vous reste que  1 article dans votre inventaire, Etes vous sur de vouloir de le  vendre ? votre article : ", p1.inventaire[0])
					fmt.Scanln(&y)
					switch y {
					case "non":
						fmt.Println("Vente annulée")
						vendre(p1, equipement1)
						//si le joueur ne veut pas vendre son objet, il est redirigé vers le menu
					case "oui":
						fmt.Println("Vous avez vendu un ", p1.inventaire[0])
						if p1.inventaire[0] == "Cv" {
							p1.coins = p1.coins + 35
							fmt.Println(p1.coins)
							fmt.Println("temp", cointemp)
							cout = p1.coins - cointemp
							fmt.Print(p1.coins)
						}
						if p1.inventaire[0] == "potion de soin" { //si l'objet acheté est une potion de soin
							p1.coins = p1.coins + 3 //on retire 3 coins au joueur
							fmt.Println("temp", cointemp)
							cout = p1.coins - cointemp
							fmt.Print(p1.coins)

						}

						if p1.inventaire[0] == "potion de poison" { //meme fonctionnement
							p1.coins = p1.coins + 3
							fmt.Println("temp", cointemp)
							cout = p1.coins - cointemp
							fmt.Print(p1.coins)
						}
						if p1.inventaire[0] == "uzi" {
							p1.coins = p1.coins + 15
							fmt.Println("temp", cointemp)
							cout = p1.coins - cointemp
							fmt.Print(p1.coins)
						}
						if p1.inventaire[0] == "mp5" {
							p1.coins = p1.coins + 25
							fmt.Println("temp", cointemp)
							cout = p1.coins - cointemp
							fmt.Print(p1.coins)
						}
						if p1.inventaire[0] == "livre de sort (boule de feu)" {
							p1.coins = p1.coins + 35
							fmt.Println("temp", cointemp)
							cout = p1.coins - cointemp
							fmt.Print(p1.coins)
						}
						if p1.inventaire[0] == "allocation" { //si l'objet acheté est une potion de soin
							p1.coins = p1.coins + 3 //on retire 3 coins au joueur
							fmt.Println("temp", cointemp)
							cout = p1.coins - cointemp
							fmt.Print(p1.coins)
						}
						fmt.Println("Vous avez vendu un ", p1.inventaire[:1], "pour un gain de ,", cout, "d or, il vous reste ", p1.coins, " pièces d or")
						p1.inventairemarchand = append(p1.inventairemarchand, p1.inventaire[:1]...)
						p1.inventaire = append(p1.inventaire[:0], p1.inventaire[1:]...) //meme fonctiomnement que pour l'achat des objets mais dans l'autre sens
						fmt.Println(p1.inventairemarchand)
						fmt.Println(p1.inventaire)
						vendre(p1, equipement1)

					default:
						fmt.Println("Vous n'avez pas choisi une option valide : Veuillez réessayer")
						vendre(p1, equipement1)
					}
				} else {
					fmt.Println("Vous avez vendu un ", p1.inventaire[:1])
					if p1.inventaire[1] == "Cv" {
						p1.coins = p1.coins + 35
						fmt.Println("temp", cointemp)
						cout = p1.coins - cointemp
						fmt.Print(p1.coins)
					}
					if p1.inventaire[1] == "potion de soin" { //si l'objet acheté est une potion de soin
						p1.coins = p1.coins + 3 //on retire 3 coins au joueur
						fmt.Println("temp", cointemp)
						cout = p1.coins - cointemp
						fmt.Print(p1.coins)

					}

					if p1.inventaire[1] == "potion de poison" { //meme fonctionnement
						p1.coins = p1.coins + 3
						fmt.Println("temp", cointemp)
						cout = p1.coins - cointemp
						fmt.Print(p1.coins)
					}
					if p1.inventaire[1] == "uzi" {
						p1.coins = p1.coins + 15
						fmt.Println("temp", cointemp)
						cout = p1.coins - cointemp
						fmt.Print(p1.coins)
					}
					if p1.inventaire[1] == "mp5" {
						p1.coins = p1.coins + 25
						fmt.Println("temp", cointemp)
						cout = p1.coins - cointemp
						fmt.Print(p1.coins)
					}
					if p1.inventaire[1] == "livre de sort (boule de feu)" {
						p1.coins = p1.coins + 35
						fmt.Println("temp", cointemp)
						cout = p1.coins - cointemp
						fmt.Print(p1.coins)
					}
					if p1.inventaire[1] == "allocation" { //si l'objet acheté est une potion de soin
						p1.coins = p1.coins + 3 //on retire 3 coins au joueur
						fmt.Println("temp", cointemp)
						cout = p1.coins - cointemp
						fmt.Print(p1.coins)
					}
					fmt.Println("Vous avez vendu un ", p1.inventaire[:1], "pour un gain de ,", cout, "d or, il vous reste ", p1.coins, " pièces d or")
					p1.inventairemarchand = append(p1.inventairemarchand, p1.inventaire[:1]...)
					p1.inventaire = append(p1.inventaire[:0], p1.inventaire[1:]...) //meme fonctiomnement que pour l'achat des objets mais dans l'autre sens
					fmt.Println(p1.inventairemarchand)
					fmt.Println(p1.inventaire)
					vendre(p1, equipement1)
				}
			} else {
				fmt.Println("Transaction Impossible")
				vendre(p1, equipement1)
			}
		case 2:

			cointemp = p1.coins
			fmt.Println(len(p1.inventaire))
			if len(p1.inventaire) >= 2 {
				fmt.Println(len(p1.inventaire))
				if p1.inventaire[1] == "allocation" { //si l'objet acheté est une potion de soin
					p1.coins = p1.coins + 3 //on retire 3 coins au joueur
					fmt.Println("temp", cointemp)
					cout = p1.coins - cointemp
					fmt.Print(p1.coins)
				}
				if p1.inventaire[1] == "potion de soin" { //si l'objet acheté est une potion de soin
					p1.coins = p1.coins + 3 //on retire 3 coins au joueur
					fmt.Println("temp", cointemp)
					cout = p1.coins - cointemp
					fmt.Print(p1.coins)

				}

				if p1.inventaire[1] == "potion de poison" { //meme fonctionnement
					p1.coins = p1.coins + 3
					fmt.Println("temp", cointemp)
					cout = p1.coins - cointemp
					fmt.Print(p1.coins)
				}
				if p1.inventaire[1] == "uzi" {
					p1.coins = p1.coins + 15
					fmt.Println("temp", cointemp)
					cout = p1.coins - cointemp
					fmt.Print(p1.coins)
				}
				if p1.inventaire[1] == "mp5" {
					p1.coins = p1.coins + 25
					fmt.Println("temp", cointemp)
					cout = p1.coins - cointemp
					fmt.Print(p1.coins)
				}
				if p1.inventaire[1] == "livre de sort (boule de feu)" {
					p1.coins = p1.coins + 35
					fmt.Println("temp", cointemp)
					cout = p1.coins - cointemp
					fmt.Print(p1.coins)
				}
				if p1.inventaire[1] == "cv" {
					p1.coins = p1.coins + 35
					fmt.Println("temp", cointemp)
					cout = p1.coins - cointemp
					fmt.Print(p1.coins)
				}
				fmt.Println("Vous avez vendu un ", p1.inventaire[1], "pour un gain de ,", cout, "d or, il vous reste ", p1.coins, " pièces d or")
				p1.inventairemarchand = append(p1.inventairemarchand, p1.inventaire[1])
				p1.inventaire = append(p1.inventaire[:1], p1.inventaire[2:]...)
				fmt.Println(p1.inventairemarchand)
				fmt.Println(p1.inventaire)
				vendre(p1, equipement1)
			} else {
				fmt.Println("Transaction Impossible")
				vendre(p1, equipement1)
			}

		case 3:
			if len(p1.inventaire) >= 3 {
				if p1.inventaire[2] == "allocation" { //si l'objet acheté est une potion de soin
					p1.coins = p1.coins + 3 //on retire 3 coins au joueur
					fmt.Println("temp", cointemp)
					cout = p1.coins - cointemp
					fmt.Print(p1.coins)
				}
				if p1.inventaire[2] == "potion de soin" { //si l'objet acheté est une potion de soin
					p1.coins = p1.coins + 3 //on retire 3 coins au joueur
					fmt.Println("temp", cointemp)
					cout = p1.coins - cointemp
					fmt.Print(p1.coins)

				}

				if p1.inventaire[2] == "potion de poison" { //meme fonctionnement
					p1.coins = p1.coins + 3
					fmt.Println("temp", cointemp)
					cout = p1.coins - cointemp
					fmt.Print(p1.coins)
				}
				if p1.inventaire[2] == "uzi" {
					p1.coins = p1.coins + 15
					fmt.Println("temp", cointemp)
					cout = p1.coins - cointemp
					fmt.Print(p1.coins)
				}
				if p1.inventaire[2] == "mp5" {
					p1.coins = p1.coins + 25
					fmt.Println("temp", cointemp)
					cout = p1.coins - cointemp
					fmt.Print(p1.coins)
				}
				if p1.inventaire[2] == "livre de sort (boule de feu)" {
					p1.coins = p1.coins + 35
					fmt.Println("temp", cointemp)
					cout = p1.coins - cointemp
					fmt.Print(p1.coins)
				}
				if p1.inventaire[2] == "Cv" {
					p1.coins = p1.coins + 35
					fmt.Println("temp", cointemp)
					cout = p1.coins - cointemp
					fmt.Print(p1.coins)
				}
				fmt.Println("Vous avez vendu un ", p1.inventaire[2], "pour un gain de ,", cout, "d or, il vous reste ", p1.coins, " pièces d or")
				p1.inventairemarchand = append(p1.inventairemarchand, p1.inventaire[2])
				p1.inventaire = append(p1.inventaire[:2], p1.inventaire[3:]...)
				fmt.Println(p1.inventairemarchand)
				fmt.Println(p1.inventaire)
				vendre(p1, equipement1)
			} else {
				fmt.Println("Transaction Impossible")
				vendre(p1, equipement1)
			}
		case 4:
			if len(p1.inventaire) >= 4 {
				if p1.inventaire[3] == "allocation" { //si l'objet acheté est une potion de soin
					p1.coins = p1.coins + 3 //on retire 3 coins au joueur
					fmt.Println("temp", cointemp)
					cout = cointemp - p1.coins //la valeur de cout est égale à la valeur de cointemp - la valeur de coins
					fmt.Print(p1.coins)
				}
				if p1.inventaire[3] == "potion de soin" { //si l'objet acheté est une potion de soin
					p1.coins = p1.coins + 3 //on retire 3 coins au joueur
					fmt.Println("temp", cointemp)
					cout = cointemp - p1.coins //la valeur de cout est égale à la valeur de cointemp - la valeur de coins
					fmt.Print(p1.coins)

				}

				if p1.inventaire[3] == "potion de poison" { //meme fonctionnement
					p1.coins = p1.coins + 3
					fmt.Println("temp", cointemp)
					cout = p1.coins - cointemp
					fmt.Print(p1.coins)
				}
				if p1.inventaire[3] == "uzi" {
					p1.coins = p1.coins + 15
					fmt.Println("temp", cointemp)
					cout = p1.coins - cointemp
					fmt.Print(p1.coins)
				}
				if p1.inventaire[3] == "mp5" {
					p1.coins = p1.coins + 25
					fmt.Println("temp", cointemp)
					cout = p1.coins - cointemp
					fmt.Print(p1.coins)
				}
				if p1.inventaire[3] == "livre de sort (boule de feu)" {
					p1.coins = p1.coins + 35
					fmt.Println("temp", cointemp)
					cout = p1.coins - cointemp
					fmt.Print(p1.coins)
				}
				if p1.inventaire[3] == "Cv" {
					p1.coins = p1.coins + 35
					fmt.Println("temp", cointemp)
					cout = p1.coins - cointemp
					fmt.Print(p1.coins)
				}
				fmt.Println("Vous avez vendu un ", p1.inventaire[3], "pour un gain de ,", cout, "d or, il vous reste ", p1.coins, " pièces d or")
				p1.inventairemarchand = append(p1.inventairemarchand, p1.inventaire[3])
				p1.inventaire = append(p1.inventaire[:3], p1.inventaire[4:]...)
				fmt.Println(p1.inventairemarchand)
				fmt.Println(p1.inventaire)
				vendre(p1, equipement1)

			} else {
				fmt.Println("Transaction Impossible")
				vendre(p1, equipement1)
			}
		case 5:
			if len(p1.inventaire) >= 5 {
				if p1.inventaire[4] == "allocation" { //si l'objet acheté est une potion de soin
					p1.coins = p1.coins + 3 //on retire 3 coins au joueur
					fmt.Println("temp", cointemp)
					cout = p1.coins - cointemp
					fmt.Print(p1.coins)
				}
				if p1.inventaire[4] == "potion de soin" { //si l'objet acheté est une potion de soin
					p1.coins = p1.coins + 3 //on retire 3 coins au joueur
					fmt.Println("temp", cointemp)
					cout = p1.coins - cointemp
					fmt.Print(p1.coins)

				}

				if p1.inventaire[4] == "potion de poison" { //meme fonctionnement
					p1.coins = p1.coins + 3
					fmt.Println("temp", cointemp)
					cout = p1.coins - cointemp
					fmt.Print(p1.coins)
				}
				if p1.inventaire[4] == "uzi" {
					p1.coins = p1.coins + 15
					fmt.Println("temp", cointemp)
					cout = p1.coins - cointemp
					fmt.Print(p1.coins)
				}
				if p1.inventaire[4] == "mp5" {
					p1.coins = p1.coins + 25
					fmt.Println("temp", cointemp)
					cout = p1.coins - cointemp
					fmt.Print(p1.coins)
				}
				if p1.inventaire[4] == "livre de sort (boule de feu)" {
					p1.coins = p1.coins + 35
					fmt.Println("temp", cointemp)
					cout = p1.coins - cointemp
					fmt.Print(p1.coins)
				}
				if p1.inventaire[4] == "Cv" {
					p1.coins = p1.coins + 35
					fmt.Println("temp", cointemp)
					cout = p1.coins - cointemp
					fmt.Print(p1.coins)
				}
				fmt.Println("Vous avez vendu un ", p1.inventaire[4], "pour un gain de ,", cout, "d or, il vous reste ", p1.coins, " pièces d or")
				p1.inventairemarchand = append(p1.inventairemarchand, p1.inventaire[4])
				p1.inventaire = append(p1.inventaire[:4], p1.inventaire[5:]...)
				vendre(p1, equipement1)
			} else {
				fmt.Println("Transaction Impossible")
				vendre(p1, equipement1)
			}
		case 6:
			marchand(p1, equipement1)
		default:
			fmt.Println("Vous n'avez pas choisi une option valide : Veuillez réessayer")
			vendre(p1, equipement1)
		}
	}

}

func inventorylimit(p1 Personnage, equipement1 Equipement) {
	if len(p1.inventaire) == p1.capacitéinventaireperso {
		p1.estsurchargé = true //renvoie true si le joueur est surchargé
		vendre(p1, equipement1)
		//fmt.Println("Vous avez vendu un ", p1.inventaire[0])
		//p1.inventairemarchand = append(p1.inventairemarchand, p1.inventaire[:1]...)
		//p1.inventaire = append(p1.inventaire[:0], p1.inventaire[len(p1.inventaire):]...)
		//fmt.Println(p1.inventairemarchand)
		//fmt.Println(p1.inventaire)

	}
}
func removeinventory(p1 Personnage) {
	p1.inventaire = nil
	p1.inventairemarchand = nil
	p1.capacitéinventaireperso = 0
	fmt.Println(p1.inventaire)
	fmt.Println(p1.inventairemarchand)

}

func removehealth(p1 Personnage, equipement1 Equipement) {
	p1.Pv = p1.Pv - 50
	fmt.Println(p1.Pv)
	iswasted(p1, equipement1)
}

func addhealth(p1 Personnage) {
	p1.Pv += 1
	fmt.Println(p1.Pv)
}

func removehealthpotion(p1 Personnage, equipement1 Equipement) {
	potioncount := len(p1.inventairepotion)
	if len(p1.inventairepotion) == 0 { //si l'inventaire potion est vide
		fmt.Println("Vous n'avez pas de potion")
		DisplayMenu(p1, equipement1)
	} else {
		fmt.Println("????", p1.Pv)
		if p1.Pv == p1.Pvmax { //si le joueur a déjà toute sa vie
			fmt.Println("================================")
			fmt.Println("Votre vie est déja au maximum")
			fmt.Println("potion(s) restante(s)", potioncount)
			fmt.Println("================================")
			DisplayMenu(p1, equipement1)
		}
		i := len(p1.inventairepotion)
		p1.inventairepotion = append(p1.inventairepotion[:0], p1.inventairepotion[:i-1]...)
		if p1.Pv+25 > p1.Pvmax { //si la vie du joueur  après le soin est supérieur à sa vie max
			p1.Pv = p1.Pvmax //la vie du joueur est égale à sa vie max
		} else {
			p1.Pv = p1.Pv + 25 //sinon la vie du joueur est égale à sa vie actuelle + 25
		}
		potioncount = potioncount - 1 //on retire une potion de l'inventaire
		fmt.Println("================================")
		fmt.Println("vous avez consommé une potion de soin")
		fmt.Println("potion(s) restante(s)", potioncount)
		fmt.Println("================================")
		DisplayMenu(p1, equipement1)

	}
}

func addhealthpotion(p1 Personnage, equipement1 Equipement) {
	if len(p1.inventairepotion) < p1.capacitéinventairepotion {
		p1.inventairepotion = append(p1.inventairepotion, "potion")
		fmt.Println(p1.inventairepotion)
	} else {
		fmt.Println("Vous ne pouvez pas avoir plus de 3 potions, pour cela améliorez votre inventaire")
	}
	DisplayMenu(p1, equipement1)
}

func iswasted(p1 Personnage, equipement1 Equipement) { // fonction isdead()
	var z int8

	if p1.Pv <= 0 { //si la vie du joueur est inférieur ou égale à 0
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
			DisplayMenu(p1, equipement1)

		case 2:
			fmt.Println("Jeu Fermé")
			fmt.Println("Vous avez quitté le jeu, Vous feriez mieux la prochaine fois")
		}
	} else {
		DisplayMenu(p1, equipement1)
	}
}

func Poisonpot(p1 Personnage, equipement1 Equipement) {
	fmt.Println("Vous subissez les effets du poison")
	for i := 0; i < 3; i++ { //boucle qui fait perdre 10 pv pendant 3 secondes
		fmt.Println("i", i)
		time.Sleep(1 * time.Second) //temps d'attente fixé à une  1 seconde
		p1.Pv = p1.Pv - 10          //-10pv pour le joueur
		fmt.Println("-10 Pv")
		if p1.Pv <= 0 {
			p1.Pv = 0
		}
		fmt.Println("vous avez ", p1.Pv, "/", p1.Pvmax, "PV")
	}

	iswasted(p1, equipement1) //on vérifie si le joueur est mort

}

func SpellBook(p1 Personnage, equipement1 Equipement) {
	if p1.islearned == 0 { //si le joueur n'a pas encore appris le sort
		p1.skill = append(p1.skill, "Boule de feu")
		p1.persoinventairesort = append(p1.persoinventairesort[:0], p1.persoinventairesort[1:]...) //on retire le sort de son inventaire
		fmt.Println("Vous avez appris le sort boule de feu ")
		p1.islearned = 1 // la valeur de islearned passe à 1: le joueur a appris le sort
		fmt.Println(p1.skill)
		DisplayMenu(p1, equipement1)
	} else { //si le joueur a déjà appris le sort
		fmt.Println("Vous avez déjà appris ce sort")
		DisplayMenu(p1, equipement1)
	}
}

func Quitgame(p1 Personnage, equipement1 Equipement) {
	var h string
	fmt.Println("=========================================")
	fmt.Print("- Voulez-vous vraiment quitter le jeu ?- ")

	fmt.Scanln(&h)
	switch h {
	case "oui":
		//p1.inventaire = nil
		//p1.inventairemarchand = nil
		//p1.capacitéinventaireperso = 0
		fmt.Println(len(p1.inventaire))
		fmt.Println(p1.inventaire)
		fmt.Println(p1.inventairemarchand)
		if p1.estsurchargé == true {
		}
		fmt.Println("====================================")
		fmt.Println("====================================")
		fmt.Println("Jeu Fermé")
		fmt.Println("Merci d'avoir joué")
		fmt.Println("====================================")
		fmt.Println("====================================")

	case "non":
		DisplayMenu(p1, equipement1)
	default:
		fmt.Println("=============================================================")
		fmt.Println("Vous n'avez pas choisi une option valide : Veuillez réessayer")
		fmt.Println("=============================================================")
		Quitgame(p1, equipement1)

	}

}

func charCreation(p1 Personnage, equipement1 Equipement) {
	var name string
	var classe int8
	fmt.Print("Quel est votre nom ? ")
	fmt.Scanln(&name)
	p1.Nom = name
	fmt.Print(name)

	for i := 0; i < len(name); i++ { //boucle qui vérifie si le nom du joueur est supérieur à 1
		if len(name) <= 1 {
			fmt.Println("Votre nom est trop court")
			charCreation(p1, equipement1)

		}
		if name[i] < 'A' || name[i] > 'Z' {
			fmt.Println("Le nom doit commencer par une majuscule") //boucle qui vérifie si le nom ne commence pas par une majuscule
			charCreation(p1, equipement1)
			break

		}

		if name[i+1] < 'a' || name[i] > 'z' {
			fmt.Println("Le nom doit continuer par une minuscule") //boucle qui vérifie si le nom ne continue pas par une minuscule
			charCreation(p1, equipement1)

			//break
		} else {
			fmt.Println("cool")

			fmt.Println("votre nom ", p1.Nom)

			fmt.Print("choisissez votre classe : ")
			fmt.Scanln(&classe)
			switch classe {
			case 1:
				p1.Classe = "Rentier"
				p1.Pvmax = 100
				fmt.Println(p1.Pv)
				fmt.Println("Vous avez choisi la classe Rentier")
				DisplayMenu(p1, equipement1)

			case 2:
				p1.Classe = "Bénéficiaire du RSA"
				p1.Pvmax = 50
				fmt.Println(p1.Pv)
				fmt.Println("Vous avez choisi la classe Bénéficiaire du RSA")
				DisplayMenu(p1, equipement1)

			case 3:
				p1.Classe = "Chomeur"
				p1.Pvmax = 50
				fmt.Println(p1.Pv)
				fmt.Println("Vous avez choisi la classe Chomeur")
				DisplayMenu(p1, equipement1)

			default:
				fmt.Println("Vous n'avez pas choisi une classe valide")
				charCreation(p1, equipement1)
				//pour sortir du switch case
			}
			//break
		}
		break // pour sortir de la boucle for
	}

}
func forge(p1 Personnage, equipement1 Equipement) {
	var choix int8
	//var choix2 int8
	fmt.Println("=========================================")
	fmt.Println("1. Forger une armure")
	fmt.Println("2. forger une arme")
	fmt.Println("3. Afficher les statistiques")
	fmt.Println("4. Quitter le jeu")
	fmt.Println("=========================================")
	fmt.Print("Veuillez choisir une option : ")
	fmt.Scanln(&choix)
	var true1 int
	var true2 string

	switch choix {
	case 1:
		if len(p1.persoinventairemateriaux) <= 1 {
			fmt.Println("Vous n'avez pas assez de matériaux pour forger une armure")
			fmt.Println("Vous n'avez pas assez de matériaux ")
			forge(p1, equipement1)
		} else {
			var choix2 int
			fmt.Println("=========================================")
			fmt.Println("1. Armure en cuir")
			fmt.Println("2. Armure en fer")
			fmt.Println("3. Armure en or")
			fmt.Println("4. Armure en diamant")
			fmt.Println("5. Armure en adamantium")
			fmt.Println("6. choissisez une option : ")
			fmt.Scanln(&choix2)
			fmt.Println("=========================================")
			switch choix2 {
			case 1:
				fmt.Scanln(&choix2)
				fmt.Println(p1.persoinventairemateriaux)
				if p1.estsurchargé == true {
					fmt.Println("Vous êtes surchargé, vous ne pouvez pas forger d'armure")
					DisplayMenu(p1, equipement1)
				} else {
					fmt.Println("etape 1")
					if p1.coins >= 25 {
						fmt.Println("etape 2")
						for i := 1; i < len(p1.persoinventairemateriaux); i++ {
							fmt.Println("efd")
							if p1.persoinventairemateriaux[i] == "cuir de sanglier" {
								true1 = i
								fmt.Println("true1", i, p1.persoinventairemateriaux[i])
								for j := 1; j < len(p1.persoinventairemateriaux); j++ {
									if p1.persoinventairemateriaux[i+1] == "plume de corbeaux" {
										true2 = p1.persoinventairemateriaux[i+1]
										fmt.Println(true2)
										fmt.Println("true2", i, p1.persoinventairemateriaux[i+1])
										equipement1.helmetinventory = append(equipement1.helmetinventory, "casquette de chomeur")
										p1.persoinventairemateriaux = append(p1.persoinventairemateriaux[:true1])
										p1.coins = p1.coins - 25
										fmt.Print(equipement1.helmetinventory)
										fmt.Print(p1.persoinventairemateriaux)
										fmt.Println("Vous avez forgé une casquette de chomeur pour le prix de 25 pièces d'or")
										forge(p1, equipement1)

									} else {
										fmt.Println("false2", i, p1.persoinventairemateriaux[i+1])
										fmt.Println("la condition est ", p1.persoinventairemateriaux[i+1] == "plume de corbeau")
										fmt.Println("Vous n'avez pas les matériaux nécessaires pour forger une armure")
									}

								}
							} else {
								fmt.Println("false1", i, p1.persoinventairemateriaux[i])
								if len(p1.persoinventairemateriaux[i]) == len(p1.persoinventairemateriaux) {
									fmt.Println("Vous n'avez pas les matériaux nécessaires pour forger une armure")
								}

							}

						}

					} else {
						fmt.Println("Vous n'avez pas assez de pièces")
						DisplayMenu(p1, equipement1)
					}

				}

			case 2:
				fmt.Println(p1.persoinventairemateriaux)
				if p1.estsurchargé == true {
					fmt.Println("Vous êtes surchargé, vous ne pouvez pas forger d'armure")
					DisplayMenu(p1, equipement1)
				} else {
					var count int
					fmt.Println("etape 1")
					if p1.coins >= 25 {
						fmt.Println("etape 2")
						for i := 1; i < len(p1.persoinventairemateriaux); i++ {
							count = count + 1
							fmt.Println(count)
							fmt.Println("efd")
							if p1.persoinventairemateriaux[i] == "peau de troll" {
								true1 = i
								fmt.Println("true1", i, p1.persoinventairemateriaux[i])
								for j := 0; j < len(p1.persoinventairemateriaux); j++ {
									fmt.Print("j", j, len(p1.persoinventairemateriaux))
									if p1.persoinventairemateriaux[j] == "fourrure de loup" {
										fmt.Println("true2", i, p1.persoinventairemateriaux[i+1])
										equipement1.chestinventory = append(equipement1.chestinventory, "sweat de chomeur")
										p1.persoinventairemateriaux = append(p1.persoinventairemateriaux[:true1])
										p1.coins = p1.coins - 25
										fmt.Print(equipement1.chestinventory)
										fmt.Print(p1.persoinventairemateriaux)
										fmt.Println("Vous avez forgé un sweat de chomeur pour le prix de 25 pièces d'or")
										forge(p1, equipement1)
										break

									} else {
										fmt.Println("false2", i, p1.persoinventairemateriaux[i+1])
										//fmt.Println("la condition est ", p1.persoinventairemateriaux[i+1] == "plume de corbeau")
										//fmt.Println("Vous n'avez pas les matériaux nécessaires pour forger une armure", len(p1.persoinventairemateriaux))
										if len(p1.persoinventairemateriaux[i]) == len(p1.persoinventairemateriaux) {
											fmt.Println(p1.persoinventairemateriaux[i])
										}
									}

								}
							} else {
								fmt.Println("false1", i, p1.persoinventairemateriaux[i])
								if len(p1.persoinventairemateriaux[i]) == len(p1.persoinventairemateriaux) {
									fmt.Println("Vous n'avez pas les matériaux nécessaires pour forger une armure")
								}

							}

						}

					} else {
						fmt.Println("Vous n'avez pas assez de pièces")
						DisplayMenu(p1, equipement1)
					}

				}

			case 3:
				fmt.Println(p1.persoinventairemateriaux)
				if p1.estsurchargé == true {
					fmt.Println("Vous êtes surchargé, vous ne pouvez pas forger d'armure")
					DisplayMenu(p1, equipement1)
				} else {
					var count int
					fmt.Println("etape 1")
					if p1.coins >= 25 {
						fmt.Println("etape 2")
						for i := 0; i < len(p1.persoinventairemateriaux); i++ {
							count = count + 1
							fmt.Println(count)
							fmt.Println("efd")
							if p1.persoinventairemateriaux[i] == "fourrure de loup" {
								true1 = i
								fmt.Println("true1", i, p1.persoinventairemateriaux[i])
								for j := 0; j < len(p1.persoinventairemateriaux); j++ {
									fmt.Print("j", j, len(p1.persoinventairemateriaux))
									if p1.persoinventairemateriaux[j] == "cuir de sanglier" {
										fmt.Println("true2", i, p1.persoinventairemateriaux[i+1])
										equipement1.bootinventory = append(equipement1.bootinventory, "claquette")
										p1.persoinventairemateriaux = append(p1.persoinventairemateriaux[:true1])
										p1.coins = p1.coins - 25
										fmt.Print(equipement1.bootinventory)
										fmt.Print(p1.persoinventairemateriaux)
										fmt.Println("Vous avez forgé une claquette pour le prix de 25 pièces d'or")
										forge(p1, equipement1)
										break
									} else {
										fmt.Println("false2", i, p1.persoinventairemateriaux[i+1])
										//fmt.Println("la condition est ", p1.persoinventairemateriaux[i+1] == "plume de corbeau")
										//fmt.Println("Vous n'avez pas les matériaux nécessaires pour forger une armure", len(p1.persoinventairemateriaux))
										if len(p1.persoinventairemateriaux[i]) == len(p1.persoinventairemateriaux) {
											fmt.Println(p1.persoinventairemateriaux[i])
										}
									}

								}
							} else {
								fmt.Println("false1", i, p1.persoinventairemateriaux[i])
								if len(p1.persoinventairemateriaux[i]) == len(p1.persoinventairemateriaux) {
									fmt.Println("Vous n'avez pas les matériaux nécessaires pour forger une armure")
								}

							}

						}

					} else {
						fmt.Println("Vous n'avez pas assez de pièces")
						DisplayMenu(p1, equipement1)
					}

				}
			}
		}

	case 4:
		DisplayMenu(p1, equipement1)

	default:
		fmt.Println("Vous n'avez pas choisi une option valide")
		forge(p1, equipement1)
	}
}

func Armor(equipement1 Equipement, p1 Personnage) {
	fmt.Println(equipement1.helmetinventory)
	fmt.Println(equipement1.helmetequiped)
	var choix int8
	var choix2 int8
	fmt.Println("===========================================")
	fmt.Println("===========================================")
	fmt.Println("casque équipé", equipement1.helmetequiped, p1.havehelmet, "  |")
	fmt.Println("===========================================")
	fmt.Println("plastron équipé", equipement1.chestequiped)
	fmt.Println("===========================================")
	fmt.Println("jambières équipé", equipement1.leggingsequiped)
	fmt.Println("===========================================")
	fmt.Println("bottes équipé", equipement1.bootequiped)
	fmt.Println("===========================================")
	fmt.Println("===========================================")
	fmt.Println("=========================================")
	fmt.Println("1. Equiper")
	fmt.Println("2. Déséquiper")
	fmt.Println("3. Quitter le jeu")
	fmt.Println("=========================================")
	fmt.Print("Veuillez choisir une option : ")
	fmt.Scanln(&choix)
	switch choix {
	case 1:
		fmt.Println()
		fmt.Print("choississez une armure à équiper")
		fmt.Println("1. Casque")
		fmt.Println("2.Plastron")
		fmt.Println("3.jambières")
		fmt.Println("4.bottes")
		fmt.Scanln(&choix2)
		switch choix2 {
		case 1:
			fmt.Println(equipement1.helmetinventory)
			if p1.havehelmet == false {
				if len(equipement1.helmetinventory) == 0 {
					fmt.Println("Vous n'avez pas de casque dans votre inventaire")
					Armor(equipement1, p1)
				} else {
					var choix3 int8
					fmt.Print("Choississez une armure à équiper")
					fmt.Println(equipement1.helmetinventory)
					fmt.Scan(&choix3)

					switch choix3 {
					case 1:
						if len(equipement1.helmetinventory) == 1 {
							equipement1.helmetequiped = append(equipement1.helmetinventory, equipement1.helmetinventory[:0]...)
							equipement1.helmetinventory = nil
							if equipement1.helmetequiped[0] == "casquette de chomeur" {
								p1.Pvmax = p1.Pvmax + 10
							}
							p1.havehelmet = true
							fmt.Println(equipement1.helmetequiped)
							var tet []string
							tet = nil
							fmt.Println(tet)
							tet = append(tet, "casque")
							fmt.Println(tet)
							Armor(equipement1, p1)

						} else {
							fmt.Println(equipement1.helmetinventory)
							equipement1.helmetequiped = append(equipement1.helmetinventory, equipement1.helmetinventory[:1]...) //on ajoute l'objet acheté dans l'inventaire du joueur
							equipement1.helmetinventory = append(equipement1.helmetinventory[:0], equipement1.helmetinventory[1:]...)
							if equipement1.helmetequiped[0] == "casquette de chomeur" {
								p1.Pvmax = p1.Pvmax + 10
							}
							p1.havehelmet = true
							fmt.Println(equipement1.helmetinventory)
							fmt.Println(equipement1.helmetequiped)
							fmt.Println("choississez une armure à équiper")
							fmt.Println("Vous avez équiper un casque")
							Armor(equipement1, p1)

						}
					}
				}
			} else {
				fmt.Println("Vous avez déjà un casque équipé")
				Armor(equipement1, p1)
			}
		case 2:
			if p1.havechestplate == false {
				fmt.Println(equipement1.chestinventory)
				if len(equipement1.chestinventory) == 0 {
					fmt.Println("Vous n'avez pas de plastron dans votre inventaire")
					Armor(equipement1, p1)
				} else {
					var choix3 int8
					fmt.Print("Choississez une armure à équiper")
					fmt.Println(equipement1.chestinventory)
					fmt.Scan(&choix3)

					switch choix3 {
					case 1:
						if len(equipement1.chestinventory) == 1 {
							equipement1.chestequiped = append(equipement1.chestinventory, equipement1.chestinventory[:0]...)
							equipement1.chestinventory = nil
							p1.havechestplate = true
							fmt.Println(equipement1.helmetequiped)
							var tet []string
							tet = nil
							fmt.Println(tet)
							tet = append(tet, "casque")
							fmt.Println(tet)
							Armor(equipement1, p1)

						} else {
							fmt.Println(equipement1.helmetinventory)
							equipement1.chestequiped = append(equipement1.chestinventory, equipement1.chestinventory[:1]...) //on ajoute l'objet acheté dans l'inventaire du joueur
							equipement1.chestinventory = append(equipement1.chestinventory[:0], equipement1.chestinventory[1:]...)

							p1.havechestplate = true
							fmt.Println(equipement1.chestinventory)
							fmt.Println(equipement1.chestequiped)
							fmt.Println("choississez une armure à équiper")
							fmt.Println("Vous avez équipé un(e) ", equipement1.chestequiped)
							Armor(equipement1, p1)

						}
					}
				}
			} else {
				fmt.Println("Vous avez déjà équipé une armure")
				Armor(equipement1, p1)
			}
		case 3:
			p1.haveleggings = true
		case 4:
			if p1.haveboots == false {

				fmt.Println(equipement1.bootinventory)
				if len(equipement1.bootinventory) == 0 {
					fmt.Println("Vous n'avez pas de botte dans votre inventaire")
					Armor(equipement1, p1)
				} else {
					var choix3 int8
					fmt.Print("Choississez une armure à équiper")
					fmt.Println(equipement1.bootinventory)
					fmt.Scan(&choix3)

					switch choix3 {
					case 1:
						if len(equipement1.bootinventory) == 1 {
							equipement1.bootequiped = append(equipement1.bootinventory, equipement1.bootinventory[:0]...)
							equipement1.chestinventory = nil
							p1.haveboots = true
							fmt.Println(equipement1.bootequiped)
							var tet []string
							tet = nil
							fmt.Println(tet)
							tet = append(tet, "casque")
							fmt.Println(tet)
							Armor(equipement1, p1)

						} else {
							fmt.Println(equipement1.bootinventory)
							equipement1.bootequiped = append(equipement1.bootinventory, equipement1.bootinventory[:1]...) //on ajoute l'objet acheté dans l'inventaire du joueur
							equipement1.bootinventory = append(equipement1.bootinventory[:0], equipement1.bootinventory[1:]...)

							p1.haveboots = true
							fmt.Println(equipement1.bootinventory)
							fmt.Println(equipement1.bootequiped)
							fmt.Println("choississez une armure à équiper")
							fmt.Println("Vous avez équiper un casque")
							Armor(equipement1, p1)

						}
					}
				}
			} else {
				fmt.Println("Vous avez déjà équipé une armure")
				Armor(equipement1, p1)
			}

		default:
			fmt.Println("Vous n'avez pas choisi une option valide")
			Armor(equipement1, p1)
		}
	case 3:
		DisplayMenu(p1, equipement1)
	default:
		fmt.Println("Vous n'avez pas choisi une option valide")
		Armor(equipement1, p1)
	}
}
