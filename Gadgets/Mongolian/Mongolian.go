/*
Mongolian - Mongo Wrapper using the OFFICIAL Mongo Driver for golang
---------------------------------------------------------------------------------------
This is based on the MONGO_Common (which uses the Legacy ...as of circa 2018) mgo Driver..
NOTE: For Functions or Variables to be globally availble. The MUST start with a capital letter.
	  (This is a GO Thing)

	v1.31	-	Feb 03, 2021	-	Revamped to run on Azure DevOps
	v1.25	-	Dec 29, 2019	-	Converted this to a GO MOdule that lives in Github/Bitbucket. No more local stuff!
	v1.23	- 	Nov 05, 2016	-	Initial Release

*/

package CUSTOM_GO_MODULE

import (
	
	// = = = = Native Library
		"context"
		"time"
		"flag"
		"strings"
		"os"

	// = = = = CUSTOM Libraries
	. "github.com/acedev0/GOGO/Gadgets"


	// = = = = 3RD Party MODULES - - - -
		"go.mongodb.org/mongo-driver/mongo"
		"go.mongodb.org/mongo-driver/mongo/options"	
		"go.mongodb.org/mongo-driver/x/bsonx"	


)

// **** DATABASE GLOBALS ***
var CONN_OBJ *mongo.Client		// This is the Server/HOST connection OBJECT

var SESS_OBJ 	*mongo.Collection		
var SESS 		*mongo.Collection  	// this is meant to be an alias to SESS_OBJ
var DBOBJ 		*mongo.Collection  	// this is meant to be an alias to SESS_OBJ
var DB_OBJ 		*mongo.Collection  	// this is meant to be an alias to SESS_OBJ

var DB_HOST = "localhost"
var DB_PORT = "27017"
var DB_NAME = ""
var DB_USER = ""
var DB_PASS = ""

var CTX = context.TODO()

var DB_ERROR error
var COLLECTION_NAME = "ABC_DEF_GHI"

/*
	Call When connecting to the Database for the first time.
	DB_Init will return some global connection objects and you can specific the DBNAME as either the GLOBALS or as parameters

	Params order:
	DB_HOST, DB_NAME_DB_COLLECTION and optionally, DB_PORT, USERNAME, PASSS
*/
func DB_INIT(INPUT_PARAMS ...string) {

	//1. Iterate through the params (if they are provided)
	for pnum, param := range INPUT_PARAMS {

		switch pnum {
			case 0:
				DB_HOST = param
			case 1:
				DB_NAME = param	
			case 2:
				COLLECTION_NAME = param
			case 3:
				DB_PORT = param
			case 4:
				DB_USER = param
			case 5:
				DB_PASS = param
		} //end of switch

	} //end of for

	//2. Some error handling
	if DB_NAME == "" || COLLECTION_NAME == "" {
		R.Println("\nERROR: You need to set the DB_NAME and COLLECTION to connect to mongo!")
		os.Exit(ERROR_EXIT_CODE)
	}

	//3. We should have all our database params by now.. Lets give a status	
	M.Print(" ** Initiating Connection to MONGO SERVER: ")
	W.Print(DB_HOST)
	M.Print(" - DATABASE: ")
	G.Print(DB_NAME)
	W.Print(" / ")
	Y.Println(COLLECTION_NAME)

	tempctx, _ := context.WithTimeout(context.Background(), 10*time.Second)			// This is how you do driver/connection settings on your DB Connection
	db_conn_line := "mongodb://" +  DB_HOST + ":" + DB_PORT

	//4. Safety, if there is NO DB_USER or DB_PASS, lets not use them in the connect line
	if DB_USER != "" && DB_PASS != "" {
		db_conn_line = "mongodb://" + DB_USER + ":" + DB_PASS + "@" +  DB_HOST + ":" + DB_PORT
	}

	CONN_OBJ, DB_ERROR = mongo.Connect(tempctx, options.Client().ApplyURI(db_conn_line))

	//4b. Error Handling
	if DB_ERROR != nil {
		R.Println(" Sorry! Cant connect to Mongo for some reason!, ", DB_ERROR)
	}

	//5. Finally,  Now make SESS_OBJ/DBOBJ available globally
	SESS_OBJ = CONN_OBJ.Database(DB_NAME).Collection(COLLECTION_NAME)
	DBOBJ = SESS_OBJ
	SESS = SESS_OBJ
	DB_OBJ = SESS_OBJ

} //end of func



func Test_Mongo_CALL() {
	C.Println(" Here is the test Mongo Call")
}


// Connect to a NEW database or Collection. This returns a NEW connection object (but does NOT replace the original)
func NEW_CONN(dbname string, collname string) *mongo.Collection {
	
	TEMP_SESS_OBJ := CONN_OBJ.Database(dbname).Collection(collname)
	return TEMP_SESS_OBJ
} //end of func



// Easy way to create indexes on Tables.. this only creates and index if it DOES NOT already exist (ver EnsureIndex)
// THere is a global that is created (And set to null when we are done) each time this is called
func INDEX_CREATE(COLLECT_NAME string, indexName string, isUNIQ bool, keys ...string) bool {

	C.Println(" ")
	C.Print(" *** Attempting to Create Index Named: ")
	G.Println(indexName)
	C.Println(" ")
	cobj := CONN_OBJ.Database(DB_NAME).Collection(COLLECT_NAME)
	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	//opts.Index().SetName(indexName)
	

	indexView := cobj.Indexes()
    keysDoc := bsonx.Doc{}
    
    // Composite index
    for _, key := range keys {
        if strings.HasPrefix(key, "-") {
            keysDoc = keysDoc.Append(strings.TrimLeft(key, "-"), bsonx.Int32(-1))
        } else {
            keysDoc = keysDoc.Append(key, bsonx.Int32(1))
        }
    } //end of for

	// Create index
	result, err := indexView.CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    keysDoc,
			Options: options.Index().SetUnique(isUNIQ),
		
		},
		opts,
	)
	if result == "" || err != nil {
		R.Println(" Error Creating INDEX!!! ", err)
		return false
	}

	return true
	
} //end of index creator



// This is an alias to INDEX_CREATE for backwards compantibility
func INDEX_CREATOR(COLLECT_NAME string, indexName string, isUNIQ bool, keys ...string) bool {

	return INDEX_CREATE(COLLECT_NAME, indexName, isUNIQ, keys...)
}

func NEW_INDEX(COLLECT_NAME string, indexName string, isUNIQ bool, keys ...string) bool {

	return INDEX_CREATE(COLLECT_NAME, indexName, isUNIQ, keys...)
}

func CREATE_INDEX(COLLECT_NAME string, indexName string, isUNIQ bool, keys ...string) bool {

	return INDEX_CREATE(COLLECT_NAME, indexName, isUNIQ, keys...)
}



func init() {

	flag.StringVar(&DB_HOST, "dbhost", DB_HOST, " Mongo DB HOST")
	flag.StringVar(&DB_NAME, "dbname", DB_NAME, " DB on the mongo Host you want to connect to")
	flag.StringVar(&COLLECTION_NAME, "coll", COLLECTION_NAME, " Name of the Collection we wanna are reading from")
	flag.StringVar(&DB_PORT, "dbport", DB_PORT, " Port to connect to the database with")

	flag.StringVar(&DB_USER, "dbuser", DB_USER, " Username (authentication User) for DB")	
	flag.StringVar(&DB_PASS, "dbpass", DB_PASS, " Password (auth password)")	
}

