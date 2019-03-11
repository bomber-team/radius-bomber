package models

import (
	"layeh.com/radius"
	"layeh.com/radius/dictionary"
)

type AttributeParams struct {
	Type dictionary.AttributeType
	OID  radius.Type
}
