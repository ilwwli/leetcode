package main

import "fmt"

func entityParser(text string) string {
	for i := 0; i < len(text); i++ {
		if text[i] == '&' {
			if i+3 < len(text) {
				if text[i:i+4] == "&gt;" {
					text = text[:i] + ">" + text[i+4:]
					continue
				}
				if text[i:i+4] == "&lt;" {
					text = text[:i] + "<" + text[i+4:]
					continue
				}
			}
			if i+4 < len(text) {
				if text[i:i+5] == "&amp;" {
					text = text[:i] + "&" + text[i+5:]
					continue
				}
			}
			if i+5 < len(text) {
				if text[i:i+6] == "&quot;" {
					text = text[:i] + "\"" + text[i+6:]
					continue
				}
				if text[i:i+6] == "&apos;" {
					text = text[:i] + "'" + text[i+6:]
					continue
				}
			}
			if i+6 < len(text) {
				if text[i:i+7] == "&frasl;" {
					text = text[:i] + "/" + text[i+7:]
					continue
				}
			}
		}
	}
	return text
}

func main() {
	fmt.Println(entityParser("x &gt; y &amp;&amp; x &lt; y is always false"))
}
