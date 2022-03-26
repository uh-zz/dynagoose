package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type Item struct {
	MyHashKey  string
	MyRangeKey int
	MyText     string
}

type MyDynamo struct {
	Db *dynamo.DB
}

var Dynamo MyDynamo

func InitDynamo() {
	dynamoDbRegion := os.Getenv("AWS_REGION")
	disableSsl := false

	dynamoDbEndpoint := os.Getenv("DYNAMO_ENDPOINT")
	if len(dynamoDbEndpoint) != 0 {
		disableSsl = true
	}

	if len(dynamoDbRegion) == 0 {
		dynamoDbRegion = "ap-northeast-1"
	}

	Dynamo.Db = dynamo.New(session.New(), &aws.Config{
		Region:     aws.String(dynamoDbRegion),
		Endpoint:   aws.String(dynamoDbEndpoint),
		DisableSSL: aws.Bool(disableSsl),
	})
}

func main() {
	InitDynamo()

	table := Dynamo.Db.Table("MyFirstTable")
	item := Item{
		MyHashKey:  "MyHash",
		MyRangeKey: 1,
		MyText:     "My First Text",
	}

	if err := Create(table, item); err != nil {
		fmt.Printf("unable to Create: %v\n", err)
		return
	}

	readItem, err := Read(table, item)
	if err != nil {
		fmt.Printf("unable to Read: %v\n", err)
		return
	}
	fmt.Printf("read item: %v\n", readItem)

	item.MyText = "My Second Text"
	updateItem, err := Update(table, item)
	if err != nil {
		fmt.Printf("unable to Update: %v\n", err)
		return
	}
	fmt.Printf("updated item: %v\n", updateItem)

	if err := ConditionalCheck(table, item); err != nil {
		fmt.Printf("unable to delete item with conditional check: %v\n", err)
		return
	}

	if err := Delete(table, item); err != nil {
		fmt.Printf("unable to Delte: %v\n", err)
		return
	}
	fmt.Printf("deleted item: %v\n", item)
}

func Create(table dynamo.Table, item Item) error {
	if err := table.Put(item).Run(); err != nil {
		return err
	}
	return nil
}

func Read(table dynamo.Table, item Item) (Item, error) {
	var result Item
	if err := table.Get("MyHashKey", item.MyHashKey).Range("MyRangeKey", dynamo.Equal, item.MyRangeKey).One(&result); err != nil {
		return Item{}, err
	}
	return result, nil
}

func Update(table dynamo.Table, item Item) (Item, error) {
	var result Item
	if err := table.Update("MyHashKey", item.MyHashKey).Range("MyRangeKey", item.MyRangeKey).Set("MyText", item.MyText).Value(&result); err != nil {
		return Item{}, err
	}
	return result, nil
}

func ConditionalCheck(table dynamo.Table, item Item) error {
	if err := table.Delete("MyHashKey", item.MyHashKey).Range("MyRangeKey", item.MyRangeKey).If("MyText = ?", item.MyText).Run(); err != nil {
		return err
	}
	return nil
}

func Delete(table dynamo.Table, item Item) error {
	if err := table.Delete("MyHashKey", item.MyHashKey).Range("MyRangeKey", item.MyRangeKey).Run(); err != nil {
		return err
	}
	return nil
}
