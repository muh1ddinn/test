package main

import "fmt"

func mergeStrings(str1, str2 string) string {
	var mergedString string
	maxlen := len(str1)
	if len(str2) > maxlen {
		maxlen = len(str2)
	}

	for i := 0; i < maxlen; i++ {
		if i < len(str1) {
			mergedString += string(str1[i])
			fmt.Println(mergedString) 
		}

		if i < len(str2) {
			mergedString += string(str2[i])
			fmt.Println(mergedString) 
		}
	}

	return mergedString
}

func main() {
	str1 := "hello"
	str2 := "world"
	result := mergeStrings(str1, str2)
	fmt.Println(result) // This will print "hweolrllod"
}


