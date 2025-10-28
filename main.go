package main

import (
	"strings"

	"golang.org/x/net/html"
)

func main() {
	maliciousHTML := strings.Repeat("<template>", 100000)

	html.Parse(strings.NewReader(maliciousHTML))
}
