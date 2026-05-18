package cyphers

import (
	"fmt"
	"log/slog"
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

type Gronsfeld struct {
	key []int
}

type Vigener struct {
	key []int
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

func (a *Atbash) ChangeParams(params []interface{}) {}

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

func (s *Scytale) ChangeParams(params []interface{}) {
	slog.Info("Got params for scytale:", "params", params)
	s.height = int(params[0].(float64))
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

func (p *Polybius) ChangeParams(params []interface{}) {
	slog.Info("Got params for polybius:", "params", params)
	p.chosenLang = int(params[0].(float64))
}

func (c *Caesar) Cypher(input string) string {
	runed := []rune(input)
	cyphered := make([]rune, len(runed))
	for i := range runed {
		upped := strings.ToUpper(string(runed[i]))
		runed[i] = []rune(upped)[0]
		if strings.Contains(string(alph), upped) {
			cyphered[i] = alph[abs(slices.Index(alph, runed[i])+c.step)%len(alph)]
		} else if strings.Contains(string(rusAlph), upped) {
			cyphered[i] = rusAlph[abs(slices.Index(rusAlph, runed[i])+c.step)%len(rusAlph)]
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
			cyphered[i] = alph[abs(slices.Index(alph, runed[i])-c.step)%len(alph)]
		} else if strings.Contains(string(rusAlph), upped) {
			cyphered[i] = rusAlph[abs(slices.Index(rusAlph, runed[i])-c.step)%len(rusAlph)]
		} else {
			cyphered[i] = runed[i]
		}
	}
	return string(cyphered)
}

func (c *Caesar) ChangeParams(params []interface{}) {
	slog.Info("Got params for caesar:", "params", params)
	c.step = int(params[0].(float64))
}

func abs(input int) int {
	for input < 0 {
		input += 26
	}
	return input
}

func (g *Gronsfeld) Cypher(input string) string {
	runed := []rune(input)
	cyphered := make([]rune, len(runed))
	for i := range runed {
		upped := strings.ToUpper(string(runed[i]))
		runed[i] = []rune(upped)[0]
		if strings.Contains(string(alph), upped) {
			cyphered[i] = alph[abs(slices.Index(alph, runed[i])+g.key[i%len(g.key)])%len(alph)]
		} else if strings.Contains(string(rusAlph), upped) {
			cyphered[i] = rusAlph[abs(slices.Index(rusAlph, runed[i])+g.key[i%len(g.key)])%len(rusAlph)]
		} else {
			cyphered[i] = runed[i]
		}
	}
	fmt.Println(cyphered)
	return string(cyphered)
}

func (g *Gronsfeld) Decypher(input string) string {
	runed := []rune(input)
	cyphered := make([]rune, len(runed))
	for i := range runed {
		upped := strings.ToUpper(string(runed[i]))
		runed[i] = []rune(upped)[0]
		if strings.Contains(string(alph), upped) {
			cyphered[i] = alph[abs(slices.Index(alph, runed[i])-g.key[i%len(g.key)])%len(alph)]
		} else if strings.Contains(string(rusAlph), upped) {
			cyphered[i] = rusAlph[abs(slices.Index(rusAlph, runed[i])-g.key[i%len(g.key)])%len(rusAlph)]
		} else {
			cyphered[i] = runed[i]
		}
	}
	return string(cyphered)
}

func (g *Gronsfeld) ChangeParams(params []interface{}) {
	slog.Info("Got params for gronsfeld:", "params", params)
	param := []rune(params[0].(string))
	g.key = make([]int, len(param))
	for i := range param {
		g.key[i], _ = strconv.Atoi(string(param[i]))
	}
}

func (v *Vigener) Cypher(input string) string {
	runed := []rune(input)
	cyphered := make([]rune, len(runed))
	for i := range runed {
		upped := strings.ToUpper(string(runed[i]))
		runed[i] = []rune(upped)[0]
		if strings.Contains(string(alph), upped) {
			cyphered[i] = alph[abs(slices.Index(alph, runed[i])+v.key[i%len(v.key)])%len(alph)]
		} else if strings.Contains(string(rusAlph), upped) {
			cyphered[i] = rusAlph[abs(slices.Index(rusAlph, runed[i])+v.key[i%len(v.key)])%len(rusAlph)]
		} else {
			cyphered[i] = runed[i]
		}
	}
	fmt.Println(cyphered)
	return string(cyphered)

}

func (v *Vigener) Decypher(input string) string {
	runed := []rune(input)
	cyphered := make([]rune, len(runed))
	for i := range runed {
		upped := strings.ToUpper(string(runed[i]))
		runed[i] = []rune(upped)[0]
		if strings.Contains(string(alph), upped) {
			cyphered[i] = alph[abs(slices.Index(alph, runed[i])-v.key[i%len(v.key)])%len(alph)]
		} else if strings.Contains(string(rusAlph), upped) {
			cyphered[i] = rusAlph[abs(slices.Index(rusAlph, runed[i])-v.key[i%len(v.key)])%len(rusAlph)]
		} else {
			cyphered[i] = runed[i]
		}
	}
	return string(cyphered)

}

func (v *Vigener) ChangeParams(params []interface{}) {
	slog.Info("Got params for vigener:", "params", params)
	letters := []rune(params[0].(string))
	v.key = make([]int, len(letters))

	for i := range letters {
		letterIndex := slices.Index(alph, letters[i])
		if letterIndex == -1 {
			letterIndex = slices.Index(rusAlph, letters[i])
		}
		v.key[i] = letterIndex
	}
}
