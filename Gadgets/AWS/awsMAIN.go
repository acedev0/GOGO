/*
Easy use AWS helper functions
*/

package CUSTOM_GO_MODULE

import (
	
	"context"
    "time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb"    
    "github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"     
    "github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"   


    . "github.com/acedev0/GOGO/Gadgets"

)


var AWS_REGION = "us-east-1"
var AWS_PROFILE = "terry"
var AWS_CONF_SESS aws.Config

var DYNAMO_TABLE = "toolchain_requestsXXXXX"
var DYNAMO_SVC  *dynamodb.Client

var AWS_ALREADY_INIT = false

func AWS_INIT() {


    if AWS_ALREADY_INIT {
        return
    }

    C.Println(" - -| Connecting to AWS ")

     // Using the SDK's default configuration, loading additional config
    // and credentials values from the environment variables, shared
    // credentials, and shared configuration files
    cfg, err := config.LoadDefaultConfig(
        context.TODO(),
        config.WithRegion(AWS_REGION),
        config.WithSharedConfigProfile(AWS_PROFILE),
    )

	/*
	// For doing this with static credentials

staticProvider := credentials.NewStaticCredentialsProvider(
    accessKey, 
    secretKey, 
    sessionToken,
)
cfg, err := config.LoadDefaultConfig(
    context.Background(), 
    config.WithCredentialsProvider(staticProvider),
)

	*/


    if err != nil {
        R.Print("AWS_INIT error: ")
        Y.Println(err)
		return
    }
	// Otherwsie save the CONF object so we can use it later
	AWS_CONF_SESS = cfg

    AWS_ALREADY_INIT = true

}



func DYNAMO_INIT() {
    
    SHOW_BOX(" Connecting to DYNAMO ..")

    AWS_INIT()

    DYNAMO_SVC = dynamodb.NewFromConfig(AWS_CONF_SESS)
	
}





func DYN_CreateTable(tableName string) {

	var PRIMARY_KEY_NAME = "id"

	C.Print(" - -| Trying to Create Dynamo Table: ")
	Y.Println(tableName)
	C.Print("      Remember Primary Key is always: ")
    Y.Println(PRIMARY_KEY_NAME)

    _, err := DYNAMO_SVC.CreateTable(context.TODO(), &dynamodb.CreateTableInput{ 
        TableName:   aws.String(tableName),
        BillingMode: types.BillingModePayPerRequest,

		// ProvisionedThroughput: &types.ProvisionedThroughput{
		// 	ReadCapacityUnits:  aws.Int64(PartitionWriteReadCap),
		// 	WriteCapacityUnits: aws.Int64(PartitionWriteReadCap),
		// },		
        KeySchema: []types.KeySchemaElement{
            {
                AttributeName: aws.String(PRIMARY_KEY_NAME),
                KeyType:       types.KeyTypeHash,
            },
        },
        AttributeDefinitions: []types.AttributeDefinition{
            {
                AttributeName: aws.String(PRIMARY_KEY_NAME),
                AttributeType: types.ScalarAttributeTypeS,
            },
        },		
    })
    if err != nil {
		M.Println(" Cant Create Table: ", err)
		return
    }

	// Now wait for table creation
	Y.Println(" - -| Now Waiting for Table to be created...")

	w := dynamodb.NewTableExistsWaiter(DYNAMO_SVC)
    err = w.Wait(context.TODO(),
        &dynamodb.DescribeTableInput{
            TableName: aws.String(tableName),
        },
        2*time.Minute,
        func(o *dynamodb.TableExistsWaiterOptions) {
            o.MaxDelay = 5 * time.Second
            o.MinDelay = 5 * time.Second
        })
    if err != nil {
		M.Print(" - -| Timed out waiting for table! ")
		Y.Println(err)
    }
}





func DNY_GetAll(tableName string) (bool, int, dynamodb.ScanOutput) {


	var EMPTY dynamodb.ScanOutput

    Y.Print(" - -| Pulling ALL Records from: ")
    W.Println(tableName)

    out, err := DYNAMO_SVC.Scan(context.TODO(), &dynamodb.ScanInput{
        TableName:                 aws.String(tableName),
    })
	
    if err != nil {
		M.Println(" --| FIND Error: ", err, err.Error())	//, err.Error())
		return false, 0, EMPTY
    }

    var found = false
    if len(out.Items) > 0 {
        found = true
    }

	return found, len(out.Items), *out
}


/*  
		TO EXTRACT AN ACTUAL DYNAMO ITEM from FIND

		total, OBJ := DYN_FindItem(DYNAMO_TABLE, "GAME_ID", gameid)
		
		var O []GAME_OBJ

		err2 := attributevalue.UnmarshalListOfMaps(OBJ.Items, &O)
		if err2 != nil {
			M.Println(" Error with DYN_FIND Unmarshal!", err2.Error())
		}
    
      Returns True (if found) total ITEMS found, and ARRAY WITH those items
*/
func DYN_FindItem(tableName string, keyname string, keyval string) (bool, int, dynamodb.ScanOutput) {


	var EMPTY dynamodb.ScanOutput

	expr, err := expression.NewBuilder().WithFilter(
        expression.And(
            expression.AttributeNotExists(expression.Name("deletedAt")),
            expression.Contains(expression.Name(keyname), keyval),
        ),
    ).Build()
    if err != nil {
		M.Println(" --| FIND Error: ", err, err.Error())	//, err.Error())
        return false, 0, EMPTY
    }

    out, err := DYNAMO_SVC.Scan(context.TODO(), &dynamodb.ScanInput{
        TableName:                 aws.String(tableName),
        FilterExpression:          expr.Filter(),
        ExpressionAttributeNames:  expr.Names(),
        ExpressionAttributeValues: expr.Values(),
    })
	
    if err != nil {
		M.Println(" --| FIND Error: ", err, err.Error())	//, err.Error())
		return false, 0, EMPTY
    }

    var found = false
    if len(out.Items) > 0 {
        found = true
    }

	return found, len(out.Items), *out
}


func DYN_InsertItem(tableName string, item interface{} )  {

    C.Print(" - -| About to Dynamo INSERT on: ")
    Y.Println(tableName)
	
	data, err := attributevalue.MarshalMap(item)

	if err != nil {
		M.Println("DYN Insert error", err)
        return
	}

	_, err = DYNAMO_SVC.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      data,
	})

	if err != nil {
		M.Println("DYN Insert error", err, err.Error())
        return		
	}

	G.Println(" - -| DYN Insert Success")
}
