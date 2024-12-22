package main

import (
	"aoc2024/util"
	"fmt"
	"strconv"
)

func main() {
	var tallyA, tallyB int
	for _, code := range util.ReadLines() {
		n, _ := strconv.Atoi(code[:len(code)-1])
		m := map[string]int{code: 1}
		for range 3 {
			m = next(m)
		}
		tallyA += n * count(m)
		for range 23 {
			m = next(m)
		}
		tallyB += n * count(m)
	}
	fmt.Println(tallyA)
	fmt.Println(tallyB)
}

func next(m map[string]int) map[string]int {
	out := make(map[string]int)
	for code, n := range m {
		code = "A" + code
		for i := range len(code) - 1 {
			out[replacements[code[i:i+2]]] += n
		}
	}
	return out
}

func count(m map[string]int) int {
	var tally int
	for code, n := range m {
		tally += n * len(code)
	}
	return tally
}

var replacements = map[string]string{
	"AA": "A", "A0": "<A", "A1": "^<<A", "A2": "<^A", "A3": "^A", "A4": "^^<<A", "A5": "<^^A", "A6": "^^A", "A7": "^^^<<A", "A8": "<^^^A", "A9": "^^^A",
	"0A": ">A", "00": "A", "01": "^<A", "02": "^A", "03": "^>A", "04": "^^<A", "05": "^^A", "06": "^^>A", "07": "^^^<A", "08": "^^^A", "09": "^^^>A",
	"1A": ">>vA", "10": ">vA", "11": "A", "12": ">A", "13": ">>A", "14": "^A", "15": "^>A", "16": "^>>A", "17": "^^A", "18": "^^>A", "19": "^^>>A",
	"2A": "v>A", "20": "vA", "21": "<A", "22": "A", "23": ">A", "24": "<^A", "25": "^A", "26": "^>A", "27": "<^^A", "28": "^^A", "29": "^^>A",
	"3A": "vA", "30": "<vA", "31": "<<A", "32": "<A", "33": "A", "34": "<<^A", "35": "<^A", "36": "^A", "37": "<<^^A", "38": "<^^A", "39": "^^A",
	"4A": ">>vvA", "40": ">vvA", "41": "vA", "42": "v>A", "43": "v>>A", "44": "A", "45": ">A", "46": ">>A", "47": "^A", "48": "^>A", "49": "^>>A",
	"5A": "vv>A", "50": "vvA", "51": "<vA", "52": "vA", "53": "v>A", "54": "<A", "55": "A", "56": ">A", "57": "<^A", "58": "^A", "59": "^>A",
	"6A": "vvA", "60": "<vvA", "61": "<<vA", "62": "<vA", "63": "vA", "64": "<<A", "65": "<A", "66": "A", "67": "<<^A", "68": "<^A", "69": "^A",
	"7A": ">>vvvA", "70": ">vvvA", "71": "vvA", "72": "vv>A", "73": "vv>>A", "74": "vA", "75": "v>A", "76": "v>>A", "77": "A", "78": ">A", "79": ">>A",
	"8A": "vvv>A", "80": "vvvA", "81": "<vvA", "82": "vvA", "83": "vv>A", "84": "<vA", "85": "vA", "86": "v>A", "87": "<A", "88": "A", "89": ">A",
	"9A": "vvvA", "90": "<vvvA", "91": "<<vvA", "92": "<vvA", "93": "vvA", "94": "<<vA", "95": "<vA", "96": "vA", "97": "<<A", "98": "<A", "99": "A",
	"A^": "<A", "A>": "vA", "Av": "<vA", "A<": "v<<A",
	"^A": ">A", "^^": "A", "^>": "v>A", "^v": "vA", "^<": "v<A",
	">A": "^A", ">^": "<^A", ">>": "A", ">v": "<A", "><": "<<A",
	"vA": "^>A", "v^": "^A", "v<": "<A", "vv": "A", "v>": ">A",
	"<A": ">>^A", "<^": ">^A", "<v": ">A", "<<": "A", "<>": ">>A",
}
