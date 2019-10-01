package resource_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/thinhlvv/resource-management/handler/resource"
	"github.com/thinhlvv/resource-management/handler/user"
	"github.com/thinhlvv/resource-management/model"
	"github.com/thinhlvv/resource-management/testhelper"
)

func TestService(t *testing.T) {
	db := testhelper.NewDB()

	resourceRepo := resource.NewRepo(db)
	service := resource.NewService(resourceRepo)

	Convey("Service Resource: Resource services", t, func() {
		userRepo := user.NewRepo(db)
		u := model.User{
			Email:          "test@gmail.com",
			HashedPassword: "123456789",
			Role:           model.RoleUser,
		}
		userID, err := userRepo.CreateUser(u)
		So(err, ShouldBeNil)

		req := resource.CreateReq{
			Name:   "Unique Name",
			UserID: userID,
		}
		actualResource, err := service.CreateResource(req)
		So(err, ShouldBeNil)
		So(actualResource.Name, ShouldEqual, req.Name)
		So(actualResource.UserID, ShouldEqual, req.UserID)
		So(actualResource.ID, ShouldNotEqual, 0)

		Reset(func() {
			testhelper.RemoveResource(db, actualResource.ID, "")
			testhelper.RemoveUser(db, userID, "")
			db.Close()
		})
	})
}
