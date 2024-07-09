package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
)

const url = "http://localhost:8000/account"

var (
	HTTP = &http.Client{}

	CreateAccountCmd = &cobra.Command{
		Use:   "create",
		Short: "Create new account",
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")

			if name == "" {
				log.Println("Name is required.")
				return
			}

			data := map[string]string{"name": name}
			dataFromJSON, err := json.Marshal(data)

			if err != nil {
				log.Println("Error marshalling data.")
				return
			}

			response, err := HTTP.Post(url, "application/json", bytes.NewBuffer(dataFromJSON))

			if err != nil {
				log.Println("Error making request.")
				return
			}

			defer response.Body.Close()
			q, err := ioutil.ReadAll(response.Body)

			if err != nil {
				log.Println("Failed reading response.")
				return
			}

			log.Printf("Response: %s\n", q)
		},
	}

	GetAccountCmd = &cobra.Command{
		Use:   "get",
		Short: "Get account by name",
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")

			if name == "" {
				log.Println("Name is required.")
				return
			}

			response, err := HTTP.Get(url + "/" + name)

			if err != nil {
				log.Println("Error making request.")
				return
			}

			defer response.Body.Close()
			q, err := ioutil.ReadAll(response.Body)

			if err != nil {
				log.Println("Failed reading response body.")
				return
			}

			log.Printf("Response: %s\n", q)
		},
	}

	UpdateAmountCmd = &cobra.Command{
		Use:   "update",
		Short: "Update account amount",
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")
			balance, _ := cmd.Flags().GetFloat64("balance")

			if name == "" {
				log.Println("Name is required.")
				return
			}

			if balance < 0 {
				log.Println("Balance must be >= 0.")
				return
			}

			data := map[string]interface{}{"balance": balance}
			dataFromJSON, err := json.Marshal(data)

			if err != nil {
				log.Println("Error marshalling data.")
				return
			}

			request, err := http.NewRequest(http.MethodPatch, url+"/"+name, bytes.NewBuffer(dataFromJSON))

			if err != nil {
				log.Println("Error making request.")
				return
			}

			request.Header.Set("Content-Type", "application/json")
			client := &http.Client{}
			response, err := client.Do(request)

			if err != nil {
				log.Println("Error making request.")
				return
			}

			defer response.Body.Close()

			q, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Println("Failed reading response body.")
				return
			}

			log.Printf("Response: %s\n", q)
		},
	}

	ChangeAccountNameCmd = &cobra.Command{
		Use:   "rename",
		Short: "Change account name",
		Run: func(cmd *cobra.Command, args []string) {
			oldName, _ := cmd.Flags().GetString("name")
			newName, _ := cmd.Flags().GetString("new_name")

			if newName == "" || oldName == "" {
				log.Println("Old and new names are required")
				return
			}

			data := map[string]string{"new_name": newName}
			dataFromJSON, err := json.Marshal(data)

			if err != nil {
				log.Printf("Error marshalling data: %v\n", err)
				return
			}

			request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/%s", url, oldName), bytes.NewBuffer(dataFromJSON))

			if err != nil {
				log.Printf("Error making request: %v\n", err)
				return
			}

			request.Header.Set("Content-Type", "application/json")
			client := &http.Client{}
			response, err := client.Do(request)

			if err != nil {
				log.Printf("Error making request: %v\n", err)
				return
			}

			defer response.Body.Close()
			q, err := ioutil.ReadAll(response.Body)

			if err != nil {
				log.Printf("Failed to read response body: %v\n", err)
				return
			}

			log.Printf("Response: %s\n", q)
		},
	}

	DeleteAccountCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete account",
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")

			if name == "" {
				log.Println("Account name is required.")
				return
			}

			req, err := http.NewRequest(http.MethodDelete, url+"/"+name, nil)

			if err != nil {
				log.Println("Error creating request.")
				return
			}

			response, err := HTTP.Do(req)

			if err != nil {
				log.Println("Error making request.")
				return
			}

			defer response.Body.Close()
			q, err := ioutil.ReadAll(response.Body)

			if err != nil {
				log.Println("Failed to read response body.")
				return
			}

			log.Printf("Response: %s\n", q)
		},
	}
)

func main() {
	CreateAccountCmd.Flags().String("name", "", "Account name")
	GetAccountCmd.Flags().String("name", "", "Account name")
	UpdateAmountCmd.Flags().String("name", "", "Account name")
	UpdateAmountCmd.Flags().Float64("balance", 0, "Balance for adding funds to your account using")
	ChangeAccountNameCmd.Flags().String("new_name", "", "New account name")
	ChangeAccountNameCmd.Flags().String("name", "", "Old account name")
	DeleteAccountCmd.Flags().String("name", "", "Account name")

	rootCmd := &cobra.Command{Use: "app"}
	rootCmd.AddCommand(CreateAccountCmd, GetAccountCmd, UpdateAmountCmd, DeleteAccountCmd, ChangeAccountNameCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
