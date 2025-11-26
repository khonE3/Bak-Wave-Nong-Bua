package main

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// TeaseRequest represents the tease API request body
type TeaseRequest struct {
	Name string `json:"name" binding:"required"`
}

// TeaseResponse represents the tease API response
type TeaseResponse struct {
	Message   string    `json:"message"`
	Name      string    `json:"name"`
	Count     int       `json:"count"`
	Timestamp time.Time `json:"timestamp"`
}

// StatsResponse represents the stats API response
type StatsResponse struct {
	TotalTeases int            `json:"totalTeases"`
	TeaseCounts map[string]int `json:"teaseCounts"`
}

// In-memory storage for tease counts
var (
	teaseCounts = make(map[string]int)
	totalTeases = 0
	mutex       = &sync.RWMutex{}
)

func main() {
	r := gin.Default()

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// API routes
	api := r.Group("/api")
	{
		api.GET("/health", healthCheck)
		api.POST("/tease", teaseHandler)
		api.GET("/stats", statsHandler)
	}

	log.Println("🌸 เกรียนเวฟ Backend Server starting on :8080")
	log.Println("🐃 หนองบัวลำภูอีสานมากๆ!")

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

// healthCheck returns server health status
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"message": "หนองบัวลำภูอีสานมากๆ! 🌸",
	})
}

// teaseHandler handles the tease button press
func teaseHandler(c *gin.Context) {
	var req TeaseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	mutex.Lock()
	teaseCounts[req.Name]++
	totalTeases++
	count := teaseCounts[req.Name]
	mutex.Unlock()

	// Generate random Isan message
	messages := []string{
		"หนองบัวลำภูอีสานมากๆ! 🌾",
		"อีสานบ้านเฮา! 🐃",
		"สิบักหลาย! 🌸",
		"ม่วนซื่นคักๆ! 🎵",
		"ลำซิ่งหนองบัวลำภู! 💃",
	}
	message := messages[time.Now().UnixNano()%int64(len(messages))]

	response := TeaseResponse{
		Message:   message,
		Name:      req.Name,
		Count:     count,
		Timestamp: time.Now(),
	}

	log.Printf("🎉 %s ถูกเกรียนครั้งที่ %d", req.Name, count)
	c.JSON(http.StatusOK, response)
}

// statsHandler returns tease statistics
func statsHandler(c *gin.Context) {
	mutex.RLock()
	defer mutex.RUnlock()

	// Create a copy of the map
	countsCopy := make(map[string]int)
	for k, v := range teaseCounts {
		countsCopy[k] = v
	}

	response := StatsResponse{
		TotalTeases: totalTeases,
		TeaseCounts: countsCopy,
	}

	c.JSON(http.StatusOK, response)
}
