package user_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/thinhlvv/resource-management/handler/user"
	"github.com/thinhlvv/resource-management/model"
	"github.com/thinhlvv/resource-management/testhelper"
)

func TestRepository(t *testing.T) {
	db := testhelper.NewDB()
	repo := user.NewRepo(db)
	u := model.User{
		Email:          "test@gmail.com",
		HashedPassword: "123456789",
		Role:           model.RoleUser,
		Quota:          model.UnlimitedQuota.Int(),
	}

	Convey("Repository Feature: User repository", t, func() {
		newID, err := repo.CreateUser(u)
		So(err, ShouldBeNil)

		u.ID = newID

		createdUser, err := repo.GetUserByEmail(u.Email)
		So(err, ShouldBeNil)
		So(createdUser, ShouldResemble, &u)

		Reset(func() {
			testhelper.RemoveUser(db, newID, "")
			db.Close()
		})
	})

}
