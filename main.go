package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/payload", func(c *gin.Context) {
		userInput := &Payload{}

		// validate the input
		if err := c.ShouldBindJSON(userInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"status":  http.StatusText(http.StatusBadRequest),
				"message": err.Error(),
			})
			return
		}

		// conditional logic if user's payload included favoriteNumber
		if userInput.FavoriteNumber != nil {
			c.JSON(http.StatusCreated, gin.H{
				"code":    http.StatusOK,
				"status":  http.StatusText(http.StatusOK),
				"message": fmt.Sprintf("Your favoriteNumber is %d", *userInput.FavoriteNumber),
			})
			return
		}

		// respond when favoriteNumber was omited from the user's payload
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"status":  http.StatusText(http.StatusOK),
			"message": fmt.Sprintf("You did not provide favoriteNumber"),
		})
		return
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// Payload holds an optional field
type Payload struct {
	FavoriteNumber *int `json:"favoriteNumber" binding:"omitempty,min=1,max=10"`
}
