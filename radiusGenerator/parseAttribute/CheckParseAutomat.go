package parseAttribute

import (
	"fmt"
	"gitlab.com/radius-tank/radiusGenerator/models"
)

type CheckAutomat struct {
}

const (
	SPACE              = byte(' ')
	GenStateKey        = byte('G')
	ValueStateKey      = byte('V')
	IntegratedStateKey = byte('I')
	NEWLINE            = byte('\n')
)


const (
	StartState      = 1
	TagState        = 2
	GenState        = 3
	ValueState      = 4
	IntegratedState = 5
)

const (
	MacGenState    = byte('M')
	WordGenState   = byte('W')
	NumberGenState = byte('N')
)

//Automate for format checking
func (automat CheckAutomat) checkFormatAndGet(s []byte) (result models.ScenarioAttribute) {
	var currentState = StartState
	currentBuf := ""
	for _, c := range s {
		currentBuf += string(c)
		switch currentState {
		case StartState:
			switch c {
			case SPACE:
				fmt.Print("Space")
				result.Name = currentBuf
				currentBuf = ""
				currentState = TagState
			}
		case TagState:
			//Move to another function
			result.Tag = currentBuf
			currentBuf = ""
			switch c {
			case GenStateKey:
				currentState = GenState
			case ValueStateKey:
				currentState = ValueState
			case IntegratedStateKey:
				currentState = IntegratedState
			}
		case GenState:
			switch c {
			case MacGenState:
				result.Value = "MAC"
			case WordGenState:
				result.Value = "WORD"
			case NumberGenState:
				result.Value = "NUMBER"
			}
			return result
		case ValueState:
			//Do nothing
		case IntegratedState:
			switch c {
			case NEWLINE:
				result.Value = currentBuf
				return result
			}
		}
	}
	//Set this is our state is value we get to that
	result.Value = currentBuf
	return result
}
