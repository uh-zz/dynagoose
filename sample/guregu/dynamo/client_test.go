package main

import (
	"testing"
)

func TestCreate(t *testing.T) {
	t.Setenv("DYNAMO_ENDPOINT", "http://localhost:8000")
	InitDynamo()

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
	t.Setenv("DYNAMO_ENDPOINT", "http://localhost:8000")
	InitDynamo()

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
	t.Setenv("DYNAMO_ENDPOINT", "http://localhost:8000")
	InitDynamo()

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
	t.Setenv("DYNAMO_ENDPOINT", "http://localhost:8000")
	InitDynamo()

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
	t.Setenv("DYNAMO_ENDPOINT", "http://localhost:8000")
	InitDynamo()

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
