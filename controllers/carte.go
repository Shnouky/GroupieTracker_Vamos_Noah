package controllers

import (
	"http/net"
	"model"
)

var Bibliotheque []model.Carte
var Nombre int

func Carte1(){
	Nombre = 1

	cart1 := model.Carte{
	personnage: iron-man,
	titre: "Trop fort",
	description: "iron man est le numero un des avengers",
	}
	Bibliotheque = append(Bibliotheque, carte1)
}

func FindCarte(c *gin.Context) {
	c.JSON(http:StatusOK, gin.H{"data": Bibliotheque})
}