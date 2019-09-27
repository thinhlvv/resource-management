package user_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/thinhlvv/resource-management/config"
	"github.com/thinhlvv/resource-management/handler/user"
	"github.com/thinhlvv/resource-management/model"
	"github.com/thinhlvv/resource-management/pkg"
	"github.com/thinhlvv/resource-management/testhelper"
)

var (
	signupJSON = `{
 	"email": "test@gmail.com",
 	"password": "12345678",
	"role": 1
 }`
	invalidSignupJSON = `{
 	"email": "testgmail.com",
 	"password": "12345678",
	"role": 1
 }`
	invalidSigninJSON = `{
 	"email": "testgmail.com",
 	"password": "12345678"
 }`
	signinJSON = `{
 	"email": "test@gmail.com",
 	"password": "12345678"
 }`
)

func TestControllerCreateArticle(t *testing.T) {

	// Setup
	cfg := config.New()
	db := testhelper.NewDB()
	hasher := pkg.NewHasher(pkg.NewConfig())
	signer := config.MustInitJWTSigner(cfg)
	requestValidator := pkg.NewRequestValidator()
	app := model.App{
		DB:               db,
		RequestValidator: requestValidator,
		Hasher:           hasher,
		JWTSigner:        signer,
	}

	e := echo.New()

	userHandler := user.New(app)
	userHandler.RegisterHTTPRouter(e)

	Convey("Controller User: User endpoints controller", t, func() {
		Convey("signup controller successfully", func() {
			req := httptest.NewRequest(http.MethodPost, "/user/signup", strings.NewReader(signupJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := userHandler.Signup(c)
			So(err, ShouldBeNil)
			So(rec.Code, ShouldEqual, http.StatusOK)

			Convey("login controller successfully", func() {
				req := httptest.NewRequest(http.MethodPost, "/user/login", strings.NewReader(signinJSON))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)

				err := userHandler.Login(c)
				So(err, ShouldBeNil)
				So(rec.Code, ShouldEqual, http.StatusOK)
			})

		})

		Convey("signup controller failed with invalid params", func() {
			req := httptest.NewRequest(http.MethodPost, "/user/signup", strings.NewReader(invalidSignupJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := userHandler.Signup(c)
			So(err, ShouldBeNil)
			So(rec.Code, ShouldEqual, http.StatusBadRequest)
		})
		Convey("signup controller failed with empty JSON", func() {
			req := httptest.NewRequest(http.MethodPost, "/user/signup", strings.NewReader(""))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := userHandler.Signup(c)
			So(err, ShouldBeNil)
			So(rec.Code, ShouldEqual, http.StatusBadRequest)
		})

		Convey("login controller failed with empty JSON", func() {
			req := httptest.NewRequest(http.MethodPost, "/user/login", strings.NewReader(""))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := userHandler.Login(c)
			So(err, ShouldBeNil)
			So(rec.Code, ShouldEqual, http.StatusBadRequest)
		})

		Reset(func() {
			testhelper.RemoveUser(db, 0, "test@gmail.com")
			db.Close()
		})

	})
}
