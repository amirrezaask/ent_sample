package main

import (
	"context"
	"ent_sample/models"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	ctx := context.Background()
	models, err := models.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		panic(err)
	}
	err = models.Schema.Create(ctx)
	if err != nil {
		panic(err)
	}
	amirreza := models.Student.Create().SetName("amirreza").SetBirthYear(1998).SaveX(ctx)

	rizmo := models.Course.Create().SetName("rizmo").SetCapacity(20).SaveX(ctx)

	amirreza.Update().AddCourses(rizmo).SaveX(ctx)
	amirreza = models.Student.GetX(ctx, 1)

	fmt.Printf("%+v\n", rizmo.QueryStudents().AllX(ctx))
	fmt.Printf("%+v\n", amirreza.QueryCourses().AllX(ctx))

}
