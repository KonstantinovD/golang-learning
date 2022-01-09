package main

import (
	"fmt"
	"strings"
	"unicode"
)

var f = fmt.Printf
var line = fmt.Println

func main() {
	upper := strings.ToUpper("Hello there!")
	f("To Upper: %s\n", upper)
	f("To Lower: %s\n", strings.ToLower("Hello THERE"))
	f("%s\n", strings.Title("tHis wiLL be A title!"))
	line()

	f("EqualFold: %v\n", // сравнение строк без учета регистра
		strings.EqualFold("Mihalis", "MIHAlis"))
	f("EqualFold: %v\n",
		strings.EqualFold("Mihalis", "MIHAli"))
	line()

	// = starts from
	f("Prefix: %v\n", strings.HasPrefix("Mihalis", "Mi"))
	f("Prefix: %v\n", strings.HasPrefix("Mihalis", "mi"))
	// = ends with
	f("Suffix: %v\n", strings.HasSuffix("Mihalis", "is"))
	f("Suffix: %v\n", strings.HasSuffix("Mihalis", "IS"))
	line()

	f("Index: %v\n", strings.Index("Mihalis", "ha"))
	f("Index: %v\n", strings.Index("Mihalis", "Ha"))
	line()

	f("Count: %v\n", strings.Count("Mihalis", "i"))
	f("Count: %v\n", strings.Count("Mihalis", "I"))
	line()

	f("Repeat: %s\n", strings.Repeat("ab", 5))
	line()

	f("TrimSpace: %s\n",
		strings.TrimSpace(" \tThis is a line. \n"))
	f("TrimLeft: %s\n",
		strings.TrimLeft(" \tThis is a\t line. \n", "\n\t "))
	f("TrimRight: %s\n",
		strings.TrimRight(" \tThis is a\t line. \n", "\n\t "))
	line()

	// --- COMPARATION
	f("Compare: %v\n", strings.Compare("Mihalis", "MIHALIS"))
	f("Compare: %v\n", strings.Compare("Mihalis", "Mihalis"))
	f("Compare: %v\n", strings.Compare("MIHALIS", "MIHalis"))
	line()

	// разбивает строку, переданную в качестве параметра, на части,
	// sep = space
	f("Fields: %v\n", strings.Fields("This is a string!"))
	f("Fields: %v\n", strings.Fields("Thisis\na\tstring!"))
	line()

	// разбивает строку, переданную в качестве параметра, на части,
	// sep can be changed
	f("%s\n", strings.Split("abcd efg", ""))
	line()

	f("Replace %s\n",
		strings.Replace("abcd efg", "", "_", -1))
	f("Replace %s\n",
		strings.Replace("abcd efg", "", "_", 4))
	f("Replace %s\n",
		strings.Replace("abcd efg", "", "_", 2))
	f("Replace %s\n",
		strings.Replace("abcd efg", "e", "k", 2))
	// last param of Replace func - max amount of replacements can be
	// done. If == -1; amount isn't limited
	line()

	lines := []string{"Line 1", "Line 2", "Line 3"}
	f("Join: %s\n", strings.Join(lines, "+++"))
	f("SplitAfter: %s\n",
		strings.SplitAfter("123++432++", "++"))
	line()

	// функция обрезки, позволяет сохранять интересующие вас руны строки
	trimFunction := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	f("TrimFunc: %s\n",
		strings.TrimFunc("123 abc ABC \t .", trimFunction))
}
