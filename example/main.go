package main

import (
	"fmt"
	"os"

	lw "github.com/onioncall/wrapt"
	"golang.org/x/term"
)

func main() {
	s := "The quick brown fox jumps over the lazy dog.\nBright vixens jump dozy fowl quack. My jocks box, get hard, very numb. Crazy Fredrick bought many very exquisite opal jewels. How razorback-jumping frogs can level six piqued gymnasts. Jaded zombies acted quaintly but kept driving their oxen forward. Sphinx of black quartz, judge my vow. Pack my box with five dozen liquor jugs. The five boxing wizards jump quickly. Jackdaws love my big sphinx of quartz. Five quacking zephyrs jolt my wax bed. Waltz, bad nymph, for quick jigs vex. Quick zephyrs blow, vexing daft Jim. How quickly daft jumping zebras vex. Two driven jocks help fax my big quiz. Sympathizing would fix quaker objectives. Pack my red box with five dozen quality jugs. Amazingly few discotheques provide jukeboxes. Junk MTV quiz graced by fox whelps."

	w, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		panic(err)
	}

	fmt.Print("Regular Print\n\n")
	fmt.Println(s)
	fmt.Print("\n")
	fmt.Print("----------------------\n\n")

	fmt.Print("Wrap\n\n")
	fmt.Println(lw.Wrap(s, w))
	fmt.Print("\n")
	fmt.Print("----------------------\n\n")

	fmt.Print("WrapArray\n\n")
	for _, s := range lw.WrapArray(s, 4, w) {
		fmt.Printf("    %s\n", s)
	}
}
