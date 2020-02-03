package widget

import (
	"context"

	"github.com/sea350/ustart_micro/backend/widget/widgetpb"
)

//Delete adds a new widget to database
func (wid *Widget) Delete(ctx context.Context, req *widgetpb.DeleteRequest) (*widgetpb.DeleteResponse, error) {
	//everything is concurrency sensitve here
	widgetLock.Lock()
	defer widgetLock.Unlock()

	wig, err := wid.strg.Lookup(ctx, req.WidgetID)
	if err != nil {
		return nil, err
	}

	//if the widget is a showcase widget special measures need to be taken
	if wig.Index != 0 {
		//find owner
		var owner string
		for _, ref := range wig.References {
			if ref.Classification == "owner" {
				owner = ref.ReferenceID
			}
		}
		//if theres no owner this is a critical problem
		if owner == "" {
			return nil, errStorageDiscrepancy
		}

		//get all widgets after the deleted one
		_, wigsAfter, err := wid.strg.GetAfter(ctx, owner, int(wig.Index))
		if err != nil {
			return nil, err
		}
		//decriment their index so that they fill the position the deleted widget will leave
		for _, wigAfter := range wigsAfter {
			newIndex := int(wigAfter.Index) - 1
			wid.strg.UpdateIndex(ctx, wigAfter.WidgetID, newIndex)
		}
	}

	//finally delete the wiget
	wid.strg.DeleteWidget(ctx, req.WidgetID)
	if err != nil {
		return nil, err
	}

	return &widgetpb.DeleteResponse{}, nil
}
