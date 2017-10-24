package Gofiles

import (
	"fmt"
	/*"math"
	"strings"*/

	"strings"
)


func main() {

	sentence := "James is such a uvavu"
	mission := strings.Replace(sentence, "uvavu", "mamene", 1)
	words := strings.Split(sentence, " ")

	words[3] = words[3][:1]
	sentence = strings.Join(words, " ")
	fmt.Println(mission )
}
