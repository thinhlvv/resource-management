package resource_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/thinhlvv/resource-management/handler/resource"
	"github.com/thinhlvv/resource-management/handler/user"
	"github.com/thinhlvv/resource-management/model"
	"github.com/thinhlvv/resource-management/testhelper"
)

func TestRepository(t *testing.T) {
	db := testhelper.NewDB()
	repo := resource.NewRepo(db)
	userRepo := user.NewRepo(db)

	Convey("Repository Feature: Resource repository", t, func() {
		u := model.User{
			Email:          "test@gmail.com",
			HashedPassword: "123456789",
			Role:           model.RoleUser,
		}
		userID, err := userRepo.CreateUser(u)
		So(err, ShouldBeNil)

		resource := model.Resource{
			Name:   "unique resource name",
			UserID: userID,
		}
		resourceID, err := repo.CreateResource(resource)
		So(err, ShouldBeNil)

		Reset(func() {
			testhelper.RemoveResource(db, resourceID, "")
			testhelper.RemoveUser(db, userID, "")
			db.Close()
		})
	})

}
