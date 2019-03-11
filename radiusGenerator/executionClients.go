package radiusGenerator

import (
	"fmt"
	"gitlab.com/radius-tank/radiusGenerator/models"
	"os/exec"
)

type ExecutiongClient struct {
	path string
}

//return attributes from client script with keys and values todo
func (ExecutiongClient) getAttributes() []models.Attribute {
	return []models.Attribute{}
}


func (client ExecutiongClient) executeScript() {
	cmd := exec.Command("go run "+client.path, "")
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic("")
	}
	fmt.Println(string(out))
}