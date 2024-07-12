package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	configFile := "./internal/config/config.pkl"
	configExampleFile := "./internal/config/config.example.pkl"

	// Check if environment variables are set
	envVars := []string{"PORT", "DB", "CLOUDFLARE_ACCOUNT_ID", "CLOUDFLARE_KV_NAMESPACE_ID", "CLOUDFLARE_TOKEN", "CLOUDFLARE_TURNSTILE_SECRET"}
	for _, envVar := range envVars {
		if _, ok := os.LookupEnv(envVar); !ok {
			fmt.Printf("Error: environment variable '%s' is not set\n", envVar)
			return
		}
	}

	// Create the config file if it doesn't exist
	_, err := os.Stat(configFile)
	if err != nil && os.IsNotExist(err) {
		f, err := os.Create(configFile)
		if err != nil {
			fmt.Println("Error creating config file:", err)
			return
		}
		f.Close()
	}

	// Copy the contents
	if err := exec.Command("cp", configExampleFile, configFile).Run(); err != nil {
		fmt.Println("Error copying config file:", err)
		return
	}

	// Read the config file content
	data, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	// Replace placeholders with environment variables
	replacedContent := string(data)
	placeholders := map[string]string{
		"port = \"port\"":              os.Getenv("PORT"),
		"dsn = \"dsn\"":                os.Getenv("DB"),
		"accountID = \"id\"":           os.Getenv("CLOUDFLARE_ACCOUNT_ID"),
		"kvNamespaceID = \"id\"":       os.Getenv("CLOUDFLARE_KV_NAMESPACE_ID"),
		"token = \"token\"":            os.Getenv("CLOUDFLARE_TOKEN"),
		"turnstileSecret = \"secret\"": os.Getenv("CLOUDFLARE_TURNSTILE_SECRET"),
	}

	for placeholder, value := range placeholders {
		replacedContent = strings.Replace(replacedContent, placeholder, fmt.Sprintf(`%s = "%s"`, strings.Split(placeholder, " = ")[0], value), -1)
	}

	// Write the updated content back to the file
	err = os.WriteFile(configFile, []byte(replacedContent), 0644)
	if err != nil {
		fmt.Println("Error writing to config file:", err)
		return
	}

	fmt.Println("Config file updated successfully!")
}
