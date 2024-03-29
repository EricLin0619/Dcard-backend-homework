package utils
import (
	"encoding/json"
)

func StructToJson (target interface{}) string {
	jsonBytes, _ := json.Marshal(target)
	jsonString := string(jsonBytes)
	return jsonString
}

func JsonToStruct (jsonString string, target *interface{}) {
	json.Unmarshal([]byte(jsonString), target)
}