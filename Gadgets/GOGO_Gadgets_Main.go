/*   GOGO_Gadgets  - Useful multi-purpose GO functions to make GO DEV easier
	 by TerryCowboy

------------------------------------------------------------------------------------
NOTE: For Functions or Variables to be globally availble. The MUST start with a capital letter. (This is a GO Thing)

	Nov 24, 2022	- MAJOR CLEANUP AND RIVSIONS 
	
	Oct 22, 2022	- Cleanup and revisions
	Apr 22, 2022	- ... check out GET_FUNC_PARAM_DYNAMIC!!!
	Aug 27, 2021	- Ripped out a bunch of stuff to make this smaller. They are 
	Jun 05, 2014    - Initial Rollout

*/

package CUSTOM_GO_MODULE

import (

	// = = = = = Native Libraries
		"flag"		
		"math"
		"math/rand"
		"os"
		"os/exec"
		"runtime"
		"strconv"
		"strings"
		"time"
		"bufio"
		"unicode/utf8"
		"net/http"
		"io/ioutil"
		"encoding/json"
		"crypto/md5"
		"encoding/hex"		
		"io"


	// = = = = = 3rd Party Libraries

		"github.com/atotto/clipboard"
		"github.com/briandowns/spinner"
		"github.com/dustin/go-humanize"
		"github.com/fatih/color"
		mini "github.com/janeczku/go-spinner"
)

/*
	- - - -
	- - - -
	- - - - START OF GLOBALS WE NEED - - - - - -
	- - - -
	- - - -
*/

var SERIAL_NUM = "" // This is unique execution sid generated everytime a program starts. useful for troubleshooting in jenkins
var SHOW_SERIAL = false 		// If set, we Generate and show a serial number

var OSTYPE=""
var CURRENT_OS = ""	
var GOOS_VALUE = "" 		// Holds the current OS as reported by runtime.GOOS

var DEBUG_MODE = false		// Useful universal flag for enabling DEBUG_MODE code blocks
var ERROR_EXIT_CODE = -9999
var ERROR_CODE = ERROR_EXIT_CODE

var PROG_START_TIME string
var PROG_START_TIMEOBJ time.Time

var GLOBAL_CURR_DATE = ""		// Current Actual Date in the Timezone we specified
var GLOBAL_CURR_TIME = ""		// alias to CURR_DATE

var GLOBAL_DATE_OBJ time.Time		// Actual Global Date OBJ
var GLOBAL_TIME_OBJ time.Time		// alias



// -=-=-= COMMON COLOR GLOBAL references =-=-=-=-

var R = color.New(color.FgRed, color.Bold)
var G = color.New(color.FgGreen, color.Bold)
var Y = color.New(color.FgYellow, color.Bold)
var B = color.New(color.FgBlue, color.Bold)
var M = color.New(color.FgMagenta, color.Bold)
var C = color.New(color.FgCyan, color.Bold)
var W = color.New(color.FgWhite, color.Bold)

var R2 = color.New(color.FgRed)
var G2 = color.New(color.FgGreen)
var Y2 = color.New(color.FgYellow)
var B2 = color.New(color.FgBlue)
var M2 = color.New(color.FgMagenta)
var C2 = color.New(color.FgCyan)
var W2 = color.New(color.FgWhite)

var R3 = color.New(color.FgRed, color.Underline)
var G3 = color.New(color.FgGreen, color.Underline)
var Y3 = color.New(color.FgYellow, color.Underline)
var B3 = color.New(color.FgBlue, color.Underline)
var M3 = color.New(color.FgMagenta, color.Underline)
var C3 = color.New(color.FgCyan, color.Underline)
var W3 = color.New(color.FgWhite, color.Underline)


// Time Zone related stuff
var DEFAULT_ZONE_LOCATION_OBJ, _ = time.LoadLocation("Local")

var ZONE_LOCAL = ""
var ZONE_HOUR_OFFSET = "" 
var ZONE_UPPER = ""
var ZONE_FULL = ""

var USE_PST = false
var USE_CST = false
var USE_EST = false
var USE_MST = false
var USE_UTC = false

// These are used by SET_TIMEZONE_DEFAULTS.....and....GET_CURRENT_TIME 
var EST_OBJ, _ = time.LoadLocation("EST")
var CST_OBJ, _ = time.LoadLocation("America/Chicago")			// aka CST	}
var MST_OBJ, _ = time.LoadLocation("MST") 		// MDT / Mountain Standard
var PST_OBJ, _ = time.LoadLocation("America/Los_Angeles")		// aka PST
var UTC_OBJ, _ = time.LoadLocation("UTC") 


var DEFAULT_FLOAT = -999999999.999999999
var DEFAULT_INT = -999999999 

/*
	= = = = = = = = = = = = = = = = = = = = =
	End of GLOBALS definitions
	= = = = = = = = = = = = = = = = = = = = =
*/



/*
	SORTING Strings and Lists.. Cheat Sheet!

	var NAMES []string

	NAMES = append(NAMES, "Jenny")
	NAMES = append(NAMES, "Alec")
	NAMES = append(NAMES, "Kendal")
	NAMES = append(NAMES, "Carter")

	//4. To Sort the names in NAMES alphabetically, 
	sort.Strings(NAMES)


	// For sorting Structs Alphabetically 
	type STOCK_OBJ struct {

		SYMBOL		string
		Date		string
		Price		float64
		TIME_OBJ	time.Time
	}

	var STOCKS []STOCK_OBJ

	var S STOCK_OBJ

	S.SYMBOL = "NFLX"
	S.Date	= "12/05/2005"
	S.Price = 353.24
	STOCKS = append(STOCKS, S)

	S.SYMBOL = "AAPL"
	S.Date	= "10/10/2010"
	S.Price = 210.24
	STOCKS = append(STOCKS, S)

	S.SYMBOL = "ZNFL"
	S.Date	= "11/11/2002"
	S.Price = 53.24
	STOCKS = append(STOCKS, S)

	S.SYMBOL = "BAC"
	S.Date	= "01/05/2003"
	S.Price 98.24
	STOCKS = append(STOCKS, S)

	// At this point the list is unordered (items are entered in the order shown above)
	// To Sort ALPHABETICALLY do this: (if you want to sort in REVERSE alpha order, use > (greater than)) 
	
	//  (Note, the same applies to INTS and Floats)

    sort.Slice(STOCKS, func(i, j int) bool {
        return STOCKS[i].SYMBOL < STOCKS[j].SYMBOL
    })


	// Finally to sort the above slice /struct by way of the time.Time TIME_OBJ:
	
	sort.Slice(STOCKS, func(i, j int) bool { return (STOCKS)[i].TIME_OBJ.Before((STOCKS)[j].TIME_OBJ)})	

	(Will sort the slice by order of time/date the record was entered in say mongo)

	
	 
*/


func STRING_to_FLOAT(input string) float64 {
	result, _ :=  strconv.ParseFloat(input, 32)
	return result
}

// Converts a float to a string.. If another number is specified, it is interpreted as decimal precision
func FLOAT_to_STRING(input float64,  ALL_PARAMS ...int) string {

	var prec = -1
	for p, VAL := range ALL_PARAMS {

		if p == 0 {
			prec = VAL
		}	
	}

	result := strconv.FormatFloat(input, 'f', prec, 64)
	
	return result
}

func INT_to_STRING(input int) string {
	return strconv.Itoa(input)
}

func INT64_to_STRING(input int64) string {
	return strconv.Itoa(int(input))
}


// If first is more recent than second, we return true
func MOST_RECENT_DATE(first time.Time, second time.Time) bool {

	if first.After(second) {
		return true
	}
	
	return false
}

// GEts the MD5 of a string
func GET_MD5(input string) string {

    hasher := md5.New()
    hasher.Write([]byte(input))

	//Get the 16 bytes hash
	hashInBytes := hasher.Sum(nil)[:16]

	//Convert the bytes to a string
	returnMD5String := hex.EncodeToString(hashInBytes)
	
	return returnMD5String
}


func GET_FILE_MD5(filePath string) (string, error) {
	//Initialize variable returnMD5String now in case an error has to be returned
	var returnMD5String string
	//Open the passed argument and check for any error
	file, err := os.Open(filePath)
	if err != nil {
		R.Println(" GET MD5 File ERR", err)
		return returnMD5String, err
	}
	//Tell the program to call the following function when the current function returns
	defer file.Close()
	//Open a new hash interface to write to
	hash := md5.New()
	//Copy the file in the hash interface and check for any error
	if _, err := io.Copy(hash, file); err != nil {
		R.Println(" GET MD5 File ERR", err)
		return returnMD5String, err
	}
	//Get the 16 bytes hash
	hashInBytes := hash.Sum(nil)[:16]
	//Convert the bytes to a string
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String, nil
}



var s *mini.Spinner

func MINI_SpinStart() {
	s = mini.StartNew("")
}

func MINI_SpinSTOP() {
	s.Stop()
}




type IP struct {
    Query string
}

func GET_PUBLIC_IP() string {
    req, err := http.Get("http://ip-api.com/json/")
    if err != nil {
        return err.Error()
    }
    defer req.Body.Close()

    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        return err.Error()
    }

    var ip IP
    json.Unmarshal(body, &ip)

    return ip.Query
}



// Determines if a file is a sym link or not
func IS_FILE_LINK(filename string) (bool, string) {

	FINFO, err := os.Lstat(filename)
	if err != nil {
		R.Println("IS_FILE_LINK: ", err)
		return false, ""
	}

	//2b. Detect if this is a symlink
	if FINFO.Mode()&os.ModeSymlink != 0 {
		originFile, err43 := os.Readlink(filename)
		if err43 != nil {
			R.Println(" IS_FILE_LINK Detect SymLink err: ", err43)
			return false, ""
		}

		// Otherwise we have asymlink!
		// Return true and the SOURCE the file points to

		return true, originFile
	}	

	return false, ""
}


func IS_DIR(filename string) bool {

	FINFO, err := os.Lstat(filename)
	if err != nil {
		R.Println("IS_DIR: ", err)
		return false
	}

	if FINFO.IsDir() {
		return true
	}

	return false
}





func TOUCH_FILE(filename string) {
	currentTime := time.Now().Local()
	err := os.Chtimes(filename, currentTime, currentTime)
	if err != nil {
		R.Println(" ERROR Touching File!!", err)
	}
}

/*
  This padds a string that is passed and makes it a total LENGH of TOTAL_FINAL_LEN
  Pass a character as a string param and it will pad with THAT char instead of space
*/
func PAD_STRING(input string, TOTAL_FINAL_LEN int, ALL_PARAMS ...string) string {

	// This is the default padd string
	var padString = " "

	//2. If you want to use another just pass it as a string
	for p, VAL := range ALL_PARAMS {

		if p == 0 {
			padString = VAL

			//2b. also add space to before and after input so the pad string char isnt running up against it
			input = " " + input + " "
		}

	}	

	// Courtesy of: https://gist.github.com/asessa/3aaec43d93044fc42b7c6d5f728cb039

	var padLength = TOTAL_FINAL_LEN
	var inputLength = len(input)
	var padStringLength = len(padString)

	length := (float64(padLength - inputLength)) / float64(2)
	repeat := math.Ceil(length / float64(padStringLength))
	output := strings.Repeat(padString, int(repeat))[:int(math.Floor(float64(length)))] + input + strings.Repeat(padString, int(repeat))[:int(math.Ceil(float64(length)))]

	return output
}


func DELETE_from_LIST[T any](slice []T, s int) []T {
    return append(slice[:s], slice[s+1:]...)
}
func REMOVE_from_LIST[T any](slice []T, s int) []T {
    return append(slice[:s], slice[s+1:]...)
}
func DELETE_ITEM[T any](slice []T, s int) []T {
    return append(slice[:s], slice[s+1:]...)
}


// This has limited color support and if it receives a string with the color in |cyan| format..it prints in that color
func helper_SHOW_with_COLOR(input string, SHOW_OUTPUT bool) (string, string) {

	// var JUST_STRING = temps[2]
	var JUST_STRING = input
	var COLOR = ""

	if strings.Count(JUST_STRING, "|") == 2 {
		temps := strings.Split(input, "|")
		COLOR = temps[1]
		JUST_STRING = temps[2]

		if SHOW_OUTPUT {

			switch COLOR {
			case "cyan":
				C.Print(JUST_STRING)
				break

			case "green":
				G.Print(JUST_STRING)
				break

			case "yellow":
				Y.Print(JUST_STRING)
				break
			case "red":
				R.Print(JUST_STRING)
				break				
			default:
				W.Print(JUST_STRING)
				break
			}
		}

		COLOR = "|" + COLOR + "|"

	} else {
		if SHOW_OUTPUT {
			W.Print(JUST_STRING)
		}
	}

	return JUST_STRING, COLOR

} // end of


// Alias to SHOW_BOX
func SHOW_BOX_MESSAGE (ALL_PARAMS ...string) {
	SHOW_BOX(ALL_PARAMS...)
}

// Alias to SHOW_BOX
func SHOW_MESSAGE_BOX (ALL_PARAMS ...string) {
	SHOW_BOX(ALL_PARAMS...)
}

/*
	This is a nice way of showing a message in a box
	Just pass each line you want in the box as a seperate parameter

╭――――――――――――――――――╮
│                  │
│                  │
│                  │
╰――――――――――――――――――╯
*/
var BOX_INDENT_SPACES = "         "
func SHOW_BOX(ALL_PARAMS ...string) {

	var lines []string

	var SPACE_PREFIX = "       "

	//1. if multiple lines are passed, lets iterate through them
	for _, VAL := range ALL_PARAMS {
		lines = append(lines, VAL)
	}

	//2. FIgure out which line is the LONGEST.. this is how we grow the box length
	largest_len := 0

	for _, l := range lines {

		var JUST_STRING, _ = helper_SHOW_with_COLOR(l, false)
		temp_len := len(JUST_STRING)

		if temp_len > largest_len {
			largest_len = temp_len
		}
	} //end of line len determine for

	largest_len += len(SPACE_PREFIX) + 4

	//3. Now drop top of box
	var BOX_TOP = "┌"
	var BOX_BOTTOM = "└"
	for x := 0; x < largest_len; x++ {
		BOX_TOP += "─"
		BOX_BOTTOM += "─"
	}
	//4. CLose the ends of the BOX
	BOX_TOP += "┐"
	BOX_BOTTOM += "┘"

	//5. MUST use the utf8 way to get the string length since it contains ASCII chars
	var BOX_LEN = utf8.RuneCountInString(BOX_TOP)
	BOX_LEN = BOX_LEN - 2 // We have to do BOXLEN-2 to account for the Right and Left angle Brakcets

	//6. Top of Box
	M.Println(SPACE_PREFIX + BOX_TOP)

	//7. Prints the Lines in between top and bottom
	for _, line := range lines {

		var temp_full_line, MCOLOR = helper_SHOW_with_COLOR(line, false)

		// Most likely the temp_full line is LESS than the BOX_LEN.. so lets padd it
		if len(temp_full_line) < BOX_LEN {
			temp_full_line = PAD_STRING(temp_full_line, BOX_LEN)
		}

		M.Print(SPACE_PREFIX + "│")

		helper_SHOW_with_COLOR(MCOLOR+temp_full_line, true)

		M.Println("│")
	}

	//8. Prints the BOTTOM of box.. we are DONE
	M.Println(SPACE_PREFIX + BOX_BOTTOM)

	//9. Add an indent so the next thing that is C.Println'd ... will be indented under the box
	C.Println("")


} //end of SHOW_BOX_MESSAGE


func getProcessOwner() string {
    stdout, err := exec.Command("ps", "-o", "user=", "-p", strconv.Itoa(os.Getpid())).Output()
    if err != nil {
        R.Println("Error with AM_I_ROOT_ getProcess Owner", err)
        os.Exit(1)
    }

	//Y.Println(" RETURNED is", string(stdout))
    return string(stdout)
}

func AM_I_ROOT() bool {

	result := getProcessOwner()
	if strings.Contains(result, "root") {
		return true
	}

	return false
	
}

func DO_EXIT(EXTRA_ARGS ...string) {

	var bequiet = false 

	//1. Parse out EXTRA_ARGS
	for _, VAL := range EXTRA_ARGS {

		if VAL != "" {
			if VAL == "silent" {
				bequiet = true
			}			
		}

	} // end of for	

	if bequiet == false {
		SHOW_BOX("|red| Forced Program Exit!")
	}
	os.Exit(ERROR_EXIT_CODE)
}

// alias
func DOEXIT(EXTRA_ARGS ...string) {
	DO_EXIT(EXTRA_ARGS...)
}



/*
   ADD_LEADING_ZERO: This takes in a number and returns a string with a leading 0
   If the number is already 10 or greater, it returns that same number as is
 
	SHOW_PRETTY_DATE is dependant on this
*/
func ADD_LEADING_ZERO( myNum int) string {

	RESULT := strconv.Itoa(myNum)

	if myNum <= 9 {
		RESULT = "0" + RESULT
	}

	return RESULT
}

// Trims the first characgter from a string
func TRIM_FIRST(s string) string {
    _, i := utf8.DecodeRuneInString(s)
    return s[i:]
}


// This makes sure we're running as run.. program exists otherwise
func MAKE_Sure_Running_As_ROOT() {
	if AM_I_ROOT() == false {

		M.Println(" WARNING: You HAVE to use sudo to run this program!")
		C.Println("")
		C.Println("")
		DOEXIT("silent")
	}

}


/* SHOW_PRETTY_DATE Takes in a time.Time DATE_OBJ and returns a PRETTY formatted based on what you specify
   - Returns a STRING and a WEEKDAY
   - needs ADD_LEADING_ZERO
*/
func SHOW_PRETTY_DATE(input_DATE time.Time, EXTRA_ARGS...string) (string, string) {
	var output_FORMAT = "short"

	//1. Parse out EXTRA_ARGS
	for _, VAL := range EXTRA_ARGS {


		//1e. only parameter this takes is the output format we want
		// If full is passed, we show this format: Wednesday, 11/20/2020 @ 13:56
		// if british or iso is passed, we show: 2015-05-30
		if VAL != "" {
			output_FORMAT = VAL
			continue
		}

	} // end of for

	//2. From this object, extract the M/D/Y HH:MM
	montemp := int(input_DATE.Month())
	daytemp := input_DATE.Day()

	hourtemp := input_DATE.Hour()
	mintemp := input_DATE.Minute()

	//3. Then, we add leading 0's as needed
	cMon := ADD_LEADING_ZERO(montemp)
	cDay := ADD_LEADING_ZERO(daytemp)
	cHour := ADD_LEADING_ZERO(hourtemp)
	cMin := ADD_LEADING_ZERO(mintemp)
	
	sectemp := input_DATE.Second()
	cSec := ADD_LEADING_ZERO(sectemp)
	nanotemp := input_DATE.Nanosecond()

	cNanoSecond := strconv.Itoa(nanotemp)



	//4. Thankfully we dont have to worry about this fuckery with the year!
	cYear := strconv.Itoa(input_DATE.Year())
	weekd := input_DATE.Weekday().String()



	//5. Update ZONE info 
	tmp_zone, tmp_offset := input_DATE.Zone()

	tmp_off_string := strconv.Itoa(tmp_offset)

	TMP_ZONE_FULL := "(" + tmp_zone + " " + tmp_off_string + ")"

	/* 7. Here is the DEFAULT Pretty format that is returned

		09/26/1978 @ 13:58

			or (if SHOW_SECONDS is passed) 

		09/26/1978 @ 13:58:05
	*/
	//result_TEXT := cMon + "/" + cDay + "/" + cYear + " @ " + cHour + ":" + cMin
	result_TEXT := "nullDATE_specified"
	
	//8. SHORT Format is:  Wednesday, 11/20/2001
	if output_FORMAT == "short" {

		result_TEXT = weekd + ", " + cMon + "/" + cDay + "/" + cYear + " @ " + cHour + ":" + cMin

	//9. FULL Format: //Wednesday, 11/20/2020 @ 13:56 EST (-5 Hours)
	} else if output_FORMAT == "full" {
		
		result_TEXT = weekd + ", " + cMon + "/" + cDay + "/" + cYear + " @ " + cHour + ":" + cMin + ":" + cSec

	} else if output_FORMAT == "zone" {
		result_TEXT = weekd + ", " + cMon + "/" + cDay + "/" + cYear + " @ " + cHour + ":" + cMin + ":" + cSec + "_" + TMP_ZONE_FULL

	} else if output_FORMAT == "nano" {
		result_TEXT = weekd + ", " + cMon + "/" + cDay + "/" + cYear + " @ " + cHour + ":" + cMin + ":" + cSec + ":n:" + cNanoSecond
	
	//10. This is the british/iso format: 2020-09-26
	} else if output_FORMAT == "british" || output_FORMAT == "iso" {

		result_TEXT = cYear + "-" + cMon + "-" + cDay

	//11. This is JUSTDATE:  09/26/1988
	} else if output_FORMAT == "justdate" {

		result_TEXT = weekd + "_" + cMon + "_" + cDay + "_" + cYear
	
	
	//12. For use as a simple timestamp for a file suffix
	} else if output_FORMAT == "timestamp" {
	
		result_TEXT = weekd + "_" + cMon + "_" + cDay + "_" + cYear + "_" + cHour + "_" + cMin
	

	} else {

		R.Println(" ERROR in SHOW_PRETTY_DATE: .. invalid output_FORMAT sent!!!")
		os.Exit(-9)
	}

	//12. As a bonus, we always return the weekday as a second variable

	return result_TEXT, weekd
} //end of func



// This utilizes the HUMANIZE library and shows a HUMAN readable number of the passed variable
func ShowNum(innum int) string {

	result := humanize.Comma(int64(innum))

	// result = strconv.Itoa(result)
	return result
}


// Shows a pretty Number based on passed FLOAT
func ShowNum_FLOAT(innum float64) string {

        result := humanize.Comma(int64(innum))

        // result = strconv.Itoa(result)
        return result
}




// 64bit version of this.. not sure why im using this yet
func ShowNum64(innum int64) string {

	result := humanize.Comma(innum)

	// result = strconv.Itoa(result)
	return result
}

// When called, copies a specified string to the users CLIPBOARD
func CLIPBOARD_COPY(instring string) {
	clipboard.WriteAll(instring)
}



func GET_VAR_TYPE(ALL_PARAMS ...interface{}) string {

	for _, param := range ALL_PARAMS {
		_, IS_INT := param.(int)
		_, IS_FLOAT := param.(float64) 
		_, IS_STRING := param.(string) 
		_, IS_BOOL := param.(bool) 

		if IS_INT {
			return "int"
		}

		if IS_FLOAT {
			return "float"
		}

		if IS_STRING {
			return "string"
		}

		if IS_BOOL {
			return "bool"
		}
	}

	return " ( Unknown Type ) "
}


/*
	= = = = = = 
	= = = = = =  Better versions of the IS_xxx functions
	= = = = = = 
*/
func IS_INT(param interface{}) bool {
	_, found_int := param.(int) 
	_, found_int32 := param.(int32) 
	_, found_int64 := param.(int64)
	
	if found_int || found_int32 || found_int64{
		return true
	}

	return false
}
func IS_STRING(param interface{}) bool {

	_, found_string := param.(string) 
	return found_string
}

func IS_FLOAT(param interface{}) bool {
	_, found_32 := param.(float32) 
	_, found_64 := param.(float64) 
	
	if found_32 || found_64 {
		return true
	}

	return false
}

func IS_BOOL(param interface{}) bool {

	_, found_bool := param.(bool) 
	return found_bool
}

/*
	FOr Dynamic Parameter support... DONT FORGET THE ... thre dots!!!

	Examples:
		1. for function call:

			WRITE_EMBED("ca.pem.gz", "verbose", "quiet", 55, "www.podshop.com")

		2. Then write the function like this:

			func WRITE_EMBED(name string, EXTRA_ARGS ...interface{}) string {

				var verbose = GET_EXTRA_ARG("verbose",  EXTRA_ARGS...).(bool)		// specify verbose will set to TRUE... case verbase was passed as a parameter

				var bequiet = GET_EXTRA_ARG(1, EXTRA_ARGS...).(string)				// parameter at index 1 is quiet.. So bequiet is set to string "quiet"
				
				var antcount = GET_EXTRA_ARG(2, EXTRA_ARGS...).(int)				// parameter at index 2 is the number 55... so antcount is set to  specify 2 to get the INT value 55 that was specified

				var mydomain = GET_EXTRA_ARG(3, EXTRA_ARGS...).(string)			  // and.. specify 3 to get the param at index 3 (which is www.podshop.com)


				C.Println(" RESULT Verbose is: ", verbose)
				C.Println(" antcount is: ", antcount)

				C.Println(" mydomain is: ", mydomain)
			}
*/
func GET_EXTRA_ARG(key interface{}, EXTRA_ARGS ...interface{}) interface{} {

	var find_by_INDEX = false
	var find_by_STRING = false


	var IND_2_find = -9
	var KEY_VAL_2_find = ""
	

	var EMPTY_RES interface{}

	/*
	
		We'll ether have an INT (implying we GET the value of the parameter passed..by INDEX
		Or we will have an explicit string ("meaning if this value is found, return true or false")
		
		if by index, doesnt matter if it is bool, float, string or whatever.. its returned explicitly
	*/
	
	if IS_INT(key) {
		find_by_INDEX = true
		IND_2_find = key.(int)

	} else if IS_STRING(key) {
		find_by_STRING = true
		KEY_VAL_2_find = key.(string)
		
	} else {
		R.Println(" Invalid Key type! Must be either INT (for index) or string (for key val)")

		return EMPTY_RES
	}

	// Now iterate through the args!
	for ind, arg := range EXTRA_ARGS {

		// Now if an index was passed.. and we match it.. we return the value at that index verbatim
		if find_by_INDEX {
			if ind == IND_2_find {
				return arg
			}
		} 		

		// If a KEy was passed (a  string) we just return true or false
		// example.. if someone passes verbose, we return true... which will turn on verbose mode
		// even tho the arg value is returned, we dont use it cause its identical to KEY_VAL_2_find
		if find_by_STRING {
			PARAM_string, FOUND_string := arg.(string)
			if FOUND_string {
				if PARAM_string == KEY_VAL_2_find {
					return arg
				}
			}
		}
	}


	// default return value if we ever get this far
	return EMPTY_RES

} //end of




/* 
        Searches the passed map[string]interface{} for a string key  .. OR a value:
        NOTE: You MUST append .(string) || .(int) || .(float64) to the return val (in what calls SearchMap)

        EXAMPLE: 

        foods := map[string]interface{}{
            "bacon": "delicious",
            "eggs": "are a protien", 
            "checkien_cost": 1.75,
            "eatsteak": true,
          }


        // Searches by specific key
        found, val := SearchMap(foods, "eatsteak")

        if found {
            // Must cast the exact type of the value you expect
            if val.(bool) == true {
                C.Println(" We like to eat steak!")
            }
        }

        // Searches by specific value
        found, val := SearchMap(foods, "byval", 1.75)

        if found {
            if val.(float64) > 1.0 {
                Y.Println(" Found a float greater than 1.0 !!")
            }
        }



     */
	 func SearchMap (m map[string]interface{}, EXTRA_ARGS ...interface{}) (bool, interface{}) {
        
        var search_type = GET_EXTRA_ARG(0, EXTRA_ARGS...).(string)

        // If the first parameter is byval or val.. means we look for a specific value in the map
        if search_type == "val" || search_type == "value" || search_type == "byval" {

            var lookfor = GET_EXTRA_ARG(1,  EXTRA_ARGS...)
            
            for _, val := range m {

                if val == lookfor {
                    return true, val
                }
            }

        /// else the default is to search for the key.. which has to be a string
        } else {

            lookfor := search_type
            if IS_STRING(lookfor) {
                v, found := m[lookfor]

                // if found, we return this value... note.. in th line that calls search map.. you MUST append
                if found {
                    return true, v
                }
            }
        }
        
        var EMPTY interface{}

        return false, EMPTY
    } 

/*

	The below items are for EMBED support.. 
	be sure to specify the files you want to embed IN the go file
	(NOT in  a function)  .. below embeds the steam_plugin_isntall_helper into the variable steam_plugin_helper

	// Also be sure to do import like this:  (dont put it in MAIN)

	_ "embed"



	// Also define your imports like this:

////: //go:embed _EMBEDDED_FILES/steam_plugin_install_helper.sh
var steam_plugin_helper string				// For TEXT files

////: //go:embed _EMBEDDED_FILES/ca.pem.gz
var pem []byte								// For BINARY files


*/


type EMBED_OBJ struct {
	Name 	        string
	TYPE			string

	DATA  			[]byte
	TEXT			string
	// BIN 			[]byte
}

var EMBED_FILES []EMBED_OBJ

// This saves items in EMBED_FILES.. and always returns totall number of items in EMBED so far
func SAVE_EMBED(filename string, fileOBJ interface{} ) int {

	var T EMBED_OBJ
	T.Name = filename

	// Determine the type of fileOBJ (either string or byte)
	STRING_TEXT, FOUND_string := fileOBJ.(string)
	BYTE_DATA, FOUND_bytes := fileOBJ.([]byte)

	if FOUND_string {
		T.TYPE = "text"
		T.DATA = []byte(STRING_TEXT)


	} else if FOUND_bytes {

		T.TYPE = "binary"
		T.DATA = BYTE_DATA
	} else {

		R.Println(" Cant Save Embed for some reason! Weird or Invalid Data Type??")
		return 0
	}

	EMBED_FILES = append(EMBED_FILES, T)	

	return len(EMBED_FILES)
}





func WRITE_EMBEDDED(filename string, destpath string, EXTRA_ARGS ...interface{}) {

	var alt_filename = GET_EXTRA_ARG(0,  EXTRA_ARGS...).(string)
	var verbose_mode  = GET_EXTRA_ARG("verbose", EXTRA_ARGS...).(bool)

	for _, f := range EMBED_FILES {

		if f.Name == filename {

			var FULL_FILE_PATH = destpath + "/" + f.Name
		
			if alt_filename != "" {
				FULL_FILE_PATH = destpath + "/" + alt_filename
			}


			// Also safety remove file
			err1 := os.RemoveAll(FULL_FILE_PATH)

			// This will only fail if we dont have permission or file isnt there
			if err1 != nil {

				if verbose_mode {
					R.Println(" WRITE_EMBED Error: ", err1)
					return
				}
			}

			 // Text written as 644
			 if f.TYPE == "text" {
				if err := os.WriteFile(FULL_FILE_PATH, f.DATA, 0644); err != nil {
					if verbose_mode {
						R.Print("ERROR! Cant save the file: ")
						Y.Println(FULL_FILE_PATH)
						R.Println(err)
					}
				}
			 }


			 // Binarys are written as 755
			 if f.TYPE == "binary" {
				if err := os.WriteFile(FULL_FILE_PATH, f.DATA, 0755); err != nil {
					if verbose_mode {
						R.Print("ERROR! Cant save the file: ")
						Y.Println(FULL_FILE_PATH)
						R.Println(err)
					}
				}
			 }


			 /// Also check to see if this was a gz file.. if so, we extract it
			 if strings.Contains(f.Name, ".gz") {

				NOGZ_FILENAME := strings.Replace(FULL_FILE_PATH, ".gz", "", 1)
				os.RemoveAll(NOGZ_FILENAME)

				RUN_COMMAND("gunzip " + FULL_FILE_PATH )
			 }

			return
		}

	}


	// else if they specified somethign we didnt have, tell them

	R.Println(" ERRROR! Cant find the specified Embedded file!!")
}




func SHOW_EMBED() {

	C.Print("\n Showing EMBEDDED Files: ")
	G.Println(len(EMBED_FILES))
	C.Println("")

	for _, v := range EMBED_FILES {
		
		if v.TYPE == "text" {
			Y.Print("     ", v.TYPE, " - ")
		} else {
			Y.Print("   ", v.TYPE, " - ")
		}
		
		C.Println(v.Name)
	}

	Y.Print("\n Total Embedded Files: ")
	C.Println(len(EMBED_FILES))	

	C.Println("")
}
func SHOW_EMBEDDED() {
	SHOW_EMBED()
}


func makeRound(num float64) int {
	return int(num + math.Copysign(0.5, num))
}
// Returns file size in MEGS and BYTES
func GET_FILE_SIZE(bytes int64) (float64, float64) {
	megs := float64(bytes / 1024000)
	gigs := float64(megs / 1024)

	precision := 2

	output := math.Pow(10, float64(precision))
	
	megs = float64(makeRound(megs*output)) / output	
	gigs = float64(makeRound(gigs*output)) / output	

	return megs, gigs
}


// This initializes some pretty necessary timezone defaults. 
func SET_TIMEZONE_DEFAULTS() {

	//1. By Deffault this will be Local
	ZONE_UPPER = "Local"		

		   if USE_EST {  DEFAULT_ZONE_LOCATION_OBJ = EST_OBJ
	} else if USE_CST {  DEFAULT_ZONE_LOCATION_OBJ = CST_OBJ
	} else if USE_MST {  DEFAULT_ZONE_LOCATION_OBJ = MST_OBJ
	} else if USE_PST {  DEFAULT_ZONE_LOCATION_OBJ = PST_OBJ
	// if needed, add more here.. before we hit UTC
	} else if USE_UTC {	 DEFAULT_ZONE_LOCATION_OBJ = UTC_OBJ   }
 
	//2. Also lets get the Timezone and OFFSET info	
	t := time.Now().In(DEFAULT_ZONE_LOCATION_OBJ)
	curr_zone, offset := t.Zone()

	//2b. see if it is negative, we need to convert this temporarily
	hprefix := "+"
	
	if offset < 0 {
		fixed_off := math.Abs(float64(offset))

		//2b. Convert fixed_off to an integar
		entry_FIXED_text := strconv.FormatFloat(fixed_off, 'f', 2, 64)
		entry_NUM_text := strings.Replace(entry_FIXED_text, ".", "", -1)
		offset, _ = strconv.Atoi(entry_NUM_text)

		hprefix = "-"
	}


	//3. NOw convert the offset seconds to hours
	off_hours := (offset / 60) / 60
	offstring := hprefix + strconv.Itoa(off_hours) + " hours"

	//4. Set ZONE_UPPER .. we use this in other places. 
	ZONE_UPPER = curr_zone
	ZONE_LOCAL = t.Location().String()
	
	ZONE_HOUR_OFFSET = offstring
	ZONE_FULL = " (" + curr_zone + " " + offstring + ")"
	
	//4. Get the current time to set the global time vaiables
	GLOBAL_CURR_DATE, GLOBAL_DATE_OBJ = GET_CURRENT_TIME("full")
	GLOBAL_CURR_TIME = GLOBAL_CURR_DATE
	GLOBAL_TIME_OBJ = GLOBAL_DATE_OBJ


} //end of func


/*

	Gets the current Time/Date ...defaults to LOCAL
	However if you specify EST/PST/MST ...you get THAT time
*/
func GET_CURRENT_TIME(EXTRA_ARGS ...string) (string, time.Time) {

	//1. Default ot the local machines time zone
	dateOBJ := time.Now()
	dateOBJ = dateOBJ.In(DEFAULT_ZONE_LOCATION_OBJ)
	var output_FORMAT = "full"		// Full is the default format we return the time string in

	//2. Now, see if flags were specified. Iterate through them
	for _, VAL := range EXTRA_ARGS {

		VAL = strings.ToLower(VAL)

		switch VAL {
			case "est":
				dateOBJ = dateOBJ.In(EST_OBJ)

			case "cst":
				dateOBJ = dateOBJ.In(CST_OBJ)

			case "mst":
				dateOBJ = dateOBJ.In(MST_OBJ)
			case "mdt":
				dateOBJ = dateOBJ.In(MST_OBJ)							

			case "pst":
				dateOBJ = dateOBJ.In(PST_OBJ)

			case "utc":
				dateOBJ = dateOBJ.In(UTC_OBJ)
	
		} //end of switch

		//3. If full, short, british or iso specified, set the output format
		// MAKE SURE SHOW_PRETTY_DATE also picks these UP!! (supports them)
		if VAL == "full" || VAL == "british" || VAL == "iso" || VAL == "justdate" || VAL == "nano" || VAL == "timestamp" {
			output_FORMAT = VAL
		}

		

	} //end of for

	result, _ := SHOW_PRETTY_DATE(dateOBJ, output_FORMAT)

	return result, dateOBJ

} //end of func


func PRETTY_STRUCT_json(input interface{}) string {
	byte_json, _ := json.MarshalIndent(input, "", "  ")

	result := string(byte_json)

	return result
}
// RETURNS a STRUCt in JSON format.. a string..which you can PRINTLN
func SHOW_STRUCT(ALL_PARAMS ...interface{}) {
	
	var tmpSTRUCT interface{}

	var COLOR = "yellow"		// whill show struct in yellow

	// Collects the input params specified... supports INT and FLOAT dynamically
	for n, param := range ALL_PARAMS {
	
		// First paramn is always the STRUCT
		if n == 0 {
			tmpSTRUCT = param
			continue

		}

		// If second paramis sepectified.. This will be the color in whichw e display thie struct
		if n == 1 {
			COLOR = param.(string)
		}
	}	
	

	if COLOR == "red" {			R.Println(PRETTY_STRUCT_json(tmpSTRUCT))	
	} else if COLOR == "magenta" {		M.Println(PRETTY_STRUCT_json(tmpSTRUCT))	
	} else if COLOR == "green" {		G.Println(PRETTY_STRUCT_json(tmpSTRUCT))	
	} else if COLOR == "white" {		W.Println(PRETTY_STRUCT_json(tmpSTRUCT))	 
	} else if COLOR == "cyan" {		C.Println(PRETTY_STRUCT_json(tmpSTRUCT))
	// The default color
	} else {
		Y.Println(PRETTY_STRUCT_json(tmpSTRUCT))
	}
}


// Shows the amount of time a program ran (and start and end time)
func SHOW_START_and_END_TIME() {

	endTime, endOBJ := GET_CURRENT_TIME()
	difftemp := endOBJ.Sub(PROG_START_TIMEOBJ)
	TIME_DIFF := difftemp.String()
	


	Y.Println("\n\n ****************************************************** ")


	W.Print("              Start Time:")
	B.Println(" " + PROG_START_TIME)
	Y.Print("                End Time:")
	M.Println(" " + endTime)
	C.Print("      Total PROGRAM DURATION: ")
	G.Println(" ", TIME_DIFF)
	C.Println("******************************************************")
}

var SPINNER_SPEED = 100
var SPINNER_CHAR = 4
var spinOBJ = spinner.New(spinner.CharSets[14], 100*time.Millisecond)

//	Creates a cool "im busy right now" status spinner so you know the program is running 
func START_Spinner() {

	sduration := time.Duration(SPINNER_SPEED)

	spinOBJ = spinner.New(spinner.CharSets[SPINNER_CHAR], sduration*time.Millisecond)
	spinOBJ.Start()
}

func STOP_Spinner() {

	spinOBJ.Stop()
}

// this is a simple sleep function
func Sleep(seconds int, ALL_PARAMS ...bool) {

	var showOutput = false

	for x, BOOL_VAL := range ALL_PARAMS {

		//1. First Param is allthat is used
		if x == 0 {
			showOutput = BOOL_VAL
			continue
		}
	} // end of for	

	if showOutput == true {
		secText := ""
		suffix := "seconds"
		sectemp := seconds

		if seconds >= 119 {
			sectemp = seconds / 60
			suffix = "minutes"
		}
		secText = strconv.Itoa(sectemp)
		C.Println("        ** Sleeping for: " + secText + " ", suffix, "...")
	}

	duration := time.Duration(seconds) * time.Second
	time.Sleep(duration)

} //end of sleep function


var SHOW_WHAT_WAS_TYPED = false
func GET_USER_INPUT() string {
	reader := bufio.NewReader(os.Stdin)
	userTEMP, _ := reader.ReadString('\n')
	userTEMP = strings.TrimSuffix(userTEMP, "\n")

	if SHOW_WHAT_WAS_TYPED {
		Y.Print("\n     You Typed: ")
		W.Print(userTEMP)
		Y.Println("**")
	}

	return userTEMP

} //end of

func GET_INPUT() string {
	return GET_USER_INPUT()
}

// This takes IN a string and returns a shuffle of the characters contained in it
func SHUFFLE_STRING(input_STRING string) string {

	//1. Get the length of the string
	slen := len(input_STRING)

	stringRUNE := []rune(input_STRING)

	shuffledString_RESULT := make([]rune, slen)

	for i := range shuffledString_RESULT {
		shuffledString_RESULT[i] = stringRUNE[rand.Intn(slen)]
	}
	return string(shuffledString_RESULT)
} // end of genSESSION

func VERIFICATION_PROMPT(warning_TEXT string, required_input string) {

	M.Println("\n      - - - - - - - - WARNING - - - - - - - - - - - - - -")
	
	for x := 0; x < 3; x++ {
		C.Println("")
		C.Println("      ", warning_TEXT)
		C.Print("       Type: ")
		G.Print(required_input)
		C.Println(" To Continue")
		Y.Print("       RESPONSE: ")
		userResponse := GET_USER_INPUT()

		if strings.Contains(userResponse, required_input) {
			return
		} else {
			R.Println("\n ! ! ! ! ! ! INVALID RESPONSE  ! ! ! ! ! !")
		    M.Println("\n     - - - - - - - - - - - - - - - - - - - - - - - - -")			
		}
	} //end of for
	

	//2. If we get this far without a valid response, we will exit the program without proceeding
	os.Exit(ERROR_EXIT_CODE)


} //end of prompt

func PROMPT(warning_TEXT string, required_input string) {
	VERIFICATION_PROMPT(warning_TEXT, required_input)
}



var PAGE_COUNT = 0
var PAGE_MAX = 5

// This is a basic Paging routine that prompts you to PressAny key
// after x number of items have been shown
func Pager(tmax int) {
	PAGE_MAX = tmax
	PAGE_COUNT++

	if PAGE_COUNT == PAGE_MAX {
		C.Print("   - - PAGER - -")
		PressAny()
		PAGE_COUNT = 0
	}

} //end of Pager

// Simple PressAny Key function
func PressAny() {

	W.Println("")
	W.Println("         ...Press Enter to Continue...")
	W.Println("")

	//1. New way of doing PAK
	b := make([]byte, 10)
	if _, err := os.Stdin.Read(b); err != nil {
		R.Println("Fatal error in PressAny Key: ", err)
	}

} // end of func


// This gets the platform we are running on (mac, linux, windows)
func GET_CURRENT_OS_INFO() {

	if runtime.GOOS == "linux" {
		OSTYPE="Linux"

	//2. Otherwise see if this is MAC
	} else if runtime.GOOS == "darwin" {
		OSTYPE="MAC"

	//3. otherwise.. its windows.. it wins by default!!	
	} else if runtime.GOOS == "windows" {
		OSTYPE="Windows"

	//4. If we get this far, means we have some weird unrecognizable OS:
	} else {
		OSTYPE="- - UNKNOWN OS - -"
	}

	//5. Another courtesy Alias
	CURRENT_OS = OSTYPE
	GOOS_VALUE = runtime.GOOS

} //end of getOsType
func Show_TOTAL_PROG_RUNTIME() {

	endTime, endOBJ := GET_CURRENT_TIME()

	dtemp := endOBJ.Sub(PROG_START_TIMEOBJ)
	DIFF := dtemp.String()

	//12. DISPPLAY Status on this Threads Performance (and metrics)
	R.Println("")
	W.Print(" = = = = = = = = = = = = = ")
	G.Print(" Program run Complete! ")
	W.Println(" = = = = = = = = = = = = = ")

	
	C.Println("")
	W.Print("         STARTED:")
	B.Println(" " + PROG_START_TIME)
	Y.Print("        ENDED on:")
	M.Println(" " + endTime)
	C.Print("  Total DURATION: ")
	G.Println(DIFF)

	W.Println(" = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = ")

}

// Returns a randomly generated number within a given range (returns a STRING AND an int)
func GenRandomRange(min int, max int) (int, string) {

	resultNum := rand.Intn(max-min) + min
	resultText := strconv.Itoa(resultNum)


	// Always return a string with a 0 prefix
	if resultNum < 10 {
		resultText = "0" + resultText
	}


	return resultNum, resultText

} //end of genRandomRange


// Return a random delay when we hit too many CALLS

func GENERATE_RETRY_SLEEP_and_WAIT(retry_count int, MAX_SLEEP_VAL int, MAX_RETRY_ATTEMPTS int) {

	rnum, sec_text := GenRandomRange(5, MAX_SLEEP_VAL)
	retry_count = retry_count + 1
	parent_FUNC := GET_PARENT_FUNC(2)	// play with the integer you ass to get the proper calling function name

	retry_text := ShowNum(retry_count)
	max_text := ShowNum(MAX_RETRY_ATTEMPTS)

	SHOW_BOX_MESSAGE("|red|RETRY ERROR in: ", "|yellow|" + parent_FUNC, "Retry in " + sec_text + " seconds..", "|yellow|" + retry_text + " of " + max_text )

	Sleep(rnum, true)

}


// This gets the name of the PARENT calling function
// pass dnum of 1 to it for starters.. if you arent getting the right caller name.. increment by 1
func GET_PARENT_FUNC(dnum int) string {

	var result = ""
	
	pc, _, _, ok := runtime.Caller(dnum)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		result = details.Name()
	}	

	return result
} //end of




// These are the characters used to generate the serial
//var sessTEMPLATE = []rune("GRZBJHUFLEKVXMNTQPSOADWYC527183469")

// This generates a serial.. usually used discern between multiple execution runs like in jenkins
func GenSerial(serial_length int) {

	result := SHUFFLE_STRING("grzbjhuflcekivxmntqpsoadwy527183469")

	part_ONE := result[0:4]
	part_TWO := result[3:serial_length]

	SERIAL_NUM = part_ONE + "-" + part_TWO

} // end of GenSerial



var ENABLE_DEFAULT_PARAMS = false

func SETUP_DEFAULT_COMMAND_LINE_PARAMS() {

	// This is too useful NOT to have everywhere
	
	flag.BoolVar(&DEBUG_MODE,       "debug", DEBUG_MODE,         "  If specified we run in DEBUG MODE. Code that checks for this run when this is set")

	//1. First lets get some default command line params we always use
	if ENABLE_DEFAULT_PARAMS {
		//2. These are variables used for modifying the TIMEZONE (if ever even needed)
		flag.BoolVar(&USE_EST, "est", USE_EST,         "  Force Timezone to be EST")
		flag.BoolVar(&USE_CST, "cst", USE_CST,         "  Force Timezone to be CST")
		flag.BoolVar(&USE_CST, "mst", USE_MST,         "  Force TZ to be MST/Mountain")
		flag.BoolVar(&USE_CST, "mdt", USE_MST,         "  Force TZ to be MST/Mountain")
		flag.BoolVar(&USE_PST, "pst", USE_PST,         "  Force Timezone to be PST")
		flag.BoolVar(&USE_UTC, "utc", USE_UTC,         "  Force Timezone to be UTC")

		//3. In case we want to show serial numbers the program automatically generates program runs	
		flag.BoolVar(&SHOW_SERIAL, "serial", SHOW_SERIAL, "  If specified we Show a RUNTIME serial number")
	}
	
	//4. And finally, Very important.. This is the final flag.Parse that is run in the program
	flag.Parse()

} //end of setup default command line params



// Much better RUN_COMMAND in 2022
func RUN_COMMAND(command_text string, ALL_PARAMS ...string) string {

	command := strings.Split(command_text, " ")


	rundir := ""

	var verbose = false

	//2. Now, see if version was passed
	for x, VAL := range ALL_PARAMS {

		// First param specifid is always the dest rundir
		if x == 0 {
			rundir = VAL
			continue

		}

		if x == 1 {
			if VAL == "verbose" {
				verbose = true
			}
		}

	} //end of for
	
	

	if len(command) < 2 {
		R.Println(" ERROR: Command specified is TOO SHORT")
	}

	cmd := exec.Command(command[0], command[1:]...)

	if rundir != "" {
		cmd.Dir = rundir
	}
	

	output, err := cmd.CombinedOutput()
	if err != nil {

		if verbose {
			R.Println(" Error in RUN COMMAND: ", err)
		}
	}
	
	return string(output)
}





/* 
  MUST ALWAYS CALL THIS in the MAIN of every program.. 
  This is how command line params get initted
   Also make sure it is the LAST 'init' type function called (for example.. BEFORE AWS_INIT
*/

func MASTER_INIT(PROGNAME string, ALL_PARAMS ...string) {
	var verbose = false
	var show_exec_serial = false



	var use_default_params = false

	var VERSION_NUM = ""

	//2. Now, see if version was passed
	for _, VAL := range ALL_PARAMS {

		if VAL == "verbose" {
			verbose = true
			continue
		}

		if VAL == "serial" {
			show_exec_serial = true
			continue
		}

		if VAL == "defaultparams" {
			use_default_params = true
			continue
		}		
		
		// Otherwise, assume the first param is the VERSION_NUM
		VERSION_NUM = VAL
		break



	} //end of for


	//1. Setup default COmmand line params
	if use_default_params {
		SETUP_DEFAULT_COMMAND_LINE_PARAMS()
	}

	//2. Load defaults we need for proper Timezone and DateMath operations
	SET_TIMEZONE_DEFAULTS()

	//3. And, Always init the random number seeder
	rand.Seed(time.Now().UTC().UnixNano())	

	//3b. And get current OS Data
	GET_CURRENT_OS_INFO()

	//4. Setup the prog start time globals
	PROG_START_TIME, PROG_START_TIMEOBJ = GET_CURRENT_TIME()


	if verbose {
	

		if VERSION_NUM != "" {
			SHOW_BOX_MESSAGE(PROGNAME, "|cyan|Ver:", "|green|" + VERSION_NUM,  "|cyan|Running On: " + CURRENT_OS + "," + GOOS_VALUE)	 
		} else {
			SHOW_BOX_MESSAGE(PROGNAME, "|cyan|Running On: " + CURRENT_OS + "," + GOOS_VALUE)
		}

		W.Print("      [ Timezone: ")
		if ZONE_LOCAL == "Local" {
			C.Print(ZONE_LOCAL)
		}
		Y.Print(" " + ZONE_UPPER)
		G.Print(" (", ZONE_HOUR_OFFSET, ") ")
		W.Println(" ]")		

	} else {



		if VERSION_NUM != "" {
			SHOW_BOX_MESSAGE(PROGNAME, "|cyan|Ver:", "|green|" + VERSION_NUM)	

		} else {
			SHOW_BOX_MESSAGE(PROGNAME)
		}


	}



	//7. If specified we show the unique execution serial. This is useful when running from within Jenkins
	if show_exec_serial {
		GenSerial(10)
		SHOW_BOX_MESSAGE("Generated EXEC Serial: ", "|cyan|" + SERIAL_NUM)
	}


	if DEBUG_MODE {
		SHOW_BOX_MESSAGE("|yellow|RUNNNING in DEBUG MODE")
	}
	

	W.Println("")

} //end of func


/* 
   Kept here as filler /template /example / Reference
.. anything you put in this function will run  when the module is imported

*/
func init() {
	
	// Add stuff here

} // end of main
