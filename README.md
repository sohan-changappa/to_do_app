{
	"info": {
		"_postman_id": "28f5de9b-d555-42cd-9247-2243c5d658bc",
		"name": "todo",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		"_exporter_id": "47027035",
		"_collection_link": "https://sohanchangappa-2502763.postman.co/workspace/Sohan's-Workspace~f51576fe-7dd6-4732-8bdc-6be5ce05d09c/collection/47027035-28f5de9b-d555-42cd-9247-2243c5d658bc?action=share&source=collection_link&creator=47027035"
	},
	"item": [
		{
			"name": "Signup",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n    \"username\":\"test1\",\n    \"password\": \"mypass\"\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "http://localhost:8080/signup",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"signup"
					]
				}
			},
			"response": [
				{
					"name": "Signup",
					"originalRequest": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\n    \"username\":\"test3\",\n    \"password\": \"mypass\"\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://localhost:8080/signup",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"signup"
							]
						}
					},
					"_postman_previewlanguage": null,
					"header": null,
					"cookie": [],
					"body": null
				}
			]
		},
		{
			"name": "Login",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n    \"username\":\"test3\",\n    \"password\": \"mypass\"\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "http://localhost:8080/login",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"login"
					]
				}
			},
			"response": [
				{
					"name": "Login",
					"originalRequest": {
						"method": "POST",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTU2ODEyMDIsInVzZXJuYW1lIjoidGVzdDIifQ.LBzLp9ybmqoo4DQgwupvsQkj7OzY-1hD_VeMmj1L3zg",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\n    \"username\":\"test3\",\n    \"password\": \"mypass\"\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://localhost:8080/login",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"login"
							]
						}
					},
					"_postman_previewlanguage": null,
					"header": null,
					"cookie": [],
					"body": null
				}
			]
		},
		{
			"name": "Get tasks",
			"request": {
				"method": "GET",
				"header": [
					{
						"key": "Authorization",
						"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTU2ODEyMDIsInVzZXJuYW1lIjoidGVzdDIifQ.LBzLp9ybmqoo4DQgwupvsQkj7OzY-1hD_VeMmj1L3zg",
						"type": "text"
					}
				],
				"url": {
					"raw": "http://localhost:8080/protect/todos",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"protect",
						"todos"
					]
				}
			},
			"response": [
				{
					"name": "Get tasks",
					"originalRequest": {
						"method": "GET",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTU2ODEyMDIsInVzZXJuYW1lIjoidGVzdDIifQ.LBzLp9ybmqoo4DQgwupvsQkj7OzY-1hD_VeMmj1L3zg",
								"type": "text"
							}
						],
						"url": {
							"raw": "http://localhost:8080/protect/todos",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"protect",
								"todos"
							]
						}
					},
					"_postman_previewlanguage": null,
					"header": null,
					"cookie": [],
					"body": null
				}
			]
		},
		{
			"name": "Add tasks",
			"request": {
				"method": "POST",
				"header": [],
				"url": {
					"raw": "http://localhost:8080/protect/todos",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"protect",
						"todos"
					]
				}
			},
			"response": [
				{
					"name": "Add tasks",
					"originalRequest": {
						"method": "POST",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTU2ODEyMDIsInVzZXJuYW1lIjoidGVzdDIifQ.LBzLp9ybmqoo4DQgwupvsQkj7OzY-1hD_VeMmj1L3zg",
								"type": "text"
							},
							{
								"key": "Content-Type",
								"value": "application/json",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\n  \"title\": \"to code\",\n  \"completed\": true\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://localhost:8080/protect/todos",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"protect",
								"todos"
							]
						}
					},
					"_postman_previewlanguage": null,
					"header": null,
					"cookie": [],
					"body": null
				}
			]
		},
		{
			"name": "Update task",
			"request": {
				"method": "PUT",
				"header": [
					{
						"key": "Authorization",
						"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTU2ODEyMDIsInVzZXJuYW1lIjoidGVzdDIifQ.LBzLp9ybmqoo4DQgwupvsQkj7OzY-1hD_VeMmj1L3zg",
						"type": "text"
					}
				],
				"url": {
					"raw": "http://localhost:8080/protect/todos/6",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"protect",
						"todos",
						"6"
					]
				}
			},
			"response": [
				{
					"name": "New Request",
					"originalRequest": {
						"method": "PUT",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTU2ODEyMDIsInVzZXJuYW1lIjoidGVzdDIifQ.LBzLp9ybmqoo4DQgwupvsQkj7OzY-1hD_VeMmj1L3zg",
								"type": "text"
							}
						],
						"url": {
							"raw": "http://localhost:8080/protect/todos/6",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"protect",
								"todos",
								"6"
							]
						}
					},
					"_postman_previewlanguage": null,
					"header": null,
					"cookie": [],
					"body": null
				}
			]
		},
		{
			"name": "Delete task",
			"request": {
				"method": "DELETE",
				"header": [
					{
						"key": "Authorization",
						"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTU2ODEyMDIsInVzZXJuYW1lIjoidGVzdDIifQ.LBzLp9ybmqoo4DQgwupvsQkj7OzY-1hD_VeMmj1L3zg",
						"type": "text"
					}
				],
				"url": {
					"raw": "http://localhost:8080/protect/todos/6",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"protect",
						"todos",
						"6"
					]
				}
			},
			"response": [
				{
					"name": "Delete task",
					"originalRequest": {
						"method": "DELETE",
						"header": [
							{
								"key": "Authorization",
								"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTU2ODEyMDIsInVzZXJuYW1lIjoidGVzdDIifQ.LBzLp9ybmqoo4DQgwupvsQkj7OzY-1hD_VeMmj1L3zg",
								"type": "text"
							}
						],
						"url": {
							"raw": "http://localhost:8080/protect/todos/6",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"protect",
								"todos",
								"6"
							]
						}
					},
					"_postman_previewlanguage": null,
					"header": null,
					"cookie": [],
					"body": null
				}
			]
		}
	]
}
