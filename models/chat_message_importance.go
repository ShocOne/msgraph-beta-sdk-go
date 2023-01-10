package models
import (
    "errors"
)
// 
type ChatMessageImportance int

const (
    NORMAL_CHATMESSAGEIMPORTANCE ChatMessageImportance = iota
    HIGH_CHATMESSAGEIMPORTANCE
    URGENT_CHATMESSAGEIMPORTANCE
)

func (i ChatMessageImportance) String() string {
    return []string{"normal", "high", "urgent"}[i]
}
func ParseChatMessageImportance(v string) (interface{}, error) {
    result := NORMAL_CHATMESSAGEIMPORTANCE
    switch v {
        case "normal":
            result = NORMAL_CHATMESSAGEIMPORTANCE
        case "high":
            result = HIGH_CHATMESSAGEIMPORTANCE
        case "urgent":
            result = URGENT_CHATMESSAGEIMPORTANCE
        default:
            return 0, errors.New("Unknown ChatMessageImportance value: " + v)
    }
    return &result, nil
}
func SerializeChatMessageImportance(values []ChatMessageImportance) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
