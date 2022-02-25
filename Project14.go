package main

import "fmt"

func main() {
	//format1
	// x := make(map[string]string)
	// x["TH"] = "THAILAND"
	// x["JP"] = "JAPAN"
	// x["EN"] = "ENGLAND"

	// fmt.Println(x["TH"])
	// fmt.Println(x["JP"])
	// fmt.Println(x["EN"])

	//format2
	x := map[string]string{
		"TH": "THAILAND",
		"JP": "JAPAN",
	}

	fmt.Println(x)
}
