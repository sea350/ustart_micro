package backend

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/sea350/ustart_micro/backend/uploader/uploaderpb"

	"github.com/sea350/ustart_micro/backend/backendpb"
	"github.com/sea350/ustart_micro/backend/profile/profilepb"
)

//ProfileChangeIcon updates a single profile's icon
func (s *Server) ProfileChangeIcon(ctx context.Context, req *backendpb.ProfileChangeIconRequest) (*backendpb.ProfileChangeIconResponse, error) {
	//Get old image
	iconRes, err := (*s.profileClient).Get(ctx, &profilepb.GetRequest{
		UUID: req.UUID,
	})
	if err != nil {
		return nil, err
	}
	if iconRes == nil {
		logger.Println("this shouldnt happen")
		return nil, err
	}

	//Need to upload new image data
	uplRes, err := (*s.uploaderClient).UploadAvatar(ctx, &uploaderpb.UploadAvatarRequest{
		Base64:       req.Image,
		UUID:         req.UUID,
		OldImageLink: iconRes.UserProfile.Avatar,
	})
	if err != nil {
		return nil, err
	}
	if uplRes == nil {
		logger.Println("this shouldnt happen")
		return nil, err
	}

	//Then update profile with the new link
	profileRes, err := (*s.profileClient).UpdateIcon(ctx, &profilepb.UpdateIconRequest{
		UUID:        req.UUID,
		NewIconLink: "",
	})
	if err != nil {
		return nil, err
	}
	if profileRes == nil {
		logger.Println("this shouldnt happen")
		return nil, err
	}

	//TODO
	//we may need to pull some data from project etc.

	return &backendpb.ProfileChangeIconResponse{}, nil
}

//EXPERIMENTAL

//ProfileChangeIconHTTP is an http wrapper for the ProfileChangeIcon function
func (s *Server) ProfileChangeIconHTTP(w http.ResponseWriter, r *http.Request) {
	if !setCORS(&w, r) {
		return
	}

	uid, err := s.sesh.UID(r)
	if err != nil {
		logger.Println("Session error | err: ", err)
		data := packageResponse(nil, err)
		fmt.Fprintln(w, string(data))
		return
	}

	req := &backendpb.ProfileChangeIconRequest{}

	if strings.Contains(r.Header.Get("Content-type"), "application/json") {
		r.Header.Set("Content-Type", "application/json")
		json.NewDecoder(r.Body).Decode(req)
	} else {
		r.ParseForm()
		req.Image = r.Form.Get("image")
	}

	resp, err := s.ProfileChangeIcon(r.Context(), req)
	if err != nil {
		logger.Println("UUID: "+uid+" | err: ", err)
	}

	data := packageResponse(resp, err)

	fmt.Fprintln(w, string(data))

}
