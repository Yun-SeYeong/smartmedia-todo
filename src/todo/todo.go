package todo

import (
	"encoding/json"
	"net/http"
	"todo/src/database"
	"todo/src/models"

	"github.com/labstack/echo/v4"
)

func CreateTodos(c echo.Context) error {
	request := new(models.RequestTodo)

	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	manager, err := database.New("name")
	if err != nil {
		return err
	}

	for i, _ := range request.TodoList {
		request.TodoList[i].UserId = c.QueryParam("userid")
	}

	result := manager.DBMS.Create(request.TodoList)

	if result.Error != nil {
		return c.String(http.StatusOK, result.Error.Error())
	}

	resultJson, _ := json.Marshal(request.TodoList)

	return c.String(http.StatusOK, string(resultJson))
}

func UpdateTodos(c echo.Context) error {
	request := new(models.RequestTodo)

	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	manager, err := database.New("manager")
	if err != nil {
		return err
	}

	for i, _ := range request.TodoList {
		request.TodoList[i].UserId = c.QueryParam("userid")
		manager.DBMS.Updates(request.TodoList[i])
	}

	return c.String(http.StatusOK, "updated")
}

func DeleteTodos(c echo.Context) error {
	request := new(models.RequestTodo)

	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	manager, err := database.New("manager")
	if err != nil {
		return err
	}

	for i, _ := range request.TodoList {
		request.TodoList[i].UserId = c.QueryParam("userid")
		manager.DBMS.Unscoped().Delete(request.TodoList[i])
	}

	return c.String(http.StatusOK, "deleted")
}

func QueryTodos(c echo.Context) error {
	fromDate := c.QueryParam("from")
	toDate := c.QueryParam("to")

	request := new(models.RequestTodo)

	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	manager, err := database.New("manager")
	if err != nil {
		return err
	}

	var todoList []models.Todo

	manager.DBMS.Where("start_date >= ? AND start_Date <= ? and user_id = ?", fromDate, toDate, c.QueryParam("userid")).Find(&todoList)

	return c.JSON(http.StatusOK, models.RequestTodo{
		Version:  "1.0",
		TodoList: todoList,
	})
}
