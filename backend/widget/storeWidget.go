package widget

import (
	"context"
	"time"

	"github.com/sea350/ustart_micro/backend/widget/widgetpb"
)

//StoreWidget adds a new widget to database
func (wid *Widget) StoreWidget(ctx context.Context, req *widgetpb.StoreRequest) (*widgetpb.StoreResponse, error) {

	var index int
	index = 0

	references := []*widgetpb.Reference{
		&widgetpb.Reference{
			ReferenceID:    req.OwnerID,
			Classification: "owner",
		},
	}

	if req.Showcase {
		//there is only concurrency sensitivity when working with showcase widgets
		widgetLock.Lock()
		defer widgetLock.Unlock()
		i, _, err := wid.strg.GetShowcase(ctx, req.OwnerID)
		if err != nil {
			return nil, err
		}
		index = i + 1

	} else if req.ReplyID != "" {
		references = append(
			references,
			&widgetpb.Reference{
				ReferenceID:    req.ReplyID,
				Classification: "replyto",
			},
		)
	}

	id, err := wid.strg.StoreWidget(
		ctx,
		index,
		req.Body,
		time.Now().Format(wid.timeFormat),
		references,
	)
	if err != nil {
		return nil, err
	}

	return &widgetpb.StoreResponse{WidgetID: id}, nil
}
