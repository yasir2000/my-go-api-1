# Go React REST API (JWT Authorization + React app Authentication + User management)

This is a full Clean Architecture REST API with all REST  four end-points and also three endpoints for setting JWT web token. Authentication of user will encrupt password and generate random Token and based on valid token, a Cookie will be saved for determined duration  

Pre-requisites:
- Golang
- MySQL or PostgreSQL with a schema (DB independent)

- gin: https://github.com/gin-gonic/gin
- mysql: https://github.com/go-sql-driver/mysql
- jwt-go: https://github.com/dgrijalva/jwt-go
- cors: https://github.com/gin-contrib/cors
- bcrypt: https://golang.org/x/crypto/bcrypt

For the frontend 

Create React app with typescript template using:
- npx create-react-app frontend --template typescript
- npm run start
- npm i react-router-dom @types/react-router-dom

For the backend API end-points :
``` go
	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	router.GET("/internal/users/search", users.Search)
	router.POST("/api/login", users.Login)
	router.POST("/api/register", users.Register)
	router.GET("/api/user", users.GetByCookie)
	router.GET("/api/logout", users.Logout)
```
