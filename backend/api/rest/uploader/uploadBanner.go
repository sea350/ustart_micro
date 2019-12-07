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

// UploadBanner wraps backend/uploader/uploadBanner.go
func (rapi *RESTAPI) UploadBanner(w http.ResponseWriter, req *http.Request) {
	regCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if !setCORS(&w, req) {
		return
	}

	uploaderReq := &uploaderpb.UploadBannerRequest{}

	if strings.Contains(req.Header.Get("Content-type"), "application/json") {
		req.Header.Set("Content-Type", "application/json")
		json.NewDecoder(req.Body).Decode(uploaderReq)
	} else {
		req.ParseForm()
		uploaderReq.Base64 = req.Form.Get("base64")
		uploaderReq.UUID = req.Form.Get("uuid")
		uploaderReq.OldImageLink = req.Form.Get("oldimagelink")
	}

	ret := make(map[string]interface{})

	resp, err := rapi.uploader.UploadBanner(regCtx, uploaderReq)
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
		logger.Println("Problem martialing return data", err)
	}

	fmt.Fprintln(w, string(data))
}
