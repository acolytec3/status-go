package protocol

import (
	"github.com/pkg/errors"

	v1protocol "github.com/status-im/status-go/protocol/v1"
)

// HandleMembershipUpdate updates a Chat instance according to the membership updates.
// It retrieves chat, if exists, and merges membership updates from the message.
// Finally, the Chat is updated with the new group events.
func HandleMembershipUpdate(chat *Chat, m *v1protocol.MembershipUpdateMessage) (*Chat, error) {
	if chat == nil {
		group, err := v1protocol.NewGroupWithMembershipUpdates(m.ChatID, m.Updates)
		if err != nil {
			return nil, err
		}
		newChat := createGroupChat()
		newChat.updateChatFromProtocolGroup(group)
		chat = &newChat
	} else {
		existingGroup, err := newProtocolGroupFromChat(chat)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create a Group from Chat")
		}
		updateGroup, err := v1protocol.NewGroupWithMembershipUpdates(m.ChatID, m.Updates)
		if err != nil {
			return nil, errors.Wrap(err, "invalid membership update")
		}
		merged := v1protocol.MergeFlatMembershipUpdates(existingGroup.Updates(), updateGroup.Updates())
		newGroup, err := v1protocol.NewGroup(chat.ID, merged)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create a group with new membership updates")
		}
		chat.updateChatFromProtocolGroup(newGroup)
	}
	return chat, nil
}
