package facade

import (
	"testing"

	"github.com/design-patterns/facade/pkg/torrentfile"
	"github.com/design-patterns/facade/pkg/torrentserver"
	"github.com/stretchr/testify/assert"
)

const (
	filename                        = "test.torrent"
	methodNameTurnDistribution      = "TurnDistribution"
	methodNameGetName               = "GetName"
	methodNamePrepareToDistribution = "PrepareToDistribution"
	methodNameUpload                = "Upload"
	unexpectedErrorMsg              = "unexpected error:"
)

func TestFacade(t *testing.T) {
	t.Run(methodNameTurnDistribution, func(t *testing.T) {
		torrentFileMock := new(torrentfile.MockTorrentFile)

		torrentFileMock.On(methodNamePrepareToDistribution, filename).Return(nil).Once()
		torrentFileMock.On(methodNameGetName).Return(filename).Once()

		torrentServerMock := new(torrentserver.MockTorrentServer)
		torrentServerMock.On(methodNameUpload, filename).Return(nil).Once()

		client := NewTorrentClient(torrentServerMock, torrentFileMock)
		err := client.TurnDistribution(filename)
		assert.NoError(t, err, unexpectedErrorMsg, err)
	})
}
