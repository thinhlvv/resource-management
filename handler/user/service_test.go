package user_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/thinhlvv/resource-management/config"
	"github.com/thinhlvv/resource-management/handler/user"
	"github.com/thinhlvv/resource-management/model"
	"github.com/thinhlvv/resource-management/pkg"
	"github.com/thinhlvv/resource-management/testhelper"
)

func TestService(t *testing.T) {

	cfg := config.New()
	db := testhelper.NewDB()
	hasher := pkg.NewHasher(pkg.NewConfig())
	signer := config.MustInitJWTSigner(cfg)
	app := model.App{
		DB:        db,
		Hasher:    hasher,
		JWTSigner: signer,
	}

	repo := user.NewRepo(db)
	service := user.NewService(repo, app)

	Convey("Service User: User services", t, func() {

		Convey("user signup", func() {
			req := user.SignupReq{
				Email:    "test@gmail.com",
				Password: "12345678",
				Role:     model.RoleUser.Int(),
			}
			token, err := service.Signup(nil, req)

			So(err, ShouldBeNil)
			So(token, ShouldNotEqual, "")

			Convey("user login", func() {
				req := user.LoginRequest{
					Email:    "test@gmail.com",
					Password: "12345678",
				}
				token, err := service.Login(nil, req)

				So(err, ShouldBeNil)
				So(token, ShouldNotEqual, "")
			})
		})

		Reset(func() {
			testhelper.RemoveUser(db, 0, "test@gmail.com")
			db.Close()
		})
	})
}
