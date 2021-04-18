package builder

import (
	artifacts "github.com/Junkes887/3bases-artifacts"
	"github.com/Junkes887/3bases-server-b/model"
)

func EncryptUsuario(usuario model.UsuarioDecrypt) model.UsuarioEncrypt {
	return model.UsuarioEncrypt{
		ID:           usuario.ID,
		Idade:        artifacts.EncryptInt(usuario.Idade),
		FonteDeRenda: artifacts.Encrypt(usuario.FonteDeRenda),
		Endereco:     artifacts.Encrypt(usuario.Endereco),
		ListBens:     encryptListBens(usuario.ListBens),
	}
}

func encryptListBens(bens []model.BenDecrypt) []model.BenEncrypt {
	var divadasEncrypt []model.BenEncrypt
	for _, ben := range bens {
		divadaEncrypt := model.BenEncrypt{
			Descricao: artifacts.Encrypt(ben.Descricao),
		}
		divadasEncrypt = append(divadasEncrypt, divadaEncrypt)
	}

	return divadasEncrypt
}

func DecryptUsuario(usuario model.UsuarioEncrypt) model.UsuarioDecrypt {
	return model.UsuarioDecrypt{
		ID:           usuario.ID,
		Idade:        artifacts.DecryptInt(usuario.Idade),
		FonteDeRenda: artifacts.Decrypt(usuario.FonteDeRenda),
		Endereco:     artifacts.Decrypt(usuario.Endereco),
		ListBens:     DecryptListBens(usuario.ListBens),
	}
}

func DecryptListBens(bens []model.BenEncrypt) []model.BenDecrypt {
	var divadasDecrypt []model.BenDecrypt
	for _, ben := range bens {
		divadaDecrypt := model.BenDecrypt{
			Descricao: artifacts.Decrypt(ben.Descricao),
		}
		divadasDecrypt = append(divadasDecrypt, divadaDecrypt)
	}

	return divadasDecrypt
}
