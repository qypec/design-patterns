package torrentserver

import (
	"errors"
	"fmt"
	"math/rand"
)

type accessor interface {
	PrepareToDistribution(filename string) (err error)
	GetName() (name string)
}

// Uploader connects and uploads the file
type Uploader interface {
	Upload(file accessor) (err error)
}

type torrentServer struct {
	address []byte
}

// Upload uploads a file to the server at its address
func (s *torrentServer)Upload(file accessor) (err error) {
	err = s.ping()
	if err != nil {
		return
	}

	fmt.Printf("%v: file has been successfully uploaded to the server %v\n", file.GetName(), string(s.address))
	return
}

func (s *torrentServer)ping() (err error) {
	fmt.Printf("Ping %v...\n", string(s.address))

	serverRequestStatus := rand.Intn(2)
	if serverRequestStatus == 0 {
		err = errors.New(fmt.Sprintf("server %v status - KO\n", string(s.address)))
		return
	}
	fmt.Println("The connection to the server is established: server_status - OK")
	return
}

// NewUploader creates an instance of the Uploader
func NewUploader(serverAddress []byte) Uploader {
	return &torrentServer{
		address: serverAddress,
	}
}
