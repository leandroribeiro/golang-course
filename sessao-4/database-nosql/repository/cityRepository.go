package repository

import (
	"context"
	"fmt"
	"github.com/leandroribeiro/go-hello-world/sessao-4/database-nosql/model"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func Get(phoneCode int) (city model.City, err error) {

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second) // Options to the database.
	coll := Db.Collection("cities")
	fmt.Println(coll.Name())

	count, _ := coll.Find(ctx, nil, nil)
	fmt.Println(count)

	findResult := coll.FindOne(ctx, bson.M{"phone_code": string(phoneCode)})
	if err := findResult.Err(); err != nil {
		fmt.Println(err)
	}

	n := model.City{}
	err = findResult.Decode(&n)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n.City) // output: Some updated text

	return

}
