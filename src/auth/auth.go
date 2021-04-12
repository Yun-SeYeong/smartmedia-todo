package auth

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"todo/src/setting"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type JwtClaims struct {
	UserId string `json:"Id"`
	jwt.StandardClaims
}

type User struct {
	gorm.Model
	UserId   string `gorm:"primaryKey;unique"`
	Password string
	Email    string
}

var OauthConf *oauth2.Config
var OauthSignUpConf *oauth2.Config

func init() {
	fmt.Println("init")
	OauthSignUpConf = &oauth2.Config{
		ClientID:     "173533637091-9qk1vidiui0j9gk4v7iruecfls1el250.apps.googleusercontent.com",
		ClientSecret: "gwyuKtyndIzwhCydmhp1cPn9",
		RedirectURL:  "http://localhost:1323/signup/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
	OauthConf = &oauth2.Config{
		ClientID:     "173533637091-9qk1vidiui0j9gk4v7iruecfls1el250.apps.googleusercontent.com",
		ClientSecret: "gwyuKtyndIzwhCydmhp1cPn9",
		RedirectURL:  "http://localhost:1323/login/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}

}

func RandToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func GoogleLogin(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, OauthConf.AuthCodeURL(RandToken()))
}

func GoogleLoginCallback(c echo.Context) error {
	code := c.FormValue("code")
	if code != "" {
		fmt.Println("code: ", code)
	}

	token, err := OauthConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		return err
	}

	client := OauthConf.Client(oauth2.NoContext, token)
	userInfoResp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return err
	}
	defer userInfoResp.Body.Close()

	userInfo, err := ioutil.ReadAll(userInfoResp.Body)
	if err != nil {
		return err
	}

	var authUser User
	json.Unmarshal(userInfo, &authUser)

	fmt.Println(authUser.Email)

	db, err := gorm.Open(mysql.Open(setting.MYSQL_INFO), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return err
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})

	var checkUser User
	result := db.Where(User{Email: authUser.Email}).First(&checkUser)

	if result.RowsAffected <= 0 {
		return c.String(http.StatusUnauthorized, "invalid auth\n")
	}

	claims := &JwtClaims{
		checkUser.UserId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := jwtToken.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func Login(c echo.Context) error {
	userid := c.FormValue("userid")
	password := c.FormValue("password")

	fmt.Println("userid", userid)
	fmt.Println("password", password)

	db, err := gorm.Open(mysql.Open(setting.MYSQL_INFO), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return err
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})

	var checkUser User
	result := db.Where(User{UserId: userid}).First(&checkUser)

	if result.RowsAffected <= 0 || checkUser.Password != password {
		return c.String(http.StatusUnauthorized, "invalid auth\n")
	}

	claims := &JwtClaims{
		userid,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func SignUp(c echo.Context) error {
	userid := c.FormValue("userid")
	password := c.FormValue("password")

	if userid == "" || password == "" {
		return c.String(http.StatusOK, "incorrect format")
	}

	db, err := gorm.Open(mysql.Open(setting.MYSQL_INFO), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return err
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&(User{}))

	user := &User{
		UserId:   userid,
		Password: password,
	}

	result := db.Create(user)

	if result.Error != nil {
		return c.String(http.StatusOK, result.Error.Error())
	}

	return c.String(http.StatusOK, "created")
}

func GoogleSignup(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, OauthSignUpConf.AuthCodeURL(RandToken()))
}

func GoogleSignUpCallback(c echo.Context) error {
	code := c.FormValue("code")
	if code != "" {
		fmt.Println("code: ", code)
	}

	token, err := OauthSignUpConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		return err
	}

	fmt.Println(token)

	client := OauthSignUpConf.Client(oauth2.NoContext, token)
	userInfoResp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return err
	}
	defer userInfoResp.Body.Close()

	userInfo, err := ioutil.ReadAll(userInfoResp.Body)
	if err != nil {
		return err
	}

	var authUser User
	json.Unmarshal(userInfo, &authUser)

	fmt.Println(authUser.Email)

	db, err := gorm.Open(mysql.Open(setting.MYSQL_INFO), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return err
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&(User{}))

	fmt.Println(c.QueryParam("userid"))

	user := &User{
		UserId: c.QueryParam("userid"),
		Email:  authUser.Email,
	}

	result := db.Updates(user)

	if result.Error != nil {
		return c.String(http.StatusOK, result.Error.Error())
	}

	return c.String(http.StatusOK, "ok")
}

func Authorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("check authorization...")

		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*JwtClaims)

		c.QueryParams().Add("userid", claims.UserId)

		return next(c)
	}
}
