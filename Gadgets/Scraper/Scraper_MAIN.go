/*   GOGO_Gadgets  - Useful multi-purpose GO functions to make GO DEV easier

---------------------------------------------------------------------------------------
NOTE: For Functions or Variables to be globally availble. The MUST start with a capital letter.
	  (This is a GO Thing)

	Feb 24, 2021    v2.0    - High time this become a real tool library. Refactored this from being Terry_COMMON
	Feb 22, 2021	v1.90	- Added SCRAPE_TOOL for screen scraping
	Feb 15, 2020	v1.81	- Major Revamp to the GET_CURRENT_TIME and also have --zone TIME_ZONE_FLAG variable available to force Timezone

	Feb 14, 2020	v1.79	- Removed a lot of redundant functions (for date in particular)
							- Added an awesome ADD_LEADING_ZERO  function
							- Remove some more stuff i dont need
							- Working on the SHOW_PRETTY_DATE

	Feb 13, 2020	v1.76	- Removed UNEEDED stuff
	Jan 26, 2020	v1.73	- Got rid of more redundant functions that got pushed to APIcebaby
	Jan 23, 2020	v1.67	- Removed redundant fucntions (stuff that is now in APIce)
	Jan 05, 2020	v1.66	- Some cosmetic changes, Updated TerryCOMMON again
	Dec 29, 2019	v1.63	- Updated TerryCOMMON again
	Jun 05, 2014    v1.23   - Initial Rollout

*/

package CUSTOM_GO_MODULE

import (

	// = = = NATIVE Libraries / Standard Library Deps
		"net/http" // Needed for the functions that send JSON back and forth
		"strings"

	//2. = = = CUSTOM Libraries
		. "github.com/acedev0/GOGO/Gadgets"

	//3. = = = 3rd Party Libraries	
		"github.com/PuerkitoBio/goquery"

)

/*
	- - - -
	- - - -
	- - - - START OF GLOBALS WE NEED - - - - - -
	- - - -
	- - - -
*/

// MULTI-PURPOSE SCREEN SCRAPE TOOL
// Params: URL, UserAgent
// Returns: bool, GOQUERY_DOC, Text of Response
// var DEFAULT_USER_AGENT = "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36"
var DEFAULT_USER_AGENT = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36"


var S_RETRY_MAX = 10
var S_RETRY_SLEEP = 15

func SCRAPE_TOOL(URL string, EXTRA_ARGS ...string) (bool, *goquery.Document, string) {

	// Just blank Goquery Doc
	var EMPTY_GOQUERY_doc *goquery.Document	
	
	var IGNORE_REDIRECTS = false
	var VERBOSE = false

	//1. Get tvars passed
	for _, VAL := range EXTRA_ARGS {

		if strings.Contains(VAL, "redir")  {
			IGNORE_REDIRECTS = true
			continue
		}

		if strings.Contains(VAL, "verbose") {
			VERBOSE = true
			continue
		}

	} //end of ARGS


	if VERBOSE {

		W.Println("")
		C.Println(" - - - - - - - - - - - - - - - - - - - - ")
		C.Println(" *** Calling SCRAPE_TOOL")
		C.Print(" *** URL: ")
		Y.Println(URL)
		C.Println(" - - - - - - - - - - - - - - - - - - - - ")
		W.Println("")

	}


	




	//3. Attempt to connect using Max Retry
	var SUCCESS_FLAG = false
	var doc *goquery.Document
	var err3 error


	var DO_HARD_EXIT = false
	for i := 1; i < S_RETRY_MAX; i++ {

		SUCCESS_FLAG = true

		//2. Now generate a NewRequest Object with http
		client := &http.Client{}
		
		if IGNORE_REDIRECTS {
			
			client = &http.Client{
				CheckRedirect: func(req *http.Request, via []*http.Request) error {
					DO_HARD_EXIT = true
					R.Println(" Error! Have Redirect: ", http.ErrUseLastResponse)
					return http.ErrUseLastResponse
				},
			}

			if DO_HARD_EXIT {				
				R.Println(" Error AGAIN! Redirect detected: ", http.ErrUseLastResponse)
				return false, EMPTY_GOQUERY_doc, ""
			}

		}

		req, err := http.NewRequest("GET", URL, nil)
		if err != nil {
			R.Println(" *** ")
			R.Println(" *** ERROR IN SCRAPE_TOOL - During http OBJECT Create: ")
			Y.Println(err)
			R.Println(" *** ")
			R.Println("")	
			return false, EMPTY_GOQUERY_doc, ""
		}

		//3. Next, Set the User Agent the client will use during the HTTP Pull
		req.Header.Set("User-Agent", DEFAULT_USER_AGENT)

		// Try setting the Connection Close header if you need it.. This forces HTTP protocol to "close quick" as it is short lived
		//req.Header.Set("Connection", "close")

		//3b. Now.. actually do the Http Client Request (with the header)
		res, err2 := client.Do(req)
		

		if err2 != nil {
			R.Println(" *** ")
			R.Println(" *** ERROR IN SCRAPE_TOOL - During CLIENT HTTP Pull: ")
			Y.Println(err2)
			R.Println(" *** ")
			R.Println("")
			return false, EMPTY_GOQUERY_doc, ""
		}

		//5. Now finally, lets create our DOM object using goquery and empty the reader into the DOM object
		doc, err3 = goquery.NewDocumentFromReader(res.Body)
		if err3 != nil {
			R.Println(" *** ")
			R.Println(" *** ERROR IN SCRAPE_TOOL - During GOQUERY: ")
			Y.Println(err3)
			R.Println(" *** ")
			R.Println("")
			return false, EMPTY_GOQUERY_doc, ""
		}	

		//6. Error Hanlding

		if err != nil || err2 != nil || err3 != nil {
			R.Println("")
			R.Println(" ERROR in scrapeTool..")
			Y.Println(" Sleeping before Attempting retry: ")
			W.Print(i)
			Y.Print(" of ")
			SUCCESS_FLAG = false
			G.Println(S_RETRY_MAX)
			Sleep(S_RETRY_SLEEP, true)

		} else {
		
			SUCCESS_FLAG = true
			break
			
		}


		//4. Excplictly Close connections:
		res.Body.Close()	

	} //end of for


	if DO_HARD_EXIT {				
		R.Println(" Error AGAIN! Redirect detected: ", http.ErrUseLastResponse)
		return false, EMPTY_GOQUERY_doc, ""
	}	

	if SUCCESS_FLAG == false {
		R.Println("")
		R.Println(" ***** SCRAPE TOOL ERRORED OUT!!!! *****")
		R.Println("")
		return false, EMPTY_GOQUERY_doc, ""
	}

	//6. Also for DEBUG purposes: Lets show the body of the ENTIRE document
	FULL_RESPONSE_TEXT := doc.Find("html").Text()


	return true, doc, FULL_RESPONSE_TEXT

} //end of func

// Alias for SCRAPE_TOOL
func SCRAPER(URL string, EXTRA_ARGS ...string) (bool, *goquery.Document, string) {
	return SCRAPE_TOOL(URL, EXTRA_ARGS...)
} //end of 

// Alias for SCRAPE_TOOL
func SCRAPER_TOOL(URL string, EXTRA_ARGS ...string) (bool, *goquery.Document, string) {
	return SCRAPE_TOOL(URL, EXTRA_ARGS...)
}




// Kept here as filler/example.. anything you put in this function will start when the module is imported
func init() {

	//1. Startup Stuff (init the command line params etc) . We need these Time ZONE Objects




} // end of main
