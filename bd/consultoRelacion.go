package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/ZachIgarz/tweetApp/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ConsultoRelacion consulta la relacion entre 2 usuarios*/
func ConsultoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("tweet")
	col := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         t.UsuarioID,
		"usuariolecacionid": t.UsuarioRelacionID,
	}

	var resultado models.Relacion
	fmt.Println(resultado)

	err := col.FindOne(ctx, condicion).Decode(&resultado)

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	return true, nil
}
