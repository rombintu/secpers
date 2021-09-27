package vault

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type SecretData struct {
	Data map[string]string
}

type RespSecretData struct {
	CreatedTime  string
	DeletionTime string
	Destroyed    bool
	Version      int
}

type Secret struct {
	Data SecretData
}

type RespNewSecret struct {
	RequestID     string
	leaseID       string
	Renewable     bool
	LeaseDuration int
	Data          SecretData
	WrapInfo      interface{}
	Warnings      interface{}
	Auth          interface{}
}

func GenerateData(keys, values []string) map[string]string {
	var data map[string]string
	for i := 0; i < len(keys); i++ {
		data[keys[i]] = values[i]
	}
	return data
}

func CreateSecret(token, path string, keys, values []string) (RespNewSecret, error) {
	addr := "http://localhost"
	port := ":8200"
	respNewSecret := RespNewSecret{}
	client := &http.Client{}
	secretData := SecretData{
		Data: GenerateData(keys, values),
	}
	secret := Secret{
		Data: secretData,
	}
	reqBody, err := json.Marshal(secret)
	if err != nil {
		return respNewSecret, err
	}
	req, err := http.NewRequest(
		"POST",
		fmt.Sprint(addr, port),
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		return respNewSecret, err
	}
	req.Header = http.Header{
		"X-Vault-Token": []string{token},
		"Content-Type":  []string{"application/json"},
	}

	resp, err := client.Do(req)
	if err != nil {
		return respNewSecret, err
	}
	err = json.NewDecoder(resp.Body).Decode(&respNewSecret)
	if err != nil {
		return respNewSecret, err
	}

	return respNewSecret, err
}
