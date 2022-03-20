package main

import (
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

func init() {
	os.Setenv("DYNAMO_ENDPOINT", "http://localhost:8000")

	dynamoDbRegion := os.Getenv("AWS_REGION")
	disableSsl := false

	dynamoDbEndpoint := os.Getenv("DYNAMO_ENDPOINT")
	if len(dynamoDbEndpoint) != 0 {
		disableSsl = true
	}

	if len(dynamoDbRegion) == 0 {
		dynamoDbRegion = "ap-northeast-1"
	}

	// TODO: ここでモック化してもいいかも
	Dynamo.Db = dynamo.New(session.New(), &aws.Config{
		Region:     aws.String(dynamoDbRegion),
		Endpoint:   aws.String(dynamoDbEndpoint),
		DisableSSL: aws.Bool(disableSsl),
	})
}

func TestCreate(t *testing.T) {
	table := Dynamo.Db.Table("MyFirstTable")
	item := Item{
		MyHashKey:  "MyHash",
		MyRangeKey: 1,
		MyText:     "My First Text",
	}

	if err := Create(table, item); err != nil {
		t.Fatalf("unable to Create: %v", err)
	}
}

func TestRead(t *testing.T) {
	table := Dynamo.Db.Table("MyFirstTable")
	item := Item{
		MyHashKey:  "MyHash",
		MyRangeKey: 1,
		MyText:     "My First Text",
	}

	if err := Create(table, item); err != nil {
		t.Fatalf("unable to Create: %v", err)
	}

	_, err := Read(table, item)
	if err != nil {
		t.Fatalf("unable to Read: %v\n", err)
	}
}

func TestUpdate(t *testing.T) {
	table := Dynamo.Db.Table("MyFirstTable")
	item := Item{
		MyHashKey:  "MyHash",
		MyRangeKey: 1,
		MyText:     "My First Text",
	}
	if err := Create(table, item); err != nil {
		t.Fatalf("unable to Create: %v", err)
	}

	item.MyText = "My Second Text"

	_, err := Update(table, item)
	if err != nil {
		t.Fatalf("unable to Update: %v\n", err)
	}
}

func TestConditionalCheck(t *testing.T) {
	table := Dynamo.Db.Table("MyFirstTable")
	item := Item{
		MyHashKey:  "MyHash",
		MyRangeKey: 1,
		MyText:     "My First Text",
	}
	if err := Create(table, item); err != nil {
		t.Fatalf("unable to Create: %v", err)
	}

	err := ConditionalCheck(table, item)
	if err != nil {
		t.Fatalf("unable to delete item with conditional check: %v\n", err)
	}
}

func TestDelete(t *testing.T) {
	table := Dynamo.Db.Table("MyFirstTable")
	item := Item{
		MyHashKey:  "MyHash",
		MyRangeKey: 1,
		MyText:     "My First Text",
	}
	if err := Create(table, item); err != nil {
		t.Fatalf("unable to Create: %v", err)
	}

	if err := Delete(table, item); err != nil {
		t.Fatalf("unable to Delte: %v\n", err)
	}
}
