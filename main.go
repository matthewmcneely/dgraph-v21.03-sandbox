package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Pallinder/go-randomdata"
)

const itemCount = 10000
const offset = 500

func main() {

	for n := 0; n < itemCount; n++ {
		//_:book_%d <nid> %q .\n", i, strconv.Itoa(i+offset)
		line := fmt.Sprintf("_:book_%d <nid> %q .", n, strconv.Itoa(n+offset))
		fmt.Println(line)
		s := strings.SplitAfterN(randomdata.Paragraph(), " ", 10)[:9]
		title := strings.Title(strings.Join(s, ""))
		title = strings.TrimRight(title, "., ")
		line = fmt.Sprintf("_:book_%d <book_name> %q .", n, title)
		fmt.Println(line)
		line = fmt.Sprintf("_:book_%d <dgraph.type> %q .", n, "Book")
		fmt.Println(line)
	}
}
