package number

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNewNumber(t *testing.T) {
	ss := []string{
		"0",
		"1",
		"1.",
		"0.1",
		"1.0",
		"1.00",
		"0.00000000000001",
		"0.0000000000000123",
		"123.0000000000000123",
	}
	for _, s := range ss {
		n, err := NewNumber(s)
		if err != nil {
			t.Errorf("new number err:%v", err)
			return
		}
		fmt.Printf("bigint:%v precise=%v\n", n.ToString(), defaultPrecise)
		fmt.Printf("src:%s\n", s)
		fmt.Printf("str:%s\n", n.ToString())
	}
}

func TestCalc(t *testing.T) {
	calc("999", "5")
	calc("5", "999")
}

func calc(s1, s2 string) {
	v1, _ := NewNumber(s1)
	v2, _ := NewNumber(s2)
	v3 := v1.Copy()
	v3.Add(v2)
	fmt.Printf("%v+%v=%v\n", v1.ToString(), v2.ToString(), v3.ToString())
	v3 = v1.Copy()
	v3.Sub(v2)
	fmt.Printf("%v-%v=%v\n", v1.ToString(), v2.ToString(), v3.ToString())
	v3 = v1.Copy()
	v3.Mul(v2)
	fmt.Printf("%vx%v=%v\n", v1.ToString(), v2.ToString(), v3.ToString())
	v3 = v1.Copy()
	v3.Div(v2)
	fmt.Printf("%v/%v=%v\n", v1.ToString(), v2.ToString(), v3.ToString())
	v3.Floor(6)
	fmt.Printf("floor 6=%v\n", v3.ToString())
	v3 = v1.Copy()
	v3.MulDivInt(9, 10)
	fmt.Printf("%v*9/10=%v\n", v1.ToString(), v3.ToString())
	v3 = v1.Copy()
	v3.Mod(v2)
	fmt.Printf("%v mod %v=%v\n", v1.ToString(), v2.ToString(), v3.ToString())
	v3 = v1.Copy()
	v3.Rem(v2)
	fmt.Printf("%v rem %v=%v\n", v1.ToString(), v2.ToString(), v3.ToString())
	v3 = v1.Copy()
	fmt.Printf("%v cmp %v=%v\n", v1.ToString(), v2.ToString(), v3.Cmp(v2))
}

type NST struct {
	Value Number `json:"value"`
}

func TestJson(t *testing.T) {
	v1, _ := NewNumber("100.30000000003")
	d := NST{
		Value: *v1,
	}
	dat, err := json.Marshal(d)
	if err != nil {
		t.Errorf("json marshal err:%v", err)
		return
	}
	fmt.Printf("json:%v\n", string(dat))
	var d2 NST
	err = json.Unmarshal(dat, &d2)
	if err != nil {
		t.Errorf("json unmarshal err:%v", err)
		return
	}
	fmt.Printf("object:%#v\n", d2)
	fmt.Printf("object.value:%v\n", d2.Value.ToString())
}

type roundCaseSt struct {
	src     string
	precise int
	floor   string
	ceil    string
	round   string
}

func TestFloor(t *testing.T) {
	cases := []roundCaseSt{
		{
			"0",
			1,
			"0",
			"0",
			"0",
		},
		{
			"0.123",
			1,
			"0.1",
			"0.2",
			"0.1",
		},
		{
			"0.153",
			1,
			"0.1",
			"0.2",
			"0.2",
		},
		{
			"-1.133",
			1,
			"-1.2",
			"-1.1",
			"-1.1",
		},
		{
			"-1.153",
			1,
			"-1.2",
			"-1.1",
			"-1.2",
		},
	}
	for _, i := range cases {
		s := New(i.src)
		f := s.Copy().Floor(i.precise)
		fmt.Printf("Floor(%v, %v)=%v\n", i.src, i.precise, f.ToString())
		if f.Cmp(New(i.floor)) != 0 {
			t.Errorf("floor error expect %v get %v", i.floor, f.ToString())
		}
		c := s.Copy().Ceil(i.precise)
		fmt.Printf("Ceil(%v, %v)=%v\n", i.src, i.precise, c.ToString())
		if c.Cmp(New(i.ceil)) != 0 {
			t.Errorf("ceil error expect %v get %v", i.ceil, c.ToString())
		}
		r := s.Copy().Round(i.precise)
		fmt.Printf("Round(%v, %v)=%v\n", i.src, i.precise, r.ToString())
		if r.Cmp(New(i.round)) != 0 {
			t.Errorf("round error expect %v get %v", i.round, r.ToString())
		}
	}
}
