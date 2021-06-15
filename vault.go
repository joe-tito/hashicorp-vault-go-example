package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/vault/api"
)

func main() {

	fmt.Println("Using Vault:", os.Getenv("VAULT_ADDR"))

	// Create a new client and configure how we connecto Vault
	client, err := api.NewClient(api.DefaultConfig())

	if err != nil {
		log.Fatal(err)
		fmt.Println("Failed to connect to Vault :(")
	}

	client.SetAddress(os.Getenv("VAULT_ADDR"))
	client.SetToken(os.Getenv("VAULT_TOKEN"))

	// Create a new secret
	var secretName string = "test/data/demo"
	var secrets = map[string]interface{}{
		"data": map[string]interface{}{
			"item1": "secrets are cool",
			"item2": "adp rocks",
			"item3": "golang is neat",
		},
	}

	// Write KV2 secret and read it back
	writeSecret(client, secretName, secrets)
	readSecret(client, secretName)

	// Update the secret, write it, and read it back
	secrets = map[string]interface{}{
		"data": map[string]interface{}{
			"item1": "protect your secrets",
			"item2": "adp is amazing",
			"item3": "golang is fun",
		},
	}

	writeSecret(client, secretName, secrets)
	readSecret(client, secretName)
}

func writeSecret(client *api.Client, secretName string, secrets map[string]interface{}) {

	fmt.Printf("Writing secret: %v\n", secretName)

	client.Logical().Write(secretName, secrets)
}

func readSecret(client *api.Client, secretName string) {

	fmt.Printf("Reading secret: %v\n", secretName)

	secret, err := client.Logical().Read(secretName)
	if err != nil {
		fmt.Println(err)
		return
	}

	m := secret.Data["data"].(map[string]interface{})
	fmt.Printf("item1: %v\n", m["item1"])
	fmt.Printf("item2: %v\n", m["item2"])
	fmt.Printf("item3: %v\n", m["item3"])
}
