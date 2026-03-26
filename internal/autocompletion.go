package internal

import (
	"strings"
	"unicode/utf8"
)

type CompletionData struct {
	theType string
	theData []string
}

var CurrentCompletionData map[string][]string

func ContextAutocompletion(line string, pos int, key rune) (newLine string, newPos int, ok bool) {
	newLine = line
	newPos = pos
	ok = false
	//defer log.Printf("defer in ContextAutocompletion '%s'", line)
	if key == rune(9) { // "Tab" at pos "pos" (count starting at 0)
		lineArray := CleanInput(line)
		if len(lineArray) < 1 {
			return
		}
		/*log.Printf("lineArray: %v, len: %d, runes in first: %d, pos: %d\n",
		lineArray, len(lineArray), utf8.RuneCountInString(lineArray[0]), pos)*/
		var commonPrefix string
		var fits bool
		if pos <= utf8.RuneCountInString(lineArray[0]) {
			// cursor in first word => command
			//log.Printf("first: %s\n", lineArray[0])
			commonPrefix, fits = searchForCompletion(lineArray[0], "command")
			newLine = commonPrefix
		} else if len(lineArray) > 1 {
			//log.Printf("runes in second: %d", utf8.RuneCountInString(lineArray[1]))
			if pos <= utf8.RuneCountInString(lineArray[0])+utf8.RuneCountInString(lineArray[1])+1 {
				//log.Printf("second: %s\n", lineArray[1])
				// cursor in second word => param; type depending on command
				commonPrefix, fits = searchForCompletion(lineArray[1], Registry[lineArray[0]].ParamType)
				//log.Println("common: ", commonPrefix)
				newLine = strings.Join([]string{lineArray[0], commonPrefix}, " ")
			}
		} else {
			//log.Printf("later or only  white spaces after first word... pos %d", pos)
			commonPrefix = ""
			fits = false
		}

		if fits {
			newPos = utf8.RuneCountInString(newLine)
			//log.Printf("yes, new pos %d, newLine %s", newPos, newLine)
		}
		ok = fits
	}
	return
}

func searchForCompletion(wordStart string, dataType string) (string, bool) {
	// for testing
	/*	testData := CompletionData{
			theType: "area",
			theData: []string{"abcde", "12345", "abctu", "234567890", "23456789"},
		}
	*/
	var fittingWords []string
	//log.Printf("searching for type %s. We have saved: %v\b", dataType, CurrentCompletionData[dataType])

	for _, word := range CurrentCompletionData[dataType] {
		if strings.HasPrefix(word, wordStart) {
			fittingWords = append(fittingWords, word)
		}
	}
	if len(fittingWords) == 0 {
		return "", false
	}
	if len(fittingWords) == 1 {
		return fittingWords[0], true
	}
	sharedPrefix := fittingWords[0]
	for i := 0; i < len(fittingWords); i++ {
		for j := i + 1; j < len(fittingWords); j++ {
			prefix := getSharedPrefix(fittingWords[i], fittingWords[j], wordStart)
			//log.Printf("%s, %s, common=%s\n", fittingWords[i], fittingWords[j], prefix)
			if len(prefix) < len(sharedPrefix) {
				sharedPrefix = prefix
			}
		}
	}
	return sharedPrefix, true
}

func getSharedPrefix(s1 string, s2 string, commonPrefix string) string {
	firstSplit := strings.Split(s1, "")
	secondSplit := strings.Split(s2, "")
	k := len(commonPrefix)
	for ; k < min(len(firstSplit), len(secondSplit)); k++ {
		if firstSplit[k] != secondSplit[k] {
			break
		}
	}
	sharedPrefix := make([]string, k)
	copy(sharedPrefix, firstSplit)
	return strings.Join(sharedPrefix, "")
}
