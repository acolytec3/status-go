package gethbridge

import (
	"github.com/status-im/status-go/eth-node/types"
	whisper "github.com/status-im/whisper/whisperv6"
)

// NewGethMailServerResponseWrapper returns a types.MailServerResponse object that mimics Geth's MailServerResponse
func NewGethMailServerResponseWrapper(mailServerResponse *whisper.MailServerResponse) *types.MailServerResponse {
	if mailServerResponse == nil {
		panic("mailServerResponse should not be nil")
	}

	return &types.MailServerResponse{
		LastEnvelopeHash: types.Hash(mailServerResponse.LastEnvelopeHash),
		Cursor:           mailServerResponse.Cursor,
		Error:            mailServerResponse.Error,
	}
}
