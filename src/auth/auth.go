package auth

import (
	"fmt"
	"net/http"
	"time"
	"todo/src/setting"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type JwtClaims struct {
	UserId string `json:"Id"`
	jwt.StandardClaims
}

type User struct {
	gorm.Model
	UserId   string `gorm:"primaryKey;unique"`
	Password string
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

func Authorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("check authorization...")

		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*JwtClaims)

		c.QueryParams().Add("userid", claims.UserId)

		return next(c)
	}
}
