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
