package main

import (
	"gitlab.com/kokegudiel2/analisisLexico/lexer"
)

func main() {
	lexer.LexElements("a = (b * c + d -- (e/100))")
	lexer.LexElements("YANCARLOS PERNILLO CASTILLO 09051619260")
	lexer.LexElements("YANCARLdOS PERNILLdO CASTILdLO 123412d1234")
}
