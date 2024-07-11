package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"proj/hw3/grpc/dto"
	ProtocolBuffer "proj/hw3/proto"
)

type server struct {
	storage *dto.AccountStorage
	ProtocolBuffer.UnimplementedBankAccountServiceServer
}

func (s *server) CreateServer(ctx context.Context, req *ProtocolBuffer.CreateRequest) (*ProtocolBuffer.CreateResponse, error) {
	name := req.GetName()
	if name == "" {
		return nil, fmt.Errorf(name)
	}
	s.storage.CreateAccount(&dto.Account{Name: name, Balance: 0})

	return &ProtocolBuffer.CreateResponse{Name: name, Balance: 0}, nil

}

func (s *server) GetServer(ctx context.Context, req *ProtocolBuffer.GetRequest) (*ProtocolBuffer.GetResponse, error) {
	name := req.GetName()
	account, err := s.storage.GetAccount(name)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return &ProtocolBuffer.GetResponse{Name: account.Name, Balance: account.Balance}, nil
}

func (s *server) RenameServer(ctx context.Context, req *ProtocolBuffer.RenameRequest) (*ProtocolBuffer.RenameResponse, error) {
	oldName := req.GetName()
	newName := req.GetNewName()
	err := s.storage.RenameAccount(newName, oldName)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return &ProtocolBuffer.RenameResponse{Name: newName}, nil
}

func (s *server) UpdateBalanceServer(ctx context.Context, req *ProtocolBuffer.UpdateBalanceRequest) (*ProtocolBuffer.UpdateBalanceResponse, error) {
	name := req.GetName()
	balance := req.GetBalance()
	account, err := s.storage.GetAccount(name)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	account.Balance = balance

	return &ProtocolBuffer.UpdateBalanceResponse{Name: account.Name, Balance: account.Balance}, nil
}

func (s *server) DeleteServer(ctx context.Context, req *ProtocolBuffer.DeleteRequest) (*ProtocolBuffer.DeleteResponse, error) {
	name := req.GetName()
	err := s.storage.DeleteAccount(name)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return &ProtocolBuffer.DeleteResponse{Name: name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	ProtocolBuffer.RegisterBankAccountServiceServer(s, &server{
		storage: dto.NewAS(),
	})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed")
	}
}
