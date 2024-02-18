package number

import (
	"errors"
	"math/big"

	"database/sql/driver"
)

var (
	ErrInvalidNumberChar = errors.New("invalid number char")
	ErrInvalidPrecise    = errors.New("invalid number precise")
	ErrInvalidScanValue  = errors.New("invalid scan value")
)

const (
	AdjustNone  = 0
	AdjustFloor = -1
	AdjustCeil  = 1
	AdjustRound = 45
)

var narr = []*big.Int{
	big.NewInt(0),
	big.NewInt(1),
	big.NewInt(2),
	big.NewInt(3),
	big.NewInt(4),
	big.NewInt(5),
	big.NewInt(6),
	big.NewInt(7),
	big.NewInt(8),
	big.NewInt(9),
	big.NewInt(10),
}

var defaultPrecise int = 16
var base *big.Int

func init() {
	base = big.NewInt(1)
	for i := 0; i < defaultPrecise; i++ {
		base.Mul(base, narr[10])
	}
}

type Number struct {
	v big.Int
}

func New(s string) (n *Number) {
	n, _ = NewNumber(s)
	return
}

func NewNumber(s string) (n *Number, err error) {
	n = &Number{}
	err = n.ParseString(s)
	if err != nil {
		return nil, err
	}
	return
}

func (n *Number) Copy() *Number {
	v := Number{}
	return v.Add(n)
}

func (n *Number) Sign() int {
	return n.v.Sign()
}

func (n *Number) ParseString(s string) (err error) {
	idx := 0
	neg := false
	if len(s) > 0 {
		if s[0] == '+' {
			idx++
		} else if s[0] == '-' {
			neg = true
			idx++
		}
	}
	for idx < len(s) {
		c := s[idx]
		if c == '.' {
			break
		} else if c < '0' || c > '9' {
			return ErrInvalidNumberChar
		}
		n.v.Mul(&n.v, narr[10])
		n.v.Add(&n.v, narr[c-'0'])
		idx += 1
	}
	l := 0
	if idx < len(s) && s[idx] == '.' {
		idx += 1
		for l < defaultPrecise && idx < len(s) {
			c := s[idx]
			if c < '0' || c > '9' {
				return ErrInvalidNumberChar
			}
			n.v.Mul(&n.v, narr[10])
			n.v.Add(&n.v, narr[c-'0'])
			l += 1
			idx += 1
		}
		for idx < len(s) {
			c := s[idx]
			if c < '0' || c > '9' {
				return ErrInvalidNumberChar
			}
			idx += 1
		}
	}
	for l < defaultPrecise {
		n.v.Mul(&n.v, narr[10])
		l += 1
	}
	if neg {
		n.v.Neg(&n.v)
	}
	return
}

func (n *Number) ToString() string {
	sign := ""
	s := n.v.Text(10)
	l := len(s)
	if l > 0 && s[0] == '-' {
		sign = "-"
		s = s[1:]
		l--
	}
	var prefix string
	var suffix string
	if defaultPrecise == 0 || s == "0" {
		return sign + s
	} else if l > defaultPrecise {
		prefix = s[:l-defaultPrecise]
		suffix = s[l-defaultPrecise:]
	} else {
		prefix = "0"
		for l < defaultPrecise {
			suffix += "0"
			l += 1
		}
		suffix += s
	}
	l = len(suffix)
	for l > 0 && suffix[l-1] == '0' {
		l -= 1
	}
	if l > 0 {
		return sign + prefix + "." + suffix[:l]
	}
	return sign + prefix
}

func (n *Number) SetInt64(i int64) *Number {
	n.v.SetInt64(i)
	n.v.Mul(&n.v, base)
	return n
}

func (n *Number) Abs() *Number {
	n.v.Abs(&n.v)
	return n
}

func (n *Number) Neg() *Number {
	n.v.Neg(&n.v)
	return n
}

func (n *Number) Add(oth *Number) *Number {
	n.v.Add(&n.v, &oth.v)
	return n
}

func (n *Number) Sub(oth *Number) *Number {
	n.v.Sub(&n.v, &oth.v)
	return n
}

func (n *Number) Mul(oth *Number) *Number {
	n.v.Mul(&n.v, &oth.v)
	n.v.Div(&n.v, base)
	return n
}

func (n *Number) MulDivInt(m int64, d int64) *Number {
	v := big.NewInt(m)
	n.v.Mul(&n.v, v)
	v.SetInt64(d)
	n.v.Div(&n.v, v)
	return n
}

func (n *Number) Div(oth *Number) *Number {
	n.v.Mul(&n.v, base)
	n.v.Div(&n.v, &oth.v)
	return n
}

func (n *Number) Mod(oth *Number) *Number {
	n.v.Mod(&n.v, &oth.v)
	return n
}

func (n *Number) Rem(oth *Number) *Number {
	n.v.Rem(&n.v, &oth.v)
	return n
}

func (n *Number) Cmp(oth *Number) int {
	return n.v.Cmp(&oth.v)
}

func (n *Number) Floor(p int) *Number {
	if p >= defaultPrecise {
		return n
	}
	g := defaultPrecise - p
	v := big.NewInt(1)
	for i := 0; i < g; i++ {
		v.Mul(v, narr[10])
	}
	var m big.Int
	m.Mod(&n.v, v)
	if m.Sign() > 0 {
		n.v.Sub(&n.v, &m)
	} else if m.Sign() < 0 {
		v.Add(v, &m)
		n.v.Sub(&n.v, v)
	}
	return n
}

func (n *Number) Ceil(p int) *Number {
	if p >= defaultPrecise {
		return n
	}
	g := defaultPrecise - p
	v := big.NewInt(1)
	for i := 0; i < g; i++ {
		v.Mul(v, narr[10])
	}
	var m big.Int
	m.Mod(&n.v, v)
	if m.Sign() > 0 {
		v.Sub(v, &m)
		n.v.Add(&n.v, v)
	} else if m.Sign() < 0 {
		n.v.Sub(&n.v, &m)
	}
	return n
}

func (n *Number) Round(p int) *Number {
	if p >= defaultPrecise {
		return n
	}
	g := defaultPrecise - p
	v := big.NewInt(1)
	for i := 0; i < g; i++ {
		v.Mul(v, narr[10])
	}
	var v2 big.Int
	v2.Div(v, narr[2])
	var m big.Int
	m.Mod(&n.v, v)
	if m.Sign() < 0 {
		m.Add(&m, v)
	}
	if m.Cmp(&v2) < 0 {
		n.v.Sub(&n.v, &m)
	} else {
		v.Sub(v, &m)
		n.v.Add(&n.v, v)
	}
	return n
}

func (n *Number) Adjust(p int, adjust int) *Number {
	switch adjust {
	case AdjustRound:
		n.Round(p)
	case AdjustCeil:
		n.Ceil(p)
	case AdjustFloor:
		n.Floor(p)
	}
	return n
}

func (n Number) MarshalJSON() ([]byte, error) {
	s := "\"" + n.ToString() + "\""
	return []byte(s), nil
}

func (n *Number) UnmarshalJSON(text []byte) error {
	l := len(text)
	if l > 1 && text[0] == '"' && text[l-1] == '"' {
		return n.ParseString(string(text[1 : l-1]))
	}
	return n.ParseString(string(text))
}

func (n Number) Value() (v driver.Value, err error) {
	return n.ToString(), nil
}

func (n *Number) Scan(value interface{}) (err error) {
	switch value.(type) {
	case []byte:
		s := string(value.([]byte))
		return n.ParseString(s)
	case string:
		s := value.(string)
		return n.ParseString(s)
	default:
		return ErrInvalidScanValue
	}
}
