package main

import (
	"fmt"

	"github.com/uh-zz/dynagoose/dynagoose"
)

func main() {
	d := dynagoose.Dynamo{}
	fmt.Println(Select(&d))
	fmt.Println(Create(&d))
	fmt.Println(Update(&d))
	fmt.Println(Delete(&d))
}

func Select(d dynagoose.Schema) string {
	return d.Select()
}

func Create(d dynagoose.Schema) string {
	return d.Create()
}

func Update(d dynagoose.Schema) string {
	return d.Update()
}

func Delete(d dynagoose.Schema) string {
	return d.Delete()
}
