package types

import (
	"crypto/md5"
	"encoding/hex"
)

type BotDataKey struct {
	BotId          string
	UserId         string
	ConversationId string
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