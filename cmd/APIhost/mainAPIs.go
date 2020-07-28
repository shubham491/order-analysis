package main

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jyotishp/order-analysis/pkg/APIUtil"
	"github.com/tamerh/jsparser"
	"log"
	"os"
	//"github.com/shubham491/order-analysis/pkg/APIUtil"
	"github.com/shubham491/order-analysis/pkg/AuthUtil"
	"github.com/shubham491/order-analysis/pkg/Grpc_client"
)

func addAPIPaths(router *gin.Engine){
	restaurantAPI := router.Group("/restaurant", gin.BasicAuth(AuthUtil.Accounts))
	cuisineAPI := router.Group("/cuisine", gin.BasicAuth(AuthUtil.Accounts))
	stateCuisineAPI := router.Group("/state", gin.BasicAuth(AuthUtil.Accounts))
	orderAPI := router.Group("/order", gin.BasicAuth(AuthUtil.Accounts))

	//restaurantAPI:=router.Group("/restaurant")
	restaurantAPI.GET("/all", Grpc_client.GetAllRestaurants)
	restaurantAPI.GET("/top/:num", Grpc_client.GetTopNumRestaurants)

	//cuisineAPI:=router.Group("/cuisine")
	cuisineAPI.GET("/all", Grpc_client.GetAllCusines)
	cuisineAPI.GET("/top/:num", Grpc_client.GetTopNumCuisines)

	//stateCuisineAPI:=router.Group("/state")
	stateCuisineAPI.GET("/all", Grpc_client.GetAllStatesCuisines)
	stateCuisineAPI.GET("/top/:state/:num", Grpc_client.GetTopNumStateCuisines)

	orderAPI.POST("/add",APIUtil.AddOrder)
}

func main() {
	router := gin.Default()

	addAPIPaths(router)

	r, _ := os.Open("outputs.json")
	br := bufio.NewReaderSize(r, 65536)
	parser := jsparser.NewJSONParser(br, "orders")

	for json := range parser.Stream() {
		if json.Err != nil {
			log.Fatal(json.Err)
		}
		//fmt.Println(json.ObjectVals["OrderId"])
		restaurant := json.ObjectVals["RestName"]
		cuisine := json.ObjectVals["Cuisine"]
		state := json.ObjectVals["State"]
		id := json.ObjectVals["Id"]

		Grpc_client.Restaurant_count[restaurant.(string)]++
		Grpc_client.Cuisine_count[cuisine.(string)]++
		Grpc_client.Orders[id.(string)]++
		statemap, ok := Grpc_client.State_cuisine_count[state.(string)]
		if ok {
			statemap[cuisine.(string)]++
		} else {
			Grpc_client.State_cuisine_count[state.(string)] = make(map[string]int)
			Grpc_client.State_cuisine_count[state.(string)][cuisine.(string)]++
		}
	}
	fmt.Println(APIUtil.Orders["2999999"])
	router.Run(":5665")

}
