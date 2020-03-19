package sender

type Sender interface {
	SendMessage(msg, recipient, recipientType string, opt ...interface{}) error
	SendCustomMessage(msg interface{}, msgType, recipient, recipientType string, opt ...interface{}) error
	SendFile(fileBytes []byte, recipient, recipientType string, opt ...interface{}) error
	SendFileByPath(filePath, recipient, recipientType string, opt ...interface{}) error
}
