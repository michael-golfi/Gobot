package types

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"errors"
)

const (
	BotConversationData = 0
	BotUserData = 1
	BotPerUserInConversationData = 2
)

type BotDataKey struct {
	BotId          string
	UserId         string
	ConversationId string
}

type BotDataEntry struct {
	BotConversationData          interface{}
	BotPerUserInConversationData interface{}
	BotUserData                  interface{}
}

func (b *BotDataKey) GetKey(keyType int) (string, error) {
	switch (keyType) {
	case BotConversationData:
		return fmt.Sprintf("conversation:%s:%s", b.BotId, b.ConversationId), nil
	case BotUserData:
		return fmt.Sprintf("user:%s:%s", b.BotId, b.UserId), nil
	case BotPerUserInConversationData:
		return fmt.Sprintf("perUserInConversation:%s:%s:%s", b.BotId, b.UserId, b.ConversationId), nil
	default:
		return nil, errors.New("Could not get key for unknown type of data")
	}
}

func AreEqual(key1, key2 BotDataKey) bool {
	return key1.BotId == key2.BotId && key1.UserId == key2.UserId && key1.ConversationId == key2.ConversationId
}

func GetHashCode(key BotDataKey) *string {
	hash := GetMD5Hash(key.UserId) + GetMD5Hash(key.ConversationId) + GetMD5Hash(key.BotId)
	return &hash
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}