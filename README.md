Сolors and styles in your terminal

Usage example:

import (
    "github.com/sgkochnev/colorout"
	"github.com/sgkochnev/colorout/color/bg"
	"github.com/sgkochnev/colorout/color/fg"
	"github.com/sgkochnev/colorout/style"
)

type character struct {
		name string
		age  uint8
	}

var chars = []character{
	{
		name: "Heihachi",
		age:  75,
	},
	{
		name: "Kazuya",
		age:  49,
	},
	{
		name: "Jin",
		age:  21,
	},
}

You can use the lib like this (1):

	fstr := colorout.New() //(1)
	fstr.Style(fg.Green).Format("%-10s") //(2)
	fstr.Style(fg.Default).Format(":") //(3)
	fstr.Style(fg.Cyan, style.Bold, style.Italic).Format(" %-3d") //(4)
	fstr.Style(fg.Default, style.NotItalic) //(5)
	format := fstr.String() //(6)


	for _, char := range chars {
		fmt.Printf(format", char.name, char.age)

		// All styles are reset
		// The string will be printed with default styles
		fmt.Println("description")
	}

Or (2):

	for _, char := range chars {
		fstr.Println(char.name, char.age)
	}

	// since the format string expects 2 arguments, 
	// but 3 arguments are passed to print, 
	// the third argument will be written separately.
	// styles for it are set in the line 5
	for _, char := range chars {
		fstr.Println(char.name, char.age, "description")
	}


	for i, char := range chars {
		fstr.Printf("%v %v %v\n", "description", "for", i+1, char.name, char.age)
	}

	// or

	for i, char := range chars {
		description := fmt.Sprintf("%v %v %v\n", "description", "for", i+1)
		fstr.Print(char.name, char.age, description)
	}


Example default options:

	// or default option
	for i, char := range chars {
		description := fmt.Sprintf("%v %v %v\n", "description", "for", i+1)
		fmt.Printf("\x1b[32m%-10s\x1b[39m:\x1b[36;1;3m %-3d\x1b[39;23m %s\x1b[0m",
			char.name, char.age, description,
	)
