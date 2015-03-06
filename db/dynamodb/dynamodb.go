package dynamodb

import (
	"errors"
	"log"
	"os"
	"reflect"

	aws "github.com/AdRoll/goamz/aws"
	dynamo "github.com/AdRoll/goamz/dynamodb"
	"github.com/joho/godotenv"
	. "github.com/obieq/goar"
)

type ArDynamodb struct {
	ActiveRecord
	ID string `json:"ID,omitempty"`
	Timestamps
}

var (
	client *dynamo.Server
)

var connectOpts = func() map[string]string {
	opts := make(map[string]string)

	if envs, err := godotenv.Read(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	} else {
		opts["accessKey"] = os.Getenv(envs["AWS_ACCESS_KEY_ID"])
		opts["secretKey"] = os.Getenv(envs["AWS_SECRET_ACCESS_KEY"])
	}

	return opts
}

func connect() *dynamo.Server {
	opts := connectOpts()

	region := aws.USEast
	auth := aws.Auth{AccessKey: opts["accessKey"], SecretKey: opts["secretKey"]}

	return dynamo.New(auth, region)
}

func init() {
	client = connect()
}

func Client() *dynamo.Server {
	return client
}

func (ar *ArDynamodb) All(models interface{}, opts map[string]interface{}) (err error) {
	return errors.New("All method not supported by Dynamodb.  Create a View instead.")
}

func (ar *ArDynamodb) Truncate() (numRowsDeleted int, err error) {
	return -1, errors.New("Truncate method not yet implemented")
}

func (ar *ArDynamodb) Find(key interface{}) (interface{}, error) {
	self := ar.Self()
	strKey := key.(string)

	primary := dynamo.NewStringAttribute("ID", "")
	pk := dynamo.PrimaryKey{KeyAttribute: primary}
	t := dynamo.Table{Server: client, Name: ar.ModelName(), Key: pk}
	dynamoKey := &dynamo.Key{HashKey: strKey}

	err := t.GetDocument(dynamoKey, self)

	// NOTE: the AdRoll sdk returns an error is the key doesn't exist
	if err == nil {
		// set the ID b/c the AdRoll sdk purposefully empties it for some reason
		pointer := reflect.Indirect(reflect.ValueOf(self))
		field := pointer.FieldByName("ID")
		field.SetString(strKey)
	} else {
		self = nil
	}

	return self, err
}

func (ar *ArDynamodb) DbSave() error {
	// primary key initialization example
	//     https://github.com/AdRoll/goamz/blob/c73835dc8fc6958baf8df8656864ee4d6d04b130/dynamodb/query_builder_test.go
	//         primary := NewStringAttribute("TestHashKey", "")
	//         secondary := NewNumericAttribute("TestRangeKey", "")
	//         key := PrimaryKey{primary, secondary}
	primary := dynamo.NewStringAttribute("ID", "")
	pk := dynamo.PrimaryKey{KeyAttribute: primary}
	t := dynamo.Table{Server: client, Name: ar.ModelName(), Key: pk}

	dynamoKey := &dynamo.Key{HashKey: ar.ID}
	return t.PutDocument(dynamoKey, ar.Self())
}

func (ar *ArDynamodb) DbDelete() (err error) {
	primary := dynamo.NewStringAttribute("ID", "")
	pk := dynamo.PrimaryKey{KeyAttribute: primary}
	t := dynamo.Table{Server: client, Name: ar.ModelName(), Key: pk}

	dynamoKey := &dynamo.Key{HashKey: ar.ID}
	return t.DeleteDocument(dynamoKey)
}

func (ar *ArDynamodb) DbSearch(models interface{}) (err error) {
	return errors.New("Search method not supported by Dynamodb.  Create a View instead.")
}
