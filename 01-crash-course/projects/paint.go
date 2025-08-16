package main

import "fmt"

// color structure
type Color struct {
    Red    string
    Green  string
    Yellow string
    Blue   string
    Purple string
    Cyan   string
    Gray   string
    White  string
}

/*

// Something abt this is off frfr

// Paint text with specific passwd via args
func paint(color, text string) string {
	return color + text
}

*/

func (color Color) red(text string) string {
	return color.Red + text
}

func (color Color) green(text string) string {
	return color.Green + text
}

func (color Color) yellow(text string) string {
	return color.Yellow + text
}

func (color Color) blue(text string) string {
	return color.Blue + text
}

func (color Color) purple(text string) string {
	return color.Purple + text
}

func (color Color) cyan(text string) string {
	return color.Cyan + text
}

func (color Color) gray(text string) string {
	return color.Gray + text
}

func (color Color) white(text string) string {
    return color.White + text
}

// MAIN
func main() {

	// Define colors
	color := Color {
		Red:    "\033[31m",
		Green:  "\032[32m",
		Yellow: "\033[33m",
		Blue:   "\033[34m",
		Purple: "\033[35m",
		Cyan:   "\033[36m",
		Gray:   "\033[37m",
		White:  "\033[97m",
	}

	/*
		fmt.Println(paint(color.Red, " I will be RED"))
		fmt.Println(paint(color.Purple, " I will be Purple"))
		fmt.Println(paint(color.Cyan, " I will be Cyan"))
		fmt.Println(paint(color.Blue, " I will be Blue"))
	*/

	fmt.Println(color.red("Roses are red"))
	fmt.Println(color.blue("Violets are blue"))
	fmt.Println(color.cyan("And I don't know what to do"))
}
