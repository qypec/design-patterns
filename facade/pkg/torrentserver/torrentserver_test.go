package torrentserver

import (
	"testing"

	"github.com/design-patterns/facade/pkg/api/v1"
	"github.com/stretchr/testify/assert"
)

const (
	filename           = "test.torrent"
	methodNameUpload   = "upload"
	rightServerAddress = v1.RightAddress
	badServerAddress   = "8.8.8.8"
	unexpectedErrorMsg = "unexpected error:"
)

func TestTorrentRightServer(t *testing.T) {
	t.Run(methodNameUpload, func(t *testing.T) {
		f := NewUploader([]byte(rightServerAddress))
		err := f.Upload(filename)
		assert.NoError(t, err, unexpectedErrorMsg, err)
	})
}

func TestTorrentBadServer(t *testing.T) {
	t.Run(methodNameUpload, func(t *testing.T) {
		f := NewUploader([]byte(badServerAddress))
		err := f.Upload(filename)
		assert.Error(t, err)
	})
}
