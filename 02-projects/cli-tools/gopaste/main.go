/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/ShahandFahad/golang-sandbox/02-projects/cli-tools/gopaste/cmd"
	"github.com/ShahandFahad/golang-sandbox/02-projects/cli-tools/gopaste/config"
)

func main() {
	config.LoadConfig()
	cmd.Execute()
}
/*
GPT: Prompt to help
Yes - Let's implement the logic.
But before that - Please take some time to explain it to me what you did earlier and why you did it. Explain it to me and logically an conceptually.
And from now on. Please also add explanation paragraphs like why we following this structure or why are we using this concept & design and and also give it a professional industry level touch and also add helpful comments as well so that i can have a better clarity.
*/
