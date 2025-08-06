package listeners

import (
	"github.com/google/uuid"
	"github.com/goravel/framework/contracts/event"
	"github.com/goravel/framework/facades"
	"github.com/spf13/cast"
	"goravel-blog/app/models"
	"time"
)

type SendCategoryChangedNotification struct {
}

func (receiver *SendCategoryChangedNotification) Signature() string {
	return "send_category_changed_notification"
}

func (receiver *SendCategoryChangedNotification) Queue(args ...any) event.Queue {
	return event.Queue{
		Enable:     true,
		Connection: "database",
		Queue:      "default",
	}
}

func (receiver *SendCategoryChangedNotification) Handle(args ...any) error {
	id, err := cast.ToStringE(args[0])
	if err != nil {
		facades.Log().Errorf("Cast id failed: %v", err)
	}
	name, err := cast.ToStringE(args[1])
	if err != nil {
		facades.Log().Errorf("Cast name failed: %v", err)
	}
	ntype, err := cast.ToStringE(args[2])
	if err != nil {
		facades.Log().Errorf("Cast type failed: %v", err)
	}
	uid, _ := uuid.NewUUID()
	time.Sleep(10 * time.Second)
	err = facades.Orm().Query().Create(&models.Notification{
		Id:             uid.String(),
		Type:           "category:" + ntype,
		Data:           "Category [" + id + "] " + name + " has been " + ntype,
		NotifiableType: "local",
		NotifiableId:   0,
	})
	if err != nil {
		facades.Log().Errorf("Create notification failed: %v", err)
	}
	return nil
}
