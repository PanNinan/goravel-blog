package commands

import (
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
	"github.com/goravel/framework/facades"
	"goravel-blog/app/models"
)

type ClearLinks struct {
}

// Signature The name and signature of the console command.
func (r *ClearLinks) Signature() string {
	return "app:clear-links"
}

// Description The console command description.
func (r *ClearLinks) Description() string {
	return "Command description"
}

// Extend The console command extend.
func (r *ClearLinks) Extend() command.Extend {
	return command.Extend{Category: "app"}
}

// Handle Execute the console command.
func (r *ClearLinks) Handle(ctx console.Context) error {
	facades.Log().Info("Clear links...")
	_, err := facades.Orm().Query().Model(&models.Link{}).OrderByDesc("id").Where("1=1").Limit(10).Delete()
	if err != nil {
		facades.Log().Errorf("Clear links failed: %v", err)
	}
	return nil
}
