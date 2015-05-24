/* Parse XML
 */

package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

const x = `
<person>
  <name>
    <family> Smith </family>
    <personal> Jim </personal>
  </name>
  <email type="personal">
    jane.doe@gmail.com
  </email>
  <email type="work">
    jon.doe@gmail.com
  </email>
</person>
`

func main() {
	parser := xml.NewDecoder(strings.NewReader(x))
	depth := 0
	for {
		token, err := parser.Token()
		if err != nil {
			break
		}
		switch t := token.(type) {
		case xml.StartElement:
			elmt := xml.StartElement(t)
			name := elmt.Name.Local
			printElmt(name, depth)
			depth++
		case xml.EndElement:
			depth--
			elmt := xml.EndElement(t)
			name := elmt.Name.Local
			printElmt(name, depth)
		case xml.CharData:
			bytes := xml.CharData(t)
			printElmt("\""+string([]byte(bytes))+"\"", depth)
		case xml.Comment:
			printElmt("Comment", depth)
		case xml.ProcInst:
			printElmt("ProcInst", depth)
		case xml.Directive:
			printElmt("Directive", depth)
		default:
			fmt.Println("Unknown")
		}
	}
}

func printElmt(s string, depth int) {
	for n := 0; n < depth; n++ {
		fmt.Print("  ")
	}
	fmt.Println(s)
}


