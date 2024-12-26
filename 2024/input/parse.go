package input

import (
	"aoc2024/util"
	"bytes"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
)

func ReadFile[T any](day int, parse func([]byte) T) func() T {
	file := "real"
	if len(os.Args) > 1 {
		file = os.Args[1]
	}
	return func() T {
		bytes, err := os.ReadFile(fmt.Sprintf("day%d/%s.txt", day, file))
		if err != nil {
			panic(err)
		}
		return parse(bytes)
	}
}

func String(b []byte) string {
	return string(b)
}

func ByteArray(b []byte) []byte {
	return b
}

func ByteMatrix(b []byte) util.ByteMatrix {
	return util.ByteMatrix(bytes.Split(b, []byte("\n")))
}

func Number(b []byte) int {
	n, err := strconv.Atoi(string(b))
	if err != nil {
		panic(err)
	}
	return n
}

func Bool(falseVal, trueVal string) func([]byte) bool {
	return func(b []byte) bool {
		switch string(b) {
		case falseVal:
			return false
		case trueVal:
			return true
		default:
			panic("invalid value")
		}
	}
}

func List[T any](sep string, parse func([]byte) T) func([]byte) []T {
	return func(b []byte) []T {
		bs := bytes.Split(b, []byte(sep))
		out := make([]T, len(bs))
		for i, b := range bs {
			out[i] = parse(b)
		}
		return out
	}
}

func Map[K comparable, V any](entrySep, kvSep string, parseKey func([]byte) K, parseValue func([]byte) V) func([]byte) map[K]V {
	return func(b []byte) map[K]V {
		m := make(map[K]V)
		for _, b := range bytes.Split(b, []byte(entrySep)) {
			kv := bytes.SplitN(b, []byte(kvSep), 2)
			m[parseKey(kv[0])] = parseValue(kv[1])
		}
		return m
	}
}

func structFields[T any](types ...reflect.Type) []int {
	var example T
	t := reflect.TypeOf(example)

	fields := make(map[int]int)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if tag, ok := field.Tag.Lookup("input"); ok {
			n, err := strconv.Atoi(tag)
			if err != nil {
				panic(fmt.Errorf("field %s has invalid input tag", field.Name))
			}
			if n < 1 || n > len(types) {
				panic(fmt.Errorf("unexpected input tag %d on field %s", n, field.Name))
			}
			if existingField, has := fields[n-1]; has {
				panic(fmt.Errorf("duplicate input tag %d for fields %s and %s", n, t.Field(existingField).Name, field.Name))
			}
			if field.Type != types[n-1] {
				panic(fmt.Errorf("type mismatch for field %s with input tag %d", field.Name, n))
			}
			fields[n-1] = i
		}
	}

	out := make([]int, len(types))
	for i := range out {
		if n, has := fields[i]; has {
			out[i] = n
		} else {
			panic(fmt.Errorf("missing field with input tag %d", i+1))
		}
	}

	return out
}

type Tuple2[T1, T2 any] struct {
	V1 T1 `input:"1"`
	V2 T2 `input:"2"`
}

func Struct2[T, T1, T2 any]() func(v1 T1, v2 T2) T {
	var v1 T1
	var v2 T2
	fields := structFields[T](reflect.TypeOf(v1), reflect.TypeOf(v2))
	return func(v1 T1, v2 T2) T {
		var val T
		v := reflect.ValueOf(&val).Elem()
		v.Field(fields[0]).Set(reflect.ValueOf(v1))
		v.Field(fields[1]).Set(reflect.ValueOf(v2))
		return val
	}
}

func Sequence2[T, T1, T2 any](sep string, parse1 func([]byte) T1, parse2 func([]byte) T2) func([]byte) T {
	create := Struct2[T, T1, T2]()
	return func(b []byte) T {
		bs := bytes.SplitN(b, []byte(sep), 2)
		return create(parse1(bs[0]), parse2(bs[1]))
	}
}

func Regex2[T, T1, T2 any](regex string, parse1 func([]byte) T1, parse2 func([]byte) T2) func([]byte) T {
	r := regexp.MustCompile("^" + regex + "$")
	if r.NumSubexp() != 2 {
		panic("invalid number of subexpressions")
	}
	create := Struct2[T, T1, T2]()
	return func(b []byte) T {
		m := r.FindSubmatch(b)
		if m == nil {
			panic(fmt.Errorf("input doesn't match %s", regex))
		}
		return create(parse1(m[1]), parse2(m[2]))
	}
}

type Tuple3[T, T1, T2, T3 any] struct {
	V1 T1 `input:"1"`
	V2 T2 `input:"2"`
	V3 T3 `input:"3"`
}

func Struct3[T, T1, T2, T3 any]() func(v1 T1, v2 T2, v3 T3) T {
	var v1 T1
	var v2 T2
	var v3 T3
	fields := structFields[T](reflect.TypeOf(v1), reflect.TypeOf(v2), reflect.TypeOf(v3))
	return func(v1 T1, v2 T2, v3 T3) T {
		var val T
		v := reflect.ValueOf(&val).Elem()
		v.Field(fields[0]).Set(reflect.ValueOf(v1))
		v.Field(fields[1]).Set(reflect.ValueOf(v2))
		v.Field(fields[2]).Set(reflect.ValueOf(v3))
		return val
	}
}

func Regex3[T, T1, T2, T3 any](regex string, parse1 func([]byte) T1, parse2 func([]byte) T2, parse3 func([]byte) T3) func([]byte) T {
	r := regexp.MustCompile("^" + regex + "$")
	if r.NumSubexp() != 3 {
		panic("invalid number of subexpressions")
	}
	create := Struct3[T, T1, T2, T3]()
	return func(b []byte) T {
		m := r.FindSubmatch(b)
		if m == nil {
			panic(fmt.Errorf("input doesn't match %s", regex))
		}
		return create(parse1(m[1]), parse2(m[2]), parse3(m[3]))
	}
}

func Sequence3[T, T1, T2, T3 any](sep string, parse1 func([]byte) T1, parse2 func([]byte) T2, parse3 func([]byte) T3) func([]byte) T {
	create := Struct3[T, T1, T2, T3]()
	return func(b []byte) T {
		bs := bytes.SplitN(b, []byte(sep), 3)
		return create(parse1(bs[0]), parse2(bs[1]), parse3(bs[2]))
	}
}

type Tuple4[T, T1, T2, T3, T4 any] struct {
	V1 T1 `input:"1"`
	V2 T2 `input:"2"`
	V3 T3 `input:"3"`
	V4 T4 `input:"4"`
}

func Struct4[T, T1, T2, T3, T4 any]() func(v1 T1, v2 T2, v3 T3, v4 T4) T {
	var v1 T1
	var v2 T2
	var v3 T3
	var v4 T4
	fields := structFields[T](reflect.TypeOf(v1), reflect.TypeOf(v2), reflect.TypeOf(v3), reflect.TypeOf(v4))
	return func(v1 T1, v2 T2, v3 T3, v4 T4) T {
		var val T
		v := reflect.ValueOf(&val).Elem()
		v.Field(fields[0]).Set(reflect.ValueOf(v1))
		v.Field(fields[1]).Set(reflect.ValueOf(v2))
		v.Field(fields[2]).Set(reflect.ValueOf(v3))
		v.Field(fields[3]).Set(reflect.ValueOf(v4))
		return val
	}
}

func Sequence4[T, T1, T2, T3, T4 any](sep string, parse1 func([]byte) T1, parse2 func([]byte) T2, parse3 func([]byte) T3, parse4 func([]byte) T4) func([]byte) T {
	create := Struct4[T, T1, T2, T3, T4]()
	return func(b []byte) T {
		bs := bytes.SplitN(b, []byte(sep), 4)
		return create(parse1(bs[0]), parse2(bs[1]), parse3(bs[2]), parse4(bs[3]))
	}
}

func Regex4[T, T1, T2, T3, T4 any](regex string, parse1 func([]byte) T1, parse2 func([]byte) T2, parse3 func([]byte) T3, parse4 func([]byte) T4) func([]byte) T {
	r := regexp.MustCompile("^" + regex + "$")
	if r.NumSubexp() != 4 {
		panic("invalid number of subexpressions")
	}
	create := Struct4[T, T1, T2, T3, T4]()
	return func(b []byte) T {
		m := r.FindSubmatch(b)
		if m == nil {
			panic(fmt.Errorf("input doesn't match %s", regex))
		}
		return create(parse1(m[1]), parse2(m[2]), parse3(m[3]), parse4(m[4]))
	}
}

type Tuple5[T, T1, T2, T3, T4, T5 any] struct {
	V1 T1 `input:"1"`
	V2 T2 `input:"2"`
	V3 T3 `input:"3"`
	V4 T4 `input:"4"`
	V5 T5 `input:"5"`
}

func Struct5[T, T1, T2, T3, T4, T5 any]() func(v1 T1, v2 T2, v3 T3, v4 T4, v5 T5) T {
	var v1 T1
	var v2 T2
	var v3 T3
	var v4 T4
	var v5 T5
	fields := structFields[T](reflect.TypeOf(v1), reflect.TypeOf(v2), reflect.TypeOf(v3), reflect.TypeOf(v4), reflect.TypeOf(v5))
	return func(v1 T1, v2 T2, v3 T3, v4 T4, v5 T5) T {
		var val T
		v := reflect.ValueOf(&val).Elem()
		v.Field(fields[0]).Set(reflect.ValueOf(v1))
		v.Field(fields[1]).Set(reflect.ValueOf(v2))
		v.Field(fields[2]).Set(reflect.ValueOf(v3))
		v.Field(fields[3]).Set(reflect.ValueOf(v4))
		v.Field(fields[4]).Set(reflect.ValueOf(v5))
		return val
	}
}

func Sequence5[T, T1, T2, T3, T4, T5 any](sep string, parse1 func([]byte) T1, parse2 func([]byte) T2, parse3 func([]byte) T3, parse4 func([]byte) T4, parse5 func([]byte) T5) func([]byte) T {
	create := Struct5[T, T1, T2, T3, T4, T5]()
	return func(b []byte) T {
		bs := bytes.SplitN(b, []byte(sep), 5)
		return create(parse1(bs[0]), parse2(bs[1]), parse3(bs[2]), parse4(bs[3]), parse5(bs[4]))
	}
}

func Regex5[T, T1, T2, T3, T4, T5 any](regex string, parse1 func([]byte) T1, parse2 func([]byte) T2, parse3 func([]byte) T3, parse4 func([]byte) T4, parse5 func([]byte) T5) func([]byte) T {
	r := regexp.MustCompile("^" + regex + "$")
	if r.NumSubexp() != 5 {
		panic("invalid number of subexpressions")
	}
	create := Struct5[T, T1, T2, T3, T4, T5]()
	return func(b []byte) T {
		m := r.FindSubmatch(b)
		if m == nil {
			panic(fmt.Errorf("input doesn't match %s", regex))
		}
		return create(parse1(m[1]), parse2(m[2]), parse3(m[3]), parse4(m[4]), parse5(m[5]))
	}
}

type Tuple6[T, T1, T2, T3, T4, T5, T6 any] struct {
	V1 T1 `input:"1"`
	V2 T2 `input:"2"`
	V3 T3 `input:"3"`
	V4 T4 `input:"4"`
	V5 T5 `input:"5"`
	V6 T6 `input:"6"`
}

func Struct6[T, T1, T2, T3, T4, T5, T6 any]() func(v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6) T {
	var v1 T1
	var v2 T2
	var v3 T3
	var v4 T4
	var v5 T5
	var v6 T6
	fields := structFields[T](reflect.TypeOf(v1), reflect.TypeOf(v2), reflect.TypeOf(v3), reflect.TypeOf(v4), reflect.TypeOf(v5), reflect.TypeOf(v6))
	return func(v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6) T {
		var val T
		v := reflect.ValueOf(&val).Elem()
		v.Field(fields[0]).Set(reflect.ValueOf(v1))
		v.Field(fields[1]).Set(reflect.ValueOf(v2))
		v.Field(fields[2]).Set(reflect.ValueOf(v3))
		v.Field(fields[3]).Set(reflect.ValueOf(v4))
		v.Field(fields[4]).Set(reflect.ValueOf(v5))
		v.Field(fields[5]).Set(reflect.ValueOf(v6))
		return val
	}
}

func Sequence6[T, T1, T2, T3, T4, T5, T6 any](sep string, parse1 func([]byte) T1, parse2 func([]byte) T2, parse3 func([]byte) T3, parse4 func([]byte) T4, parse5 func([]byte) T5, parse6 func([]byte) T6) func([]byte) T {
	create := Struct6[T, T1, T2, T3, T4, T5, T6]()
	return func(b []byte) T {
		bs := bytes.SplitN(b, []byte(sep), 6)
		return create(parse1(bs[0]), parse2(bs[1]), parse3(bs[2]), parse4(bs[3]), parse5(bs[4]), parse6(bs[5]))
	}
}

func Regex6[T, T1, T2, T3, T4, T5, T6 any](regex string, parse1 func([]byte) T1, parse2 func([]byte) T2, parse3 func([]byte) T3, parse4 func([]byte) T4, parse5 func([]byte) T5, parse6 func([]byte) T6) func([]byte) T {
	r := regexp.MustCompile("^" + regex + "$")
	if r.NumSubexp() != 6 {
		panic("invalid number of subexpressions")
	}
	create := Struct6[T, T1, T2, T3, T4, T5, T6]()
	return func(b []byte) T {
		m := r.FindSubmatch(b)
		if m == nil {
			panic(fmt.Errorf("input doesn't match %s", regex))
		}
		return create(parse1(m[1]), parse2(m[2]), parse3(m[3]), parse4(m[4]), parse5(m[5]), parse6(m[6]))
	}
}
