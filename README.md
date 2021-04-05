# smartmidea-todo# Todo List API



회원가입 및 로그인 기능과 Todo 정보를 생성, 수정, 삭제, 조회를 할 수 있다.



#### 기본정보

> Language: Go
>
> Framework:
>
> - echo (github.com/labstack/echo/v4 [v4.2.0])
> - gorm (gorm.io/gorm [v1.20.12])
> - mysql (gorm.io/driver/mysql [v1.0.4])
> - jwt-go (github.com/dgrijalva/jwt-go [v3.2.0+incompatible])



#### DB 구조

users

|       | Field      | Type                     | Comment |
| ----- | ---------- | ------------------------ | ------- |
| :key: | id         | bigint unsigned NOT NULL |         |
|       | created_at | datetime(3) NULL         |         |
|       | updated_at | datetime(3) NULL         |         |
|       | deleted_at | datetime(3) NULL         |         |
| :key: | user_id    | varchar(191) NOT NULL    |         |
|       | password   | longtext NULL            |         |



todos

|       | Field      | Type                     | Comment |
| ----- | ---------- | ------------------------ | ------- |
| :key: | id         | bigint unsigned NOT NULL |         |
|       | created_at | datetime(3) NULL         |         |
|       | updated_at | datetime(3) NULL         |         |
|       | deleted_at | datetime(3) NULL         |         |
| :key: | user_id    | varchar(191) NOT NULL    |         |
|       | start_date | datetime(3) NULL         |         |
|       | end_date   | datetime(3) NULL         |         |
|       | title      | longtext NULL            |         |
|       | status     | longtext NULL            |         |





API 기능 (PostMan)

```json
{
	"info": {
		"_postman_id": "da1b0407-a2f5-43a7-8b49-1d0d15041328",
		"name": "todo",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
	},
	"item": [
		{
			"name": "todo조회",
			"protocolProfileBehavior": {
				"disableBodyPruning": true
			},
			"request": {
				"method": "GET",
				"header": [
					{
						"key": "Authorization",
						"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6InRlc3QzIiwiZXhwIjoxNjE3NjY4NDgzfQ.dFfnGrZ51wAi7j7EiZs65kjulQCYLM8f8Z6IlXWfkRA",
						"type": "text"
					}
				],
				"body": {
					"mode": "raw",
					"raw": "",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:1323/todo?userid=test3&from=2021-02-19&to=2021-02-21",
					"host": [
						"localhost"
					],
					"port": "1323",
					"path": [
						"todo"
					],
					"query": [
						{
							"key": "userid",
							"value": "test3"
						},
						{
							"key": "from",
							"value": "2021-02-19"
						},
						{
							"key": "to",
							"value": "2021-02-21"
						}
					]
				}
			},
			"response": []
		},
		{
			"name": "todo생성",
			"request": {
				"method": "PUT",
				"header": [
					{
						"key": "Authorization",
						"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6InRlc3QzIiwiZXhwIjoxNjE3NjY4NDgzfQ.dFfnGrZ51wAi7j7EiZs65kjulQCYLM8f8Z6IlXWfkRA",
						"type": "text"
					}
				],
				"body": {
					"mode": "raw",
					"raw": "{\r\n    \"version\": \"1.0\",\r\n    \"todolist\": [\r\n        {\r\n            \"ID\": 4,\r\n            \"CreatedAt\": \"2021-02-20T00:11:42.171Z\",\r\n            \"UpdatedAt\": \"2021-02-20T00:11:42.171Z\",\r\n            \"DeletedAt\": null,\r\n            \"startdate\": \"2021-02-20T02:20:00Z\",\r\n            \"enddate\": \"2021-02-20T03:20:00Z\",\r\n            \"title\": \"play game\",\r\n            \"status\": \"past\"\r\n        },\r\n        {\r\n            \"ID\": 5,\r\n            \"CreatedAt\": \"2021-02-20T00:11:42.171Z\",\r\n            \"UpdatedAt\": \"2021-02-20T00:11:42.171Z\",\r\n            \"DeletedAt\": null,\r\n            \"startdate\": \"2021-02-20T02:20:00Z\",\r\n            \"enddate\": \"2021-02-20T03:20:00Z\",\r\n            \"title\": \"eat lunch\",\r\n            \"status\": \"todo\"\r\n        },\r\n        {\r\n            \"ID\": 6,\r\n            \"CreatedAt\": \"2021-02-20T00:11:42.171Z\",\r\n            \"UpdatedAt\": \"2021-02-20T00:11:42.171Z\",\r\n            \"DeletedAt\": null,\r\n            \"startdate\": \"2021-02-20T02:20:00Z\",\r\n            \"enddate\": \"2021-02-20T03:20:00Z\",\r\n            \"title\": \"sleep\",\r\n            \"status\": \"done\"\r\n        }\r\n    ]\r\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:1323/todo",
					"host": [
						"localhost"
					],
					"port": "1323",
					"path": [
						"todo"
					]
				}
			},
			"response": []
		},
		{
			"name": "todo수정",
			"request": {
				"method": "PATCH",
				"header": [
					{
						"key": "Authorization",
						"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6InRlc3QzIiwiZXhwIjoxNjE3NjY4NDgzfQ.dFfnGrZ51wAi7j7EiZs65kjulQCYLM8f8Z6IlXWfkRA",
						"type": "text"
					}
				],
				"body": {
					"mode": "raw",
					"raw": "{\r\n    \"version\": \"1.0\",\r\n    \"todolist\": [\r\n        {\r\n            \"startDate\": \"2021-02-20T02:20:00Z\",\r\n            \"endDate\": \"2021-02-20T03:20:00Z\",\r\n            \"title\": \"play game\",\r\n            \"status\": \"done\"\r\n        },\r\n        {\r\n            \"startDate\": \"2021-02-20T02:20:00Z\",\r\n            \"endDate\": \"2021-02-20T03:20:00Z\",\r\n            \"title\": \"eat lunch\",\r\n            \"status\": \"done\"\r\n        },\r\n        {\r\n            \"startDate\": \"2021-02-20T02:20:00Z\",\r\n            \"endDate\": \"2021-02-20T03:20:00Z\",\r\n            \"title\": \"sleep\",\r\n            \"status\": \"done\"\r\n        }\r\n    ]\r\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:1323/todo",
					"host": [
						"localhost"
					],
					"port": "1323",
					"path": [
						"todo"
					]
				}
			},
			"response": []
		},
		{
			"name": "todo삭제",
			"request": {
				"method": "DELETE",
				"header": [
					{
						"key": "Authorization",
						"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6InRlc3QzIiwiZXhwIjoxNjE3NjY4NDgzfQ.dFfnGrZ51wAi7j7EiZs65kjulQCYLM8f8Z6IlXWfkRA",
						"type": "text"
					}
				],
				"body": {
					"mode": "raw",
					"raw": "{\r\n    \"version\": \"1.0\",\r\n    \"todolist\": [\r\n        {\r\n            \"id\": 4\r\n        },\r\n        {\r\n            \"id\": 5\r\n        },\r\n        {\r\n            \"id\": 6\r\n        }\r\n    ]\r\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:1323/todo",
					"host": [
						"localhost"
					],
					"port": "1323",
					"path": [
						"todo"
					]
				}
			},
			"response": []
		},
		{
			"name": "회원가입",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "formdata",
					"formdata": [
						{
							"key": "userid",
							"value": "test4",
							"type": "text"
						},
						{
							"key": "password",
							"value": "test4",
							"type": "text"
						}
					]
				},
				"url": {
					"raw": "localhost:1323/signup",
					"host": [
						"localhost"
					],
					"port": "1323",
					"path": [
						"signup"
					]
				}
			},
			"response": []
		},
		{
			"name": "로그인",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "formdata",
					"formdata": [
						{
							"key": "userid",
							"value": "test4",
							"type": "text"
						},
						{
							"key": "password",
							"value": "test4",
							"type": "text"
						}
					]
				},
				"url": {
					"raw": "localhost:1323/login",
					"host": [
						"localhost"
					],
					"port": "1323",
					"path": [
						"login"
					]
				}
			},
			"response": []
		}
	]
}
```

