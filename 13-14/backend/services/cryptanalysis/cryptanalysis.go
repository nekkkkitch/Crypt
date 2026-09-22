package crypanalysis

import (
	"log/slog"
	"regexp"
	"slices"
	"strings"
)

type Analysator struct {
	Freq    Frequency
	Index   Index
	AC      AutoCorrelation
	Kasiski Kasiski
	IsRus   bool
}

type Frequency struct {
	currentData FreqData
}

type FreqData struct {
	Frequencies map[string]Place
	Sorted      []SymbolFreq
	InStupid    string
}

type Place struct {
	part   float64
	number int
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
	Correlations []Correlation
}

type Correlation struct {
	Amount      int
	Coefficient float64
	T           int
}

type Kasiski struct {
}

type KasiskiData struct {
	PossibleDividers map[int]struct{}
}

var engFreq = []rune("ETAOINSHRDLCUMWFGYPBVKJXQZ")
var rusFreq = []rune("ОЕАИНТСЛВРКМДПЫУБЯЬГЗЧЙЖХШЮЦЭЩФЁЪ")

// TODO доделать приём файлов, хз пока в каком виде их бэк получает
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
		place.part += 1
		data.Frequencies[string(runedText[i])] = place
	}
	slog.Info("Before sorted", "data", data, "len", len([]rune(text)))

	data.Sorted = make([]SymbolFreq, 0, len(data.Frequencies))

	for k, v := range data.Frequencies {
		v.part = v.part / float64(len([]rune(text)))
		data.Frequencies[k] = v
		data.Sorted = append(data.Sorted, SymbolFreq{k, data.Frequencies[k].part, data.Frequencies[k].number})
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
		data.Frequencies[data.Sorted[i].Symbol] = Place{data.Frequencies[data.Sorted[i].Symbol].part, i}
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
		runed[i] = rusFreq[f.currentData.Frequencies[string(runed[i])].number]
	}

	return string(runed)
}

func (i *Index) Analyse(data FreqData) float64 {
	ind := 0.0
	for i := range data.Frequencies {
		ind += data.Frequencies[i].part * data.Frequencies[i].part
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
		data.Correlations[i].Coefficient = float64(v.Amount) / float64(len(runed)-(i+1))
		data.Correlations[i].T = i + 1
	}

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
			}
		}
	}

	var data KasiskiData
	data.PossibleDividers = make(map[int]struct{})
	for _, v := range gs {
		dividers := getDividers(v)
		for j := range dividers {
			data.PossibleDividers[dividers[j]] = struct{}{}
		}
	}

	return data
}

var ForTS = regexp.MustCompile(`[^\p{Latin}\p{Cyrillic}]`)

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
