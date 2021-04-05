package todo

import (
	"net/http"
	"time"
	"todo/src/setting"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	UserId    string    `json:"userid" gorm:"primaryKey"`
	StartDate time.Time `json:"startdate"`
	EndDate   time.Time `json:"enddate"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
}

type RequestTodo struct {
	Version  string `json:"version"`
	TodoList []Todo `json:"todolist"`
}

func CreateTodos(c echo.Context) error {
	request := new(RequestTodo)

	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	db, err := gorm.Open(mysql.Open(setting.MYSQL_INFO), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return err
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Todo{})

	for i, _ := range request.TodoList {
		request.TodoList[i].UserId = c.QueryParam("userid")
	}

	result := db.Create(request.TodoList)

	if result.Error != nil {
		return c.String(http.StatusOK, result.Error.Error())
	}

	return c.String(http.StatusOK, "created")
}

func UpdateTodos(c echo.Context) error {
	request := new(RequestTodo)

	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	db, err := gorm.Open(mysql.Open(setting.MYSQL_INFO), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return err
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Todo{})

	for i, _ := range request.TodoList {
		request.TodoList[i].UserId = c.QueryParam("userid")
		db.Updates(request.TodoList[i])
	}

	return c.String(http.StatusOK, "updated")
}

func DeleteTodos(c echo.Context) error {
	request := new(RequestTodo)

	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	db, err := gorm.Open(mysql.Open(setting.MYSQL_INFO), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return err
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Todo{})

	for i, _ := range request.TodoList {
		request.TodoList[i].UserId = c.QueryParam("userid")
		db.Unscoped().Delete(request.TodoList[i])
	}

	return c.String(http.StatusOK, "deleted")
}

func QueryTodos(c echo.Context) error {
	fromDate := c.QueryParam("from")
	toDate := c.QueryParam("to")

	request := new(RequestTodo)

	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	db, err := gorm.Open(mysql.Open(setting.MYSQL_INFO), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return err
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Todo{})

	var todoList []Todo

	db.Where("start_date >= ? AND start_Date <= ? and user_id = ?", fromDate, toDate, c.QueryParam("userid")).Find(&todoList)

	return c.JSON(http.StatusOK, RequestTodo{
		Version:  "1.0",
		TodoList: todoList,
	})
}
