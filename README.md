# Go React REST API

This is a full Clean Architecture REST API with all REST  four end-points and also three endpoints for setting JWT web token. Authentication of user will encrupt password and generate random Token and based on valid token, a Cookie will be saved for determined duration  

Pre-requisites:
- Golang
- MySQL with a schema

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