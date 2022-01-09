package notify

import (
	"bytes"
	"fmt"
	"io"
	"net/textproto"
	"time"

	"github.com/go-playground/validator/v10"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// MailSender 投递消息器
type MailSender interface {
	Send(*SendMailRequest) error
}

// NewSendMailRequest todo
func NewSendMailRequest() *SendMailRequest {
	return &SendMailRequest{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

// SendMailRequest todo
type SendMailRequest struct {
	To      string `json:"to,omitempty" validate:"required"`
	Cc      string `json:"cc,omitempty"`
	Subject string `json:"subject,omitempty" validate:"required"`
	Content string `json:"content,omitempty" validate:"required"`

	buffer *bytes.Buffer `json:"-"`
}

// Validate todo
func (req *SendMailRequest) Validate() error {
	return validate.Struct(req)
}

// PrepareBody todo
func (req *SendMailRequest) PrepareBody(from string) ([]byte, error) {
	// 设置邮件Header
	headers := textproto.MIMEHeader{}
	headers.Set("From", from)
	headers.Set("To", req.To)
	headers.Set("Cc", req.Cc)

	headers.Set("Subject", req.Subject)
	headers.Set("Date", time.Now().Format(time.RFC1123Z))
	headers.Set("MIME-Version", "1.0")
	headers.Set("Content-Type", "text/html; charset=UTF-8;")
	if err := headerToBytesSeq(req.buffer, headers); err != nil {
		return nil, fmt.Errorf("Failed to render message headers: %s", err)
	}
	req.buffer.WriteString("\r\n")
	req.buffer.WriteString(req.Content)
	return req.buffer.Bytes(), nil
}

// hack here 原来的headerToBytes有个严重bug，textproto.MIMEHeader是个map，for range出来是无序的
func headerToBytesSeq(w io.Writer, t textproto.MIMEHeader) (err error) {
	var seq = []string{`From`, `To`, `Cc`, `Subject`, `Disposition-Notification-To`, `Date`, `MIME-Version`, `Content-Type`}
	for _, mapKey := range seq {
		// Write the header key
		if `` == t.Get(mapKey) {
			continue
		}
		_, err = fmt.Fprintf(w, "%s:", mapKey)
		if err != nil {
			return err
		}
		_, err = fmt.Fprintf(w, " %s\r\n", t.Get(mapKey))
		if err != nil {
			return err
		}
	}
	return nil
}
