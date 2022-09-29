package main

import (
	"fmt"
	"time"
)

type Personnage struct {
	Nom                         string
	Class                       string   //Classe du joueur
	Level                       int      //Niveau du joueur
	Pvmax                       int      //Santé Max du joueur
	Pv                          int      //Santé du joueur
	ptsdamages                  int      //Dégats du joueur
	xpaquired                   int      //xp gagnée du joueur
	inventory                   []string //Inventaire du joueur
	inventorypotion             []string // Inventaire des potions
	capacityinventorypotionchar int      // Capacité de l'inventaire du joueur s'augmente grace avec la fonction upgradeinventoryslot()
	capacityinventorypotion     int      // Capacité de l'inventaire d
	isoverloaded                bool     //Renvoie true si l'inventaire est plein
	inventorytrader             []string // Inventaire du marchand
	skill                       []string // Compétences apprises du joueur
	deathcount                  int      // Compteur de mort du joueur
	islearned                   int      // Renvoie 1 si le joueur a appris le sort
	coins                       int      // Nombre de pièces du joueur
	traderinventoryspell        []string // Inventaire du marchand pour les sorts
	persoinventairesort         []string // Inventaire du joueur pour les sorts
	tradermaterialsinventory    []string // Inventaire du marchand pour les matériaux
	charmeterialsinventory      []string // Inventaire du joueur pour les matériaux
	armor                       int      // pts d' armure du joueur
	havehelmet                  bool     // Renvoie true si le joueur a un casque
	havechestplate              bool     // Renvoie true si le joueur a une armure
	haveleggings                bool     // Renvoie true si le joueur a des jambières
	haveboots                   bool     // Renvoie true si le joueur a des bottes
	useupgradeinventoryslot     int      // Nombre de fois que le joueur a utilisé l'amelioration de l'inventaire
	initiative                  int      //   points d'Initiative du joueur
	hasattacked                 bool     // Renvoie true si le joueur a attaqué
	xpneeded                    int      // xp nécessaire pour passer au niveau supérieur
	mana                        int      // Mana du joueur
	manamax                     int      // Mana Max du joueur
	isgameshutdown              bool     // Renvoie true si le joueur a quitté le jeu
}

type Equipement struct {
	helmet            []string
	chest             []string
	leggings          []string
	boot              []string
	helmeteffect      []string //Effet du casque
	helmetequiped     []string // inventaire du casque équipé
	chestequiped      []string // inventaire de l'armure  équipé
	leggingsequiped   []string // inventaire des jambières  équipé
	bootequiped       []string // inventaire des bottes  équipé
	helmetinventory   []string // inventaire des casque
	chestinventory    []string // inventaire des armures
	leggingsinventory []string // inventaire des jambières
	bootinventory     []string // inventaire des bottes
	resistanceeffect  bool
}

func main() {
	var p1 Personnage

	var equipement1 Equipement
	equipement1.helmet = []string{} //déclaration des champs de la structures equipement
	equipement1.chest = []string{}
	equipement1.leggings = []string{}
	equipement1.boot = []string{}
	equipement1.helmetinventory = []string{}
	equipement1.chestinventory = []string{}
	equipement1.leggingsinventory = []string{}
	equipement1.bootinventory = []string{}

	var Monster1 Monster
	Monster1.Nom = "Voyou du guetto cammé" //Nom du "monstre"
	Monster1.Pvmax = 40                    // Santé Max du "monstre"
	Monster1.Pv = 400
	Monster1.ptsdamages = 5
	Monster1.initiative = 0      //pts initiative du "monstre"
	Monster1.turn = 0            //Tour du combat
	Monster1.attackturn = -25    // variable qui permet de determiner si c'est au tour du monstre d'attaquer
	Monster1.hasattacked = false // variable qui permet de determiner si le monstre a attaqué
	Monster1.turnbuff = 0        // variable qui permet de determiner si le monstre son attaque de 200% de dégats

	//fmt.Println(p1)
	//p1.display()
	//p1.setPower()
	//p1.LessPower()
	fmt.Println("-------------------------------------------------------------------------------")
	fmt.Println("===============================================================================")
	fmt.Println("                 •BIENVENUE SUR  A CHOMAGE ADVENTURE• ", p1.isgameshutdown)
	fmt.Println("===============================================================================")
	fmt.Println("-------------------------------------------------------------------------------")
	charCreation(p1, equipement1, Monster1)

	//lance la fonction menu du jeu avec en argument les valeurs de la structure Personnage
}
func display(p1 Personnage, equipement1 Equipement, Monster1 Monster) { //Affiche les stats du personnage

	fmt.Println("-----------------------")
	fmt.Println("Nom du joueur:", p1.Nom)
	fmt.Println("Classe:", p1.Class)
	fmt.Println("Santé:", p1.Pv)
	fmt.Println("Niveau:", p1.Level)
	fmt.Println("Progression:", p1.xpaquired, "/", p1.xpneeded)
	fmt.Println("Santé Max :", p1.Pvmax)
	fmt.Println("Mana:", p1.mana)
	fmt.Println("Mana max:", p1.manamax)
	fmt.Println("Dégats de base:", p1.ptsdamages)
	fmt.Println("----------------")
	fmt.Println("")
	DisplayMenu(p1, equipement1, Monster1)
	//DisplayMenu(p1,Equipement{"bob de kevin","débardeur","short de foot","claquette-chaussette",[]string{},[]string{},[]string{},[]string{},[]string{},[]string{},[]string{},[]string{}})

}

func DisplayMenu(p1 Personnage, equipement1 Equipement, Monster1 Monster) { //Affiche le menu du jeu
	p1.mana = 100 //réinitialise le mana du joueur apres un potentiel combat
	var i int8
	if p1.xpaquired >= p1.xpneeded { //Si le joueur a assez d'xp voir plus pour passer au niveau supérieur
		p1.Level += 1                      //gain d'un niveau du joueur
		p1.xpaquired = p1.xpaquired - 1000 //retrait de l'xp nécessaire pour passer au niveau supérieur
		if p1.Class == "Bénéficiaire du RSA" {
			p1.Pvmax += 15
			p1.Pv = p1.Pvmax
			p1.ptsdamages = p1.ptsdamages + 5
		}
		if p1.Class == "Chomeur" {
			p1.Pvmax += 25
			p1.Pv = p1.Pvmax
			p1.ptsdamages = p1.ptsdamages + 2
		}
		if p1.Class == "Rentier" {
			p1.Pvmax += 2
			p1.Pv = p1.Pvmax
			p1.manamax = p1.manamax + 10

		}
	}
	//Menu du jeu avec les différentes options qui sont supportées par un switch case

	fmt.Println("=====================================================================")
	fmt.Println(p1.Nom, "                                                             |")
	fmt.Println("=====================================================================")
	fmt.Println("Santé:", p1.Pv, "♥/", p1.Pvmax, "♥                                                  |")
	fmt.Println("Potion(s) Restante(s):", len(p1.inventorypotion), "                                           |")
	fmt.Println("Niveau:", p1.Level, "                                                          |")
	fmt.Println("Progression:", p1.xpaquired, "/ ", p1.xpneeded, "                                             |")
	fmt.Println("Morts:", p1.deathcount, "                                                           |")
	fmt.Println("=====================================================================")
	fmt.Println("-----")
	fmt.Println("Menu|")
	fmt.Println("-----")
	fmt.Println("----------------------------------------")
	fmt.Println("----------------------------------------")
	fmt.Println("1. Afficher l'inventaire 💼             |")
	fmt.Println("2. Afficher les stats du Personnage 📈  |")
	fmt.Println("3. Vendre/Acheter💲                     |")
	fmt.Println("4. Ingurgiter une potion de poison ☠️   |")
	fmt.Println("5. Arene d'entrainement 🗡️              |")
	fmt.Println("6. Qui sont ils ?❓                     |")
	fmt.Println("7. Réinitialiser 🚪                     |")
	fmt.Println("8. Quitter le jeu❌                     |")
	fmt.Println("-----------------------------------------")
	fmt.Println("-----------------------------------------")
	fmt.Println("►choisissez parmis ces options◄")

	//fmt.Print("Type a number:")
	fmt.Scanln(&i)
	switch i { //Switch permettant de choisir une option dans le menu
	case 1:
		fmt.Println("Votre inventaire contient les objets suivants:")
		AcessInventory(p1, equipement1, Monster1) //Affiche l'inventaire du joueur
	case 2:
		display(p1, equipement1, Monster1) //appel de la fonction display

	case 3:
		trader(p1, equipement1, Monster1) //appel de la fonction marchand
	case 4:
		//SpellBook(p1)
		Poisonpot(p1, equipement1, Monster1)
		//removehealth(p1)
	case 5:
		InitGobelin(p1, equipement1) //appel de la fonction InitGobelin qui initialise le champs de la structure Monster pour le "monstre"

	case 6:
		fmt.Println("================================================================================")
		fmt.Println("                            Auteurs:                                           |")
		fmt.Println("                     ABBA,Steven Spielberg                                     |") // affiche les artistes cachés
		fmt.Println("================================================================================")
		DisplayMenu(p1, equipement1, Monster1)
	case 7:
		fmt.Println("Réinitialisation du personnage en cours.")
		fmt.Println("Réinitialisation du personnage en cours..")
		fmt.Println("Réinitialisation du personnage en cours...")
		fmt.Println("Réinitialisation du personnage en cours.")
		fmt.Println("Réinitialisation du personnage en cours..")
		fmt.Println("Réinitialisation du personnage en cours...")
		charCreation(p1, equipement1, Monster1) //appel de la fonction charcreation qui permet de créer un nouveau personnage
	case 8:
		Quitgame(p1, equipement1, Monster1) //appel de la fonction permettant de quitter le jeu

	default: //Dans le cas ou  l'utilisateur rentre une valeur autre que 1,2,3,4 ou 5
		fmt.Println("❌Vous n'avez pas choisi une option valide : Veuillez réessayer❌")
		fmt.Println("=============================================================")
		DisplayMenu(p1, equipement1, Monster1) //appel de la fonction DisplayMenu

	}
}

func (p1 *Personnage) AddItem(item string) {
	p1.inventory = append(p1.inventory, item)
}

func AcessInventory(p1 Personnage, equipement1 Equipement, Monster1 Monster) { //Affiche l'inventaire du joueur et permet d'intéragir avec  un objet à utiliser tel qu'une potion ou un sort
	var j int8
	fmt.Println(p1.deathcount)
	fmt.Println("============================================")
	fmt.Println("============================================")
	fmt.Println("|										    |")
	fmt.Println("|                                          |")
	fmt.Println("|         Inventaire de", p1.Nom, "        |")
	fmt.Println("|                                          |")
	fmt.Println("|                                          |")
	fmt.Println("============================================")
	fmt.Println("============================================")
	fmt.Println("------------------------------------")
	fmt.Println("Santé de", p1.Nom, ":", p1.Pv)
	fmt.Println("------------------------------------")
	fmt.Println("potion(s) de soin:")
	fmt.Println(len(p1.inventorypotion))
	fmt.Println("================================")
	fmt.Println("Sorts:")
	fmt.Println(p1.persoinventairesort)
	fmt.Println("================================")
	fmt.Println("Objets:")
	fmt.Println(p1.inventory, len(p1.inventory), "|")
	fmt.Println("================================")
	fmt.Println("--------------------------------")
	fmt.Println("choisissez parmis ces options")
	fmt.Println("1. Utiliser une potion de soin🩹")
	fmt.Println("2. Apprendre un sort 📖")
	fmt.Println("3.Salles des armures 🛡️")
	fmt.Println("4. Quitter l'Inventaire❌")
	fmt.Println("----------------------------")
	fmt.Print("choississez une option: ")
	fmt.Scanln(&j)
	fmt.Println("===================================")
	switch j {
	case 1:
		removehealthpotion(p1, equipement1, Monster1) //appel de la fonction permettant d'utiliser une potion et de la retirer de l'inventaire
	case 2:
		if len(p1.persoinventairesort) == 0 { //Si le joueur n'a pas de sort dans son inventaire
			fmt.Println("========================================")
			fmt.Println("❌Vous n'avez pas de sort à apprendre❌")
			fmt.Println("========================================")
			AcessInventory(p1, equipement1, Monster1)
		} else {
			SpellBook(p1, equipement1, Monster1) //appel de la fonction permettant d'utiliser un sort et de le retirer de l'inventaire
		}

	case 3:
		Armor(equipement1, p1, Monster1) //appel de la fonction Armor
	case 4:
		DisplayMenu(p1, equipement1, Monster1) //appel de la fonction permettant de quitter l'inventaire
	case 5:
		charCreation(p1, equipement1, Monster1)
	default:
		fmt.Println("===================================================================")
		fmt.Println("❌Vous n'avez pas choisi une option valide : Veuillez réessayer❌")
		fmt.Println("====================================================================")
	}

}

func trader(p1 Personnage, equipement1 Equipement, Monster1 Monster) { //Fonction permettant d'acheter et de vendre des objets
	//p1.inventairemarchand = []string {"uzi,mp5,ermaemp,ruby,remington700"}
	var i int8
	if p1.isgameshutdown == false { //Si le jeu n'est pas éteint
		fmt.Println("================================")
		fmt.Println("choisissez parmis ces options")
		fmt.Println("================================")
		fmt.Println("1. Acheter💲")
		fmt.Println("2. Vendre 💰")
		fmt.Println("3. Forger un objet 🔨")
		fmt.Println("4. Quitter →")
		fmt.Println("================================")
		fmt.Print("$ Entrez un nombre $ ")
		fmt.Scanln(&i)
		fmt.Println("Your number is:", i)
		switch i {
		case 1:
			buy(p1, equipement1, Monster1) //appel de la fonction permettant d'acheter un objet
		case 2:
			sell(p1, equipement1, Monster1) //appel de la fonction permettant de vendre un objet
		case 3:
			forge(p1, equipement1, Monster1) //appel de la fonction permettant de forger un objet
		case 4:
			DisplayMenu(p1, equipement1, Monster1) //retour au menu principal
		default: //dans les autres cas
			fmt.Println("=====================================================================")
			fmt.Println(" ❌Vous n'avez pas choisi une option valide :Veuillez réessayer ❌")
			fmt.Println("=====================================================================")
			trader(p1, equipement1, Monster1) //retour au menu du marchand
		}
	} else if p1.isgameshutdown == true { //Si le jeu est éteint

	}
}

func buy(p1 Personnage, equipement1 Equipement, Monster1 Monster) { //Fonction permettant d'acheter un objet

	inventorylimit(p1, equipement1, Monster1)
	if p1.isoverloaded == false {
		if len(p1.inventorytrader) == 0 { //si l'inventaire du marchand est vide
			fmt.Println("Le marchand n'a plus rien à vendre")
			DisplayMenu(p1, equipement1, Monster1) //retour au menu principal

		} else {
			var t int8
			var cout int
			var cointemp int
			var nocoin bool
			fmt.Println("____________________________________")
			fmt.Println("|                                  |")
			fmt.Println("|=============ACHETER💲============|")
			fmt.Println("|__________________________________|")
			fmt.Println("________________________________________________________")
			fmt.Println("                                                        ")
			fmt.Println("      Commandes disponibles :1, 2, 3, 4, 5, 6 , 7 , 8, 9")
			fmt.Println("________________________________________________________")
			fmt.Println("____________________________________________________________________________")
			fmt.Println("                                                        ")
			fmt.Println(" Inventaire du marchand:|", p1.inventorytrader)
			fmt.Println("               Votre argent:", p1.coins, "                                    ")
			fmt.Println("____________________________________________________________________________")
			fmt.Println("choisissez parmis ces options")
			fmt.Println("____________________________________")
			fmt.Println(" ACHAT:")
			fmt.Println("____________________________________")
			fmt.Println("1. 1er objet")
			fmt.Println("2. 2eme objet")
			fmt.Println("3. 3eme objet")
			fmt.Println("4. 4eme objet")
			fmt.Println("5. 5eme objet")
			fmt.Println("____________________________________")
			fmt.Println("AUTRES:")
			fmt.Println("____________________________________")
			fmt.Println("6. Objet Magiques")
			fmt.Println("7. Matériaux")
			fmt.Println("8. Augmentations inventaire")
			fmt.Println("9. Achat Potion de soin")
			fmt.Println("_____________________________________")
			fmt.Println("10.  Quitter →")
			fmt.Println("_____________________________________")
			fmt.Println(len(p1.inventorytrader))
			fmt.Println(p1.inventorytrader)
			fmt.Print("$Choississez le produit que vous souhaitez acheter$ ")
			fmt.Scanln(&t)
			switch t {

			case 1:
				//fmt.Println(p1.inventorytrader[0] == "potion de soin")
				cointemp = p1.coins               //on stocke le nombre de pièces du joueur dans une variable temporaire
				if len(p1.inventorytrader) >= 1 { //si l'inventaire du marchand contient au moins 1 objet( permet de gerer l'erreur "index out of range")

					if p1.inventorytrader[0] == "potion de soin" { //si l'objet acheté est une potion de soin
						if p1.coins >= 3 { //si le joueur a assez de pièces
							p1.coins = p1.coins - 3    //on retire 3 coins au joueur
							cout = cointemp - p1.coins //la valeur de cout est égale à la valeur de cointemp - la valeur de coins
						} else {
							nocoin = true //si le joueur n'a pas assez de pièces, on passe la variable nocoin à true
						}
					}

					if p1.inventorytrader[0] == "potion de poison" { //meme fonctionnement
						if p1.coins >= 6 {
							p1.coins = p1.coins - 6 //
							cout = cointemp - p1.coins

						} else {

							nocoin = true

						}
					}
					if p1.inventorytrader[0] == "uzi" { //meme fonctionnement
						if p1.coins >= 25 {
							p1.coins = p1.coins - 25
							cout = cointemp - p1.coins

						} else {

							nocoin = true

						}
					}
					if p1.inventorytrader[0] == "mp5" {
						if p1.coins >= 30 {
							p1.coins = p1.coins - 30 //meme fonctionnement
							cout = cointemp - p1.coins
							fmt.Print(p1.coins)
						} else {

							nocoin = true

						}
					}
					if p1.inventorytrader[0] == "Cv" {
						if p1.coins >= 30 {
							p1.coins = p1.coins - 30 //meme fonctionnement
							cout = cointemp - p1.coins
							fmt.Print(p1.coins)
						} else {

							nocoin = true

						}
					}
					if p1.inventorytrader[0] == "allocation" {
						if p1.coins >= 30 {
							p1.coins = p1.coins - 30 //meme fonctionnement
							cout = cointemp - p1.coins
							fmt.Print(p1.coins)
						} else {

							nocoin = true

						}
					}
					if p1.inventorytrader[0] == "ak47" {
						if p1.coins >= 30 {
							p1.coins = p1.coins - 30
							cout = cointemp - p1.coins //meme fonctionnement
							fmt.Print(p1.coins)
						} else {

							nocoin = true

						}
					}
					if nocoin == true {
						fmt.Println("=====================================")
						fmt.Println(("❌Vous n'avez pas assez de pièces❌"))
						fmt.Println("=====================================")
						trader(p1, equipement1, Monster1)
						break
					} else {
						fmt.Println("==============================================================================================================================")
						fmt.Println("Vous avez acheté un ", p1.inventorytrader[:1], "pour un cout de ,", cout, "d or 💸, il vous reste ", p1.coins, " 💵 pièces d or")
						fmt.Println("==============================================================================================================================")
						p1.inventory = append(p1.inventory, p1.inventorytrader[:1]...)                 //on ajoute l'objet acheté dans l'inventaire du joueur
						p1.inventorytrader = append(p1.inventorytrader[:0], p1.inventorytrader[1:]...) //on supprime l'objet acheté de l'inventaire du marchand le principe est le meme pour la vente d'objets
						t = 0
						buy(p1, equipement1, Monster1)
					}

				} else {
					fmt.Println("=============================")
					fmt.Println("❌Transaction Impossible❌") //condition gerant l'execption "index out of range"
					fmt.Println("=============================")
					buy(p1, equipement1, Monster1)
				}

			case 2: //meme fonctionnement mais pour l'objet a la deuxième place dans l'inventaire du marchand
				cointemp = p1.coins
				if len(p1.inventorytrader) >= 2 {
					if p1.inventorytrader[1] == "potion de soin" {
						if p1.coins >= 30 {
							p1.coins = p1.coins - 3
							cout = cointemp - p1.coins
							fmt.Print(p1.coins)
						} else {
							fmt.Println("=====================================")
							fmt.Println("❌Vous n'avez pas assez de pièces❌")
							fmt.Println("=====================================")
							nocoin = true

						}
					}
					if p1.inventorytrader[1] == "potion de poison" {
						if p1.coins >= 30 {
							p1.coins = p1.coins - 6
							cout = cointemp - p1.coins
							fmt.Print(p1.coins)
						} else {
							nocoin = true

						}
					}
					if p1.inventorytrader[1] == "uzi" {
						if p1.coins >= 30 {
							p1.coins = p1.coins - 25
							cout = cointemp - p1.coins
							fmt.Print(p1.coins)
						} else {
							nocoin = true

						}
					}
					if p1.inventorytrader[1] == "mp5" {
						if p1.coins >= 30 {
							p1.coins = p1.coins - 30
							cout = cointemp - p1.coins

						} else {
							nocoin = true

						}
					}
					if p1.inventorytrader[1] == "Cv" {
						if p1.coins >= 30 {
							p1.coins = p1.coins - 30
							cout = cointemp - p1.coins

						} else {
							nocoin = true

						}
					}
					if p1.inventorytrader[1] == "allocation" {
						if p1.coins >= 30 {
							p1.coins = p1.coins - 30
							cout = cointemp - p1.coins

						} else {
							nocoin = true

						}
					}
					if p1.inventorytrader[1] == "ak47" {
						if p1.coins >= 30 {
							p1.coins = p1.coins - 30
							cout = cointemp - p1.coins

						} else {
							nocoin = true

						}
					}
					if nocoin == true {
						fmt.Println("=====================================")
						fmt.Println(("❌Vous n'avez pas assez de pièces❌"))
						fmt.Println("=====================================")
						trader(p1, equipement1, Monster1)
						break
					} else {
						fmt.Println("==============================================================================================================================")
						fmt.Println("Vous avez acheté un ", p1.inventorytrader[1], "pour un cout de", cout, " pièces d'or 💸, il vous reste ", p1.coins, " 💵 pièces d or")
						fmt.Println("==============================================================================================================================")
						p1.inventory = append(p1.inventory, p1.inventorytrader[1])
						p1.inventorytrader = append(p1.inventorytrader[:1], p1.inventorytrader[2:]...)
						buy(p1, equipement1, Monster1)
					}
				} else {
					fmt.Println("===========================")
					fmt.Println("❌Transaction Impossible❌")
					fmt.Println("===========================")
					buy(p1, equipement1, Monster1)
					break
				}

				//idem mais pour le 3ème objet
			case 3:
				cointemp = p1.coins
				if len(p1.inventorytrader) >= 3 {

					if p1.inventorytrader[2] == "potion de soin" {
						if p1.coins >= 3 {
							p1.coins = p1.coins - 3

							cout = cointemp - p1.coins

						} else {
							nocoin = true

						}
					}

					if p1.inventorytrader[2] == "potion de poison" {
						if p1.coins >= 6 {
							p1.coins = p1.coins - 6
							fmt.Println("temp", cointemp)
							cout = cointemp - p1.coins
							fmt.Print(p1.coins)
						} else {
							nocoin = true

						}
					}
					if p1.inventorytrader[2] == "uzi" {
						if p1.coins >= 25 {
							p1.coins = p1.coins - 25
							cout = cointemp - p1.coins

						} else {
							nocoin = true

						}
					}
					if p1.inventorytrader[2] == "mp5" {
						if p1.coins >= 30 {
							p1.coins = p1.coins - 30

							cout = cointemp - p1.coins

						} else {
							nocoin = true

						}
					}

					if p1.inventorytrader[2] == "Cv" {
						if p1.coins >= 30 {
							p1.coins = p1.coins - 30

							cout = cointemp - p1.coins

						} else {
							nocoin = true

						}
					}
					if p1.inventorytrader[2] == "allocation" {
						if p1.coins >= 30 {
							p1.coins = p1.coins - 30

							cout = cointemp - p1.coins

						} else {
							nocoin = true
						}
					}
					if p1.inventorytrader[2] == "ak47" {
						if p1.coins >= 30 {
							p1.coins = p1.coins - 30
							cout = cointemp - p1.coins

						} else {
							nocoin = true

						}
					}
					if nocoin == true {
						fmt.Println("==================================")
						fmt.Println(("❌Vous n'avez pas assez de pièces❌"))
						fmt.Println("==================================")
						trader(p1, equipement1, Monster1)
						break
					} else {
						fmt.Println("==============================================================================================================================")
						fmt.Println("Vous avez acheté un ", p1.inventorytrader[2], "pour un cout de", cout, " pièces d'or 💸, il vous reste ", p1.coins, " 💵 pièces d or")
						fmt.Println("==============================================================================================================================")
						p1.inventory = append(p1.inventory, p1.inventorytrader[2])
						p1.inventorytrader = append(p1.inventorytrader[:2], p1.inventorytrader[3:]...)
						buy(p1, equipement1, Monster1)
					}
				} else {
					fmt.Println("Transaction Impossible")
					buy(p1, equipement1, Monster1)
				}
			case 4: // idem mais pour le 4ème objet

				cointemp = p1.coins
				if len(p1.inventorytrader) >= 4 {

					if p1.inventorytrader[3] == "potion de soin" {
						if p1.coins >= 3 {
							p1.coins = p1.coins - 3
							cout = cointemp - p1.coins

						} else {
							nocoin = true
						}
					}

					if p1.inventorytrader[3] == "potion de poison" {
						if p1.coins >= 6 {
							p1.coins = p1.coins - 6
							cout = cointemp - p1.coins

						} else {
							nocoin = true
						}
					}

					if p1.inventorytrader[3] == "uzi" {
						if p1.coins >= 25 {
							p1.coins = p1.coins - 25
							cout = cointemp - p1.coins

						} else {
							nocoin = true

						}
					}

					if p1.inventorytrader[3] == "mp5" {
						if p1.coins >= 30 {
							p1.coins = p1.coins - 30
							cout = cointemp - p1.coins
							fmt.Print(p1.coins, nocoin, p1.coins >= 30)
						} else {
							nocoin = true
						}
					}

					if p1.inventorytrader[3] == "Cv" {
						if p1.coins >= 30 {
							p1.coins = p1.coins - 30
							cout = cointemp - p1.coins
							fmt.Print(p1.coins)
						} else {
							nocoin = true
						}
					}

					if p1.inventorytrader[3] == "allocation" {
						if p1.coins >= 30 {
							p1.coins = p1.coins - 30
							cout = cointemp - p1.coins
							fmt.Print(p1.coins)
						} else {
							nocoin = true

						}
					}

					if p1.inventorytrader[3] == "ak47" {
						if p1.coins >= 30 {
							p1.coins = p1.coins - 30
							cout = cointemp - p1.coins

						} else {
							nocoin = true

						}
					}

					if nocoin == true {
						fmt.Println("==================================")
						fmt.Println("❌Vous n'avez pas assez de pièces❌")
						fmt.Println("==================================")
						trader(p1, equipement1, Monster1)
						break
					} else {
						fmt.Println("==============================================================================================================================")
						fmt.Println("Vous avez acheté un ", p1.inventorytrader[3], "pour un cout de", cout, " pièces d'or 💸, il vous reste ", p1.coins, " 💵 pièces d or")
						fmt.Println("==============================================================================================================================")
						p1.inventorytrader = append(p1.inventorytrader[:3], p1.inventorytrader[4:]...)
						buy(p1, equipement1, Monster1)
					}
				} else {
					fmt.Println("======================")
					fmt.Println("❌Transaction Impossible❌")
					fmt.Println("======================")
					buy(p1, equipement1, Monster1)
				}

			case 5: //de meme pour le 5ème objet
				cointemp = p1.coins
				if len(p1.inventorytrader) >= 5 {
					if p1.inventorytrader[4] == "potion de soin" {
						if p1.coins >= 3 {
							p1.coins = p1.coins - 3
							cout = cointemp - p1.coins

						} else {
							nocoin = true

						}
					}

					if p1.inventorytrader[4] == "potion de poison" {
						if p1.coins >= 6 {
							p1.coins = p1.coins - 6
							cout = cointemp - p1.coins

						} else {
							nocoin = true

						}
					}
					if p1.inventorytrader[4] == "uzi" {
						if p1.coins >= 25 {
							p1.coins = p1.coins - 25
							cout = cointemp - p1.coins
							fmt.Print(p1.coins)
						} else {
							nocoin = true

						}
					}
					if p1.inventorytrader[4] == "mp5" {
						if p1.coins >= 30 {
							p1.coins = p1.coins - 30
							cout = cointemp - p1.coins
							fmt.Print(p1.coins)
						} else {
							nocoin = true

						}
					}

					if p1.inventorytrader[4] == "Cv" {
						if p1.coins >= 30 {
							p1.coins = p1.coins - 30
							cout = cointemp - p1.coins
							fmt.Print(p1.coins)
						} else {
							nocoin = true

						}
					}
					if p1.inventorytrader[4] == "allocation" {
						if p1.coins >= 30 {
							p1.coins = p1.coins - 30
							cout = cointemp - p1.coins
							fmt.Print(p1.coins)
						} else {
							nocoin = true

						}
					}
					if p1.inventorytrader[4] == "ak47" {
						if p1.coins >= 30 {
							p1.coins = p1.coins - 30
							cout = cointemp - p1.coins

						} else {
							nocoin = true

						}
					}

					if nocoin == true {
						fmt.Println("=================================")
						fmt.Println(("❌Vous n'avez pas assez de pièces❌"))
						fmt.Println("=================================")
						trader(p1, equipement1, Monster1)
						break
					} else {
						fmt.Println("==============================================================================================================================")
						fmt.Println("Vous avez acheté un ", p1.inventorytrader[4], "pour un cout de", cout, " pièces d'or 💸, il vous reste ", p1.coins, " 💵 pièces d or")
						fmt.Println("==============================================================================================================================")
						fmt.Println("Vous avez acheté un ", p1.inventorytrader[4])
						p1.inventory = append(p1.inventory, p1.inventorytrader[4])
						p1.inventorytrader = append(p1.inventorytrader[:4], p1.inventorytrader[5:]...)
						buy(p1, equipement1, Monster1)
					}
				} else {
					fmt.Println("======================")
					fmt.Println("❌Transaction Impossible❌")
					fmt.Println("======================")
					buy(p1, equipement1, Monster1)
				}
			case 6:
				buyperk(p1, equipement1, Monster1) //appel de la fonction d'achat d'objet magique
			case 7:
				materials(p1, equipement1, Monster1) //appel de la fonction d'achat de matériaux pour le craft
			case 8:
				upgradeinventoryslot(p1, equipement1, Monster1) //appel de la fonction d'augmentation de l'inventaire
				trader(p1, equipement1, Monster1)
			case 9:
				buyhealthpotion(p1, equipement1, Monster1) //appel de la fonction d'achat de potion de vie
			case 10:
				trader(p1, equipement1, Monster1)
			default:
				fmt.Println("=============================================================")
				fmt.Println("❌Vous n'avez pas choisi une option valide : Veuillez réessayer❌")
				fmt.Println("=============================================================")
				buy(p1, equipement1, Monster1)

				break
			}
		}
	} else {
		fmt.Println("Vous etes surchargé, vous ne pouvez pas acheter d'objet, veuillez vendre un objet de votre inventaire pour continuer")
		trader(p1, equipement1, Monster1)
	}
}

func buyperk(p1 Personnage, equipement1 Equipement, Monster1 Monster) {

	if len(p1.traderinventoryspell) == 0 { //si le marchand n'a plus de sort

		fmt.Println("Le marchand n'a plus rien à vendre")
		DisplayMenu(p1, equipement1, Monster1)
	} else {
		var cout int
		var cointemp int
		var j int8
		var nocoin bool
		fmt.Println("____________________________________")
		fmt.Println("|                                  |")
		fmt.Println("|=============💲ACHETER💲=========|")
		fmt.Println("|__________________________________|")
		fmt.Println("                                                      ")
		fmt.Println("________________________________________________________")
		fmt.Println("                                                      ")
		fmt.Println("      Commandes disponibles :1 , 2 , 3 , 4 , 5 ")
		fmt.Println("________________________________________________________")
		fmt.Println("                                                      ")
		fmt.Println("--------------------------------------------------------")
		fmt.Println("========================================================")
		fmt.Println("objets magiques à vendre:", p1.traderinventoryspell)
		fmt.Println("========================================================")
		fmt.Println("--------------------------------------------------------")
		fmt.Println("========================================================")
		fmt.Println("Votre inventaire:|", p1.inventory, " Vous avez ", p1.coins, " pièces d'or") //affiche l'inventaire du joueur
		fmt.Println("========================================================")
		fmt.Println("--------------------------------------------------------")
		fmt.Println("                COMMANDES :                                       ")
		fmt.Println("========================================================")
		fmt.Println("1. Premier Objet")
		fmt.Println("2. Second Objet")
		fmt.Println("3. Troisième Objet")
		fmt.Println("4. Quatrième Objet")
		fmt.Println("5. Quitter")
		fmt.Println("========================================================")
		fmt.Print("$Choississez l'objet que vous souhaitez acheter$ ")
		fmt.Scanln(&j)
		switch j {
		case 1: // meme principe que pour l'achat des objets normaux
			cointemp = p1.coins               //variable qui stocke le  montant d'or du joueur avant l'achat
			if len(p1.inventorytrader) >= 1 { //si l'inventaire du joueur ne contient qu'un item
				if p1.traderinventoryspell[0] == "livre de sort (boule de feu)" { // verifie si le l'item est bien un livre de sort de boule de feu
					if p1.coins >= 30 {
						p1.coins = p1.coins - 30 //  le montant cout de l'objet est  retiré de l'inevntaire du joueur : exemple -> si 100pcs - 30 = 70pcs restants
						fmt.Println("temp", cointemp)
						cout = cointemp - p1.coins // le cout de l'objet  est égal au montant de départ - le montant restant
						fmt.Print(p1.coins)        // Imprime l'argent du joueur (restant)
					} else {
						nocoin = true
					}
				}
				if nocoin == true {
					fmt.Println("==================================================")
					fmt.Println("❌Vous n'avez pas assez d'or pour acheter cet objet❌")
					fmt.Println("==================================================")
					buyperk(p1, equipement1, Monster1)
					break
				}
				fmt.Println("==============================================================================================================================")
				fmt.Println("Vous avez acheté un ", p1.traderinventoryspell[0], "pour un cout de", cout, " pièces d'or 💸, il vous reste ", p1.coins, " 💵 pièces d or")
				fmt.Println("==============================================================================================================================")
				p1.persoinventairesort = append(p1.persoinventairesort, p1.traderinventoryspell[:1]...)       //ajoute l'objet acheté dans l'inventaire du joueur
				p1.traderinventoryspell = append(p1.traderinventoryspell[:0], p1.traderinventoryspell[1:]...) // retire l'objet acheté de l'inventaire du marchand
				fmt.Println(p1.traderinventoryspell)
				fmt.Println(p1.inventory)

				fmt.Println(p1.persoinventairesort)
				buyperk(p1, equipement1, Monster1) //retour au menu d'achat
			} else {
				fmt.Print("===============================================================")
				fmt.Println("❌Transaction Impossible❌")
				fmt.Println("===============================================================")
				buyperk(p1, equipement1, Monster1)
			}
		case 5:
			buy(p1, equipement1, Monster1) //retour au menu d'achat
		default:
			fmt.Println("===============================================================")
			fmt.Println("Vous n'avez pas choisi une option valide : Veuillez réessayer")
			fmt.Println("===============================================================")
			buyperk(p1, equipement1, Monster1)
		}
	}
}
func materials(p1 Personnage, equipement1 Equipement, Monster1 Monster) {
	var h int //variable pour le choix
	var cointemp int
	var cout int
	var nocoin bool
	if len(p1.tradermaterialsinventory) == 0 { //si l'inventaire du marchand est vide
		fmt.Println("Le marchand n'a plus rien à vendre")
		DisplayMenu(p1, equipement1, Monster1) //retour au menu principal
	} else {
		fmt.Println("____________________________________")
		fmt.Println("|                                  |")
		fmt.Println("|===========💲ACHETER💲===========|")
		fmt.Println("|__________________________________|")
		fmt.Println("                                                      ")
		fmt.Println("________________________________________________________")
		fmt.Println("                                                      ")
		fmt.Println("      Commandes disponibles :1 , 2 , 3 , 4 , 5 ")
		fmt.Println("________________________________________________________")
		fmt.Println("                                                      ")

		fmt.Println("========================================================")
		fmt.Println("Matériaux à vendre:", p1.tradermaterialsinventory)
		fmt.Println("========================================================")
		fmt.Println("--------------------------------------------------------")
		fmt.Println("========================================================")
		fmt.Println("Votre inventaire:|", p1.inventory, " Vous avez ", p1.coins, " pièces d'or")
		fmt.Println("========================================================")
		fmt.Println("--------------------------------------------------------")
		fmt.Println("Commandes disponibles ")
		fmt.Println("1. Premier Objet")
		fmt.Println("2. Second Objet")
		fmt.Println("3. Troisième Objet")
		fmt.Println("4. Quatrième Objet")
		fmt.Println("5. Quitter")
		fmt.Print("Choisissez votre action  : ")
		fmt.Scanln(&h)

		switch h {
		case 1: // meme principe  que pour la vente  des objets normaux
			cointemp = p1.coins
			if len(p1.tradermaterialsinventory) >= 1 { //si l'inventaire du marchand contient au moins un item
				if p1.tradermaterialsinventory[0] == "fourrure de loup" {
					if p1.coins >= 3 {
						p1.coins = p1.coins - 3
						cout = cointemp - p1.coins
					} else {
						nocoin = true
					}

				}

				if p1.tradermaterialsinventory[0] == "peau de troll" {
					if p1.coins >= 45 {
						p1.coins = p1.coins - 45

						cout = cointemp - p1.coins
					} else {
						nocoin = true
					}

				}
				//fmt.Println(p1.marchandinventairemateriaux[0] == "cuir de sanglier")
				if p1.tradermaterialsinventory[0] == "cuir de sanglier" {
					if p1.coins >= 15 {
						p1.coins = p1.coins - 15

						cout = cointemp - p1.coins
					} else {
						nocoin = true
					}

				}
				//fmt.Println(p1.marchandinventairemateriaux[0] == "cuir de sanglier")
				if p1.tradermaterialsinventory[0] == "plume de corbeaux" {
					if p1.coins >= 9 {
						p1.coins = p1.coins - 9

						cout = cointemp - p1.coins
					} else {
						nocoin = true
					}
				}

				if nocoin == true {
					fmt.Println("==================================================")
					fmt.Println("❌Vous n'avez pas assez d'or pour acheter cet objet❌")
					fmt.Println("==================================================")
					materials(p1, equipement1, Monster1)
					break
				}
				fmt.Println("========================================================================================================================================")
				fmt.Println("Vous avez acheté un ", p1.tradermaterialsinventory[0], "pour un cout de", cout, "💸  pièces d or, il vous reste ", p1.coins, "  💵 pièces d or")
				fmt.Println("========================================================================================================================================")
				p1.charmeterialsinventory = append(p1.charmeterialsinventory, p1.tradermaterialsinventory[0])             //ajoute l'objet acheté dans l'inventaire du joueur
				p1.tradermaterialsinventory = append(p1.tradermaterialsinventory[:0], p1.tradermaterialsinventory[1:]...) // retire l'objet acheté de l'inventaire du marchand
				fmt.Println(p1.tradermaterialsinventory)
				fmt.Println(p1.charmeterialsinventory)
				materials(p1, equipement1, Monster1)
			} else {
				fmt.Println("Transaction Impossible")
				materials(p1, equipement1, Monster1)
			}

		case 2: // meme chose que pour le cas 1
			cointemp = p1.coins
			if len(p1.tradermaterialsinventory) >= 2 {
				fmt.Print("Vous avez acheté un", p1.tradermaterialsinventory[0])
				if p1.tradermaterialsinventory[1] == "fourrure de loup" {
					if p1.coins >= 3 {
						p1.coins = p1.coins - 3
						cout = cointemp - p1.coins
					} else {
						nocoin = true
					}

				}

				if p1.tradermaterialsinventory[1] == "peau de troll" {
					if p1.coins >= 45 {
						p1.coins = p1.coins - 45
						cout = cointemp - p1.coins
					} else {
						nocoin = true
					}

				}

				if p1.tradermaterialsinventory[1] == "cuir de sanglier" {
					if p1.coins >= 15 {
						p1.coins = p1.coins - 15

						cout = cointemp - p1.coins
					} else {
						nocoin = true
					}

				}
				if p1.tradermaterialsinventory[1] == "plume de corbeaux" {
					if p1.coins >= 9 {
						p1.coins = p1.coins - 9

						cout = cointemp - p1.coins
					} else {
						nocoin = true
					}

				}
				if nocoin == true {
					fmt.Println("==================================================")
					fmt.Println("❌Vous n'avez pas assez d'or pour acheter cet objet❌")
					fmt.Println("==================================================")
					buyperk(p1, equipement1, Monster1)
					break
				}
				fmt.Println("Vous avez acheté un ", p1.tradermaterialsinventory[1], "pour un cout de", cout, " pièces d or, il vous reste ", p1.coins, " pièces d or")
				p1.charmeterialsinventory = append(p1.charmeterialsinventory, p1.tradermaterialsinventory[1])
				p1.tradermaterialsinventory = append(p1.tradermaterialsinventory[:1], p1.tradermaterialsinventory[2:]...)
				fmt.Println(p1.tradermaterialsinventory)
				fmt.Println(p1.charmeterialsinventory)
				materials(p1, equipement1, Monster1)
			} else {
				fmt.Println("❌Transaction Impossible❌")
				materials(p1, equipement1, Monster1)
			}
		case 3: // meme chose que pour le cas 1
			cointemp = p1.coins
			if len(p1.tradermaterialsinventory) >= 3 {
				fmt.Print("Vous avez acheté un", p1.tradermaterialsinventory[0])
				if p1.tradermaterialsinventory[2] == "fourrure de loup" {
					if p1.coins >= 9 {
						p1.coins = p1.coins - 9

						cout = cointemp - p1.coins
					} else {
						nocoin = true
					}

				}

				if p1.tradermaterialsinventory[2] == "peau de troll" {
					if p1.coins >= 45 {
						p1.coins = p1.coins - 45

						cout = cointemp - p1.coins
					} else {
						nocoin = true
					}
				}

				if p1.tradermaterialsinventory[2] == "cuir de sanglier" {
					if p1.coins >= 15 {
						p1.coins = p1.coins - 15
						cout = cointemp - p1.coins
					} else {
						nocoin = true
					}
				}

				if p1.tradermaterialsinventory[2] == "plume de corbeaux" {
					if p1.coins >= 9 {
						p1.coins = p1.coins - 9

						cout = cointemp - p1.coins
					} else {
						nocoin = true
					}
				}
				if nocoin == true {
					fmt.Println("==================================================")
					fmt.Println("❌Vous n'avez pas assez d'or pour acheter cet objet❌")
					fmt.Println("==================================================")
					buyperk(p1, equipement1, Monster1)
					break
				}
				fmt.Println("Vous avez acheté un ", p1.tradermaterialsinventory[2], "pour un cout de", cout, " pièces d or, il vous reste ", p1.coins, " pièces d or")
				p1.charmeterialsinventory = append(p1.charmeterialsinventory, p1.tradermaterialsinventory[2])
				p1.tradermaterialsinventory = append(p1.tradermaterialsinventory[:2], p1.tradermaterialsinventory[3:]...)
				fmt.Println(p1.tradermaterialsinventory)
				fmt.Println(p1.charmeterialsinventory)
				materials(p1, equipement1, Monster1)
			} else {
				fmt.Println("============================")
				fmt.Println("❌Transaction Impossible❌")
				fmt.Println("============================")
				materials(p1, equipement1, Monster1)
			}
		case 4: // meme chose que pour le cas 1
			cointemp = p1.coins
			if len(p1.tradermaterialsinventory) >= 4 {
				fmt.Print("Vous avez acheté un", p1.tradermaterialsinventory[0])

				if p1.tradermaterialsinventory[3] == "fourrure de loup" {
					if p1.coins >= 3 {
						p1.coins = p1.coins - 3
						cout = cointemp - p1.coins
					} else {
						nocoin = true
					}
				}
				if p1.tradermaterialsinventory[3] == "peau de troll" {
					if p1.coins >= 45 {
						p1.coins = p1.coins - 45
						cout = cointemp - p1.coins
					} else {
						nocoin = true
					}
				}

				if p1.tradermaterialsinventory[3] == "cuir de sanglier" {
					if p1.coins >= 15 {
						p1.coins = p1.coins - 15
						cout = cointemp - p1.coins
					} else {
						nocoin = true
					}
				}
				if p1.tradermaterialsinventory[3] == "plume de corbeaux" {
					if p1.coins >= 9 {
						p1.coins = p1.coins - 9
						cout = cointemp - p1.coins
					} else {
						nocoin = true
					}
				}
				if nocoin == true {
					fmt.Println("==================================================")
					fmt.Println("❌Vous n'avez pas assez d'or pour acheter cet objet❌")
					fmt.Println("==================================================")
					buyperk(p1, equipement1, Monster1)
					break
				}
				fmt.Println("Vous avez acheté un ", p1.tradermaterialsinventory[3], "pour un cout de", cout, " pièces d or, il vous reste ", p1.coins, " pièces d or")
				p1.charmeterialsinventory = append(p1.charmeterialsinventory, p1.tradermaterialsinventory[3])
				p1.tradermaterialsinventory = append(p1.tradermaterialsinventory[:3], p1.tradermaterialsinventory[4:]...)
				fmt.Println(p1.tradermaterialsinventory)
				fmt.Println(p1.charmeterialsinventory)
				materials(p1, equipement1, Monster1)
			} else {
				fmt.Println("===========================")
				fmt.Println("❌Transaction Impossible❌")
				fmt.Println("===========================")
				materials(p1, equipement1, Monster1)
			}
		case 5:
			trader(p1, equipement1, Monster1)

		default:
			fmt.Println("=========================================")
			fmt.Println("❌Vous n'avez pas choisi une action valide❌")
			fmt.Println("=========================================")
			materials(p1, equipement1, Monster1)
		}

	}
}

func upgradeinventoryslot(p1 Personnage, equipement1 Equipement, Monster1 Monster) {
	if p1.useupgradeinventoryslot < 3 { // on verifie que le joueur n'a pas deja utilisé toutes ses amélioration
		if p1.coins >= 100 { // on verifie que le joueur a assez d'argent pour acheter l'amélioration
			fmt.Println(p1.capacityinventorypotionchar, p1.capacityinventorypotion)
			p1.coins = p1.coins - 100                                   // on retire les 100 pieces d'or
			p1.useupgradeinventoryslot = p1.useupgradeinventoryslot + 1 // ajoute 1 au compteur d'utilisation de l'amélioration d'inventaire
			p1.capacityinventorypotionchar = p1.capacityinventorypotionchar + 10
			p1.capacityinventorypotion = p1.capacityinventorypotion + 5
			fmt.Println("vous avez amélioré votre inventaire de 10 cases, pour un total de", p1.capacityinventorypotionchar, "il vous reste ", p1.coins, " pièces d or")
			buy(p1, equipement1, Monster1) // on retourne au menu d'achat
		} else {
			fmt.Println("❌Vous n'avez pas assez d'argent, vous avez ", p1.coins, " pièces d'or ,  il vous en faut 100")
			buy(p1, equipement1, Monster1) // on retourne au menu d'achat
		}
	}
}

func sell(p1 Personnage, equipement1 Equipement, Monster1 Monster) { // fonction de vente: fait l'inverse de la fonction d'achat (retire des objets de l'inventaire du joueur et les ajoute à l'inventaire du marchand)
	var cout int
	fmt.Println("____________________________________")
	fmt.Println("|                                  |")
	fmt.Println("|============💰VENDRE💰===========|")
	fmt.Println("|__________________________________|")
	fmt.Println("____________________________________")
	var u int8
	var y string
	var cointemp int
	if len(p1.inventory) == 0 { // on verifie que l'inventaire du joueur n'est pas vide
		fmt.Println("Vous n'avez plus rien à vendre")
		DisplayMenu(p1, equipement1, Monster1)
	} else {

		fmt.Println("========================================================")
		fmt.Println("Votre inventaire:", p1.inventory, " Vous avez ", p1.coins, " pièces d'or")
		fmt.Println("========================================================")
		fmt.Println("--------------------------------------------------------")
		fmt.Println("========================================================")
		fmt.Println("Objet à vendre:", p1.inventorytrader)
		fmt.Println("========================================================")
		fmt.Println("--------------------------------------------------------")

		fmt.Println("________________________________________________________")
		fmt.Println("                                                      ")
		fmt.Println("      Commandes disponibles :1 , 2 , 3 , 4 , 5, 6 ")
		fmt.Println("________________________________________________________")
		fmt.Println("                                                      ")

		fmt.Println("Articles du Marchand:", p1.inventorytrader)
		fmt.Println("Votre inventaire", p1.inventory)
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
			if len(p1.inventory) >= 1 {
				if len(p1.inventory) == 1 { //avertit le joueur qu'il ne lui reste que 1 objet dans son inventaire
					fmt.Print("il ne vous reste que  1 article dans votre inventaire, Etes vous sur de vouloir de le  vendre ? votre article : ", p1.inventory[0])
					fmt.Scanln(&y)
					switch y {
					case "non":
						fmt.Println("==================")
						fmt.Println("❌Vente annulée❌")
						fmt.Println("==================")
						sell(p1, equipement1, Monster1)
						//si le joueur ne veut pas vendre son objet, il est redirigé vers le menu
					case "oui":
						if p1.inventory[0] == "Cv" {
							p1.coins = p1.coins + 35 // on ajoute 35 pieces d'or au joueur
							fmt.Println(p1.coins)

							cout = p1.coins - cointemp // on calcule le gain de l'achat qui est égal au nombre de pieces d'or du joueur apres l'achat moins le nombre de pieces d'or du joueur avant l'achat

						}
						if p1.inventory[0] == "potion de soin" { //si l'objet acheté est une potion de soin
							p1.coins = p1.coins + 3 //on retire 3 coins au joueur

							cout = p1.coins - cointemp

						}

						if p1.inventory[0] == "potion de poison" { //meme fonctionnement
							p1.coins = p1.coins + 3

							cout = p1.coins - cointemp

						}
						if p1.inventory[0] == "uzi" { //meme fonctionnement
							p1.coins = p1.coins + 15

							cout = p1.coins - cointemp

						}
						if p1.inventory[0] == "mp5" {
							p1.coins = p1.coins + 25

							cout = p1.coins - cointemp

						}
						if p1.inventory[0] == "livre de sort (boule de feu)" {
							p1.coins = p1.coins + 35 //meme fonctionnement

							cout = p1.coins - cointemp

						}
						if p1.inventory[0] == "allocation" { //si l'objet acheté est une potion de soin
							p1.coins = p1.coins + 3 //on retire 3 coins au joueur

							cout = p1.coins - cointemp

						}
						fmt.Println("======================================================================================================================")
						fmt.Println("Vous avez vendu un ", p1.inventory[:1], "pour un gain de ", cout, "d'or 💰, il vous reste ", p1.coins, " 💵 pièces d or")
						fmt.Println("======================================================================================================================")
						p1.inventorytrader = append(p1.inventorytrader, p1.inventory[:1]...)
						p1.inventory = append(p1.inventory[:0], p1.inventory[1:]...) //meme fonctiomnement que pour l'achat des objets mais dans l'autre sens
						fmt.Println(p1.inventorytrader)
						fmt.Println(p1.inventory)
						sell(p1, equipement1, Monster1)

					default:
						fmt.Println("===================================================================")
						fmt.Println("❌Vous n'avez pas choisi une option valide : Veuillez réessayer❌")
						fmt.Println("===================================================================")
						sell(p1, equipement1, Monster1)
					}
				} else {

					if p1.inventory[1] == "Cv" {
						p1.coins = p1.coins + 35

						cout = p1.coins - cointemp

					}
					if p1.inventory[1] == "potion de soin" { //si l'objet acheté est une potion de soin
						p1.coins = p1.coins + 3 //on retire 3 coins au joueur

						cout = p1.coins - cointemp

					}

					if p1.inventory[1] == "potion de poison" { //meme fonctionnement
						p1.coins = p1.coins + 3

						cout = p1.coins - cointemp

					}
					if p1.inventory[1] == "uzi" {
						p1.coins = p1.coins + 15

						cout = p1.coins - cointemp

					}
					if p1.inventory[1] == "mp5" {
						p1.coins = p1.coins + 25

						cout = p1.coins - cointemp

					}
					if p1.inventory[1] == "livre de sort (boule de feu)" {
						p1.coins = p1.coins + 35

						cout = p1.coins - cointemp

					}
					if p1.inventory[1] == "allocation" { //si l'objet acheté est une potion de soin
						p1.coins = p1.coins + 3 //on retire 3 coins au joueur

						cout = p1.coins - cointemp

					}
					fmt.Println("======================================================================================================================")
					fmt.Println("Vous avez vendu un ", p1.inventory[:1], "pour un gain de ", cout, "d'or 💰, il vous reste ", p1.coins, " 💵 pièces d or")
					fmt.Println("======================================================================================================================")
					p1.inventorytrader = append(p1.inventorytrader, p1.inventory[:1]...)
					p1.inventory = append(p1.inventory[:0], p1.inventory[1:]...) //meme fonctiomnement que pour l'achat des objets mais dans l'autre sens
					fmt.Println(p1.inventorytrader)
					fmt.Println(p1.inventory)
					sell(p1, equipement1, Monster1)
				}
			} else {
				fmt.Println("============================")
				fmt.Println("❌Transaction Impossible❌")
				fmt.Println("============================")
				sell(p1, equipement1, Monster1)
			}
		case 2: //indetique au cas 1  mais pour le deuxieme objet de l'inventaire

			cointemp = p1.coins
			fmt.Println(len(p1.inventory))
			if len(p1.inventory) >= 2 {
				fmt.Println(len(p1.inventory))
				if p1.inventory[1] == "allocation" { //si l'objet acheté est une potion de soin
					p1.coins = p1.coins + 3 //on retire 3 coins au joueur

					cout = p1.coins - cointemp

				}
				if p1.inventory[1] == "potion de soin" { //si l'objet acheté est une potion de soin
					p1.coins = p1.coins + 3 //on retire 3 coins au joueur

					cout = p1.coins - cointemp

				}

				if p1.inventory[1] == "potion de poison" { //meme fonctionnement
					p1.coins = p1.coins + 3

					cout = p1.coins - cointemp

				}
				if p1.inventory[1] == "uzi" {
					p1.coins = p1.coins + 15

					cout = p1.coins - cointemp

				}
				if p1.inventory[1] == "mp5" {
					p1.coins = p1.coins + 25

					cout = p1.coins - cointemp

				}
				if p1.inventory[1] == "livre de sort (boule de feu)" {
					p1.coins = p1.coins + 35

					cout = p1.coins - cointemp

				}
				if p1.inventory[1] == "cv" {
					p1.coins = p1.coins + 35

					cout = p1.coins - cointemp

				}
				fmt.Println("======================================================================================================================")
				fmt.Println("Vous avez vendu un ", p1.inventory[1], "pour un gain de ", cout, "d'or 💰, il vous reste ", p1.coins, " 💵 pièces d or")
				fmt.Println("======================================================================================================================")
				p1.inventorytrader = append(p1.inventorytrader, p1.inventory[1])
				p1.inventory = append(p1.inventory[:1], p1.inventory[2:]...)
				fmt.Println(p1.inventorytrader)
				fmt.Println(p1.inventory)
				sell(p1, equipement1, Monster1)
			} else {
				fmt.Println("======================")
				fmt.Println("❌Transaction Impossible❌")
				fmt.Println("======================")
				sell(p1, equipement1, Monster1)
			}

		case 3: //indetique au cas 1  mais pour le 3 eme objet de l'inventaire

			if len(p1.inventory) >= 3 {
				if p1.inventory[2] == "allocation" { //si l'objet acheté est une potion de soin
					p1.coins = p1.coins + 3 //on retire 3 coins au joueur

					cout = p1.coins - cointemp

				}
				if p1.inventory[2] == "potion de soin" { //si l'objet acheté est une potion de soin
					p1.coins = p1.coins + 3 //on retire 3 coins au joueur

					cout = p1.coins - cointemp

				}

				if p1.inventory[2] == "potion de poison" { //meme fonctionnement
					p1.coins = p1.coins + 3

					cout = p1.coins - cointemp

				}
				if p1.inventory[2] == "uzi" {
					p1.coins = p1.coins + 15

					cout = p1.coins - cointemp

				}
				if p1.inventory[2] == "mp5" {
					p1.coins = p1.coins + 25

					cout = p1.coins - cointemp

				}
				if p1.inventory[2] == "livre de sort (boule de feu)" {
					p1.coins = p1.coins + 35

					cout = p1.coins - cointemp

				}
				if p1.inventory[2] == "Cv" {
					p1.coins = p1.coins + 35

					cout = p1.coins - cointemp

				}
				fmt.Println("======================================================================================================================")
				fmt.Println("Vous avez vendu un ", p1.inventory[2], "pour un gain de ", cout, "d'or 💰, il vous reste ", p1.coins, " 💵 pièces d or")
				fmt.Println("======================================================================================================================")
				p1.inventorytrader = append(p1.inventorytrader, p1.inventory[2])
				p1.inventory = append(p1.inventory[:2], p1.inventory[3:]...)
				fmt.Println(p1.inventorytrader)
				fmt.Println(p1.inventory)
				sell(p1, equipement1, Monster1)
			} else {
				fmt.Println("======================")
				fmt.Println("❌Transaction Impossible❌")
				fmt.Println("======================")
				sell(p1, equipement1, Monster1)
			}
		case 4: //indetique au cas 1  mais pour le 4 eme objet de l'inventaire

			if len(p1.inventory) >= 4 {
				if p1.inventory[3] == "allocation" { //si l'objet acheté est une potion de soin
					p1.coins = p1.coins + 3 //on retire 3 coins au joueur

					cout = cointemp - p1.coins //la valeur de cout est égale à la valeur de cointemp - la valeur de coins
					fmt.Print(p1.coins)
				}
				if p1.inventory[3] == "potion de soin" { //si l'objet acheté est une potion de soin
					p1.coins = p1.coins + 3 //on retire 3 coins au joueur

					cout = cointemp - p1.coins //la valeur de cout est égale à la valeur de cointemp - la valeur de coins

				}

				if p1.inventory[3] == "potion de poison" { //meme fonctionnement
					p1.coins = p1.coins + 3

					cout = p1.coins - cointemp

				}
				if p1.inventory[3] == "uzi" {
					p1.coins = p1.coins + 15

					cout = p1.coins - cointemp

				}
				if p1.inventory[3] == "mp5" {
					p1.coins = p1.coins + 25

					cout = p1.coins - cointemp

				}
				if p1.inventory[3] == "livre de sort (boule de feu)" {
					p1.coins = p1.coins + 35

					cout = p1.coins - cointemp

				}
				if p1.inventory[3] == "Cv" {
					p1.coins = p1.coins + 35

					cout = p1.coins - cointemp

				}
				fmt.Println("======================================================================================================================")
				fmt.Println("Vous avez vendu un ", p1.inventory[3], "pour un gain de ", cout, "d'or 💰, il vous reste ", p1.coins, " 💵 pièces d or")
				fmt.Println("======================================================================================================================")
				p1.inventorytrader = append(p1.inventorytrader, p1.inventory[3])
				p1.inventory = append(p1.inventory[:3], p1.inventory[4:]...)
				fmt.Println(p1.inventorytrader)
				fmt.Println(p1.inventory)
				sell(p1, equipement1, Monster1)

			} else {
				fmt.Println("======================")
				fmt.Println("❌Transaction Impossible❌")
				fmt.Println("======================")
				sell(p1, equipement1, Monster1)
			}
		case 5: //indetique au cas 1  mais pour le 5 eme objet de l'inventaire

			if len(p1.inventory) >= 5 {
				if p1.inventory[4] == "allocation" { //si l'objet acheté est une potion de soin
					p1.coins = p1.coins + 3 //on retire 3 coins au joueur
					cout = p1.coins - cointemp

				}
				if p1.inventory[4] == "potion de soin" { //si l'objet acheté est une potion de soin
					p1.coins = p1.coins + 3 //on retire 3 coins au joueur
					cout = p1.coins - cointemp

				}

				if p1.inventory[4] == "potion de poison" { //meme fonctionnement
					p1.coins = p1.coins + 3
					cout = p1.coins - cointemp

				}
				if p1.inventory[4] == "uzi" {
					p1.coins = p1.coins + 15
					cout = p1.coins - cointemp

				}
				if p1.inventory[4] == "mp5" {
					p1.coins = p1.coins + 25
					cout = p1.coins - cointemp

				}
				if p1.inventory[4] == "livre de sort (boule de feu)" {
					p1.coins = p1.coins + 35
					cout = p1.coins - cointemp

				}
				if p1.inventory[4] == "Cv" {
					p1.coins = p1.coins + 35
					cout = p1.coins - cointemp

				}
				fmt.Println("======================================================================================================================")
				fmt.Println("Vous avez vendu un ", p1.inventory[4], "pour un gain de ", cout, "d'or 💰, il vous reste ", p1.coins, " 💵 pièces d or")
				fmt.Println("======================================================================================================================")
				p1.inventorytrader = append(p1.inventorytrader, p1.inventory[4])
				p1.inventory = append(p1.inventory[:4], p1.inventory[5:]...)
				sell(p1, equipement1, Monster1)
			} else {
				fmt.Println("===========================")
				fmt.Println("❌Transaction Impossible❌")
				fmt.Println("===========================")
				sell(p1, equipement1, Monster1)
			}
		case 6: // quitte le menu
			trader(p1, equipement1, Monster1)
		default:
			fmt.Println("==================================================================")
			fmt.Println("❌Vous n'avez pas choisi une option valide : Veuillez réessayer❌")
			fmt.Println("==================================================================")
			sell(p1, equipement1, Monster1)
		}
	}

}

func inventorylimit(p1 Personnage, equipement1 Equipement, Monster1 Monster) {
	if len(p1.inventory) == p1.capacityinventorypotionchar { //si l'inventaire est plein
		p1.isoverloaded = true //renvoie true si le joueur est surchargé
		sell(p1, equipement1, Monster1)
	}

	//fmt.Println("Vous avez vendu un ", p1.inventaire[0])
	//p1.inventairemarchand = append(p1.inventairemarchand, p1.inventaire[:1]...)
	//p1.inventaire = append(p1.inventaire[:0], p1.inventaire[len(p1.inventaire):]...)
	//fmt.Println(p1.inventairemarchand)
	//fmt.Println(p1.inventaire)

}
func removeinventory(p1 Personnage) {
	p1.inventory = nil
	p1.inventorytrader = nil
	p1.capacityinventorypotionchar = 0
	fmt.Println(p1.inventory)
	fmt.Println(p1.inventorytrader)

}

func removehealthpotion(p1 Personnage, equipement1 Equipement, Monster1 Monster) {
	potioncount := len(p1.inventorypotion)
	if len(p1.inventorypotion) == 0 { //si l'inventaire potion est vide
		fmt.Println("===============================")
		fmt.Println("❌Vous n'avez pas de potion❌")
		fmt.Println("===============================")
		DisplayMenu(p1, equipement1, Monster1)
	} else {
		if p1.Pv == p1.Pvmax { //si le joueur a déjà toute sa vie
			fmt.Println("================================")
			fmt.Println("❌Votre vie est déja au maximum❌")
			fmt.Println("potion(s) restante(s)", potioncount)
			fmt.Println("================================")
			DisplayMenu(p1, equipement1, Monster1)
		}
		i := len(p1.inventorypotion)
		p1.inventorypotion = append(p1.inventorypotion[:0], p1.inventorypotion[:i-1]...) //on retire une potion de l'inventaire
		if p1.Pv+25 > p1.Pvmax {                                                         //si la vie du joueur  après le soin est supérieur à sa vie max
			p1.Pv = p1.Pvmax //la vie du joueur est égale à sa vie max
		} else {
			p1.Pv = p1.Pv + 25 //sinon la vie du joueur est égale à sa vie actuelle + 25
		}
		potioncount = potioncount - 1 //on retire une potion du compteur
		fmt.Println("================================")
		fmt.Println("vous avez consommé une potion de soin")
		fmt.Println("potion(s) restante(s)", potioncount)
		fmt.Println("================================")
		DisplayMenu(p1, equipement1, Monster1)

	}
}

func buyhealthpotion(p1 Personnage, equipement1 Equipement, Monster1 Monster) { //Fonction permettant d'acheter un objet
	if len(p1.inventorytrader) == 0 { //si l'inventaire du marchand est vide
		fmt.Println("===================================")
		fmt.Println("❌Le marchand n'a plus rien à vendre❌")
		fmt.Println("===================================")
		DisplayMenu(p1, equipement1, Monster1) //retour au menu principal

	} else {
		var t int8
		fmt.Println("____________________________________")
		fmt.Println("|                                  |")
		fmt.Println("|=============ACHETER💲============|")
		fmt.Println("|__________________________________|")
		fmt.Println("                                                      ")
		fmt.Println("________________________________________________________")
		fmt.Println("                                                        ")
		fmt.Println("      Commandes disponibles :1, 2                        ")
		fmt.Println("________________________________________________________")
		fmt.Println("                                                      ")
		fmt.Println("____________________________________________________________________________")
		fmt.Println("                                                        ")
		fmt.Println(" Inventaire du marchand:  Potion de Soin quantité: ∞")
		fmt.Println("               Votre argent:", p1.coins, "                                    ")
		fmt.Println("____________________________________________________________________________")
		fmt.Println("choisissez parmis ces options")
		fmt.Println("____________________________________")
		fmt.Println(" ACHAT:")
		fmt.Println("____________________________________")
		fmt.Println("1. Acheter une potion de soin (vous en avez")
		fmt.Println("____________________________________")
		fmt.Println("AUTRES:")
		fmt.Println("____________________________________")
		fmt.Println("2.  Quitter →")
		fmt.Print("$Choississez le produit que vous souhaitez acheter$ ")
		fmt.Scanln(&t)
		switch t {

		case 1:
			if p1.coins >= 10 {
				if len(p1.inventorypotion) < p1.capacityinventorypotion { //si la limite de l'innventaire potion est atteinte
					p1.inventorypotion = append(p1.inventorypotion, "potion")
					p1.coins = p1.coins - 10
					fmt.Println("========================================================")
					fmt.Println("Vous avez acheté une potion de soin pour 10 pièces d'or")
					fmt.Println("========================================================")
					fmt.Println("Nombre de potion de soin possédées", len(p1.inventorypotion))
					buyhealthpotion(p1, equipement1, Monster1)

				} else {
					fmt.Println("====================================================================================================================")
					fmt.Println("❌Vous ne pouvez pas avoir plus de ", p1.capacityinventorypotion, " potions, pour cela améliorez votre inventaire❌")
					fmt.Println("====================================================================================================================")
					buyhealthpotion(p1, equipement1, Monster1)
				}
			} else {
				fmt.Println("Vous n'avez pas assez d'argent : il vous faut 10 pièces d'or, vous n'en avez que", p1.coins)
				buyhealthpotion(p1, equipement1, Monster1)
			}

		case 2:
			DisplayMenu(p1, equipement1, Monster1)

		default:
			fmt.Println("============================================")
			fmt.Println("❌Vous n'avez pas choisi une option valide❌")
			fmt.Println("============================================")
			buyhealthpotion(p1, equipement1, Monster1)
		}

	}
}

func iswasted(p1 Personnage, equipement1 Equipement, Monster1 Monster) { // fonction isdead()
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
			DisplayMenu(p1, equipement1, Monster1)

		case 2:
			fmt.Println("Jeu Fermé")
			fmt.Println("Vous avez quitté le jeu, Vous feriez mieux la prochaine fois")
		default:
			fmt.Println("Vous n'avez pas choisi une option valide")
			iswasted(p1, equipement1, Monster1)
		}
	} else {
		DisplayMenu(p1, equipement1, Monster1)
	}
}

func Poisonpot(p1 Personnage, equipement1 Equipement, Monster1 Monster) {
	fmt.Println("Vous subissez les effets du poison")
	for i := 0; i < 3; i++ { //boucle qui fait perdre 10 pv pendant 3 secondes
		time.Sleep(1 * time.Second) //temps d'attente fixé à une  1 seconde
		p1.Pv = p1.Pv - 10          //-10pv pour le joueur
		fmt.Println("-10 Pv 🤢 vous etes mainetant à ", p1.Pv, "♥/", p1.Pvmax, "♥")
		if p1.Pv <= 0 { //si la vie du joueur est inférieur ou égale à 0
			p1.Pv = 0 // evite d'affiche un nombre  de vie négatif
		}
		fmt.Println("vous avez ", p1.Pv, "/", p1.Pvmax, "PV")
	}

	iswasted(p1, equipement1, Monster1) //on vérifie si le joueur est mort

}

func SpellBook(p1 Personnage, equipement1 Equipement, Monster1 Monster) {
	if p1.islearned == 0 { //si le joueur n'a pas encore appris le sort
		p1.skill = append(p1.skill, "Boule de feu")                                                //on ajoute le sort à la liste des sorts appris
		p1.persoinventairesort = append(p1.persoinventairesort[:0], p1.persoinventairesort[1:]...) //on retire le sort de son inventaire
		fmt.Println("Vous avez appris le sort boule de feu ")
		p1.islearned = 1 // la valeur de islearned passe à 1: le joueur a appris le sort
		fmt.Println(p1.skill)
		DisplayMenu(p1, equipement1, Monster1)
	} else { //si le joueur a déjà appris le sort
		fmt.Println("====================================")
		fmt.Println("❌Vous avez déjà appris ce sort❌")
		fmt.Println("====================================")
		DisplayMenu(p1, equipement1, Monster1)
	}
}

func Quitgame(p1 Personnage, equipement1 Equipement, Monster1 Monster) {
	var h string
	fmt.Println("=========================================")
	fmt.Print("- Voulez-vous vraiment quitter le jeu ?- ")

	fmt.Scanln(&h)
	switch h {
	case "oui": //si le joueur répond oui
		//p1.inventaire = nil
		//p1.inventairemarchand = nil
		//p1.capacitéinventaireperso = 0
		p1.isgameshutdown = true

		fmt.Println("====================================")
		fmt.Println("====================================")
		fmt.Println("Jeu Fermé") //informe le joueur que le jeu est fermé
		fmt.Println("Merci d'avoir joué")
		fmt.Println("====================================")
		fmt.Println("====================================")

	case "non":
		DisplayMenu(p1, equipement1, Monster1) //si le joueur répond non, on revient au menu
	default:
		fmt.Println("=============================================================")
		fmt.Println("❌Vous n'avez pas choisi une option valide : Veuillez réessayer❌")
		fmt.Println("=============================================================")
		Quitgame(p1, equipement1, Monster1)

	}

}

func charCreation(p1 Personnage, equipement1 Equipement, Monster1 Monster) { //fonction de création de personnage
	var name string
	var classe int8
	fmt.Print("Quel est votre nom ? ")
	fmt.Scanln(&name)

	for i := 0; i < len(name); i++ { //boucle qui vérifie si le nom du joueur est supérieur à 1
		if len(name) <= 1 {
			fmt.Println("Votre nom est trop court")
			charCreation(p1, equipement1, Monster1)
			break
		}
		if name[i] < 'A' || name[i] > 'Z' {
			fmt.Println("Le nom doit commencer par une majuscule") //boucle qui vérifie si le nom ne commence pas par une majuscule
			charCreation(p1, equipement1, Monster1)
			break

		}

		if name[i+1] < 'a' || name[i] > 'z' {
			fmt.Println("Le nom doit continuer par une minuscule") //boucle qui vérifie si le nom ne continue pas par une minuscule
			charCreation(p1, equipement1, Monster1)

			break
		} else {

			fmt.Println("Bonjour", name, "!")

			fmt.Print("choisissez votre classe : ")
			fmt.Scanln(&classe)
			switch classe {
			case 1:
				p1.Class = "Rentier"
				p1.Pvmax = 100
				p1.Pv = p1.Pvmax / 2 //le joueur commence avec la moitié de ses points de vie
				fmt.Println(p1.Pv)
				fmt.Println("Vous avez choisi la classe Rentier") //ces 3 lignes initialise les  différentes classes
				DisplayMenu(Personnage{p1.Nom, "Rentier", 1, p1.Pvmax, p1.Pv, 10, 0, []string{"Cv", "allocation"}, []string{"potion", "potion", "potion"}, 10, 3, false, []string{"potion de soin", "potion de poison", "uzi", "mp5", "ak47"}, []string{"Coup de poing"}, 0, 0, 100, []string{"livre de sort (boule de feu)"}, []string{}, []string{"fourrure de loup", "peau de troll", "cuir de sanglier", "plume de corbeaux", "fourrure de warg"}, []string{}, 0, false, false, false, false, 0, 10, false, 1000, 100, 100, false}, Equipement{}, Monster{})

			case 2:
				p1.Class = "Bénéficiaire du RSA"
				p1.Pvmax = 50
				p1.Pv = p1.Pvmax / 2 //le joueur commence avec la moitié de ses points de vie
				fmt.Println(p1.Pv)
				fmt.Println("Vous avez choisi la classe Bénéficiaire du RSA") //ces 3 lignes initialise les  différentes classes
				DisplayMenu(Personnage{p1.Nom, "Bénéficiaire du RSA", 1, p1.Pvmax, p1.Pv, 10, 0, []string{"Cv", "allocation"}, []string{"potion", "potion", "potion"}, 10, 5, false, []string{"potion de soin", "potion de poison", "uzi", "mp5", "ak47"}, []string{"Coup de poing"}, 0, 0, 100, []string{"livre de sort (boule de feu)"}, []string{}, []string{"fourrure de loup", "peau de troll", "cuir de sanglier", "plume de corbeaux", "fourrure de warg"}, []string{}, 0, false, false, false, false, 0, 10, false, 1000, 100, 100, false}, Equipement{}, Monster{})

			case 3:
				p1.Class = "Chomeur"
				p1.Pvmax = 600
				p1.Pv = p1.Pvmax / 2 //le joueur commence avec la moitié de ses points de vie
				fmt.Println(p1.Pv)
				fmt.Println("Vous avez choisi la classe Chomeur") //ces 3 lignes initialise les  différentes classes
				DisplayMenu(Personnage{p1.Nom, "Chômeur", 1, p1.Pvmax, p1.Pv, 10, 0, []string{"Cv", "allocation"}, []string{"potion", "potion", "potion"}, 10, 5, false, []string{"potion de soin", "potion de poison", "uzi", "mp5", "ak47"}, []string{"Coup de poing"}, 0, 0, 100, []string{"livre de sort (boule de feu)"}, []string{}, []string{"fourrure de loup", "peau de troll", "cuir de sanglier", "plume de corbeaux", "fourrure de warg"}, []string{}, 0, false, false, false, false, 0, 10, false, 1000, 100, 100, false}, Equipement{}, Monster{})

			default:

				fmt.Println("❌Vous n'avez pas choisi une classe valide❌")
				fmt.Println("⚠️La création du personnage a échoué ⚠️ ")
				fmt.Println("Vous allez entre renvoyé au menu principal")
				fmt.Println("votre personnage a été crée automatiquement:")
				fmt.Println("Retour au menu principal.")
				time.Sleep(1 * time.Second)
				fmt.Println("Retour au menu principal..")
				time.Sleep(1 * time.Second)
				fmt.Println("Retour au menu principal...")
				time.Sleep(1 * time.Second)
				fmt.Println("Retour au menu principal.")
				time.Sleep(1 * time.Second)
				fmt.Println("Retour au menu principal...")
				time.Sleep(1 * time.Second)
				fmt.Println("Retour au menu principal...")
				fmt.Println("si nous n'etes pas satisfait de votre personnage, vous pouvez le réinitialiser en appuyant sur 7 dans le menu principal")
				p1.Nom = "Joueur"
				p1.Pvmax = 100
				p1.Pv = p1.Pvmax / 2 //le joueur commence avec la moitié de ses points de vie
				DisplayMenu(Personnage{p1.Nom, "Rentier", 1, p1.Pvmax, p1.Pv, 10, 0, []string{"Cv", "allocation"}, []string{"potion", "potion", "potion"}, 10, 3, false, []string{"potion de soin", "potion de poison", "uzi", "mp5", "ak47"}, []string{"Coup de poing"}, 0, 0, 100, []string{"livre de sort (boule de feu)"}, []string{}, []string{"fourrure de loup", "peau de troll", "cuir de sanglier", "plume de corbeaux", "fourrure de warg"}, []string{}, 0, false, false, false, false, 0, 10, false, 1000, 100, 100, false}, Equipement{}, Monster{})

				//pour sortir du switch case
			}
			//break
		}
		break // pour sortir de la boucle for
	}

}
func forge(p1 Personnage, equipement1 Equipement, Monster1 Monster) { //fonction de craft
	var choix int8
	//var choix2 int8
	fmt.Println("____________________________________")
	fmt.Println("|                                  |")
	fmt.Println("|=============FORGE🔨==============|")
	fmt.Println("|__________________________________|")
	fmt.Println("____________________________________")
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
		if len(p1.charmeterialsinventory) <= 1 { //si le joueur n'a pas assez de matériaux pour forger une armure
			fmt.Println("=========================================================")
			fmt.Println("Vous n'avez pas assez de matériaux pour forger une armure❌")
			fmt.Println("=========================================================")
			forge(p1, equipement1, Monster1)
		} else {
			var choix2 int
			fmt.Println("=========================================")
			fmt.Println("1. Casque")
			fmt.Println("2. Plastron")
			fmt.Println("3. Bottes")
			fmt.Println("6. choissisez une option : ")
			fmt.Scanln(&choix2)
			fmt.Println("=========================================")
			switch choix2 {
			case 1:
				fmt.Scanln(&choix2)
				fmt.Println(p1.charmeterialsinventory)
				if p1.isoverloaded == true { //si le joueur est surchargé
					fmt.Println("=======================================================")
					fmt.Println("Vous êtes surchargé, vous ne pouvez pas forger d'armure")
					fmt.Println("=======================================================")
					DisplayMenu(p1, equipement1, Monster1)
				} else {

					fmt.Println("etape 1")
					if p1.coins >= 25 { //si le joueur a assez d'argent pour forger une armure
						fmt.Println("etape 2")
						for i := 1; i < len(p1.charmeterialsinventory); i++ { //on parcourt le tableau de matériaux du joueur
							if p1.charmeterialsinventory[i] == "cuir de sanglier" { //si le joueur a assez de cuir de sanglier (s'il est est présent dans son inventaire)
								true1 = i
								for j := 1; j < len(p1.charmeterialsinventory); j++ { //on parcourt le tableau de matériaux du joueur
									if p1.charmeterialsinventory[i+1] == "plume de corbeaux" { //si le joueur a assez de plume de corbeaux (s'il est est présent dans son inventaire)
										true2 = p1.charmeterialsinventory[i+1] //on stocke la valeur de la plume de corbeaux dans une variable
										fmt.Println(true2)
										fmt.Println("true2", i, p1.charmeterialsinventory[i+1])
										equipement1.helmetinventory = append(equipement1.helmetinventory, "casquette de chomeur") //on ajoute l'armure a l'inventaire  d'armure du joueur
										p1.charmeterialsinventory = append(p1.charmeterialsinventory[:true1])                     //on supprime  avant l'index true1
										fmt.Print(equipement1.helmetinventory)
										fmt.Print(p1.charmeterialsinventory)
										fmt.Println("--------------------------------------------------------------------------------")
										fmt.Println("Vous avez forgé une 🔨casquette de chomeur🔨 pour le prix de 25 pièces d'or 💵")
										fmt.Println("--------------------------------------------------------------------------------")
										forge(p1, equipement1, Monster1)

									} else {
										fmt.Println("false2", i, p1.charmeterialsinventory[i+1])
										fmt.Println("la condition est ", p1.charmeterialsinventory[i+1] == "plume de corbeau")
										fmt.Println("Vous n'avez pas les matériaux nécessaires pour forger une armure")
									}

								}
							} else {
								fmt.Println("false1", i, p1.charmeterialsinventory[i])
								if len(p1.charmeterialsinventory[i]) == len(p1.charmeterialsinventory) {
									fmt.Println("Vous n'avez pas les matériaux nécessaires pour forger une armure")
								}

							}

						}

					} else {
						fmt.Println("Vous n'avez pas assez de pièces") //cas ou le joueur n'a pas assez d'argent pour forger une armure
						DisplayMenu(p1, equipement1, Monster1)
					}

				}

			case 2:
				fmt.Println(p1.charmeterialsinventory)
				if p1.isoverloaded == true {
					fmt.Println("🏋Vous êtes surchargé, vous ne pouvez pas forger d'armure 🏋")
					DisplayMenu(p1, equipement1, Monster1)
				} else {
					var count int
					fmt.Println("etape 1")
					if p1.coins >= 25 {
						fmt.Println("etape 2") // meme principe que pour la casquette de chomeur mais avec  la fourrure de loup et la peau de Troll
						for i := 1; i < len(p1.charmeterialsinventory); i++ {
							count = count + 1
							fmt.Println(count)
							fmt.Println("efd")
							if p1.charmeterialsinventory[i] == "peau de troll" {
								true1 = i
								fmt.Println("true1", i, p1.charmeterialsinventory[i])
								for j := 0; j < len(p1.charmeterialsinventory); j++ {
									fmt.Print("j", j, len(p1.charmeterialsinventory))
									if p1.charmeterialsinventory[j] == "fourrure de loup" {
										fmt.Println("true2", i, p1.charmeterialsinventory[i+1])
										equipement1.chestinventory = append(equipement1.chestinventory, "sweat de chomeur")
										p1.charmeterialsinventory = append(p1.charmeterialsinventory[:true1])
										p1.coins = p1.coins - 25
										fmt.Print(equipement1.chestinventory)
										fmt.Print(p1.charmeterialsinventory)
										fmt.Println("--------------------------------------------------------------------------------")
										fmt.Println("Vous avez forgé 🔨un sweat de chomeur🔨 pour le prix de 25 pièces d'or 💵")
										fmt.Println("--------------------------------------------------------------------------------")
										forge(p1, equipement1, Monster1)
										break

									} else {
										fmt.Println("false2", i, p1.charmeterialsinventory[i+1])
										//fmt.Println("la condition est ", p1.charmeterialsinventory[i+1] == "plume de corbeau")
										//fmt.Println("Vous n'avez pas les matériaux nécessaires pour forger une armure", len(p1.persoinventairemateriaux))
										if len(p1.charmeterialsinventory[i]) == len(p1.charmeterialsinventory) {
											fmt.Println(p1.charmeterialsinventory[i])
										}
									}

								}
							} else {
								fmt.Println("false1", i, p1.charmeterialsinventory[i])
								if len(p1.charmeterialsinventory[i]) == len(p1.charmeterialsinventory) {
									fmt.Println("Vous n'avez pas les matériaux nécessaires pour forger une armure")
								}

							}

						}

					} else {
						fmt.Println("===============================") //cas ou le joueur n'a pas assez d'argent pour forger une armure
						fmt.Println("❌Vous n'avez pas assez de pièces❌")
						fmt.Println("===============================")
						DisplayMenu(p1, equipement1, Monster1)
					}

				}

			case 3:
				fmt.Println(p1.charmeterialsinventory)
				if p1.isoverloaded == true {
					fmt.Println("Vous êtes surchargé, vous ne pouvez pas forger d'armure")
					DisplayMenu(p1, equipement1, Monster1)
				} else {
					var count int
					fmt.Println("etape 1")
					if p1.coins >= 25 {
						fmt.Println("etape 2")
						for i := 0; i < len(p1.charmeterialsinventory); i++ {
							count = count + 1
							fmt.Println(count) // meme principe que pour la casquette de chomeur mais avec  la fourrure de loup et le cuir de sanglier
							fmt.Println("efd")
							if p1.charmeterialsinventory[i] == "fourrure de loup" {
								true1 = i
								fmt.Println("true1", i, p1.charmeterialsinventory[i])
								for j := 0; j < len(p1.charmeterialsinventory); j++ {
									fmt.Print("j", j, len(p1.charmeterialsinventory))
									if p1.charmeterialsinventory[j] == "cuir de sanglier" {
										fmt.Println("true2", i, p1.charmeterialsinventory[i+1])
										equipement1.bootinventory = append(equipement1.bootinventory, "claquette")
										p1.charmeterialsinventory = append(p1.charmeterialsinventory[:true1])
										p1.coins = p1.coins - 25
										fmt.Print(equipement1.bootinventory)
										fmt.Print(p1.charmeterialsinventory)
										fmt.Println("--------------------------------------------------------------------------------")
										fmt.Println("Vous avez forgé 🔨 une claquette 🔨 pour le prix de 25 pièces d'or 💵")
										fmt.Println("--------------------------------------------------------------------------------")
										forge(p1, equipement1, Monster1)
										break
									} else {
										fmt.Println("false2", i, p1.charmeterialsinventory[i+1])
										//fmt.Println("la condition est ", p1.persoinventairemateriaux[i+1] == "plume de corbeau")
										//fmt.Println("Vous n'avez pas les matériaux nécessaires pour forger une armure", len(p1.persoinventairemateriaux))
										if len(p1.charmeterialsinventory[i]) == len(p1.charmeterialsinventory) {
											fmt.Println(p1.charmeterialsinventory[i])
										}
									}

								}
							} else {
								fmt.Println("false1", i, p1.charmeterialsinventory[i])
								if len(p1.charmeterialsinventory[i]) == len(p1.charmeterialsinventory) {
									fmt.Println("Vous n'avez pas les matériaux nécessaires pour forger une armure")
								}

							}

						}

					} else {
						fmt.Println("===============================") //cas ou le joueur n'a pas assez d'argent pour forger une armure
						fmt.Println("❌Vous n'avez pas assez de pièces❌")
						fmt.Println("===============================")
						DisplayMenu(p1, equipement1, Monster1)
					}

				}
			}
		}

	case 4:
		DisplayMenu(p1, equipement1, Monster1) // retour au menu principal

	default:
		fmt.Println("Vous n'avez pas choisi une option valide")
		forge(p1, equipement1, Monster1)
	}
}

func Armor(equipement1 Equipement, p1 Personnage, Monster1 Monster) { // fonction qui permet de choisir l'armure que l'on veut porter
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
	fmt.Println("2. Quitter le jeu")
	fmt.Println("=========================================")
	fmt.Print("Veuillez choisir une option : ")
	fmt.Scanln(&choix)
	switch choix {
	case 1: // on choisit le type d'armure que l'on veut équiper
		fmt.Println()
		fmt.Println("=========================================")
		fmt.Print("choississez une armure à équiper")
		fmt.Println("=========================================")
		fmt.Println("1. Casque")
		fmt.Println("2.Plastron")
		fmt.Println("3.bottes")
		fmt.Println("==========================================")
		fmt.Scanln(&choix2)
		switch choix2 {
		case 1:
			fmt.Println(equipement1.helmetinventory)
			if p1.havehelmet == false { // si le joueur n'a pas de casque
				if len(equipement1.helmetinventory) == 0 {
					fmt.Println("Vous n'avez pas de casque dans votre inventaire")
					Armor(equipement1, p1, Monster1)
				} else {
					var choix3 int8
					fmt.Print("Choississez une armure à équiper")
					fmt.Println(equipement1.helmetinventory)
					fmt.Scan(&choix3)

					switch choix3 {
					case 1: // on équipe le casque  à la première position de l'inventaire
						if len(equipement1.helmetinventory) == 1 { // si l'inventaire contient qu'un seul casque  ( meme condition que pour acheter et vendre permet de gerer l'exception "index out of range")
							equipement1.helmetequiped = append(equipement1.helmetinventory, equipement1.helmetinventory[:0]...)
							equipement1.helmetinventory = nil                           // on vide l'inventaire,  nil permet de gerer l'exception "index out of range"
							if equipement1.helmetequiped[0] == "casquette de chomeur" { // on ajoute les stats propres au  casque à celles du joueur
								p1.Pvmax = p1.Pvmax + 10
							}
							p1.havehelmet = true // on indique que le joueur a un casque équipé
							fmt.Println(equipement1.helmetequiped)
							Armor(equipement1, p1, Monster1)

						} else { // si l'inventaire contient plus d'un casque
							fmt.Println(equipement1.helmetinventory)
							equipement1.helmetequiped = append(equipement1.helmetinventory, equipement1.helmetinventory[:1]...)       //on equipe l'objet choisi dans l'inventaire du joueur
							equipement1.helmetinventory = append(equipement1.helmetinventory[:0], equipement1.helmetinventory[1:]...) // on supprime l'objet choisi de l'inventaire du joueur
							if equipement1.helmetequiped[0] == "casquette de chomeur" {
								p1.Pvmax = p1.Pvmax + 10
							}
							p1.havehelmet = true // on indique que le joueur a un casque équipé
							fmt.Println(equipement1.helmetinventory)
							fmt.Println(equipement1.helmetequiped)
							fmt.Println("choississez une armure à équiper")
							fmt.Println("Vous avez équiper un casque")
							Armor(equipement1, p1, Monster1)
						}
					case 2:
						if len(equipement1.chestinventory) >= 2 {
							fmt.Println(equipement1.helmetinventory)
							equipement1.chestequiped = append(equipement1.chestinventory, equipement1.chestinventory[:1]...) //on equipe l'objet choisi dans l'inventaire du joueur
							equipement1.chestinventory = append(equipement1.chestinventory[:0], equipement1.chestinventory[1:]...)
							p1.havechestplate = true
							if equipement1.chestequiped[0] == "sweat de chomeur" {
								p1.Pvmax = p1.Pvmax + 20
							} // meme principe pour ce cas
							fmt.Println(equipement1.chestinventory)
							fmt.Println(equipement1.chestequiped)
							fmt.Println("choississez une armure à équiper")
							fmt.Println("Vous avez équipé un(e) ", equipement1.chestequiped)
							Armor(equipement1, p1, Monster1)

						} else {
							fmt.Println("=================")
							fmt.Println("❌Action impossible❌")
							fmt.Println("=================")
							Armor(equipement1, p1, Monster1)
						}
					case 3:
						if len(equipement1.chestinventory) >= 3 {
							fmt.Println(equipement1.helmetinventory)
							equipement1.chestequiped = append(equipement1.chestinventory, equipement1.chestinventory[:1]...) //on ajoute l'objet acheté dans l'inventaire du joueur
							equipement1.chestinventory = append(equipement1.chestinventory[:0], equipement1.chestinventory[1:]...)

							p1.havechestplate = true
							if equipement1.chestequiped[0] == "sweat de chomeur" {
								p1.Pvmax = p1.Pvmax + 20
							}
							fmt.Println(equipement1.chestinventory) // meme principe pour ce cas
							fmt.Println(equipement1.chestequiped)
							fmt.Println("choississez une armure à équiper")
							fmt.Println("Vous avez équipé un(e) ", equipement1.chestequiped)
							Armor(equipement1, p1, Monster1)

						} else {
							fmt.Println("Action impossible")
							Armor(equipement1, p1, Monster1)
						}
					default:
						fmt.Println("========================================")
						fmt.Println("❌Vous n'avez pas choisi une option valide❌")
						fmt.Println("========================================")
					}
				}
			} else {
				fmt.Println("Vous avez déjà un casque équipé") //  cas si le joueur a déjà un casque équipé
				Armor(equipement1, p1, Monster1)
			}
		case 2: //meme principe que  pour les casques
			if p1.havechestplate == false {
				fmt.Println(equipement1.chestinventory)
				if len(equipement1.chestinventory) == 0 {
					fmt.Println("Vous n'avez pas de plastron dans votre inventaire")
					Armor(equipement1, p1, Monster1)
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
							if equipement1.chestequiped[0] == "sweat de chomeur" {
								p1.Pvmax = p1.Pvmax + 20
							}
							fmt.Println(equipement1.helmetequiped)
							var tet []string
							tet = nil
							fmt.Println(tet)
							tet = append(tet, "casque")
							fmt.Println(tet)
							Armor(equipement1, p1, Monster1)

						} else {
							fmt.Println(equipement1.helmetinventory)
							equipement1.chestequiped = append(equipement1.chestinventory, equipement1.chestinventory[:1]...) //on ajoute l'objet acheté dans l'inventaire du joueur
							equipement1.chestinventory = append(equipement1.chestinventory[:0], equipement1.chestinventory[1:]...)

							p1.havechestplate = true
							if equipement1.chestequiped[0] == "sweat de chomeur" {
								p1.Pvmax = p1.Pvmax + 20
							}
							fmt.Println(equipement1.chestinventory)
							fmt.Println(equipement1.chestequiped)
							fmt.Println("choississez une armure à équiper")
							fmt.Println("Vous avez équipé un(e) ", equipement1.chestequiped)
							Armor(equipement1, p1, Monster1)

						}
					case 2:
						if len(equipement1.chestinventory) >= 2 {
							fmt.Println(equipement1.helmetinventory)
							equipement1.chestequiped = append(equipement1.chestinventory, equipement1.chestinventory[:1]...) //on ajoute l'objet acheté dans l'inventaire du joueur
							equipement1.chestinventory = append(equipement1.chestinventory[:0], equipement1.chestinventory[1:]...)

							p1.havechestplate = true
							if equipement1.chestequiped[0] == "sweat de chomeur" {
								p1.Pvmax = p1.Pvmax + 20
							}
							fmt.Println(equipement1.chestinventory)
							fmt.Println(equipement1.chestequiped)
							fmt.Println("choississez une armure à équiper")
							fmt.Println("Vous avez équipé un(e) ", equipement1.chestequiped)
							Armor(equipement1, p1, Monster1)

						} else {
							fmt.Println("Action impossible")
							Armor(equipement1, p1, Monster1)
						}
					case 3:
						if len(equipement1.chestinventory) >= 3 {
							fmt.Println(equipement1.helmetinventory)
							equipement1.chestequiped = append(equipement1.chestinventory, equipement1.chestinventory[:1]...) //on ajoute l'objet acheté dans l'inventaire du joueur
							equipement1.chestinventory = append(equipement1.chestinventory[:0], equipement1.chestinventory[1:]...)

							p1.havechestplate = true
							if equipement1.chestequiped[0] == "sweat de chomeur" {
								p1.Pvmax = p1.Pvmax + 20

								fmt.Println(equipement1.chestinventory)
								fmt.Println(equipement1.chestequiped)
								fmt.Println("choississez une armure à équiper")
								fmt.Println("Vous avez équipé un(e) ", equipement1.chestequiped)
								Armor(equipement1, p1, Monster1)
							}
						} else {
							fmt.Println("Action impossible")
							Armor(equipement1, p1, Monster1)
						}
					default:
						fmt.Println("Vous n'avez pas choisi une option valide")
						Armor(equipement1, p1, Monster1)
					}
				}
			} else {
				fmt.Println("Vous avez déjà équipé une armure")
				Armor(equipement1, p1, Monster1)
			}
		case 3: //meme principe que  pour les casques
			if p1.haveboots == false {

				fmt.Println(equipement1.bootinventory)
				if len(equipement1.bootinventory) == 0 {
					fmt.Println("Vous n'avez pas de botte dans votre inventaire")
					Armor(equipement1, p1, Monster1)
				} else {
					var choix3 int8
					fmt.Print("Choississez une armure à équiper")
					fmt.Println(equipement1.bootinventory)
					fmt.Scan(&choix3)

					switch choix3 {
					case 1:
						if len(equipement1.bootinventory) == 1 {
							equipement1.bootequiped = append(equipement1.bootinventory, equipement1.bootinventory[:0]...)
							equipement1.bootinventory = nil
							p1.haveboots = true
							if equipement1.bootequiped[0] == "claquette" {
								p1.Pvmax = p1.Pvmax + 20
							}
							fmt.Println(equipement1.bootequiped)
							var tet []string
							tet = nil
							fmt.Println(tet)
							tet = append(tet, "casque")
							fmt.Println(tet)
							Armor(equipement1, p1, Monster1)

						} else {
							fmt.Println(equipement1.bootinventory)
							equipement1.bootequiped = append(equipement1.bootinventory, equipement1.bootinventory[:1]...) //on ajoute l'objet acheté dans l'inventaire du joueur
							equipement1.bootinventory = append(equipement1.bootinventory[:0], equipement1.bootinventory[1:]...)

							p1.haveboots = true
							if equipement1.bootequiped[1] == "claquette" {
								p1.Pvmax = p1.Pvmax + 20
							}
							fmt.Println(equipement1.bootinventory)
							fmt.Println(equipement1.bootequiped)
							fmt.Println("choississez une armure à équiper")
							fmt.Println("Vous avez équiper un casque")
							Armor(equipement1, p1, Monster1)
						}
					case 2:
						if len(equipement1.bootinventory) >= 2 {
							fmt.Println(equipement1.bootinventory)
							equipement1.bootequiped = append(equipement1.bootinventory, equipement1.bootinventory[:1]...) //on ajoute l'objet acheté dans l'inventaire du joueur
							equipement1.bootinventory = append(equipement1.bootinventory[:0], equipement1.bootinventory[1:]...)

							p1.haveboots = true
							if equipement1.bootequiped[1] == "claquette" {
								p1.Pvmax = p1.Pvmax + 20
							}
							fmt.Println(equipement1.bootinventory)
							fmt.Println(equipement1.bootequiped)
							fmt.Println("choississez une armure à équiper")
							fmt.Println("Vous avez équiper un casque")
							Armor(equipement1, p1, Monster1)

						} else {
							fmt.Println("Action impossible")
							Armor(equipement1, p1, Monster1)
						}
					case 3:
						if len(equipement1.bootinventory) >= 3 {
							fmt.Println(equipement1.bootinventory)
							equipement1.bootequiped = append(equipement1.bootinventory, equipement1.bootinventory[:1]...) //on ajoute l'objet acheté dans l'inventaire du joueur
							equipement1.bootinventory = append(equipement1.bootinventory[:0], equipement1.bootinventory[1:]...)

							p1.haveboots = true
							if equipement1.bootequiped[1] == "claquette" {
								p1.Pvmax = p1.Pvmax + 20
							}

							fmt.Println(equipement1.bootinventory)
							fmt.Println(equipement1.bootequiped)
							fmt.Println("choississez une armure à équiper")
							fmt.Println("Vous avez équiper un casque")
							Armor(equipement1, p1, Monster1)
						} else {
							fmt.Println("======================")
							fmt.Println("❌Action impossible❌")
							fmt.Println("======================")
							Armor(equipement1, p1, Monster1)
						}
					}
				}
			} else {
				fmt.Println("Vous avez déjà équipé une armure")
				Armor(equipement1, p1, Monster1)
			}

		default:
			fmt.Println("Vous n'avez pas choisi une option valide")
			Armor(equipement1, p1, Monster1)
		}
	case 2:
		DisplayMenu(p1, equipement1, Monster1)
	default:
		fmt.Println("Vous n'avez pas choisi une option valide")
		Armor(equipement1, p1, Monster1)
	}
}

type Monster struct { //structure monstre
	Nom         string
	Pvmax       int
	Pv          int
	ptsdamages  int
	initiative  int
	turn        int
	attackturn  int
	hasattacked bool
	turnbuff    int
}

func InitGobelin(p1 Personnage, equipement1 Equipement) { //initialisation des valeur  du "monstre"

	var Monster1 Monster
	Monster1.Nom = "Voyou du guetto cammé"
	Monster1.Pvmax = 40
	Monster1.Pv = 40
	Monster1.ptsdamages = 10
	Monster1.initiative = 30
	Monster1.turn = 0
	Monster1.attackturn = -25
	Monster1.hasattacked = false
	Monster1.turnbuff = 0

	fmt.Println(Monster1)
	confirm(p1, equipement1, Monster1)
}

func confirm(p1 Personnage, equipement1 Equipement, Monster1 Monster) { //confirmation de lancement du combat
	var confirm string
	fmt.Println("'⚔' Voulez vous vraiment lancer un combat d'entraînement'⚔'  ?")
	fmt.Scanln(&confirm)
	switch confirm {
	case "oui":
		trainingfight(Monster1, p1, equipement1) //lancement du combat
	case "non":
		DisplayMenu(p1, equipement1, Monster1) //retour au menu
	default:
		fmt.Println("========================================")
		fmt.Println("❌Vous n'avez pas choisi une option valide❌")
		fmt.Println("========================================")
		DisplayMenu(p1, equipement1, Monster1)
	}
}
func goblinpattern(Monster1 Monster, p1 Personnage, equipement1 Equipement) {
	if p1.Pv > 0 && Monster1.Pv > 0 { //si le joueur et le monstre sont en vie
		fmt.Println(p1.hasattacked, Monster1.hasattacked)
		if p1.hasattacked == true && Monster1.hasattacked == true { //si les deux ont attaqué
			Monster1.turn = Monster1.turn + 1 //le tour est terminé on ajoute alors 1 au tour du combat
			p1.hasattacked = false            //on remet les booléens à false pour le prochain tour
			Monster1.hasattacked = false
			Monster1.turnbuff = Monster1.turnbuff + 1 //on ajoute 1  à la variable qui permet  au  monstre de lancer son attaque spéciale  tous les 3  tours
			fmt.Println("===========================")
			fmt.Println("===========================")
			fmt.Println("Tour n°", Monster1.turn, "|")
			fmt.Println("===========================")
			fmt.Println("===========================")
		}

		if Monster1.attackturn == 0 { //si ce n'est pas au  monstre  d'attaquer
		}
		if Monster1.attackturn == 1 { //si c'est au monstre d'attaquer
			fmt.Println(Monster1.turn)
			if Monster1.turnbuff == 3 { //si le monstre peut lancer son attaque spéciale
				Monster1.ptsdamages = Monster1.ptsdamages * 2 //on double les dégats du monstre
				p1.Pv = p1.Pv - Monster1.ptsdamages           //on enlève les dégats au joueur
				Monster1.ptsdamages = Monster1.ptsdamages / 2 //on remet les dégats du monstre à leur valeur d'origine
				Monster1.turnbuff = 0                         //on remet la variable à 0
				fmt.Println("====================================================================================================================")
				fmt.Println("Votre ennemis:", Monster1.Nom, "vous inflige", Monster1.ptsdamages*2, "💥il vous reste", p1.Pv, "/", p1.Pvmax, "❤️")
				fmt.Println("====================================================================================================================")
				fmt.Println(p1.hasattacked, Monster1.hasattacked)

			} else {
				Monster1.ptsdamages = Monster1.ptsdamages * 1 //s les dégats du monstre restent à 100%
				fmt.Println(Monster1.ptsdamages)
				p1.Pv = p1.Pv - Monster1.ptsdamages //on enlève les dégats au joueur
				fmt.Println("====================================================================================================================")
				fmt.Println("Votre ennemis:", Monster1.Nom, "vous inflige", Monster1.ptsdamages, "💥il vous reste", p1.Pv, "/", p1.Pvmax, "❤️")
				fmt.Println("====================================================================================================================")
				fmt.Println(p1.hasattacked, Monster1.hasattacked)
			}
		}

		Monster1.hasattacked = true // le monster a terminé son tour

		fmt.Println(p1.hasattacked, Monster1.hasattacked)
		charattack(p1, equipement1, Monster1) //on lance la fonction qui permet au joueur d'attaquer

	} else { //si le joueur ou le monstre est mort
		if p1.Pv <= 0 { //si le joueur est mort
			fmt.Println("================================================")
			fmt.Println("Vous etes mort 💀 face à un(e)", Monster1.Nom)
			fmt.Println("================================================")
			iswasted(p1, equipement1, Monster1) //on lance la fonction qui permet de savoir si  le joueur est mort et s'ik veut recommencer ou non
		}
		if Monster1.Pv <= 0 { //si le monstre est mort
			fmt.Println("===========================================================================")
			fmt.Println("Vous avez gagné le combat contre un(e)", Monster1.Nom, "🏆 félicitation !")
			fmt.Println("===========================================================================")
			p1.xpaquired = p1.xpaquired + 250 //on ajoute 250 xp au joueur
			DisplayMenu(p1, equipement1, Monster1)
		}
	}
}
func charattack(p1 Personnage, equipement1 Equipement, Monster1 Monster) {
	var attackchoice int
	if p1.Pv > 0 && Monster1.Pv > 0 { //si le joueur et le monstre sont en vie
		if p1.hasattacked == true && Monster1.hasattacked == true { //meme condition que dans la fonction goblinpattern
			Monster1.turn = Monster1.turn + 1
			Monster1.hasattacked = false
			p1.hasattacked = false
			Monster1.turnbuff = Monster1.turnbuff + 1
			fmt.Println("===========================")
			fmt.Println("===========================")
			fmt.Println("TOUR n°", Monster1.turn, " |") //affichage du tour
			fmt.Println("===========================")
			fmt.Println("===========================")
		}
		fmt.Println("🗡️========================================================🗡️")
		fmt.Println("Choissisez vos actions ") //on demande au joueur de choisir son action
		fmt.Println("🗡️========================================================🗡️")
		fmt.Println("1-⚔ Attaquer ⚔")
		fmt.Println("2- Inventaire ")
		fmt.Scanln(&attackchoice)
		switch attackchoice {
		case 1:
			if p1.skill[0] == "Coup de poing" { //si le joueur a le coup de poing  dans ses skills
				Monster1.Pv = Monster1.Pv - p1.ptsdamages //on enlève les dégats au monstre
				if Monster1.Pv <= 0 {                     //si le monstre est mort
					Monster1.Pv = 0 //on met les pv du monstre à 0
				}
				fmt.Println("====================================================================================================================================================================")
				fmt.Println("Vous avez choisi de frapper avec vos poings 👊 occasionant ", p1.ptsdamages, "💥 de dégats à", Monster1.Nom, "il lui reste", Monster1.Pv, "/", Monster1.Pvmax, "❤️")
				fmt.Println("====================================================================================================================================================================")
				p1.coins = p1.coins + 10         //on ajoute 10 pièces au joueur
				p1.xpaquired = p1.xpaquired + 50 //on ajoute 50 xp au joueur
				fmt.Println("========================================================================")
				fmt.Println("Gain de 50 xp, progression: ", p1.xpaquired, "/", p1.xpneeded)
				fmt.Println("Gain de 10 pièces d'or: vous avaez desormais", p1.coins, "pièces d'or ")
				fmt.Println("========================================================================")
				Monster1.attackturn = 1 //on passe au tour du monstre
				p1.hasattacked = true   //le joueur a terminé son tour

				goblinpattern(Monster1, p1, equipement1) //on relance la fonction goblinpattern

			}
		case 2:
			AcessInventoryinCombat(p1, equipement1, Monster1)
		default:
			fmt.Println("Vous n'avez pas choisi une option valide")
			charattack(p1, equipement1, Monster1)

		}

	} else {
		if p1.Pv <= 0 {
			fmt.Println("================================================")
			fmt.Println("Vous etes mort 💀 face à un(e)", Monster1.Nom)
			fmt.Println("================================================")
			iswasted(p1, equipement1, Monster1)

		}
		if Monster1.Pv <= 0 {
			fmt.Println(Monster1.Pv)
			fmt.Println("Vous avez gagné le combat")
			DisplayMenu(p1, equipement1, Monster1)
		}
	}
}
func trainingfight(Monster1 Monster, p1 Personnage, equipement1 Equipement) {
	var turn int = 0
	fmt.Println("Vous avez choisi de combattre un", Monster1.Nom)
	fmt.Println(Monster1.Pv)
	fmt.Println(Monster1.Pvmax)
	fmt.Println(p1.Pv)
	if p1.Pv > 0 && Monster1.Pv > 0 { //si le joueur et le monstre sont en vie
		fmt.Print("tour", turn)
		if p1.initiative > Monster1.initiative { //si l'initiative du joueur est supérieure à celle du monstre
			Monster1.attackturn = 0 //ce n'est pas au tour du monstre
			turn = 1                // le premier tour commence
			Monster1.turn = 1       // le premier tour commence
			Monster1.turnbuff = 1   // +1 au compteur de tour de l'attaque spéciale du monstre
			fmt.Println("Vous avez plus d'initiative (", p1.initiative, "pts) que votre adversaire: vous  commencez le combat")
			charattack(p1, equipement1, Monster1)
		}
		if p1.initiative < Monster1.initiative {
			fmt.Println("Vous avez moins d'initiative(", p1.initiative, "pts) que votre adversaire: il commence le combat") //si l'initiative du joueur est inférieure à celle du monstre
			turn = 1
			Monster1.turn = 1
			Monster1.attackturn = 1
			Monster1.turnbuff = 1
			goblinpattern(Monster1, p1, equipement1)
		}

	} else {
		if p1.Pv <= 0 {
			fmt.Println("================================================")
			fmt.Println("Vous etes mort 💀 face à un(e)", Monster1.Nom)
			fmt.Println("================================================")
			iswasted(p1, equipement1, Monster1)
		}
		if Monster1.Pv <= 0 {
			fmt.Println("Vous avez gagné le combat")
			DisplayMenu(p1, equipement1, Monster1)
			//} else {
			//fmt.Println("issue indécise: vous avez tous les deux perdu")
			//DisplayMenu(p1, equipement1, Monster1)
		}
	}
}

func AcessInventoryinCombat(p1 Personnage, equipement1 Equipement, Monster1 Monster) { //Affiche l'inventaire du joueur et permet d'intéragir avec  un objet à utiliser tel qu'une potion ou un sort (meme fonction que celle de l'inventaire mais avec quelques différences)
	var j int8
	fmt.Println(p1.deathcount)
	fmt.Println("=========================")
	fmt.Println("Inventaire de", p1.Nom, "|")
	fmt.Println("=========================")
	fmt.Println(p1.Pv)
	fmt.Println("potion(s) de soin:")
	fmt.Println(len(p1.inventorypotion))
	fmt.Println("================================")
	fmt.Println("Sorts:")
	fmt.Println(p1.persoinventairesort)
	fmt.Println("================================")
	fmt.Println("Objets:")
	fmt.Println(p1.inventory, len(p1.inventory), "|")
	fmt.Println("================================")
	fmt.Println("--------------------------------")
	fmt.Println("choisissez parmis ces options")
	fmt.Println("1. Utiliser une potion de soin")
	fmt.Println("2. Utiliser un sort")
	fmt.Println("3. Plain")
	fmt.Println("4. Quitter l'Inventaire")
	fmt.Println("----------------------------")
	fmt.Print("choississez une option: ")
	fmt.Scanln(&j)
	fmt.Println("===================================")
	switch j {
	case 1:
		removehealthpotioninCombat(p1, equipement1, Monster1) //appel de la fonction permettant d'utiliser une potion et de la retirer de l'inventaire
	case 2:
		var choix int8
		fmt.Println("1.Boule de feu")
		fmt.Print("Choissisez un sort")
		fmt.Scanln(&choix)
		switch choix {
		case 1:
			if len(p1.skill) == 1 { //si le joueur a un sort dans son inventaire (qui est par défaut le coup de poing ce veut qu'il n'a pas encore appris le sort de boule feu  etant donné que les skills ne peuvent etre retirés)
				fmt.Println("Vous n'avez pas appris ce sort")
				AcessInventoryinCombat(p1, equipement1, Monster1)
			} else { //si le joueur a appris le sort de boule de feu
				if p1.mana >= 50 { // si le joueur a assez de mana pour lancer le sort
					Monster1.Pv = Monster1.Pv - 18 //le monstre perd 18 pv
					p1.mana = p1.mana - 50         //le joueur perd 10 mana
					fmt.Println("======================================================================================================================================================")
					fmt.Println("Vous avez choisi de lancer une boule de feu 🔥occasionant 18 points de dégats💥 à votre adversaire", Monster1.Nom, "pv pour le prix de 50 points de mana")
					fmt.Println("Il lui reste", Monster1.Pv, "points de vie")
					fmt.Println("Il vous reste", p1.mana, " ⚡points de mana")
					fmt.Println("======================================================================================================================================================")
					p1.coins = p1.coins + 15          //le joueur gagne 15 pièces d'or
					p1.xpaquired = p1.xpaquired + 150 //le joueur gagne 150 xp
					Monster1.attackturn = 1           //c'est au tour du monstre de jouer
					fmt.Println("Gain de 150 xp, Progression:", p1.xpaquired, "/", p1.xpneeded)
					fmt.Println("Gain de 15 pièces d'or: vous avaez desormais", p1.coins, "pièces d'or")
					fmt.Println("C'est au tour de votre adversaire de jouer")
					p1.hasattacked = true //le joueur a  terminé son tour
					goblinpattern(Monster1, p1, equipement1)
				} else { //si le joueur n'a pas assez de mana
					fmt.Println("Vous n'avez pas assez de mana, votre mana est de", p1.mana, "points")
					AcessInventoryinCombat(p1, equipement1, Monster1)
				}
			}
		default:
			fmt.Println("=====================================")
			fmt.Println("Vous n'avez pas choisi un sort valide")
			fmt.Println("=====================================")
			AcessInventoryinCombat(p1, equipement1, Monster1)
		}
	case 3:
		fmt.Println("=====================================================")
		fmt.Println("Vous ne pouvez pas vous équiper d'une armure en combat")
		fmt.Println("=======================================================")
		AcessInventoryinCombat(p1, equipement1, Monster1)
	case 4:
		goblinpattern(Monster1, p1, equipement1) //appel de la fonction permettant de quitter l'inventaire

	default:
		fmt.Println("=============================================================")
		fmt.Println("Vous n'avez pas choisi une option valide : Veuillez réessayer")
		fmt.Println("=============================================================")

	}
}

func removehealthpotioninCombat(p1 Personnage, equipement1 Equipement, Monster1 Monster) { // fonction proche dans le fonctionnement  de  removehealthpotion mais avec quelques différences
	potioncount := len(p1.inventorypotion)
	if len(p1.inventorypotion) == 0 { //si l'inventaire potion est vide
		fmt.Println("Vous n'avez pas de potion")
		AcessInventoryinCombat(p1, equipement1, Monster1)
	} else {
		fmt.Println("????", p1.Pv)
		if p1.Pv == p1.Pvmax { //si le joueur a déjà toute sa vie
			fmt.Println("================================")
			fmt.Println("Votre vie est déja au maximum")
			fmt.Println("potion(s) restante(s)", potioncount)
			fmt.Println("================================")
			AcessInventoryinCombat(p1, equipement1, Monster1)
		}
		i := len(p1.inventorypotion)
		p1.inventorypotion = append(p1.inventorypotion[:0], p1.inventorypotion[:i-1]...)
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
		p1.hasattacked = true
		goblinpattern(Monster1, p1, equipement1)

	}
}
