package main

import (
	"fmt"
)

var (
	intChan chan int
	counter int
)

func main() {
	mission := `Ecrivez un programme qui :

	* crée un channel avec tampon mémoire de 5 ints
	
	* écrit 5 valeurs dedans
	
	* ferme le canal

	* passe ce canal à une fonction qui va imprimer les valeurs enfilées dans le canal
	  tant qu'il est possible d'en lire`

	fmt.Println(mission)

	intChan = intChan
	counter = counter
}
