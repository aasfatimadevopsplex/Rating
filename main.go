package main
import (
	"Rating/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	)
func main(){
	fmt.Println("hello")
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/get_rating", func(c *gin.Context) {

			ISBN := c.GetHeader("ISBN")
			rat := utils.GetRating(ISBN)
			c.JSON(http.StatusOK, gin.H{"status": "success", "Rating": &rat})

		})


	}

	_ = router.Run(":8083")


}
