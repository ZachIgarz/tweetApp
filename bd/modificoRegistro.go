package bd

import (
	"context"
	"time"

	"github.com/ZachIgarz/tweetApp/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*ModificoRegistro permite modificar el perfil del usuario*/
func ModificoRegistro(u models.Usuario, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("tweet")
	col := db.Collection("usuarios")

	registro := make(map[string]interface{})

	if len(u.Nombre) > 0 {
		registro["nombre"] = u.Nombre
	}

	if len(u.Apellido) > 0 {
		registro["apellido"] = u.Apellido

	}

	registro["fechaNacimiento"] = u.FechaNacimiento

	if len(u.Avatar) > 0 {
		registro["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		registro["banner"] = u.Banner
	}
	if len(u.Biografia) > 0 {
		registro["biografia"] = u.Biografia
	}

	if len(u.Ubicacion) > 0 {
		registro["ubicacion"] = u.Ubicacion
	}

	if len(u.SitioWeb) > 0 {
		registro["sitioWeb"] = u.SitioWeb
	}

	//para actualizar en mongo se usa '$set'
	updatString := bson.M{
		"$set": registro,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)

	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	//Update registro
	_, err := col.UpdateOne(ctx, filtro, updatString)

	if err != nil {
		return false, err
	}

	return true, nil
}
