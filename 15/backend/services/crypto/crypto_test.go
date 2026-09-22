package crypto

import (
	"fmt"
	"testing"
)

func TestGamma(t *testing.T) {
	var g Gamma
	g.base = 3313
	g.modulo = 2347

	text := "hello world привет мир"
	cyphered := g.Cypher(text)
	decyphered := g.Decypher(cyphered)

	fmt.Println(cyphered, decyphered)
}
