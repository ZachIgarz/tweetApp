package bd

import (
	"context"
	"time"

	"github.com/ZachIgarz/tweetApp/models"
)

/*BorroRelacion Borra relaccion entre usuarioss*/
func BorroRelacion(t models.Relacion) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()
	db := MongoCN.Database("tweet")
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, t)

	if err != nil {
		return false, err
	}
	return true, nil

}
