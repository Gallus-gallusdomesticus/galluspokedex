package main

import("strings")

func cleanInput(text string) []string {
    lowered:=strings.ToLower(text) //lower all uppercase
	nowhitespace := strings.TrimSpace(lowered) //remove all trailing whitespace

	separated := strings.Fields(nowhitespace) //separate into slice
	
	return separated
}
