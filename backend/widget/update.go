package widget

import (
	"context"

	"github.com/sea350/ustart_micro/backend/widget/widgetpb"
)

//Update reindexes the body of a widget
func (wid *Widget) Update(ctx context.Context, req *widgetpb.UpdateRequest) (*widgetpb.UpdateResponse, error) {

	err := wid.strg.UpdateBody(ctx, req.WidgetID, req.NewBody)
	if err != nil {
		return nil, err
	}

	return &widgetpb.UpdateResponse{}, nil
}
