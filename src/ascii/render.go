package ascii

import("strings")


func Render(text string, banner map[rune][]string)string{
	var result strings.Builder

	input:=strings.Split(text, "\r\n")

	for _, word:=range input{

		for row:= range 8{
			var lines strings.Builder
			for _, char:= range word{
				font:=banner[char][row]
				lines.WriteString(font)
			}
			result.WriteString(lines.String()); result.WriteString("\n")
		}
	}
	return result.String()
}