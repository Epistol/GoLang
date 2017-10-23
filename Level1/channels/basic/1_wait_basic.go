package main

import (
	"fmt"
)

var (
	intChan chan int
	counter int
)

func main() {
	mission := `Pour illustrer la détection de la complétion d'un job sur un canal, 
	écrivez un programme qui :

  * utilise un canal non bufferisé nommé intChan
  * A une fonction SendValue qui doit juste envoyer 5 valeurs sur ce canal (qui lui a été passé en argument)
  * Lance SendValue en tant que goroutine
  * Utilise une boucle for qui boucle à l'infini sur la réception d'une valeur sur le canal
  
  => Faites en sorte que la boucle soit interrompue dès lors que la goroutine a fini d'envoyer ses valeurs

  Allez-y, à votre tour de coder maintenant :)`

	fmt.Println(mission)

	intChan = intChan
	counter = counter
}
