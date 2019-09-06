package emailer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sea350/ustart_micro/backend/emailer/resources"
)

// Authenticate wraps backend/auth/authenticate.go
func (rapi *RESTAPI) SendMail(w http.ResponseWriter, req *http.Request) {
	regCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if !setCORS(&w, req) {
		return
	}

	req.ParseForm()
	sendTo := req.Form.Get("email")
	subject := req.Form.Get("subject")
	body := req.Form.Get("body")

	ret := make(map[string]interface{})

	sendToReq := &emailerpb.EmailerRequest{
		From:    resources.SenderAddress,
		To:      sendTo,
		Subject: subject,
		Body:    body,
	}

	resp, err := rapi.emailer.SendMail(regCtx, sendToReq)
	if resp != nil {
		ret["response"] = resp
	} else {
		ret["response"] = ""
	}
	if err != nil {
		ret["error"] = err.Error()
	} else {
		ret["error"] = ""
	}

	data, err := json.Marshal(ret)
	if err != nil {
		logger.Panic(err)
	}

	fmt.Fprintln(w, string(data))
}
