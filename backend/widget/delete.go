package widget

import (
	"context"

	"github.com/sea350/ustart_micro/backend/widget/widgetpb"
)

//Delete adds a new widget to database
func (wid *Widget) Delete(ctx context.Context, req *widgetpb.DeleteRequest) (*widgetpb.DeleteResponse, error) {

	//TODO:
	//All

	return &widgetpb.DeleteResponse{}, nil
}
