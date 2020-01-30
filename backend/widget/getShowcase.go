package widget

import (
	"context"

	"github.com/sea350/ustart_micro/backend/widget/widgetpb"
)

//LookupShowcase pulls the showcase of a user
func (wid *Widget) LookupShowcase(ctx context.Context, req *widgetpb.LookupShowcaseRequest) (*widgetpb.LookupShowcaseResponse, error) {

	_, wigs, err := wid.strg.GetShowcase(ctx, req.UID)
	if err != nil {
		return nil, err
	}

	widPoints := []*widgetpb.Widget{}

	for _, wig := range wigs {
		widPoints = append(widPoints, &wig)
	}

	return &widgetpb.LookupShowcaseResponse{WidgetData: widPoints}, nil
}
