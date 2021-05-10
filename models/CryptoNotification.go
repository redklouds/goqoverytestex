package models

type CryptoNotification struct {
	UserId           string `bson:"omitempty"`
	Symbol           string `json:"symbol"`
	Price            float32
	Direction        PriceDirectionType
	NotificationInfo NotificationInfo `json:"notificationinfo"`
}
type NotificationInfo struct {
	Type NotificationType
	Data string `json:"data"`
}

type NotificationType int
type PriceDirectionType int

const (
	UP PriceDirectionType = iota
	DOWN
)
const (
	SMS NotificationType = iota
	EMAIL
	PUSH
)
