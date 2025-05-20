package facade

import (
	"context"
	
	"common/pkg/service/database_service"
)

func DB(c context.Context) database_service.Client {
	service, ok := c.Value(database_service.Token).(database_service.Client)
	if !ok {
		panic(database_service.Token + " not found in context")
	}
	return service
}
