package torrentserver

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/design-patterns/facade/pkg/api/v1"
)

type accessor interface {
	PrepareToDistribution(filename string) (err error)
	GetName() (name string)
}

// Uploader connects and uploads the file
type Uploader interface {
	Upload(filename string) (err error)
}

type torrentServer struct {
	address []byte
}

// Upload uploads a file to the server at its address
func (s *torrentServer)Upload(filename string) (err error) {
	err = s.ping()
	if err != nil {
		return
	}

	fmt.Println(filename + v1.FileUploadedOnServerMsg)
	return
}

func (s *torrentServer)ping() (err error) {
	fmt.Println(v1.PingMsg + string(s.address))

	serverRequestStatus := rand.Intn(2)
	if serverRequestStatus == 0 {
		err = errors.New(v1.ServerStatusKO)
		return
	}
	fmt.Println(v1.ServerStatusOK)
	return
}

// NewUploader creates an instance of the Uploader
func NewUploader(serverAddress []byte) Uploader {
	return &torrentServer{
		address: serverAddress,
	}
}
