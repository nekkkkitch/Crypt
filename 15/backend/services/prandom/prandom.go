package prandom

type Randomer struct {
	mul  int64
	con  int64
	mask int64
	seed int64
}

func NewRandomer() Randomer {
	return Randomer{
		229,
		2347,
		47717,
		67128091,
	}
}

func (r *Randomer) roll() {
	r.seed = (r.mul*r.seed + r.con) % r.mask
}

// Returns random number in [0:n) diapason. Return -1 if n <= 0
func (r *Randomer) RandN(n int) int {
	if n <= 0 {
		return -1
	}
	r.roll()
	return int(AbsoluteModulo(r.seed, int64(n)))
}

func (r *Randomer) RandNoLimit() int64 {
	toRet := r.seed
	r.roll()
	return toRet
}

func (r *Randomer) RandPositive() int64 {
	toRet := r.seed
	if toRet < 0 {
		toRet *= -1
	}
	r.roll()
	return toRet
}

func AbsoluteModulo(num, mod int64) int64 {
	num = num % mod
	if num < 0 {
		num += mod
	}
	return num
}
