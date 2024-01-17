package auth

import "github.com/gin-gonic/gin"

func AuthLogin(c *gin.Context) {
	// [todo]
	// 1. get account and pwd from form-data
	// 2. query db to get the stored-pwd
	// 3. compared the pwd
	// 4. [generate jwt token and return] or [tell the failure]
}

func AuthRegister(c *gin.Context) {
	// [todo]
	// 1. get account, username, pwd from form-data
	// 2. crypt and insert into db (check if dup-account)
	// 3. [generate jwt token and return] or [tell the failure]
}
