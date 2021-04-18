package repository

import (
	"fmt"
	"log"

	"github.com/Junkes887/3bases-server-b/builder"
	"github.com/Junkes887/3bases-server-b/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (client Client) FindAllMongo() []model.UsuarioDecrypt {
	var usuarios []model.UsuarioEncrypt
	var usuariosDecrypt []model.UsuarioDecrypt
	cur, err := client.DB_MONGO.Find(client.CTX, bson.D{})

	if err != nil {
		fmt.Print(err)
	}

	for cur.Next(client.CTX) {
		var u model.UsuarioEncrypt
		err := cur.Decode(&u)
		if err != nil {
			fmt.Println(err)
		}
		usuarios = append(usuarios, u)
	}

	if err := cur.Err(); err != nil {
		fmt.Println(err)
	}

	cur.Close(client.CTX)

	for _, usuario := range usuarios {
		usuarioDecrypt := builder.DecryptUsuario(usuario)
		usuariosDecrypt = append(usuariosDecrypt, usuarioDecrypt)
	}

	return usuariosDecrypt
}

func (client Client) FindMongo(id string) model.UsuarioDecrypt {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}

	var usuario model.UsuarioEncrypt
	filter := bson.M{
		"_id": objectId,
	}
	err = client.DB_MONGO.FindOne(client.CTX, filter).Decode(&usuario)
	if err != nil {
		log.Println(err)
	}

	if usuario.ID == primitive.NilObjectID {
		return model.UsuarioDecrypt{}
	}

	usuarioDecrypt := builder.DecryptUsuario(usuario)

	client.SaveRedis(id, usuarioDecrypt)

	return usuarioDecrypt
}

func (client Client) SaveMongo(usuario model.UsuarioDecrypt) interface{} {
	usuarioEncrypt := builder.EncryptUsuario(usuario)

	res, err := client.DB_MONGO.InsertOne(client.CTX, usuarioEncrypt)
	if err != nil {
		fmt.Println(err)
	}
	return res.InsertedID
}

func (client Client) UpadateMongo(id string, usuario model.UsuarioDecrypt) string {
	usuarioEncrypt := builder.EncryptUsuario(usuario)

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Println("Invalid id")
	}

	update := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "idade", Value: usuarioEncrypt.Idade},
		primitive.E{Key: "fonteDeRenda", Value: usuarioEncrypt.FonteDeRenda},
		primitive.E{Key: "endereco", Value: usuarioEncrypt.Endereco},
		primitive.E{Key: "listBens", Value: usuarioEncrypt.ListBens},
	}}}

	res, err := client.DB_MONGO.UpdateByID(client.CTX, objectId, update)

	if err != nil {
		log.Println("Invalid id")
	}

	if res.ModifiedCount == 0 {
		return "Usuario não encontrado"
	}

	return "Usuario atualizado"
}

func (cliente Client) DeleteMongo(id string) string {
	objectId, err := primitive.ObjectIDFromHex(id)
	filter := bson.D{primitive.E{Key: "_id", Value: objectId}}

	res, err := cliente.DB_MONGO.DeleteOne(cliente.CTX, filter)
	if err != nil {
		log.Println(err)
	}

	if res.DeletedCount == 0 {
		return "Usuario não encontrado"
	}

	return "Usuario removido"
}
