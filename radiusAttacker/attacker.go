package radiusAttacker

import (
	"context"
	"gitlab.com/radius-tank/radiusGenerator"
	"gitlab.com/radius-tank/radiusGenerator/models"
	"layeh.com/radius"
	"layeh.com/radius/dictionary"
	"log"
	"strconv"
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
		packet.Add(radiusGenerator.Dictionary[attribute.Name].OID, attacker.addAttribute(packet, attribute))
	}

	response, err := radius.Exchange(context.Background(), packet, attacker.address)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("response code", response.Code)
}

/**
Form radius attribute from generated attributes
 */
func (attacker Attacker) addAttribute(p *radius.Packet, attr models.Attribute) radius.Attribute {
	params := radiusGenerator.Dictionary[attr.Name]
	switch params.Type {
	case dictionary.AttributeString:
		//magic number means userPassword
		if params.OID != 2 {
			a, err := radius.NewString(attr.Value)
			if err != nil {
				panic(err)
			}
			return a
		} else {
			a, err := radius.NewUserPassword([]byte(attr.Value), p.Secret, p.Authenticator[:])
			if err != nil {
				panic(err)
			}
			return a
		}
	case dictionary.AttributeIPAddr:
		a, err := radius.NewIPAddr([]byte(attr.Value))
		if err != nil {
			panic(err)
		}
		return a
	case dictionary.AttributeInteger:
		number, errConvert := strconv.Atoi(attr.Value)
		if errConvert != nil {
			log.Fatalf("Can't convert genereted number")
			panic(errConvert)
		}

		return radius.NewInteger(uint32(number))
	}
	return nil
}
