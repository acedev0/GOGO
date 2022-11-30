/*   GOGO_Math / Date / Conversion Gadget - Useful math and Calucation code to make Go Dev Easier

---------------------------------------------------------------------------------------
NOTE: For Functions or Variables to be globally availble. The MUST start with a capital letter.
	  (This is a GO Thing)

	
	Aug 28, 2021    v1.23   - Initial Rollout

*/

package CUSTOM_GO_MODULE


import (

	// = = = = = Native Libraries

		"time"
		"strconv"
		"strings"
		"math"
		//"math/rand"

	// = = = = = CUSTOM Libraries

	. "github.com/acedev0/GOGO/Gadgets"
	. "github.com/acedev0/GOGO/Gadgets/StringOPS"



	// = = = = = 3rd Party Libraries

)



func MakeRound(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

// identical to FIX_FLOAT
func FORCE_FLOAT_PRECISION(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(MakeRound(num*output)) / output
}
func FIX_FLOAT_PRECISION(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(MakeRound(num*output)) / output
}

/*
	PERCENTAGE MAGIK
*/
/*
 FIXED GET_PERCENT as of 7/2022
 SMALLIST number goes first.. then BIGGEST number
 // specify a 3rd parameter of 2 or 3...to get a percent returned with THAT much precision
*/
func GET_PERCENT(ALL_PARAMS ...interface{}) (float64, string) {

	var firstNUM = 0.0
	var secNUM = 0.0

	var PRECISION = 1		// Default preceission set to return a percent with 1 dec points... specified

	var FORCE_RETURN_POSITIVE = false
	// Collects the input params specified... supports INT and FLOAT dynamically
	for n, param := range ALL_PARAMS {
		// First paramn is always FIRSTNUM
		if n == 0 {

			if IS_INT(param) { 
				firstNUM = float64(param.(int)) 
			} else {
				firstNUM = param.(float64)
			}
			continue
		}
		
		if n == 1 {
			if IS_INT(param) { 
				secNUM = float64(param.(int)) 
			} else {
				secNUM = param.(float64)
			}		
			continue	
		}
		
		// IF a 3 param is specified, this implies the PRECISION of the percentage we return
		if n == 2 && IS_INT(param) {
			PRECISION = param.(int)
		}

		// If a bool is specified, this controls ensuring we return a positive number
		if IS_BOOL(param) {
			FORCE_RETURN_POSITIVE = true
		}

	} //end of for

	// This is same as the TradingView percent function in Terry_CCOMMON_LIB
	var smallNUM = firstNUM
	var largeNUM = secNUM
	
	if smallNUM == 0.0 || largeNUM == 0.0 {
		return 0.0, "0.0%"
	}

	// if numbers are equal
	if largeNUM == smallNUM {
		return 100.0, "100.0%"
	}
	
	if FORCE_RETURN_POSITIVE {
		if secNUM > firstNUM {
			largeNUM = secNUM
			smallNUM = firstNUM
		}
	}
	// SMALLIST number goes first.. then BIGGEST number
	calc_VAL := (100.00 * smallNUM) / largeNUM

	percSTRING := strconv.FormatFloat(calc_VAL, 'f', PRECISION, 64)		// Make a string of calc_VAL with specific # of PRECISION dec points
	percNUM, _ := strconv.ParseFloat(percSTRING, 64) // this NUMERICALLY reformats the percentage to have just the # decimal points that PRECISIONS specifies
	
	// make pretty string	
	percSTRING = percSTRING + "%"

	return percNUM, percSTRING



}

/*
	This takes in a number (float) and shows it as pretty PERCENT string
	// Specify another INT to define PRECISION
*/
func SHOW_PRETTY_PERCENT(ALL_PARAMS ...interface{}) string {

	var inputNUM = 0.0
	var PRECISION = 1

	for n, param := range ALL_PARAMS {
		// First paramn is always FIRSTNUM
		if n == 0 {

			if IS_INT(param) { 
				inputNUM = float64(param.(int)) 

			} else {
				inputNUM = param.(float64)
			}
			continue
		}
		
		if n == 1 {
			if IS_INT(param) { 
				PRECISION = param.(int)
			}		
		}
	} //end of for	

	percSTRING := strconv.FormatFloat(inputNUM, 'f', PRECISION, 64)		// Make a string of num with specific # of PRECISION dec points
	percSTRING = percSTRING + "%"

	return percSTRING
}



// Alias for Get_PERCENT
func GET_PERCENTAGE(ALL_PARAMS ...interface{}) (float64, string) {
	return GET_PERCENT(ALL_PARAMS...)
}
/*
	END OF PERCENTAGE MAGIK
*/




// Easy way to get the UTC form of DATE_OBJ so there is no confusion.. Returns Time, String(pretty date) and Weekday all converted from the orig time
func GET_DB_DATE_UTC(input_DATE_OBJ time.Time) (time.Time, string, string, string) {

	result_DATE_OBJ := input_DATE_OBJ.In(UTC_OBJ)

	pretty, weekday := SHOW_PRETTY_DATE(result_DATE_OBJ)

	result_as_STRING := result_DATE_OBJ.String()

	return result_DATE_OBJ, result_as_STRING, pretty, weekday

}


func CALC_RATIO(smallNUM float64, bigNUM float64, FORCE_DETECT_LARGEST_NUM ...bool) float64 {

	divid := smallNUM / bigNUM
	if FORCE_DETECT_LARGEST_NUM[0] == true {
		if smallNUM > bigNUM {
			divid = bigNUM / smallNUM
		}
	}

	perc := divid * 100.0
	fixed_PERC := FORCE_FLOAT_PRECISION(perc, 2)

	return fixed_PERC
} //end of func
	
func GET_RATIO(smallNUM float64, bigNUM float64, FORCE_DETECT_LARGEST_NUM ...bool) float64 {
	return CALC_RATIO(smallNUM, bigNUM, FORCE_DETECT_LARGEST_NUM...)
}

// Gets the difference between to integer values
func GET_DIFF(num_A int, num_B int) int {

	result := num_A - num_B

	if num_A == num_B {
		return 0
	
	} else if num_B > num_A {		

		result = num_B - num_A
	}

	return result

}


// Easy way to find if FIRST date is AFTER the PREV DATE
func DATE_IS_AFTER(first, prev time.Time) bool {
    return first.After(prev)
}

// Correspondingly a date that is BEFORE
func DATE_IS_BEFORE(first, prev time.Time) bool {
    return first.Before(prev)
}



// Takes in Two Time periods.. and returns the duration in DAYS, Hours and Minutes (and comprable strings)
// Returns MINS, HOURS, DAYS (in float first, then strings)
func GET_DURATION(startTIME time.Time, endTIME time.Time, EXTRA_ARGS ...string) (float64, string, string) {
	var precision = 1

	var interval = ""

	//1. First parameter is always the interval. We use this to "force" the value returned
	for n, VAL := range EXTRA_ARGS {

		//1b. If short or full was passed, we format the output date that way
		if n == 0 && VAL != "" {
			interval = VAL
			continue
		}		

		if n == 1 && VAL != "" {			
			precision, _ = strconv.Atoi(VAL)
			continue
		}				
	} //end of for

	temp_secs := endTIME.Sub(startTIME).Seconds()
	temp_mins := endTIME.Sub(startTIME).Minutes()
	temp_hours := endTIME.Sub(startTIME).Hours()	
	temp_Days := temp_hours / 24

	DIFF_SECS := FORCE_FLOAT_PRECISION(temp_secs, 1)
	DIFF_MINS := FORCE_FLOAT_PRECISION(temp_mins, 1)
	DIFF_Hours := FORCE_FLOAT_PRECISION(temp_hours, 1)
	DIFF_Days := FORCE_FLOAT_PRECISION(temp_Days, 1)
	
	// TEXT versions:
	sec_text := strconv.FormatFloat(DIFF_SECS, 'f', precision, 64)
	min_text := strconv.FormatFloat(DIFF_MINS, 'f', precision, 64)
	hour_text := strconv.FormatFloat(DIFF_Hours, 'f', precision, 64)
	day_text := strconv.FormatFloat(DIFF_Days, 'f', precision, 64)

	var num_val = 0.0	
	var text_value = ""
	var pretty = ""

	if interval == "hour" || interval == "hours" {
		num_val = DIFF_Hours
		text_value = hour_text
		pretty = hour_text + " Hours"

	} else if interval == "min" || interval == "mins" {
		num_val = DIFF_MINS
		text_value = min_text
		pretty = min_text + " Mins"
	

	} else if interval == "day" || interval == "days" {
		num_val = DIFF_Days
		text_value = day_text
		pretty = day_text + " Days"

	
	} else if interval == "sec" || interval == "secs" {
		num_val = DIFF_SECS
		text_value = sec_text
		pretty = sec_text + " Seconds"

	}
	

	return num_val, text_value, pretty
}



// Gets the difference between two dates (by days, hour or minutes)
func GET_DATE_DIFF(mtype string, currDATE time.Time, prevDATE time.Time) int {

	if strings.Contains(mtype, "day") {

		days := currDATE.Sub(prevDATE).Hours() / 24	
		return int(days)
	
	} else if strings.Contains(mtype, "hour") {

		delta := currDATE.Sub(prevDATE)
		result := int(delta.Hours())

		return result
		
	} else if strings.Contains(mtype, "min") {
		delta := currDATE.Sub(prevDATE)
		result := int(delta.Minutes())

		return result
	}

	return 0
}


// This takes in a string and converts it to a float: input must be XX.XXXXX (no chars)
func CONVERT_FLOAT(input string, precision int) (float64, string) {

	f_NUM, _ := strconv.ParseFloat(input, 64)
	NUM_result := FORCE_FLOAT_PRECISION(f_NUM, precision)
	FIXED_text := strconv.FormatFloat(NUM_result, 'f', precision, 64)
	
	return NUM_result, FIXED_text

} //end of func


// This converts a float to a WHOLE number
func CONVERT_FLOAT_TO_WHOLE(infloat float64, CV_PRECISION int) (int, string) {
	entry_FIXED_text := strconv.FormatFloat(infloat, 'f', CV_PRECISION, 64)
	entry_NUM_text := strings.Replace(entry_FIXED_text, ".", "", -1)
	entry_NUM, _ := strconv.Atoi(entry_NUM_text)

	return entry_NUM, entry_NUM_text
}

// Makes a floating point number rounded up and returns integer
func MakeRound(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func FORCE_FLOAT_PRECISION(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(MakeRound(num*output)) / output
}





// Tells you if an INT is even 
func IS_EVEN(input_NUM int) bool {

	if input_NUM%2 == 0 {
		return true

	}

	return false
}
// Tells you if an INT is ODD
func IS_ODD(input_NUM int) bool {

	if input_NUM%2 == 0 {

	} else {

		return true
	}

	return false
}





// Returns Percentages INCREASE DECREASE for stocks etc... Takes in floats
func GET_INC_DEC_PERCENT(ALL_PARAMS ...interface{}) (string, float64) {
	
	var firstNUM = 0.0
	var secNUM = 0.0


	var SHOW_OUTPUT = false

	// If true is passed, we show the output of this function 
	for n, param := range ALL_PARAMS {
		// First paramn is always FIRSTNUM
		if n == 0 {

			if IS_INT(param) { 
				firstNUM = float64(param.(int)) 
			} else {
				firstNUM = param.(float64)
			}			
		}
		
		if n == 1 {
			if IS_INT(param) { 
				secNUM = float64(param.(int)) 
			} else {
				secNUM = param.(float64)
			}			
		}

		if n == 2 {
			SHOW_OUTPUT = true
			continue
		}
	}

	//1. First get the diff... if they are equal, we return
	if firstNUM == secNUM {

		return "0.0% No Change", 0.0
	}

	//2. Also if one number is 0.0 .. we return.. obviously this is 100%
	var bugfix_num = 100.0
	var bugfix_TEXT = "100.0"

	if firstNUM == 0.0 {
		if secNUM > 0.0 {
			return bugfix_TEXT + "% INCREASE ( from ZERO ) ", bugfix_num
		}
		if secNUM < 0.0 {
			return bugfix_TEXT + "% DECREASE ( TO ZERO ) ", bugfix_num
		}

	}
	if secNUM == 0.0 {
		if firstNUM > 0.0 {
			return bugfix_TEXT + "% DECREASE ( TO ZERO ) ", bugfix_num
		}
		if firstNUM < 0.0 {
			return bugfix_TEXT + "% INCREASE ( from ZERO ) ", bugfix_num
		}
	}	

	// This is for convenience and makes it easier to remember what number is what
	smallNUM := firstNUM
	largeNUM := secNUM


	//2b BUG FIX and hanlding for dealing with NEGATIVE numbers

	var BUGFIX_mode = "" 
	
	//c. first scenario is two negative numbers.. convert them to positive
	if smallNUM < 0.0 && largeNUM < 0.0 {

		smallNUM = math.Abs(smallNUM) 
		largeNUM = math.Abs(largeNUM) 

		
		if largeNUM > smallNUM {

			BUGFIX_mode = "decrease"

		} else if largeNUM < smallNUM {
			BUGFIX_mode = "increase"
		}
		

	//d. If smallNUM is negative and largenum is NOT
	} else if smallNUM < 0.0 {

		my_initial_INVESTMENT := math.Abs(smallNUM)
		
		smallNUM = my_initial_INVESTMENT
		largeNUM = largeNUM + smallNUM
		//largeNUM = largeNUM + (my_initial_INVESTMENT * 2)

	//e. if largeNUM is negative (assuming smallNUM is not)
	} else if largeNUM < 0.0 {

		largeNUM = math.Abs(largeNUM)
		my_initial_INVESTMENT := math.Abs(smallNUM)

		largeNUM = largeNUM + (my_initial_INVESTMENT * 2 )

		BUGFIX_mode = "decrease"

	}	

	//3. Now that we have the numbers in the correct place, lets do the math
	// CORRECT PERCENTAGE CALCULATION:
	// Percentage Increase = [ (Final Value - Starting Value) / |Starting Value| ] × 100
	mperc := ((largeNUM - smallNUM) / smallNUM) * 100

	//4. Determine if this is increase or decrease (based on if we have a negative number returned)
	mode := "increase"
	
	if mperc < 0 {
		mperc = math.Abs(mperc) 
		mode = "decrease"
	}

	if BUGFIX_mode != "" {
		mode = BUGFIX_mode
	}

	//5. Convert the percentages into readable objects
	percSTRING := strconv.FormatFloat(mperc, 'f', 2, 64)
	fixed_percNUM, _ := strconv.ParseFloat(percSTRING, 64) // this reformats the percentage to have just 2 decimals
	

	// Now, lets adjust percSTRING so we have a leading % char
	// (must be done AFTER the call to parseFLoat)
	percSTRING = percSTRING + "%"                           

	var details = ""
	
	if mode == "increase" {
		details = "INCREASED by " + percSTRING

	} else if mode == "decrease" {
		details = "DECREASED by " + percSTRING
	}

	firstSTRING := strconv.FormatFloat(firstNUM, 'f', 2, 64)
	secSTRING := strconv.FormatFloat(secNUM, 'f', 2, 64)

	mode = strings.ToUpper(mode)
	if SHOW_OUTPUT {
		SHOW_BOX_MESSAGE("GET_PERCENTAGE", "|cyan|" + firstSTRING + " --> " + secSTRING, "|yellow|" + mode + " by", "|green|" + percSTRING)
	}

	//5. return all the magic!!
	return details, fixed_percNUM
}




/* This allows you to add/subtract a percentage ie 1.05 TO a specified value */
func PERCENT_MATH(input_FLOAT float64, command string, perc float64) float64 {
	/*

		  Percentage calculation formula is:
			- To ADD a percantage to a value:

				 1530.56 * ( (100 + 0.65) / 100 )

			- And to SUBTRACT:

				1530.56 * ( (100 + 0.65) / 100 )
	*/

	if command == "add" || command == "ADD" {

		result := input_FLOAT * ((100 + perc) / 100)
		return result
	}

	if command == "sub" || command == "SUB" || command == "subtract" {

		result := input_FLOAT * ((100 - perc) / 100)
		return result
	}

	return 0.0
}






/* UPDATED: Converts TEXT date strings passed in any of the following formats:
	
	- MM-DD-YYYY
	- YYYY-MM-DD		(ISO / British format)
	
	- MM/DD/YYYY
	- YYYY/MM/DD

	Also accepts TIME.. Which must be apppended as:

	- XXXXX_18:05
	- XXXXX@18:05   

	  ...or with space

	- XXXXX 14:30	

	NOTE: AM/PM is also detected.. If so, the HOUR is automatically converted to UTC (+12 hours).. if it is PM
	

   and then returns a normalized  M/D/Y format...as well as the weekday... and a time.Time DateOBJ
   If you pass "short" or "full" you will receive a date formatted that way (uses SHOW_PRETTY_DATE)
*/
func CONVERT_DATE(inputDate string, EXTRA_ARGS ...string) (string, string, time.Time) {

	var sMon, sDay, sYear, sHour, sMin string
	var num_Mon, num_Day, num_Year int
	
	var num_Hour = 0		// so we default to 00:00, midnight if no time is passed
	var num_Min  = 0

	var output_FORMAT = ""

	var TZ_to_use = ""

	var TIMEZONE_OBJ = DEFAULT_ZONE_LOCATION_OBJ

	//1. First parameter is always the input date in the proper format
	for n, VAL := range EXTRA_ARGS {

		//1b. If short or full was passed, we format the output date that way
		if n == 0 && VAL != "" {
			output_FORMAT = VAL
			continue
		}

		//1c. If we wawnt to FORCE the TimeZone to be used, 
		if n == 1 {
			TZ_to_use = VAL
			continue
		}		

	} //end of for	


	//2. Time Zone logic
	if TZ_to_use != "" {
		switch TZ_to_use {
			case "est":
				TIMEZONE_OBJ = EST_OBJ
				break

			case "cst":
				TIMEZONE_OBJ = CST_OBJ
				break

			case "mdt":
				TIMEZONE_OBJ = MST_OBJ
				break

			case "mst":
				TIMEZONE_OBJ = MST_OBJ
				break				

			case "pst":
				TIMEZONE_OBJ = PST_OBJ
				break

			case "utc":
				TIMEZONE_OBJ = UTC_OBJ
				break				
		}
	}





	//2. Next run UBER split.. which splits on a bunch of Delimiters

	inputDate = strings.TrimSpace(inputDate)
	sd := UBER_Split(inputDate)

	/* 3. Now, Determine which format the data is in. 
		(m/d/y    or british, yyyy/mm/dd)

	*/
	//b. for mm/dd/yyyy
	if len(sd[0]) == 2 {

		sMon = strings.TrimSpace(sd[0])
		sDay = strings.TrimSpace(sd[1])
		sYear = strings.TrimSpace(sd[2])

	//c. for yyyy/mm/dd
	} else if len(sd[0]) == 4 {

		sYear = strings.TrimSpace(sd[0])
		sMon = strings.TrimSpace(sd[1])
		sDay = strings.TrimSpace(sd[2])		


	// Else if it is the SHORT format SHOW_PRETTY kicks out: Wednesday, 09/09/2009
	} else if strings.Contains(inputDate, ", ") {

		temps := strings.Split(inputDate, " ")
		t_date := temps[1]

		dates := strings.Split(t_date, "/")

		sMon = strings.TrimSpace(dates[0])
		sDay = strings.TrimSpace(dates[1])
		sYear = strings.TrimSpace(dates[2])



	}

	//4. Convert to numerics (we need this for the dateOBJ)
	num_Mon, _ = strconv.Atoi(sMon)
	num_Day, _ = strconv.Atoi(sDay)	
	num_Year, _ = strconv.Atoi(sYear)


	//5. Now determine if we have a TIME appended
	if strings.Contains(inputDate, ":") {

		sHour = sd[3]
		sMin = strings.ToUpper(sd[4])

		//b. Lets see if AM/PM is here

		if strings.Contains(sMin, "PM") {
			hourNUM, _ := strconv.Atoi(sHour)
			hourNUM = hourNUM + 12
			sHour = strconv.Itoa(hourNUM)
			sMin = TrimSuffix(sMin, "PM")

		//c. If AM is detected, just trim it off
		}  else if strings.Contains(sMin, "AM") {

			sMin = TrimSuffix(sMin, "AM")

		}

		//d. Get numberic versions for hour and min

		num_Hour, _ = strconv.Atoi(sHour)
		num_Min, _ = strconv.Atoi(sMin)

	}


	//6. Now Create the dateOBJ,  Format is: Y, MonOBJ, D, H, M, S, Nano
	monthObj := time.Month(num_Mon)	
	dOBJ := time.Date(num_Year, monthObj, num_Day, num_Hour, num_Min, 0, 0, TIMEZONE_OBJ)

	//12. Now pass to show_Pretty_Date with the output format if specified
	OUTPUT, weekday := SHOW_PRETTY_DATE(dOBJ, output_FORMAT)

	return OUTPUT, weekday, dOBJ
} //end of func


// Another Alias for CONVERT_DATE
func CONVERT_TIME(inputDate string) (string, string, time.Time) {
	return CONVERT_DATE(inputDate)
}




/* Takes in two date objects and returns the TIME DIFFERNCE between them in the 5m40s format
 */
 func DISPLAY_TIME_DIFF(startTime time.Time, endTime time.Time) string {
	
	diff := endTime.Sub(startTime)
	return diff.String()
}

// Alias for DISPLAY_TIME_DIFF (which lives in GO_GO_Gadgets)
func GET_TIME_DIFF(startTime time.Time, endTime time.Time) string {
	return DISPLAY_TIME_DIFF(startTime, endTime)
}




func DATE_MATH(dateObj time.Time, operation string, v_amount int, interval string) (string, time.Time) {
	return DateMath(dateObj, operation, v_amount, interval)
}
/* Takes in a date object and adds or subtracts
based on the number and whatever operation you specify
returns a date object
*/
func DateMath(dateObj time.Time, operation string, v_amount int, interval string) (string, time.Time) {

	//dateObj = dateObj.UTC()


	//1. If we are subtracting, we change amount to a negative number/// otherwise we default to adding
	if operation == "sub" || operation == "subtract" {

		v_amount = -v_amount

	}

	//2. Now we do the add or subtract operattion based on the time.Duration that is interval
	// Default is minute

	timeINT := time.Minute

	if interval == "hour" || interval == "hours" {

		timeINT = time.Hour

	} else if interval == "min" || interval == "mins" || interval == "minute" || interval == "minutes" {

		timeINT = time.Minute

	} else if interval == "sec" || interval == "secs" || interval == "second" || interval == "seconds" {

		timeINT = time.Second

	} else if interval == "day" || interval == "days" {

		timeINT = (time.Hour * 24)

	}


	//3. Finally do the "date math" on the incoming dateObj
	result_DATE_OBJ := dateObj.Add(time.Duration(v_amount) * timeINT)
	prettyDATE, _ := SHOW_PRETTY_DATE(result_DATE_OBJ)

	return prettyDATE, result_DATE_OBJ

} //end of dateMath



// Kept here as filler/example.. anything you put in this function will start when the module is imported
func init() {

	//1. Startup Stuff (init the command line params etc) . We need these Time ZONE Objects




} // end of main
