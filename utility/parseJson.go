package utility

import (
	"easyget/models"
	"encoding/json"
	"io"
	"log"
	"os"
)

func ParseJson(dest *models.Aliases) {
	jsonFile, err := os.Open("aliases/depAliases.json")

	if err != nil {
		log.Fatalf("Error occured: %v", err)
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &dest)

	
}
