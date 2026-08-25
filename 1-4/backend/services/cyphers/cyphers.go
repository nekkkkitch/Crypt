package cyphers

import (
	"fmt"
	"log/slog"
	"math"
	"math/big"
	"slices"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
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

type Playfair struct {
	table [][]rune
	isRus bool
}

var plAlph = []rune(strings.ToLower(string(alph)))
var plAlphRus = []rune(strings.ToLower(string(rusAlph)))

type Hill struct {
	key   *mat.Dense
	isRus bool
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

func (pl *Playfair) Cypher(input string) string {
	word := []rune(strings.ToLower(input))
	word = []rune(strings.ReplaceAll(string(word), " ", ""))
	for i := range word {
		if word[i] == 'j' {
			word[i] = 'i'
		}
	}
	if !pl.CheckWord(word) {
		return "bad input"
	}
	cyphered := make([]rune, len(word))
	for i := 0; i < len(word); i += 2 {
		var firstI, firstJ, secondI, secondJ int
		for x := range pl.table {
			if y := slices.Index(pl.table[x], word[i]); y != -1 {
				firstI = x
				firstJ = y
			}
			if y := slices.Index(pl.table[x], word[i+1]); y != -1 {
				secondI = x
				secondJ = y
			}
		}
		if firstI == secondI {
			cyphered[i] = pl.table[(firstI)][(firstJ+1)%len(pl.table[firstI])]
			cyphered[i+1] = pl.table[secondI][(secondJ+1)%len(pl.table[secondI])]
		} else if firstJ == secondJ {
			cyphered[i] = pl.table[(firstI+1)%len(pl.table)][firstJ]
			cyphered[i+1] = pl.table[(secondI+1)%len(pl.table)][secondJ]
		} else {
			cyphered[i] = pl.table[firstI][secondJ]
			cyphered[i+1] = pl.table[secondI][firstJ]
		}
	}
	return string(cyphered)
}

func (pl *Playfair) Decypher(input string) string {
	word := []rune(strings.ToLower(input))
	word = []rune(strings.ReplaceAll(string(word), " ", ""))
	for i := range word {
		if word[i] == 'j' {
			word[i] = 'i'
		}
	}
	if !pl.CheckWord(word) {
		return "bad input"
	}
	cyphered := make([]rune, len(word))
	for i := 0; i < len(word); i += 2 {
		var firstI, firstJ, secondI, secondJ int
		for x := range pl.table {
			if y := slices.Index(pl.table[x], word[i]); y != -1 {
				firstI = x
				firstJ = y
			}
			if y := slices.Index(pl.table[x], word[i+1]); y != -1 {
				secondI = x
				secondJ = y
			}
		}
		if firstI == secondI {
			cyphered[i] = pl.table[(firstI)][abs(firstJ-1)%len(pl.table[firstI])]
			cyphered[i+1] = pl.table[secondI][abs(secondJ-1)%len(pl.table[secondI])]
		} else if firstJ == secondJ {
			cyphered[i] = pl.table[abs(firstI-1)%len(pl.table)][firstJ]
			cyphered[i+1] = pl.table[abs(secondI-1)%len(pl.table)][secondJ]
		} else {
			cyphered[i] = pl.table[firstI][secondJ]
			cyphered[i+1] = pl.table[secondI][firstJ]
		}
	}
	return string(cyphered)
}

func (pl *Playfair) ChangeParams(params []interface{}) {
	slog.Info("Got params for Playfair", "params", params)
	word := []rune(strings.ToLower(params[0].(string)))
	consistent := true
	isRus := slices.Contains(plAlphRus, word[0])
	if isRus {
		for i := range word {
			if !slices.Contains(plAlphRus, word[i]) {
				consistent = false
				break
			}
		}
	} else {
		for i := range word {
			if !slices.Contains(plAlph, word[i]) {
				slog.Info("ebebeb", "word[i]", word[i])
				consistent = false
				break
			}
		}
	}
	if !consistent {
		slog.Error("Inconsistent key for playfair:", "key", params[0], "isRus", isRus)
		return
	}
	pl.isRus = isRus
	usedChars := make(map[rune]interface{})
	if pl.isRus {
		sq := [6][6]rune{}
		indexI := 0
		indexJ := 0
		for i := range word {
			if _, ok := usedChars[word[i]]; !ok {
				sq[indexI][indexJ] = word[i]
				usedChars[word[i]] = struct{}{}
				indexJ++
				if indexJ >= 6 {
					indexI++
					indexJ = 0
				}
			}
		}
		rusAlph := plAlphRus
		rusAlph = append(plAlphRus, '.', ',', '-')
		for i := range rusAlph {
			if _, ok := usedChars[rusAlph[i]]; !ok {
				sq[indexI][indexJ] = rusAlph[i]
				usedChars[rusAlph[i]] = struct{}{}
				indexJ++
				if indexJ >= 6 {
					indexI++
					indexJ = 0
				}
			}
		}
		pl.table = make([][]rune, 6)
		for i := range 6 {
			pl.table[i] = make([]rune, 6)
			for j := range 6 {
				pl.table[i][j] = sq[i][j]
			}
		}
	} else {
		sq := [5][5]rune{}
		indexI := 0
		indexJ := 0
		for i := range word {
			if _, ok := usedChars[word[i]]; !ok {
				sq[indexI][indexJ] = word[i]
				usedChars[word[i]] = struct{}{}
				indexJ++
				if indexJ >= 5 {
					indexI++
					indexJ = 0
				}
			}
		}
		for i := range plAlph {
			if plAlph[i] == 'j' {
				continue
			}
			if _, ok := usedChars[plAlph[i]]; !ok {
				sq[indexI][indexJ] = plAlph[i]
				usedChars[plAlph[i]] = struct{}{}
				indexJ++
				if indexJ >= 5 {
					indexI++
					indexJ = 0
				}
			}
		}
		pl.table = make([][]rune, 5)
		for i := range 5 {
			pl.table[i] = make([]rune, 5)
			for j := range 5 {
				pl.table[i][j] = sq[i][j]
			}
		}
	}
	slog.Info("Result", "table", pl.table)
}

func (pl *Playfair) CheckWord(word []rune) bool {
	if len(word)%2 != 0 {
		return false
	}
	if pl.isRus {
		for i := range word {
			if !slices.Contains(append(plAlphRus, '.', ',', '-'), word[i]) {
				return false
			}
		}
	} else {
		for i := range word {
			if !slices.Contains(plAlph, word[i]) {
				return false
			}
		}
	}
	return true
}

func (h *Hill) Cypher(input string) string {
	cons, isRus := checkConsistency(input)
	if !cons || isRus != h.isRus {
		slog.Error("word not consistent or its language is not the same as key")
		return ""
	}

	runed := []rune(strings.ToUpper(input))
	if len(runed) != h.key.RawMatrix().Cols {
		slog.Error("word length should be the same as square root of key length")
		return ""
	}

	inputNums := make([]float64, len(runed))

	if !isRus {
		for i := range inputNums {
			inputNums[i] = float64(slices.Index(alph, runed[i]))
		}
	} else {
		for i := range inputNums {
			inputNums[i] = float64(slices.Index(rusAlph, runed[i]))
		}
	}

	inputMatrix := mat.NewDense(len(runed), 1, inputNums)
	resMatrix := mat.NewDense(len(runed), 1, nil)
	resMatrix.Mul(h.key, inputMatrix)
	data := resMatrix.RawMatrix().Data
	for i := range data {
		if !isRus {
			data[i] = float64(divMod(int(data[i]), len(alph)))
		} else {
			data[i] = float64(divMod(int(data[i]), len(rusAlph)))
		}
	}
	slog.Info("Hill cypher", "key", h.key.RawMatrix().Data, "input", inputMatrix.RawMatrix().Data)

	resultWord := make([]rune, len(runed))
	for i := range resultWord {
		if !isRus {
			resultWord[i] = alph[int(data[i])]
		} else {
			resultWord[i] = rusAlph[int(data[i])]
		}
	}

	return string(resultWord)
}

func (h *Hill) Decypher(input string) string {
	cons, isRus := checkConsistency(input)
	if !cons || isRus != h.isRus {
		slog.Error("word not consistent or its language is not the same as key")
		return ""
	}

	runed := []rune(input)
	if len(runed) != h.key.RawMatrix().Cols {
		slog.Error("word length should be the same as square root of key length")
		return ""
	}

	inputNums := make([]float64, len(runed))

	if !isRus {
		for i := range inputNums {
			inputNums[i] = float64(slices.Index(alph, runed[i]))
		}
	} else {
		for i := range inputNums {
			inputNums[i] = float64(slices.Index(rusAlph, runed[i]))
		}
	}

	inputMatrix := mat.NewDense(len(runed), 1, inputNums)
	resMatrix := mat.NewDense(len(runed), 1, nil)
	inverse := mat.NewDense(len(runed), len(runed), nil)
	inverse.Inverse(h.key)

	a := big.NewInt(int64(mat.Det(h.key)))
	var m *big.Int

	if !isRus {
		m = big.NewInt(int64(len(alph)))
	} else {
		m = big.NewInt(int64(len(rusAlph)))
	}

	k := new(big.Int).ModInverse(a, m)
	scale := k.Int64()

	inverse.Scale(float64(scale), inverse)
	inverse.Scale(mat.Det(h.key), inverse)
	slog.Info("Hill", "inverse key", inverse.RawMatrix().Data)
	inverseData := inverse.RawMatrix().Data
	for i := range inverseData {
		if !isRus {
			inverseData[i] = float64(divMod(int(math.Round(inverseData[i])), len(alph)))
		} else {
			inverseData[i] = float64(divMod(int(math.Round(inverseData[i])), len(rusAlph)))
		}
	}
	slog.Info("Hill", "inverse key", inverseData)
	resMatrix.Mul(inverse, inputMatrix)

	data := resMatrix.RawMatrix().Data
	for i := range data {
		if !isRus {
			data[i] = float64(divMod(int(data[i]), len(alph)))
		} else {
			data[i] = float64(divMod(int(data[i]), len(rusAlph)))
		}
	}

	resultWord := make([]rune, len(runed))
	for i := range resultWord {
		if !isRus {
			resultWord[i] = alph[int(data[i])]
		} else {
			resultWord[i] = rusAlph[int(data[i])]
		}
	}

	return string(resultWord)
}
func (h *Hill) ChangeParams(params []interface{}) {
	keyWord := []rune(strings.ToUpper(params[0].(string)))
	if math.Sqrt(float64(len(keyWord))) != float64(int(math.Sqrt(float64(len(keyWord))))) {
		slog.Error("Len of key is not square of int", "len", len(keyWord))
		return
	}

	cons, isRus := checkConsistency(string(keyWord))
	if !cons {
		slog.Error("Not consistent key(use only rus or eng symbols)")
		return
	}

	h.isRus = isRus

	tableSide := int(math.Sqrt(float64(len(keyWord))))
	key := make([]float64, tableSide*tableSide)

	if !isRus {
		for i := range key {
			key[i] = float64(slices.Index(alph, keyWord[i]))
		}
		matrix := *mat.NewDense(tableSide, tableSide, key)
		det := int(mat.Det(&matrix))
		if det == 0 {
			slog.Error("Det == 0 or has similar dividers")
			return
		}
		a := big.NewInt(int64(mat.Det(&matrix)))
		m := big.NewInt(int64(len(alph)))

		k := new(big.Int).ModInverse(a, m)
		if k == nil {
			slog.Error("Bad determinant")
			return
		}
		h.key = &matrix
	} else {
		for i := range key {
			key[i] = float64(slices.Index(rusAlph, keyWord[i]))
		}
		matrix := *mat.NewDense(tableSide, tableSide, key)
		det := int(mat.Det(&matrix))
		if det == 0 {
			slog.Error("Det == 0 or has similar dividers")
			return
		}
		a := big.NewInt(int64(mat.Det(&matrix)))
		m := big.NewInt(int64(len(rusAlph)))

		k := new(big.Int).ModInverse(a, m)
		if k == nil {
			slog.Error("Bad determinant")
			return
		}
		h.key = &matrix
	}
}

func checkConsistency(input string) (bool, isRus bool) {
	isConsistent := true
	word := []rune(strings.ToUpper(input))
	isRus = slices.Contains(rusAlph, word[0])
	for i := range word {
		if isRus != slices.Contains(rusAlph, word[i]) {
			isConsistent = false
		}
	}
	return isConsistent, isRus
}

func divMod(num, step int) int {
	if num >= 0 {
		return num % step
	}
	for num < 0 {
		num += step
	}
	return num
}
