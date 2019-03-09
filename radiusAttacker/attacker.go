package radiusAttacker

import (
	"context"
	"gitlab.com/radius-tank/radiusGenerator"
	"gitlab.com/radius-tank/radiusGenerator/models"
	"layeh.com/radius"
	"log"
)

type Attacker struct {
	code    radius.Code
	secret  []byte
	address string
}

//Outer func
func (attacker Attacker) Attack(attributes [][]models.Attribute) {
	for _, attribute := range attributes {
		go attacker.attack(attribute)
	}
}

//form a packet and send a request
func (attacker Attacker) attack(attributes []models.Attribute) {
	packet := radius.New(attacker.code, attacker.secret)

	for _, attribute := range attributes {
		attr, err := radius.NewString(attribute.Value)
		if err != nil {
			log.Fatal("Attr value is tooo long")
		}
		packet.Add(radiusGenerator.Dictionary[attribute.Name], attr)
	}

	response, err := radius.Exchange(context.Background(), packet, attacker.address)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("response code", response.Code)
}
