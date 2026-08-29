package crypanalysis

import (
	"regexp"
	"slices"
	"strings"
)

type FreqCryptAnalysator struct {
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

var engFreq = []rune("ETAOINSHRDLCUMWFGYPBVKJXQZ")
var rusFreq = []rune("ОЕАИНТСЛВРКМДПЫУБЯЬГЗЧЙЖХШЮЦЭЩФЁЪ")

// TODO доделать приём файлов, хз пока в каком виде их бэк получает
func (f *FreqCryptAnalysator) Analyse(text string) FreqData {
	f.currentData = FreqData{}
	var data FreqData
	data.Frequencies = make(map[string]Place, len(rusFreq))

	text = Beautify(text)
	runedText := []rune(text)

	for i := range runedText {
		if runedText[i] == ' ' {
			continue
		}
		place := data.Frequencies[string(runedText[i])]
		place.part += 1
		data.Frequencies[string(runedText[i])] = place
	}

	data.Sorted = make([]SymbolFreq, 0, len(data.Frequencies))

	for k, v := range data.Frequencies {
		v.part = v.part / float64(len(text))
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
		data.Frequencies[data.Sorted[i].Symbol] = Place{data.Frequencies[data.Sorted[i].Symbol].part, i}
		inStupid = append(inStupid, []rune(data.Sorted[i].Symbol)[0])
	}
	data.InStupid = string(inStupid)
	f.currentData = data

	return data
}

func (f *FreqCryptAnalysator) Decrypt(text string) string {
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

var ForTS = regexp.MustCompile(`[^\p{Latin}\p{Cyrillic} ]`)

func Beautify(text string) string {
	text = strings.ToUpper(text)
	return ForTS.ReplaceAllString(text, "")
}
