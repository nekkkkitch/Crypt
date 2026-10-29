package cryptanalysis

import (
	"fmt"
	"log/slog"
	"maps"
	"os"
	"regexp"
	"slices"
	"strings"
)

type Analysator struct {
	Freq        Frequency
	Index       Index
	AC          AutoCorrelation
	Kasiski     Kasiski
	IsRus       bool
	currentText string
}

type Data struct {
	F     FreqData    `json:"freq_data"`
	A     ACData      `json:"ac_data"`
	K     KasiskiData `json:"k_data"`
	Index float64     `json:"index_data"`
}

type Frequency struct {
	currentData FreqData
}

type FreqData struct {
	Frequencies map[string]Place `json:"frequencies"`
	Sorted      []SymbolFreq     `json:"sorted"`
	InStupid    string           `json:"in_stupid"`
}

type Place struct {
	Part   float64 `json:"part"`
	Number int     `json:"number"`
}

type SymbolFreq struct {
	Symbol string
	Freq   float64
	Place  int
}

type Index struct {
}

type AutoCorrelation struct {
}

type ACData struct {
	Correlations []Correlation `json:"correlations"`
}

type Correlation struct {
	Amount      int     `json:"amount"`
	Coefficient float64 `json:"coefficient"`
	T           int     `json:"t"`
}

type Kasiski struct {
}

type KasiskiData struct {
	PossibleDividers map[int]struct{} `json:"possible_dividers"`
}

var engFreq = []rune("ETAOINSHRDLCUMWFGYPBVKJXQZ")
var rusFreq = []rune("ОЕАИНТСЛВРКМДПЫУБЯЬГЗЧЙЖХШЮЦЭЩФЁЪ")

func (a *Analysator) Analyse() Data {
	var d Data
	d.A = a.AC.Analyse(a.currentText)
	d.F = a.Freq.Analyse(a.currentText, a.IsRus)
	d.K = a.Kasiski.Analyse(a.currentText)
	d.Index = Analyse(d.F)
	return d
}

func (a *Analysator) SetText(text string, isRus bool) {
	a.currentText = text
	a.IsRus = isRus
	slog.Info("Got text from file", "text part", text)
}

func (a *Analysator) SetFile(filePath string, isRus bool) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return
	}
	text := string(file)
	a.IsRus = isRus
	a.currentText = text
	slog.Info("Got text from file", "text part", string([]rune(text)[:50]))
}

func (f *Frequency) Analyse(text string, isRus bool) FreqData {
	f.currentData = FreqData{}
	var data FreqData

	if isRus {
		data.Frequencies = make(map[string]Place, len(rusFreq))
	} else {
		data.Frequencies = make(map[string]Place, len(engFreq))
	}

	text = Beautify(text)
	runedText := []rune(text)
	slog.Info("Text", "text", text)
	for i := range runedText {
		if runedText[i] == ' ' {
			continue
		}
		place := data.Frequencies[string(runedText[i])]
		place.Part += 1
		data.Frequencies[string(runedText[i])] = place
	}
	slog.Info("Before sorted", "data", data, "len", len([]rune(text)))

	data.Sorted = make([]SymbolFreq, 0, len(data.Frequencies))

	for k, v := range data.Frequencies {
		v.Part = v.Part / float64(len([]rune(text)))
		data.Frequencies[k] = v
		data.Sorted = append(data.Sorted, SymbolFreq{k, data.Frequencies[k].Part, data.Frequencies[k].Number})
	}

	slices.SortFunc(data.Sorted, func(a, b SymbolFreq) int {
		if a.Freq >= b.Freq {
			return -1
		} else {
			return 1
		}
	})

	inStupid := make([]rune, 0, len(data.Sorted))

	for i := range data.Sorted {
		data.Sorted[i].Place = i
		data.Frequencies[data.Sorted[i].Symbol] = Place{data.Frequencies[data.Sorted[i].Symbol].Part, i}
		inStupid = append(inStupid, []rune(data.Sorted[i].Symbol)[0])
	}
	data.InStupid = string(inStupid)
	f.currentData = data

	return data
}

func (f *Frequency) Decrypt(text string) string {
	text = Beautify(text)
	runed := []rune(text)

	for i := range runed {
		if runed[i] == ' ' {
			continue
		}
		runed[i] = rusFreq[f.currentData.Frequencies[string(runed[i])].Number]
	}

	return string(runed)
}

func Analyse(data FreqData) float64 {
	ind := 0.0
	for i := range data.Frequencies {
		ind += data.Frequencies[i].Part * data.Frequencies[i].Part
	}
	return ind
}

func (a *AutoCorrelation) Analyse(text string) ACData {
	text = Beautify(text)
	runed := []rune(text)
	var data ACData
	data.Correlations = make([]Correlation, len(runed))

	for i := 1; i < len(runed); i++ {
		var cor Correlation
		for j := 0; j < len(runed)-i; j++ {
			if runed[j] == runed[i+j] {
				cor.Amount++
			}
		}
		data.Correlations[i-1] = cor
	}

	for i, v := range data.Correlations {
		if len(runed) == i+1 {
			break
		}
		data.Correlations[i].Coefficient = float64(v.Amount) / float64(len(runed)-(i+1))
		data.Correlations[i].T = i + 1
	}

	slog.Info("AutoCorrelation", "data", data.Correlations)

	return data
}

func (k *Kasiski) Analyse(text string) KasiskiData {
	text = Beautify(text)
	runed := []rune(text)

	gs := []int{}
	for i := len(runed) / 2; i >= 3; i-- {
		substrings := make(map[string]int)
		for j := 0; j < len(runed)-i; j++ {
			substring := string(runed[j : j+i])
			if ind, ok := substrings[substring]; !ok {
				substrings[substring] = j
			} else {
				gs = append(gs, j-ind)
				substrings[substring] = j
			}
		}
	}

	var data KasiskiData
	possibleDivs := make(map[int]int)
	for _, v := range gs {
		dividers := getDividers(v)
		for j := range dividers {
			if dividers[j] != 1 {
				possibleDivs[dividers[j]] += 1

			}
		}
	}

	keys := slices.Collect(maps.Keys(possibleDivs))
	slices.SortFunc(keys, func(key1, key2 int) int {
		switch {
		case possibleDivs[key1] < possibleDivs[key2]:
			return 1
		case possibleDivs[key1] == possibleDivs[key2]:
			if key1 > key2 {
				return 1
			}
			return -1
		default:
			return -1
		}
	})

	fmt.Println(keys)
	data.PossibleDividers = make(map[int]struct{})
	for i, v := range keys {
		if i <= 4 {
			data.PossibleDividers[v] = struct{}{}
		} else {
			break
		}

	}

	return data
}

var ForTS = regexp.MustCompile(`[^\p{Latin}\p{Cyrillic};.:,']`)

func Beautify(text string) string {
	text = strings.ToUpper(text)
	return ForTS.ReplaceAllString(text, "")
}

func getDividers(num int) []int {
	dividers := []int{}
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			dividers = append(dividers, i)
		}
	}
	return dividers
}
