package controllers

import (
	"github.com/ElizeuS/gouser/database"
	"github.com/ElizeuS/gouser/models"
	"github.com/gin-gonic/gin"
)

//ShowUser() retorna os dados de um cliente cadastrado baseado em um UUID passado pelo path
func ShowUser(c *gin.Context) {
	id := c.Param("uuid")

	db := database.GetDaabase()

	var user models.User

	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find user",
		})

		return
	}
	c.JSON(200, user)

}

//CreateUser cadastra um novo cliente no banco de dados e retorna o UUID(ID) do usuário cadastrado.
func CreateUser(c *gin.Context) {
	db := database.GetDaabase()
	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(400, gin.H{
			"Error": "cannot bind JSON: " + err.Error(),
		})

		return
	}

	err = db.Create(&user).Error

	if err != nil {
		c.JSON(400, gin.H{
			"Error": "cannot create user: " + err.Error(),
		})

		return
	}
	/*TO-DO: Verifdicar um modo de passar os valores de (user)
	para o método submit(striing), contido em /microservice/ */
	c.JSON(200, user.ID)

}

//ShowUsers() retorna um JSON com todos os usuários cadastrados no banco.
func ShowUsers(c *gin.Context) {
	db := database.GetDaabase()

	var users []models.User

	err := db.Find(&users).Error
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "cannot list users: " + err.Error(),
		})
	}

	c.JSON(200, users)
}

//UpdateUser() atualiza os valores do registro do usuário, cujo o UUID é passado como parâmetro pelo path..
func UpdateUser(c *gin.Context) {
	id := c.Param("uuid")

	db := database.GetDaabase()

	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	//Similar ao err = db.Save(&user, id).Error
	err = db.Model(models.User{}).Where("id = ?", id).Updates(&user).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot update user: " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"msg": "update performed successfully"})
}

//DeleteUser() exclui um usuário cadastrado no banco de dados cujo o UUID é correspondente ao parâmetro passado pelo path.
func DeleteUser(c *gin.Context) {
	id := c.Param("uuid")

	db := database.GetDaabase()

	err := db.Where("id = ?", id).Delete(models.User{}).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete user: " + err.Error(),
		})
		return
	}

	c.Status(204)

}
