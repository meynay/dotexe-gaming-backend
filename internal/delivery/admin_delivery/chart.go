package admin_delivery

import (
	"net/http"
	"store/internal/entities"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (ad *AdminDelivery) GetChart(c *gin.Context) {
	time1 := c.Query("from")
	time2 := c.Query("to")
	from, err := time.Parse(entities.TimeLayout, time1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad 'from' format"})
		return
	}
	to, err := time.Parse(entities.TimeLayout, time2)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad 'to' format"})
		return
	}
	showType, err := strconv.Atoi(c.Query("show"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad 'show' format"})
		return
	}
	info := ad.adminusecase.GetChartInfo(entities.ChartFilter{From: from, To: to, ShowType: showType})
	c.JSON(http.StatusOK, gin.H{"chart-data": info})
}
