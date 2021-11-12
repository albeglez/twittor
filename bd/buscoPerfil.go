package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/albeglez/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BuscoPerfil(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	var perfil models.User
	ObjectID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": ObjectID,
	}

	err := col.FindOne(ctx, condicion).Decode(&perfil)
	perfil.Password = ""
	if err != nil {
		fmt.Println("Registro no encontrado" + err.Error())
		return perfil, err
	}
	return perfil, nil
}
