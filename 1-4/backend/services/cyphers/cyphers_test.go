package cyphers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAtbash(t *testing.T) {
	a := Atbash{}

	t.Log("English")
	inputDecyphered := "HELLO WORLD"
	expectedCyphered := "SVOOL DLIOW"
	realCyphered := a.Cypher(inputDecyphered)
	realDecyphered := a.Decypher(expectedCyphered)
	require.Equal(t, expectedCyphered, realCyphered)
	require.Equal(t, inputDecyphered, realDecyphered)
	t.Log("Russian")
	inputDecyphered = "ПРИВЕТ МИР"
	expectedCyphered = "ПОЦЭЪМ ТЦО"
	realCyphered = a.Cypher(inputDecyphered)
	realDecyphered = a.Decypher(expectedCyphered)
	require.Equal(t, expectedCyphered, realCyphered)
	require.Equal(t, inputDecyphered, realDecyphered)
}
func TestScytale(t *testing.T) {
	a := Scytale{height: 4}
	t.Log("English version")
	inputDecyphered := "HELLOWORLD"
	expectedCyphered := "HLODEOR*LWL*"
	realCyphered := a.Cypher(inputDecyphered)
	realDecyphered := a.Decypher(expectedCyphered)
	require.Equal(t, expectedCyphered, realCyphered)
	require.Equal(t, inputDecyphered, realDecyphered)
	t.Log("Russian version")
	inputDecyphered = "ПРИВЕТМИР"
	expectedCyphered = "ПВМ*РЕИ*ИТР*"
	realCyphered = a.Cypher(inputDecyphered)
	realDecyphered = a.Decypher(expectedCyphered)
	require.Equal(t, expectedCyphered, realCyphered)
	require.Equal(t, inputDecyphered, realDecyphered)
}
func TestPolybius(t *testing.T) {
	a := Polybius{}

	t.Log("English")
	inputDecyphered := "abc"
	expectedCyphered := "fca"
	realCyphered := a.Cypher(inputDecyphered)
	realDecyphered := a.Decypher(expectedCyphered)
	require.Equal(t, expectedCyphered, realCyphered)
	require.Equal(t, inputDecyphered, realDecyphered)
	t.Log("Russian")
	a.chosenLang = 1
	inputDecyphered = "абв"
	expectedCyphered = "ёва"
	realCyphered = a.Cypher(inputDecyphered)
	realDecyphered = a.Decypher(expectedCyphered)
	require.Equal(t, expectedCyphered, realCyphered)
	require.Equal(t, inputDecyphered, realDecyphered)
}
func TestCaesar(t *testing.T) {
	a := Caesar{step: 5}

	t.Log("English")
	inputDecyphered := "ABC"
	expectedCyphered := "FGH"
	realCyphered := a.Cypher(inputDecyphered)
	realDecyphered := a.Decypher(expectedCyphered)
	require.Equal(t, expectedCyphered, realCyphered)
	require.Equal(t, inputDecyphered, realDecyphered)
	t.Log("Russian")
	inputDecyphered = "АБВ"
	expectedCyphered = "ЕЁЖ"
	realCyphered = a.Cypher(inputDecyphered)
	realDecyphered = a.Decypher(expectedCyphered)
	require.Equal(t, expectedCyphered, realCyphered)
	require.Equal(t, inputDecyphered, realDecyphered)
}

func TestGronsfeld(t *testing.T) {
	g := Gronsfeld{key: []int{1, 2, 3}}

	t.Log("English")
	inputDecyphered := "ABC"
	expectedCyphered := "BDF"
	realCyphered := g.Cypher(inputDecyphered)
	realDecyphered := g.Decypher(expectedCyphered)
	require.Equal(t, expectedCyphered, realCyphered)
	require.Equal(t, inputDecyphered, realDecyphered)
	t.Log("Russian")
	inputDecyphered = "АБВ"
	expectedCyphered = "БГЕ"
	realCyphered = g.Cypher(inputDecyphered)
	realDecyphered = g.Decypher(expectedCyphered)
	require.Equal(t, expectedCyphered, realCyphered)
	require.Equal(t, inputDecyphered, realDecyphered)
}

func TestVigener(t *testing.T) {
	v := Vigener{key: []int{0, 1, 2}} // KEY

	t.Log("English")
	inputDecyphered := "ABC"
	expectedCyphered := "ACE"
	realCyphered := v.Cypher(inputDecyphered)
	realDecyphered := v.Decypher(expectedCyphered)
	require.Equal(t, expectedCyphered, realCyphered)
	require.Equal(t, inputDecyphered, realDecyphered)
	t.Log("Russian")
	inputDecyphered = "АБВ"
	expectedCyphered = "АВД"
	realCyphered = v.Cypher(inputDecyphered)
	realDecyphered = v.Decypher(expectedCyphered)
	require.Equal(t, expectedCyphered, realCyphered)
	require.Equal(t, inputDecyphered, realDecyphered)
}
