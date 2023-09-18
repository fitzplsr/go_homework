package uniq

import (
	"fmt"
	"testing"
)

var tests = []struct {
	in      []string
	out     []string
	options Options
}{
	{in: []string{
		"I love music.",
		"I love music.",
		"I love music.",
		"",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
		"I love music of Kartik.",
		"I love music of Kartik."},
		out: []string{
			"I love music.",
			"",
			"I love music of Kartik.",
			"Thanks.",
			"I love music of Kartik."},
		options: Options{},
	},
	{in: []string{
		"I love music.",
		"I love music.",
		"I love music.",
		"",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
		"I love music of Kartik.",
		"I love music of Kartik."},
		out: []string{
			"3 I love music.",
			"1 ",
			"2 I love music of Kartik.",
			"1 Thanks.",
			"2 I love music of Kartik."},
		options: Options{C: true}},
	{in: []string{
		"I love music.",
		"I love music.",
		"I love music.",
		"",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
		"I love music of Kartik.",
		"I love music of Kartik."},
		out: []string{
			"I love music.",
			"I love music of Kartik.",
			"I love music of Kartik."},
		options: Options{D: true}},
	{in: []string{
		"I love music.",
		"I love music.",
		"I love music.",
		"",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
		"I love music of Kartik.",
		"I love music of Kartik."},
		out: []string{
			"",
			"Thanks."},
		options: Options{U: true}},
	{in: []string{
		"I LOVE MUSIC.",
		"I love music.",
		"I LoVe MuSiC.",
		"",
		"I love MuSIC of Kartik.",
		"I love music of kartik.",
		"Thanks.",
		"I love music of kartik.",
		"I love MuSIC of Kartik."},
		out: []string{
			"I LOVE MUSIC.",
			"",
			"I love MuSIC of Kartik.",
			"Thanks.",
			"I love music of kartik."},
		options: Options{I: true}},
	{in: []string{
		"We love music.",
		"I love music.",
		"They love music.",
		"",
		"I love music of Kartik.",
		"We love music of Kartik.",
		"Thanks."},
		out: []string{
			"We love music.",
			"",
			"I love music of Kartik.",
			"Thanks."},
		options: Options{F: 1}},
	{in: []string{
		"I love music.",
		"A love music.",
		"C love music.",
		"",
		"I love music of Kartik.",
		"We love music of Kartik.",
		"Thanks."},
		out: []string{
			"I love music.",
			"",
			"I love music of Kartik.",
			"Thanks."},
		options: Options{F: 1}},
	{in: []string{
		"I LOVE MUSIC.",
		"We love music.",
		"me LoVe MuSiC.",
		"",
		"I love MuSIC of Kartik.",
		"I love music of kartik.",
		"Thanks.",
		"I love music of kartik.",
		"I love MuSIC of Kartik."},
		out: []string{
			"3 I LOVE MUSIC.",
			"1 ",
			"2 I love MuSIC of Kartik.",
			"1 Thanks.",
			"2 I love music of kartik."},
		options: Options{C: true, I: true, F: 1}},
	{in: []string{
		"I LOVE MUSIC.",
		"We love music.",
		"me LoVe MuSiC.",
		"",
		"I love MuSIC of Kartik.",
		"I love music of kartik.",
		"Thanks.",
		"I love music of kartik.",
		"I love MuSIC of Kartik."},
		out: []string{
			"I LOVE MUSIC.",
			"I love MuSIC of Kartik.",
			"I love music of kartik."},
		options: Options{D: true, I: true, F: 1}},
	{in: []string{
		"I LOVE MUSIC.",
		"We love music.",
		"me LoVe MuSiC.",
		"",
		"I love MuSIC of Kartik.",
		"I love music of kartik.",
		"Thanks.",
		"I love music of kartik.",
		"I love MuSIC of Kartik."},
		out: []string{
			"",
			"Thanks."},
		options: Options{U: true, I: true, F: 1}},
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestUniq(t *testing.T) {
	for n, tt := range tests {
		t.Run(fmt.Sprintf("TEST#%d", n), func(t *testing.T) {
			res := Uniq(tt.in, tt.options)
			if !equal(res, tt.out) {
				t.Errorf("got %s, want %s", res, tt.out)
			}
		})
	}
}
