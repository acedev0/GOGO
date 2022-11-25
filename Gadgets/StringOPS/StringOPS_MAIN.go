/*   GOGO_Gadgets  - Useful string and conversion functions

---------------------------------------------------------------------------------------
NOTE: For Functions or Variables to be globally availble. The MUST start with a capital letter. (This is a GO Thing)

	Aug 27, 2021	-  initial Rollout

*/

package CUSTOM_GO_MODULE

import (
	// = = = = = Native Libraries
		// "crypto/md5"
		// "encoding/hex"
		"regexp"
		"strings"
		"unicode"
		
	// = = = = = CUSTOM Libraries
		. "github.com/acedev0/GOGO/Gadgets"

	// = = = = = 3rd Party Libraries

)






// Returns a JSON output of a struct  



// func GET_MD5_HASH(inval string) string {


// 	hasher := md5.New()
// 	hasher.Write([]byte(inval))
// 	result := hex.EncodeToString(hasher.Sum(nil))

// 	C.Print("Now Generation MD5 Hash: ")
// 	Y.Println(result)

// 	// color.Green(" Hey the hasval is: " + result)

// 	return result
// }

// // Alias for GET_MD5_HASH
// func GENERATE_MD5(inval string) string {
// 	return GET_MD5_HASH(inval)
// }





// this acts as a "delimiter" matrix..It returns true if any of the below delimiters exists
// UBER_SPLIT is dependant on this
func helper_Delimiter_Matrix_CATCHER(r rune) bool {

	/* In this case return true if there is ANY of the following:

	   - HYPHEN
	   _ UNDERSCORE
	   : Colon
	   / Forward Slash
	   | Pipe
	   & Amphersand
	   ' ' Just Space


	*/
	return r == ':' || r == '-' || r == '/' || r == '_' || r == '|' || r == '=' || r == '&' || r == ' '
}

/*
 	This is an UBER split on delimiter routine i made that
	makes it easy to return delimited values on MULTIPLE delimiters
	(check the above items in helper_Delimiter_Matrix_CATCHER)

	It also trims spaces before and AFTER each element
*/
func UBER_Split(myText string) []string {

	ptempVals := strings.FieldsFunc(myText, helper_Delimiter_Matrix_CATCHER)

	for x := 0; x < len(ptempVals); x++ {

		ptempVals[x] = strings.TrimSpace(ptempVals[x])

	}

	return ptempVals
} //end of function


// This splits only on PIPE....and trims space before and after
func PIPE_SPLIT(incoming string) []string {

	ptempVals := strings.Split(incoming, "|")

	// Lets go through each element and trim the spaces from the end
	for x := 0; x < len(ptempVals); x++ {

		ptempVals[x] = strings.TrimSpace(ptempVals[x])

	}

	return ptempVals

} //end of PIPE_SPLIT


// Make sthe first character of a string UPPER CASE
func UpperFirst(inString string) string {

	a := []rune(inString)
	a[0] = unicode.ToUpper(a[0])

	return string(a)
}




// Returns true if the string contains ONLY numbers
func HasOnlyNumbers(s string) bool {
    for _, r := range s {
        if (r < '0' || r > '9') {
            return false
        }
    }
    return true
} //end of func

// This removes all spaces from a string via unicode
func UNICODE_REMOVE_ALL_SPACES(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}

		return r
	}, str)
}


// This takes in a string and removes all non alphanumeric chars from it.. and extra spaces
func CLEAN_STRING(input string) string {
	C.Println("\nCleansing string of no-alpha chars")
	justAlpha, _ := regexp.Compile("[^a-zA-Z0-9_ ]")
	killExtraSpace := regexp.MustCompile(`[\s\p{Zs}]{2,}`)

	PASS_1 := justAlpha.ReplaceAllString(input, "")
	FINAL_PASS := killExtraSpace.ReplaceAllString(PASS_1, " ")

	return FINAL_PASS
} //end of func

//  This removes all extra spaces in a string 
func REMOVE_Extra_Spaces(input string) string {

	re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	final := re_leadclose_whtsp.ReplaceAllString(input, "")
	final = re_inside_whtsp.ReplaceAllString(final, " ")

	return final
}

func TrimSuffix(s, suffix string) string {
    if strings.HasSuffix(s, suffix) {
        s = s[:len(s)-len(suffix)]
    }
    return s
}


/*
	 = = = = END OF SECTION  = = = = = = = = = = = = = = = = = = = = = = = = = == = 
*/

