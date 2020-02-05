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

//ProfileChangeBanner updates a single profile's icon
func (s *Server) ProfileChangeBanner(ctx context.Context, req *backendpb.ProfileChangeBannerRequest) (*backendpb.ProfileChangeBannerResponse, error) {
	//Get old image
	iconRes, err := (*s.profileClient).Get(ctx, &profilepb.GetRequest{
		UUID: req.UUID,
	})
	if err != nil {
		return nil, err
	}
	if iconRes == nil {
		logger.Println("this shouldnt happen")
		return nil, errNilResponse
	}

	//Need to upload new image data
	uplRes, err := (*s.uploaderClient).UploadBanner(ctx, &uploaderpb.UploadBannerRequest{
		Base64:       req.Image,
		UUID:         req.UUID,
		OldImageLink: iconRes.UserProfile.Banner,
	})
	if err != nil {
		return nil, err
	}
	if uplRes == nil {
		logger.Println("this shouldnt happen")
		return nil, errNilResponse
	}

	//Then update profile with the new link
	profileRes, err := (*s.profileClient).UpdateBanner(ctx, &profilepb.UpdateBannerRequest{
		UUID:          req.UUID,
		NewBannerLink: uplRes.NewLink,
	})
	if err != nil {
		return nil, err
	}
	if profileRes == nil {
		logger.Println("this shouldnt happen")
		return nil, errNilResponse
	}

	//TODO
	//we may need to pull some data from project etc.

	return &backendpb.ProfileChangeBannerResponse{}, nil
}

//EXPERIMENTAL

//ProfileChangeBannerHTTP is an http wrapper for the ProfileChangeBanner function
func (s *Server) ProfileChangeBannerHTTP(w http.ResponseWriter, r *http.Request) {
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

	req := &backendpb.ProfileChangeBannerRequest{}

	if strings.Contains(r.Header.Get("Content-type"), "application/json") {
		r.Header.Set("Content-Type", "application/json")
		json.NewDecoder(r.Body).Decode(req)
	} else {
		r.ParseForm()
		req.Image = r.Form.Get("image")
	}

	resp, err := s.ProfileChangeBanner(r.Context(), req)
	if err != nil {
		logger.Println("UUID: "+uid+" | err: ", err)
	}

	data := packageResponse(resp, err)

	fmt.Fprintln(w, string(data))

}
