package shutdown

import "context"

var Ctx, Shutdown = context.WithCancel(context.Background())
