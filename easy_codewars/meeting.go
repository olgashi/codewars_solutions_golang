package main

import (
	"fmt"
	"strings"
	"sort"
)

func Meeting(s string) string {
  namesList := strings.Split(strings.ReplaceAll(strings.ToUpper(s), ":", " "), ";")

	for index, name := range namesList {
		fullName := strings.Split(name, " ") 
		fullName[0], fullName[1] = fullName[1], fullName[0]
		namesList[index] = "(" + fullName[0] + ", " + fullName[1] +")"
	}

	sort.Strings(namesList)
	return strings.Join(namesList, "")
}

func main() {
	fmt.Println(Meeting("Fred:Corwill;Wilfred:Corwill;Barney:Tornbull;Betty:Tornbull;Bjon:Tornbull;Raphael:Corwill;Alfred:Corwill")) // (CORWILL, ALFRED)(CORWILL, FRED)(CORWILL, RAPHAEL)(CORWILL, WILFRED)(TORNBULL, BARNEY)(TORNBULL, BETTY)(TORNBULL, BJON)
	fmt.Println(Meeting("Alexis:Wahl;John:Bell;Victoria:Schwarz;Abba:Dorny;Grace:Meta;Ann:Arno;Madison:STAN;Alex:Cornwell;Lewis:Kern;Megan:Stan;Alex:Korn")) // "(ARNO, ANN)(BELL, JOHN)(CORNWELL, ALEX)(DORNY, ABBA)(KERN, LEWIS)(KORN, ALEX)(META, GRACE)(SCHWARZ, VICTORIA)(STAN, MADISON)(STAN, MEGAN)(WAHL, ALEXIS)"
	fmt.Println(Meeting("Abc:xyz;Abc:xzz")) // ("ABC,XYZ")("ABC,XZZ")
	fmt.Println(Meeting("John:Gates;Michael:Wahl;Megan:Bell;Paul:Dorries;James:Dorny;Lewis:Steve;Alex:Meta;Elizabeth:Russel;Anna:Korn;Ann:Kern;Amber:Cornwell")) // "(BELL, MEGAN)(CORNWELL, AMBER)(DORNY, JAMES)(DORRIES, PAUL)(GATES, JOHN)(KERN, ANN)(KORN, ANNA)(META, ALEX)(RUSSEL, ELIZABETH)(STEVE, LEWIS)(WAHL, MICHAEL)"	
}