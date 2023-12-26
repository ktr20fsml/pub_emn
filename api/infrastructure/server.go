package infrastructure

import (
	"api/infrastructure/database"
	"api/interface/router"
	"api/registry"
	"fmt"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	// "github.com/gin-contrib/sessions/redis"
)

var (
	db    *sqlx.DB
	errDB error
)

func Run() {
	// get database connection
	for {
		db, errDB = database.NewDatabase()
		if errDB != nil {
			fmt.Println(errDB)
			time.Sleep(time.Second * 5)
		}
		if db != nil {
			break
		}
	}
	defer db.Close()

	g := gin.Default()

	g.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"http://localhost:8880", "http://192.168.3.4:8880", "http://localhost:3000", "http://192.168.3.4:3000", "http://localhost:9090", "http://localhost:8081"},
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "X-Custom-Header", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	store := cookie.NewStore([]byte("secret"))
	g.Use(sessions.Sessions("emn-session", store))
	// store, err := redis.NewStore(10, "tcp", "redis:6379", "", []byte("secret"))
	// if err != nil {
	// 	panic(err)
	// }
	g.Use(sessions.Sessions("emn-session", store))

	// resolve dependencies
	r := registry.NewInteractor(db)

	router.NewRouter(g, r)

	g.Run(":8080")
}
