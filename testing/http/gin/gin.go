package gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
type (
	User struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
)

func NewGin() *gin.Engine {
	r := gin.Default()

	r.POST("/users", createUser)

	return r
}

func main() {
	r := NewGin()
	r.Run()
}

func createUser(c *gin.Context) {
	u := new(User)
	if err := c.ShouldBind(u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, u)
}
