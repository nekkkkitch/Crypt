package cyphers

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

var alph = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var rusAlph = []rune("АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ")

type Atbash struct{}

type Scytale struct {
	height int
}

type Polybius struct {
	chosenLang int
}

var engSq = [][]rune{
	{'a', 'b', 'c', 'd', 'e'},
	{'f', 'g', 'h', 'i', 'k'},
	{'l', 'm', 'n', 'o', 'p'},
	{'q', 'r', 's', 't', 'u'},
	{'v', 'w', 'x', 'y', 'z'},
}

var rusSq = [][]rune{
	{'а', 'б', 'в', 'г', 'д', 'е'},
	{'ё', 'ж', 'з', 'и', 'й', 'к'},
	{'л', 'м', 'н', 'о', 'п', 'р'},
	{'с', 'т', 'у', 'ф', 'х', 'ц'},
	{'ч', 'ш', 'щ', 'ъ', 'ы', 'ь'},
	{'э', 'ю', 'я', '.', ',', '-'},
}

var bro = map[rune]rune{
	'j': 'i',
}

type Caesar struct {
	step int
}

func (a *Atbash) Cypher(input string) string {
	runed := []rune(input)
	cyphered := make([]rune, len(runed))
	for i := range runed {
		if strings.Contains(string(alph), strings.ToUpper(string(runed[i]))) {
			cyphered[i] = alph[len(alph)-1-slices.Index(alph, []rune(strings.ToUpper(string(runed[i])))[0])]
		} else if strings.Contains(string(rusAlph), strings.ToUpper(string(runed[i]))) {
			cyphered[i] = rusAlph[len(rusAlph)-1-slices.Index(rusAlph, []rune(strings.ToUpper(string(runed[i])))[0])]
		} else {
			cyphered[i] = runed[i]
		}
	}
	return string(cyphered)
}

func (a *Atbash) Decypher(input string) string {
	return a.Cypher(input)
}

func (a *Atbash) ChangeParams(params int) {}

func (s *Scytale) Cypher(input string) string {
	runed := []rune(input)
	width := int(math.Ceil(float64(len(runed)) / float64(s.height)))
	index := 0
	matrix := make([][]rune, s.height)
	for i := range s.height {
		line := make([]rune, 0, width)
		for range width {
			if index >= len(runed) {
				line = append(line, '*')
				continue
			}
			line = append(line, runed[index])
			index++
		}
		matrix[i] = line
	}
	fmt.Println("Scytale table")
	for i := range matrix {
		fmt.Println(string(matrix[i]))
	}
	result := make([]rune, 0, len(runed))
	indexX := 0
	indexY := 0
	for range len(matrix) * len(matrix[0]) {
		result = append(result, matrix[indexY][indexX])
		indexY++
		if indexY >= s.height || indexX >= len(matrix[indexY]) {
			indexY = 0
			indexX++
		}
	}
	fmt.Println("Scytale ciphered")
	fmt.Println(string(result))
	return string(result)
}

func (s *Scytale) Decypher(input string) string {
	runed := []rune(input)
	width := int(math.Ceil(float64(len(runed)) / float64(s.height)))
	save := s.height
	s.height = width
	res := s.Cypher(input)
	runedRes := []rune(res)
	for i := len(runedRes) - 1; i >= 0; i-- {
		if runedRes[i] != '*' {
			runedRes = runedRes[:i+1]
			break
		}
	}
	res = string(runedRes)
	s.height = save
	return res
}

func (s *Scytale) ChangeParams(params int) {
	s.height = params
}

func (p *Polybius) Cypher(input string) string {
	input = strings.ToLower(input)
	runed := []rune(input)
	for i := range runed {
		if _, ok := bro[runed[i]]; ok {
			runed[i] = bro[runed[i]]
		}
	}
	var sq [][]rune
	if p.chosenLang == 0 {
		sq = engSq
	} else {
		sq = rusSq
	}

	horIndex := make([]rune, 0, len(runed))
	verIndex := make([]rune, 0, len(runed))

	for i := range runed {
		for ver := range sq {
			if hor := slices.Index(sq[ver], runed[i]); hor != -1 {
				horIndex = append(horIndex, []rune(strconv.Itoa(hor))[0])
				verIndex = append(verIndex, []rune(strconv.Itoa(ver))[0])
				break
			}
		}
	}

	endIndex := append(horIndex, verIndex...)
	newRuned := make([]rune, 0, len(runed)*2)

	for i := 0; i < len(endIndex); i += 2 {
		hor, _ := strconv.Atoi(string(endIndex[i]))
		ver, _ := strconv.Atoi(string(endIndex[i+1]))
		newRuned = append(newRuned, sq[ver][hor])
	}

	return string(newRuned)
}

func (p *Polybius) Decypher(input string) string {
	input = strings.ToLower(input)
	runed := []rune(input)
	for i := range runed {
		if _, ok := bro[runed[i]]; ok {
			runed[i] = bro[runed[i]]
		}
	}

	endIndex := make([]rune, 0, len(runed)*2)

	var sq [][]rune
	if p.chosenLang == 0 {
		sq = engSq
	} else {
		sq = rusSq
	}

	for i := range runed {
		for ver := range sq {
			if hor := slices.Index(sq[ver], runed[i]); hor != -1 {
				endIndex = append(endIndex, []rune(strconv.Itoa(hor))[0])
				endIndex = append(endIndex, []rune(strconv.Itoa(ver))[0])
				break
			}
		}
	}

	horIndex := endIndex[:len(endIndex)/2]
	verIndex := endIndex[len(endIndex)/2:]

	newRuned := make([]rune, 0, len(runed))
	for i := range horIndex {
		hor, _ := strconv.Atoi(string(horIndex[i]))
		ver, _ := strconv.Atoi(string(verIndex[i]))
		newRuned = append(newRuned, sq[ver][hor])
	}

	return string(newRuned)
}

func (p *Polybius) ChangeParams(params int) {
	p.chosenLang = params
}

func (c *Caesar) Cypher(input string) string {
	runed := []rune(input)
	cyphered := make([]rune, len(runed))
	for i := range runed {
		upped := strings.ToUpper(string(runed[i]))
		runed[i] = []rune(upped)[0]
		if strings.Contains(string(alph), upped) {
			cyphered[i] = alph[(slices.Index(alph, runed[i])+c.step)%len(alph)]
		} else if strings.Contains(string(rusAlph), upped) {
			cyphered[i] = rusAlph[(slices.Index(rusAlph, runed[i])+c.step)%len(rusAlph)]
		} else {
			cyphered[i] = runed[i]
		}
	}
	fmt.Println(cyphered)
	return string(cyphered)
}

func (c *Caesar) Decypher(input string) string {
	runed := []rune(input)
	cyphered := make([]rune, len(runed))
	for i := range runed {
		upped := strings.ToUpper(string(runed[i]))
		runed[i] = []rune(upped)[0]
		if strings.Contains(string(alph), upped) {
			cyphered[i] = alph[abs(slices.Index(alph, runed[i])-c.step)]
		} else if strings.Contains(string(rusAlph), upped) {
			cyphered[i] = rusAlph[abs(slices.Index(rusAlph, runed[i])-c.step)]
		} else {
			cyphered[i] = runed[i]
		}
	}
	return string(cyphered)
}

func (c *Caesar) ChangeParams(params int) {
	c.step = params
}

func abs(input int) int {
	for input < 0 {
		input += 26
	}
	return input
}
