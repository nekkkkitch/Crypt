package crypanalysis

import (
	"fmt"
	"log/slog"
	"testing"
)

func TestFreqAnalysis(t *testing.T) {
	var a Analysator

	text := `LFWKI MJCLP SISWK HJOGL KMVGU RAGKM KMXMA MJCVX WUYLG GIISW
ALXAE YCXMF KMKBQ BDCLA EFLFW KIMJC GUZUG SKECZ GBWYM OACFV
MQKYF WXTWM LAIDO YQBWF GKSDI ULQGV SYHJA VEFWB LAEFL FWKIM
JCFHS NNGGN WPWDA VMQFA AXWFZ CXBVE LKWML AVGKY EDEMJ XHUXD
AVYXL`
	data := a.Freq.Analyse(text, false)
	slog.Info("Data", "data", data.Frequencies)
	fmt.Println()
	slog.Info("Data", "data", data.Sorted)
	fmt.Println()
	slog.Info("Data", "data", data.InStupid)
	fmt.Println()
	try := a.Freq.Decrypt(text)
	slog.Info("Try", "try", try)
	ind := a.Index.Analyse(data)
	slog.Info("Data", "ind", ind)
	ac := a.AC.Analyse(text)
	slog.Info("ACData", "ac", ac)
	ka := a.Kasiski.Analyse(text)
	slog.Info("KAData", "ka", ka)
}
