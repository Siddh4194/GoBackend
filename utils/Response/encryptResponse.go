package response

import (
	"encoding/base64"
)

func EncryptThePayload(payload string) (string, error) {
	encrypted := base64.StdEncoding.EncodeToString([]byte(payload))
	return encrypted, nil
}

func respondEncrypted(data interface{})(interface{},error){
	// jsonData, err := json.Marshal(data)
	// if err != nil {
	// 	return "", err
	// }

	// encryptedData, err := EncryptThePayload(string(jsonData))
	// if err != nil {
	// 	return "", err
	// }
	return data, nil
}
