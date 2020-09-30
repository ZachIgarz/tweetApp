package bd

import (
	"context"
	"time"

	"github.com/ZachIgarz/tweetApp/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*LeoTweetsSeguidores Lee los tweets de mis seguidores*/
func LeoTweetsSeguidores(ID string, pagina int) ([]models.DevuelvoTweetsSeguidores.DevuelvoTweetsSeguidores, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("tweet")
	col := db.Collection("relacion")

	skip := (pagina - 1) * 20

	condiciones := make([]bson.M, 0)
	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}})
	/*UNION DE TABLAS*/
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "usuariorelacionid",
			"foreingField": "userid",
			"as":           "tweet",
		}})
	/*permite que vengan los "documentos iguales"*/
	condiciones = append(condiciones, bson.M{"$unwind": "$tweet"})

	/* 1 asc  -1 desc*/
	condiciones = append(condiciones, bson.M{"$sort": bson.M{
		"fecha": "-1",
	}})
	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit": 20})

	/*Aggregate framework, con este no se tiene que recorer el cursor*/
	cursor, err := col.Aggregate(ctx, condiciones)
	var result []models.DevuelvoTweetsSeguidores

	/*Recorre el cursos y lo formatea en el model*/
	err = cursor.All(ctx, &result)
	if err != nil {

		return result, false
	}

	return result, true

}
