package torrentfile

import (
	"testing"

	"github.com/design-patterns/facade/pkg/api/v1"
	"github.com/stretchr/testify/assert"
)

const (
	methodNamePrepareToDistribution = "PrepareToDistribution"
	dummySize                       = 64
	filename                        = "test.torrent"
	badFilename                     = "bad.torrent"
	fileSizeEndGame                 = v1.MaxFileSizeForEndGame
	expectedModeEndGame             = v1.Endgame
	fileSizeSeparation              = v1.MaxFileSizeForEndGame + 1
	expectedModeSeparation          = v1.Separation
	unexpectedErrorMsg              = "unexpected error:"
)

func TestTorrentFileBadName(t *testing.T) {
	t.Run(methodNamePrepareToDistribution, func(t *testing.T) {
		f := NewAccessor(filename, dummySize)
		err := f.PrepareToDistribution(badFilename)
		assert.Error(t, err)
	})
}

func TestTorrentFileEndGame(t *testing.T) {
	t.Run(methodNamePrepareToDistribution, func(t *testing.T) {
		f := NewAccessor(filename, fileSizeEndGame)
		err := f.PrepareToDistribution(filename)
		actualMode := f.GetMode()
		assert.NoError(t, err, unexpectedErrorMsg, err)
		assert.EqualValues(t, expectedModeEndGame, actualMode)
	})
}

func TestTorrentFileSeparation(t *testing.T) {
	t.Run(methodNamePrepareToDistribution, func(t *testing.T) {
		f := NewAccessor(filename, fileSizeSeparation)
		err := f.PrepareToDistribution(filename)
		actualMode := f.GetMode()
		assert.NoError(t, err, unexpectedErrorMsg, err)
		assert.EqualValues(t, expectedModeSeparation, actualMode)
	})
}
