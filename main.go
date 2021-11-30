package main

import (
	"context"
	"ent_sample/models"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	ctx := context.Background()
	models, err := models.Open("sqlite3", "file:ent?&_fk=1")
	if err != nil {
		panic(err)
	}
	start := time.Now()
	for i := 0; i < 10000; i++ {
		err = models.Schema.Create(ctx)
		if err != nil {
			panic(err)
		}
		old := models.Student.Create().SetName("amirreza" + fmt.Sprint(i)).SetBirthYear(1998).SaveX(ctx)

		amirreza := models.Student.GetX(ctx, old.ID)
		_ = amirreza.Update().AddBirthYear(1).SaveX(ctx)
		models.Student.DeleteOne(amirreza).ExecX(ctx)
	}
	fmt.Println(time.Now().Sub(start).Milliseconds())
}
