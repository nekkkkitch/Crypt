package crypto

import (
	"15/backend/services/prandom"
	"encoding/binary"
	"log/slog"
)

type Gamma struct {
	Biba   int64
	Boba   int64
	modulo int
	base   int
	gamma  []int
}

type DES struct {
	key      string
	keyBytes []byte
}

var IP = []int{58, 50, 42, 34, 26, 18, 10, 2, 60, 52, 44, 36, 28, 20, 12, 4, 62, 54, 46, 38, 30, 22, 14, 6, 64, 56, 48, 40, 32, 24, 16, 8, 57, 49, 41, 33, 25, 17, 9, 1, 59, 51, 43, 35, 27, 19, 11, 3, 61, 53, 45, 37, 29, 21, 13, 5, 63, 55, 47, 39, 31, 23, 15, 7}
var FP = []int{40, 8, 48, 16, 56, 24, 64, 32, 39, 7, 47, 15, 55, 23, 63, 31, 38, 6, 46, 14, 54, 22, 62, 30, 37, 5, 45, 13, 53, 21, 61, 29, 36, 4, 44, 12, 52, 20, 60, 28, 35, 3, 43, 11, 51, 19, 59, 27, 34, 2, 42, 10, 50, 18, 58, 26, 33, 1, 41, 9, 49, 17, 57, 25}
var E = []int{32, 1, 2, 3, 4, 5, 4, 5, 6, 7, 8, 9, 8, 9, 10, 11, 12, 13, 12, 13, 14, 15, 16, 17, 16, 17, 18, 19, 20, 21, 20, 21, 22, 23, 24, 25, 24, 25, 26, 27, 28, 29, 28, 29, 30, 31, 32, 1}
var P = []int{16, 7, 20, 21, 29, 12, 28, 17, 1, 15, 23, 26, 5, 18, 31, 10, 2, 8, 24, 14, 32, 27, 3, 9, 19, 13, 30, 6, 22, 11, 4, 25}
var PC1 = []int{57, 49, 41, 33, 25, 17, 9, 1, 58, 50, 42, 34, 26, 18, 10, 2, 59, 51, 43, 35, 27, 19, 11, 3, 60, 52, 44, 36, 63, 55, 47, 39, 31, 23, 15, 7, 62, 54, 46, 38, 30, 22, 14, 6, 61, 53, 45, 37, 29, 21, 13, 5, 28, 20, 12, 4}
var PC2 = []int{14, 17, 11, 24, 1, 5, 3, 28, 15, 6, 21, 10, 23, 19, 12, 4, 26, 8, 16, 7, 27, 20, 13, 2, 41, 52, 31, 37, 47, 55, 30, 40, 51, 45, 33, 48, 44, 49, 39, 56, 34, 53, 46, 42, 50, 36, 29, 32}

func (g *Gamma) Cypher(input string) string {
	runed := []rune(input)
	newRuned := make([]rune, len(runed))
	g.generateGamma(len(runed))
	slog.Info("Cyphering", "Gamma", g.gamma)
	for i := range runed {
		newRuned[i] = rune(int(runed[i]) ^ (g.gamma[i]))
		slog.Info("Cyphering", "rune", newRuned[i])
	}
	return string(newRuned)
}

func (g *Gamma) Decypher(input string) string {
	runed := []rune(input)
	newRuned := make([]rune, len(runed))
	for i := range runed {
		newRuned[i] = rune(int(runed[i]) ^ (g.gamma[i]))
	}
	return string(newRuned)
}

func (g *Gamma) ChangeParams(params []interface{}) {
	g.modulo = params[0].(int)
	g.base = params[1].(int)
}

func (g *Gamma) generateGamma(L int) {
	gamma := make([]int, L)
	r := prandom.NewRandomer()
	for i := range gamma {
		g.Biba = r.RandPositive()
		g.Boba = r.RandPositive()
		slog.Info("gamma gen", "biba", g.Biba, "boba", g.Boba)

		BigBoba := ModExp(int64(g.base), g.Boba, int64(g.modulo))
		slog.Info("gamma gen", "BigBoba", BigBoba)

		K := int64(ModExp(int64(BigBoba), g.Biba, int64(g.modulo)))
		slog.Info("gamma gen", "K", K)

		gamma[i] = int(K)
	}
	g.gamma = gamma
}

func ModExp(a, b, mod int64) int64 {
	base := a
	for range b - 1 {
		a = prandom.AbsoluteModulo(a*base, mod)
	}
	return a
}

func (d *DES) Cypher(input []byte) []byte {
	malo := len(input) % 8
	if malo != 0 {
		for range malo {
			input = append(input, byte(0)) // добавить дополнительные 8 байт для длины malo?
		}
	}

	for i := 0; i < len(input)/8; i++ {
		current := input[8*i : 8*(i+1)]
		u := binary.BigEndian.Uint64(current)
		toIP := PermutateBits(u, 64, IP)

	}
	return nil
}

func (d *DES) Decypher(input string) string {
	return ""
}

func (d *DES) ChangeParams(params []interface{}) {
	key := params[0].(string)
	if key == "" {
		slog.Error("Key is empty")
		return
	}
	bytes := []byte(key)
	if len(bytes) > 8 {
		bytes = bytes[:8]
	}
	for len(bytes) < 8 {
		bytes = append(bytes, 0)
	}

	permKey := PermutateBits(binary.BigEndian.Uint64(bytes), 64, PC1)
	left := uint32(permKey >> 36)
	right := uint32((permKey << 28) >> 36)
}

func PermutateBits(bits uint64, length int, perm []int) uint64 {
	var result uint64
	for pos, mov := range perm {
		bit := (bits >> (length - mov)) & 1

		result |= bit << (63 - pos)
	}

	return result
}
