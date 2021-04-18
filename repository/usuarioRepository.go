package repository

import (
	"context"

	"github.com/Junkes887/3bases-server-b/model"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Client struct {
	DB_MONGO           *mongo.Collection
	DB_REDIS           *redis.Client
	CTX                context.Context
	TIME_MINUTES_REDIS int
}

func (client Client) SetDataRedis() {
	usuarios := client.FindAllMongo()
	for _, usuario := range usuarios {
		client.SaveRedis(usuario.ID.Hex(), usuario)
	}
}

func (client Client) FindAll() []model.UsuarioDecrypt {
	return client.FindAllRedis()
}

func (client Client) Find(id string) model.UsuarioDecrypt {
	if client.ExistsRedis(id) {
		return client.FindRedis(id)
	} else {
		return client.FindMongo(id)
	}
}

func (client Client) Save(usuario model.UsuarioDecrypt) interface{} {
	res := client.SaveMongo(usuario)

	objectID := res.(primitive.ObjectID)

	usuario.ID = objectID

	client.SaveRedis(objectID.Hex(), usuario)

	return res
}

func (client Client) Upadate(id string, usuario model.UsuarioDecrypt) string {
	client.UpdateRedis(id, usuario)

	return client.UpadateMongo(id, usuario)
}

func (cliente Client) Delete(id string) string {
	cliente.DeleteRedis(id)
	return cliente.DeleteMongo(id)
}
