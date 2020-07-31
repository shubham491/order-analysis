package main

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"

	//"net/http"

	//"github.com/shubham491/order-analysis/pkg/Grpc_server"
	"github.com/shubham491/order-analysis/pkg/AuthUtil"
	"github.com/shubham491/order-analysis/pkg/Grpc_client"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	
)

//func log(h http.Handler) gin.HandlerFunc {
//	return gin.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		fmt.Println("Before")
//		h.ServeHTTP(w, r) // call original
//		fmt.Println("After")
//	})
//}



func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(Grpc_client.CpuTemp)
	prometheus.MustRegister(Grpc_client.HdFailures)
	prometheus.MustRegister(Grpc_client.ApiHits)
}

func addAPIPaths(router *gin.Engine){
	Grpc_client.CpuTemp.Set(65.3)
	Grpc_client.HdFailures.With(prometheus.Labels{"device":"/dev/sda"}).Inc()

	restaurantAPI := router.Group("/restaurant", gin.BasicAuth(AuthUtil.Accounts))
	cuisineAPI := router.Group("/cuisine", gin.BasicAuth(AuthUtil.Accounts))
	stateCuisineAPI := router.Group("/state", gin.BasicAuth(AuthUtil.Accounts))
	orderAPI := router.Group("/order", gin.BasicAuth(AuthUtil.Accounts))
	homeAPI :=router.Group("/")

	homeAPI.GET("/metrics",gin.WrapH(promhttp.Handler()))

	//restaurantAPI:=router.Group("/restaurant")
	restaurantAPI.GET("/all", Grpc_client.GetAllRestaurants)
	restaurantAPI.GET("/top/:num", Grpc_client.GetTopNumRestaurants)

	//cuisineAPI:=router.Group("/cuisine")
	cuisineAPI.GET("/all", Grpc_client.GetAllCusines)
	cuisineAPI.GET("/top/:num", Grpc_client.GetTopNumCuisines)

	//stateCuisineAPI:=router.Group("/state")
	stateCuisineAPI.GET("/all", Grpc_client.GetAllStatesCuisines)
	stateCuisineAPI.GET("/top/:state/:num", Grpc_client.GetTopNumStateCuisines)

	orderAPI.POST("/add",Grpc_client.AddOrder)
}

func main() {
	router := gin.Default()

	addAPIPaths(router)


	router.Run(":5665")

}
