package bd

import (
	"context"
	"time"

	"github.com/albeglez/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*CheckExistingUser recibe un email de parametro y chequea si ya está la BD */
func CheckExistingUser(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	condicion := bson.M{"email": email}

	var resultado models.User

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
