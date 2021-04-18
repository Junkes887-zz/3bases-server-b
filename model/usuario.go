package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UsuarioDecrypt struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Idade        int                `json:"idade"`
	FonteDeRenda string             `json:"fonteDeRenda"`
	Endereco     string             `json:"endereco"`
	ListBens     []BenDecrypt       `json:"listBens"`
}

type BenDecrypt struct {
	Descricao string `json:"descricao"`
}

type UsuarioEncrypt struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Idade        []byte             `bson:"idade"`
	FonteDeRenda []byte             `bson:"fonteDeRenda"`
	Endereco     []byte             `bson:"endereco"`
	ListBens     []BenEncrypt       `bson:"listBens"`
}

type BenEncrypt struct {
	Descricao []byte `bson:"descricao"`
}
