package elastic

import (
	"context"
	//"time"
	//"github.com/sea350/ustart_mono/backend/widget/wigetspb"
)

//Register creates a new ES document for a new registering user
func (estor *ElasticStore) StoreWidget(ctx context.Context, widgetID string) error {

	//Lock just to make sure no two people can sign up with the same email at the same time
	newWidgetLock.Lock()
	defer newWidgetLock.Unlock()

	//Will continue this later
}
