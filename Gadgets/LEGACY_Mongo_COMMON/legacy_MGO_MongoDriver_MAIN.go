/*	MONGO_Common Terrys Common Mongo Functions.. 
   ----------------------------------

   	Jul 26, 2021	- Added some fixes to Mongo Common to prevent the i/o timeout 127.0.0.1 errors
	Jul 26, 2020	v1.40	- Added some fixes to Mongo Common to prevent the i/o timeout 127.0.0.1 errors
	Feb 14, 2020	v1.33	- Changed my mind, when in a thread be sure to use the much simpler:

							TEMP_SESSION := DBSession.Copy()
							defer TEMP_SESSION.Close()		

							dbOBJ = TEMP_SESSION.DB(DBName).C(COLLECTION_NAME)
							bulk := dbOBJ.Bulk() 


	Feb 13, 2020	v1.30	- Added NEW_THREAD_DB_SESSION for dealing with DB sessions within multi thread routines
	Aug 05, 2019	v1.25	- Convert to Go Modules
	Sep 05, 2018	v1.23	- Initial Rollout

*/

package CUSTOM_GO_MODULE

import (
	// = = = NATIVE Libraries
		"flag"
		"time"
		"os"
		
	// = = = CUSTOM Libraries
	. "github.com/acedev0/GOGO/Gadgets"


	// = = = 3rd Party Libraries
		"gopkg.in/mgo.v2"

)



// **** DATABASE GLOBALS ***

var DBHOST = ""						//THis is the hostname we need to connect to the database
var DBUSER = ""
var DBPASS = ""

var DB_CONNECTION_TIMEOUT = 300		// 300 seconds (5 mins) before we stop trying to connect to mongo

var DBName = ""					// Name of the mongo database we work with
var DBNAME = ""					// alt alias for DBName

var COLLECTION_NAME = ""				// Default Name of the COLLECTION within the database we need

var MAX_DB_RETRY = 10		// MAx # of times we retry a database OP

var DB_RETRY_DELAY = 15		// Wait 15 seconds before trying to do database OP again

var INITIAL_TEST_flag = false

var INDEX_CREATE_FLAG = false



// NOTE: You have MAY have to do this connection stuff in your MAIN section so that defer clsoe gets called at end of the program
// instead of the end of this function
// If you get "Session already closed" errors re evaluate this .. something might be getting Closed early
// RETURNS a DBSESSION (copy)
func DB_INIT(EXTRA_ARGS ...string) *mgo.Session {

	//1. Get the extra VARS. this is basically the database name
	for x, VAL := range EXTRA_ARGS {

		//1b. First param is always DB NAME
		if x == 0 {
			DBName = VAL
			DBNAME = VAL
			continue
		}
		
		//2c.Second is always USERNAME
		if x == 1 {
			DBUSER = VAL
			continue
		}

		//3c. Third is always PASS
		if x == 2 {
			DBPASS = VAL
			continue
		}		

	}//end of for


	//1b. Check DBHOST ... if its passed by COMMAND LINE it takes precedence

	if DBHOST == "" {

		//1c. Otherwise.. lets check for an environment variable:
		DBHOST = os.Getenv("MONGO_DBHOST")


		//1d. Finally.. if nothing is there.. lets default to localhost.. but post a notification and delay

		if DBHOST == "" {
			Y.Println("")

			Y.Println(" ***** WARNING ******")
			C.Println(" No --dbhost was passed.. also, no MONGO_DBHOST environment variable was found!!!")
			W.Println(" I am DEFAULTING to localhost!!!")
			Y.Println("")
			Sleep(15, true)

			DBHOST = "localhost"
		} else {

			Y.Println("")
			Y.Print(" ** NOTE: Using MONGO_DBHOST environment variable: ")
			G.Println(DBHOST)
			Y.Println("")

		}
		
	}

	//1b. Some Error handling for weird situations
	if DBHOST == "" || DBNAME == "" {

		R.Println(" DBINIT ERror, DBHOST or DBNAME is empty!!")
		DO_EXIT()

	}

	//2. Fix for timeout
	timeVAL := time.Duration(DB_CONNECTION_TIMEOUT * int(time.Second))
        M.Println("")
	M.Print("       [  Connecting to MONGO DB: '")
	C.Print(DBHOST)
	M.Print("'")


	//2. Setup the dial info. If credentials are specified, we use them
	// We need this object to establish a session to our MongoDB.
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{DBHOST},
		Timeout:  timeVAL,
		Database: DBName,
	}	

	if DBUSER != "" {
		mongoDBDialInfo = &mgo.DialInfo{
			Addrs:    []string{DBHOST},
			Timeout:  timeVAL,
			Database: DBName,
			Username: DBUSER,
			Password: DBPASS,
		}	
	}

	
	// Dial and return to RESULT_SESSION
	RESULT_SESSION, RES_err := mgo.DialWithInfo(mongoDBDialInfo)

	if RES_err != nil {
		R.Println(" ERROR after Initial Mongo Connection: ")
		Y.Println(RES_err)
		DO_EXIT()
	}

	defer RESULT_SESSION.Close()


	RESULT_SESSION.SetMode(mgo.Monotonic, true)

	// This makes sockets timeout after Four HOURS
	RESULT_SESSION.SetSocketTimeout(4 * time.Hour)	// Needed to fix those i/o timeout 127.0.0.1 errors weve been seeing

	G.Print("  -  Success!!")
	M.Println("  ]")
	M.Println("")


	/*
	 if you need to, this is how you COPY a DBSESSION that will work (dont just assign it with := Assign to the global DB session object
	 DBSession = temp_DBSess.Copy()

	 Also if you ever need to drop a collection:
	 	db.DropCollection()

	*/

	
	//4. COMPLETED!

	return RESULT_SESSION.Copy()
} //end of func


var INDEX_KEYLIST_TEMP []string

// when called, it adds the index key to the TEMP_INDEX_KEYS
func INDEX_KEY(index_name string)  {

	INDEX_KEYLIST_TEMP = append(INDEX_KEYLIST_TEMP, index_name)

}//end of func

//just an alias to INDEX_KEY
func ADD_KEY(ind_name string) {
	INDEX_KEY(ind_name)
}

//just an alias to INDEX_KEY
func ADD_KEYS(ind_name string) {
	INDEX_KEY(ind_name)
}

// Easy way to create indexes on Tables.. this only creates and index if it DOES NOT already exist (ver EnsureIndex)
// THere is a global that is created (And set to null when we are done) each time this is called
func INDEX_CREATE(index_NAME string, DBName string, COLLECT_NAME string, isUNIQ bool) bool {


	// dbtemp2 := DBSession.DB(DBName).C(COLLECT_NAME)

	SESS := DB_INIT(DBName)
	defer SESS.Close()

	dbtemp2 := SESS.DB(DBName).C(COLLECT_NAME)




	Y.Print("\n ** Attempting to create Index '")
	W.Print(index_NAME)
	Y.Print("' on '")
	W.Println(COLLECT_NAME)
		
	// Note, the index defaults to ASCENDING (oldest to newest)
	// If you want DESCENDING, add a - (hyphen) before the below field name(s)

	index2 := mgo.Index{
		Name:       index_NAME,
		Key:        INDEX_KEYLIST_TEMP,
		Unique:     isUNIQ,
		DropDups:   false,
		Background: false,
		Sparse:     false,
	}


	//2. This means we will create the index ONLY if it DOESNT already exist
	err := dbtemp2.EnsureIndex(index2)

	if err != nil {
		R.Print(" ERROR in INDEX_CREATE: ")
		Y.Println(err)

		return false
	}

	G.Print(" ** Index Creation SUCCESSFUL!!\n\n")


	//5. WE MUST ALWAYS PURGE THE TEMP_KEYS array... CRITICAL
	INDEX_KEYLIST_TEMP = nil

	return true
} //end of index creator





// This smartly saves a record (or records) to a database using upsert:



var LEGACY_MONGO_SETUP_DEFAULT_PARAMS = false



func init() {

	flag.StringVar(&DBHOST,               "dbhost", DBHOST,                        "      Hostname of the Mongo DATABASE")
	flag.BoolVar(&INDEX_CREATE_FLAG,   "indexes", INDEX_CREATE_FLAG,  "       Flag for Creating the Indexes a Collection needs at startup")

	if LEGACY_MONGO_SETUP_DEFAULT_PARAMS {

		flag.StringVar(&DBUSER,               "dbuser", DBUSER,                    "       DB Username Credentials")
		flag.StringVar(&DBPASS,               "dbpass", DBPASS,                    "       DB Passs credentials")
		flag.IntVar(&DB_CONNECTION_TIMEOUT,   "dbtimeout", DB_CONNECTION_TIMEOUT,  "       DB Passs credentials")
	}


	

}

