package uploader

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/sea350/ustart_micro/backend/uploader/uploaderpb"
)

// UploadAvatar wraps backend/uploader/uploadAvatar.go
func (rapi *RESTAPI) UploadAvatar(w http.ResponseWriter, req *http.Request) {
	regCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if !setCORS(&w, req) {
		return
	}

	uploadReq := &uploaderpb.UploadAvatarRequest{}

	if strings.Contains(req.Header.Get("Content-type"), "application/json") {
		req.Header.Set("Content-Type", "application/json")
		json.NewDecoder(req.Body).Decode(uploadReq)
	} else {
		req.ParseForm()
		uploadReq.Base64 = req.Form.Get("base64")
		logger.Println(uploadReq.Base64)
		uploadReq.OldImageLink = req.Form.Get("oldimagelink")
		uploadReq.UUID = req.Form.Get("uuid")
	}

	ret := make(map[string]interface{})

	resp, err := rapi.uploader.UploadAvatar(regCtx, uploadReq)
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
		logger.Println("Problem marshaling return data", err)
	}

	fmt.Fprintln(w, string(data))
}
