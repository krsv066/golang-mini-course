package main

import (
	"context"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	ProtocolBuffer "proj/hw3/proto"
	"time"
)

var (
	CreateAccountClient = &cobra.Command{
		Use: "create",
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")
			if name == "" {
				log.Fatalf("Name is required.")
				return
			}

			client, conn := getClient()
			defer conn.Close()

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			req := &ProtocolBuffer.CreateRequest{Name: name}
			_, err := client.CreateServer(ctx, req)
			if err != nil {
				log.Fatalf("Could not create account.")
			}

			log.Println("Account created.")
		},
	}

	GetAccountClient = &cobra.Command{
		Use: "get",
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")
			if name == "" {
				log.Fatalf("Account name is required")
				return
			}

			client, conn := getClient()
			defer conn.Close()

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			req := &ProtocolBuffer.GetRequest{Name: name}
			res, err := client.GetServer(ctx, req)
			if err != nil {
				log.Fatalf("Could not get account.")
			}

			log.Printf("Account: %s, balance: %f", res.Name, res.Balance)
		},
	}

	RenameAccountClient = &cobra.Command{
		Use: "rename",
		Run: func(cmd *cobra.Command, args []string) {
			oldName, _ := cmd.Flags().GetString("name")
			newName, _ := cmd.Flags().GetString("new_name")
			if newName == "" || oldName == "" {
				log.Fatalf("Old and new names are required")
				return
			}

			client, conn := getClient()
			defer conn.Close()

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			req := &ProtocolBuffer.RenameRequest{Name: oldName, NewName: newName}
			_, err := client.RenameServer(ctx, req)
			if err != nil {
				log.Fatalf("could not update name.")
			}

			log.Println("Account name updated.")
		},
	}

	UpdateBalanceClient = &cobra.Command{
		Use: "update",
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")
			balance, _ := cmd.Flags().GetFloat64("balance")
			if name == "" {
				log.Fatalf("Account name is required.")
				return
			}
			if balance < 0 {
				log.Fatalf("Balance must be >= 0")
				return
			}

			client, conn := getClient()
			defer conn.Close()

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			req := &ProtocolBuffer.UpdateBalanceRequest{Name: name, Balance: balance}
			res, err := client.UpdateBalanceServer(ctx, req)
			if err != nil {
				log.Fatalf("Could not update balance.")
			}

			log.Printf("Account updated: %s, balance: %f", res.Name, res.Balance)
		},
	}

	DeleteAccountClient = &cobra.Command{
		Use: "delete",
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")
			if name == "" {
				log.Fatalf("Account name is required.")
				return
			}

			client, conn := getClient()
			defer conn.Close()

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			req := &ProtocolBuffer.DeleteRequest{Name: name}
			_, err := client.DeleteServer(ctx, req)
			if err != nil {
				log.Fatalf("Could not delete account.")
			}

			log.Printf("Deleted: %s", name)
		},
	}
)

func getClient() (ProtocolBuffer.BankAccountServiceClient, *grpc.ClientConn) {
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed connection.")
	}
	client := ProtocolBuffer.NewBankAccountServiceClient(conn)

	return client, conn
}

func main() {
	CreateAccountClient.Flags().String("name", "", "Account name")
	GetAccountClient.Flags().String("name", "", "Account name")
	RenameAccountClient.Flags().String("name", "", "Account old name")
	RenameAccountClient.Flags().String("new_name", "", "Account new name")
	UpdateBalanceClient.Flags().String("name", "", "Account name")
	UpdateBalanceClient.Flags().Float64("balance", 0, "Balance for adding funds to your account using")
	DeleteAccountClient.Flags().String("name", "", "Account name")

	var rootCmd = &cobra.Command{Use: "app"}

	rootCmd.AddCommand(
		CreateAccountClient,
		GetAccountClient,
		RenameAccountClient,
		UpdateBalanceClient,
		DeleteAccountClient,
	)

	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
