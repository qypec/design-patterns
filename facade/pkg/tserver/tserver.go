package tserver

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/design-patterns/facade/pkg/tfile"
)

// Uploader connects and uploads the file
type Uploader interface {
	Upload(file tfile.TorrentFile) error
}

type server struct {
	address []byte
}

func (s *server)Upload(file tfile.TorrentFile) (err error) {
	err = s.ping()
	if err != nil {
		return
	}

	fmt.Printf("%v: file has been successfully uploaded to the server %v\n", file.GetName(), string(s.address))
	return
}

func (s *server)ping() error {
	fmt.Printf("Ping %v...\n", string(s.address))

	serverRequestStatus := rand.Intn(2)
	if serverRequestStatus == 0 {
		return errors.New(fmt.Sprintf("server %v status - KO\n", string(s.address)))
	}
	fmt.Println("The connection to the server is established: server_status - OK")
	return nil
}

func NewUploader(serverAddress string) Uploader {
	return &server{
		address: []byte(serverAddress),
	}
}
