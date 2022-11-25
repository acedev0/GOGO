/*
Easy use AWS helper functions
*/

package CUSTOM_GO_MODULE

import (
	
	"context"

	. "github.com/acedev0/GOGO/Gadgets"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)


var AWS_REGION = "us-east-1"
var AWS_PROFILE = "terry"
var AWS_CONF_SESS aws.Config

func AWS_INIT() {

    C.Println(" **| Connecting to AWS ")

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


	// also lets init other AWS tools (S3 and Dynamodb)
	//S3_INIT()
	//Dynamo_INIT()
}
