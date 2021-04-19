package testing

import (
	"course-chart/config"
	"course-chart/models"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestDatabaseConnection(t *testing.T) {
	t.Run("it connects via DATABASE_URL in release mode", func(t *testing.T) {
		gin.SetMode(gin.ReleaseMode)
		config.Connect()

		err := config.Conn.Find(&models.Activity{}).Where("activities.custom = FALSE").Error
		if err.Error() != "pq: SSL is not enabled on the server" {
			t.Errorf("database connection via SSL not attempted")
		}

		gin.SetMode(gin.TestMode)
		config.Connect()
	})
}
