package repository

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/Junkes887/3bases-server-b/model"
)

func (client Client) ExistsRedis(id string) bool {
	value := client.DB_REDIS.Get(id)

	return value.Val() != ""
}

func (client Client) FindAllRedis() []model.UsuarioDecrypt {
	var usuarios []model.UsuarioDecrypt

	do := client.DB_REDIS.Do("KEYS", "*")
	res, _ := do.Val().([]interface{})

	for _, in := range res {
		key := in.(string)

		usuarios = append(usuarios, client.FindRedis(key))
	}
	return usuarios
}

func (client Client) FindRedis(id string) model.UsuarioDecrypt {
	var usuario model.UsuarioDecrypt
	idRedis := client.DB_REDIS.Get(id)

	b, _ := idRedis.Bytes()
	json.Unmarshal(b, &usuario)

	return usuario
}

func (client Client) SaveRedis(id string, usuario model.UsuarioDecrypt) {
	p, err := json.Marshal(usuario)
	if err != nil {
		fmt.Println(err)
	}
	client.DB_REDIS.Set(id, p, time.Duration(rand.Intn(client.TIME_MINUTES_REDIS))*time.Minute)
}

func (client Client) UpdateRedis(id string, usuario model.UsuarioDecrypt) {
	client.DB_REDIS.Del(id)

	p, err := json.Marshal(usuario)
	if err != nil {
		fmt.Println(err)
	}
	client.DB_REDIS.Set(id, p, time.Duration(rand.Intn(client.TIME_MINUTES_REDIS))*time.Minute)
}

func (client Client) DeleteRedis(id string) {
	client.DB_REDIS.Del(id)
}
