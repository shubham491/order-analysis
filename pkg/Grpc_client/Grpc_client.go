package Grpc_client

import (
	"github.com/gin-gonic/gin"
	"github.com/shubham491/order-analysis/pkg/AuthUtil"
	"github.com/shubham491/order-analysis/pkg/services/orders/orderspb"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var Restaurant_count = make(map[string] int64)
var Cuisine_count = make(map[string] int64)
var State_cuisine_count = make(map[string]map[string]int64)
var Orders = make(map[string] int64)

func GetAllRestaurants(c *gin.Context) {

	user := c.MustGet(gin.AuthUserKey).(string)


	if _, ok := AuthUtil.Secrets[user]; ok {
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Sorry client cannot talk to server: %v: ", err)
		}

		defer conn.Close();
		oc := orderspb.NewOrdersServiceClient(conn)
		req := &orderspb.AllRestaurantRequest{RestaurantCount: Restaurant_count}
		res, err := oc.GetAllRestaurant(c, req)
		c.JSON(200,res.GetAllRestaurant())
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
	}

}

func GetAllCusines(c *gin.Context) {

	user := c.MustGet(gin.AuthUserKey).(string)


	if _, ok := AuthUtil.Secrets[user]; ok {
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Sorry client cannot talk to server: %v: ", err)
		}

		defer conn.Close();
		oc := orderspb.NewOrdersServiceClient(conn)
		req := &orderspb.AllCuisineRequest{CuisineCount: Cuisine_count}
		res, err := oc.GetAllCuisine(c, req)
		c.JSON(200,res.GetAllCuisine())
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
	}

}


func GetAllStatesCuisines(c *gin.Context) {

	user := c.MustGet(gin.AuthUserKey).(string)


	if _, ok := AuthUtil.Secrets[user]; ok {
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Sorry client cannot talk to server: %v: ", err)
		}

		defer conn.Close();
		oc := orderspb.NewOrdersServiceClient(conn)
		var res *orderspb.AllStateResponse
		var res1 *orderspb.AllCuisine
		var tempMap=make(map[string]*orderspb.AllCuisine)
		for k,v:= range State_cuisine_count{
			res1= &orderspb.AllCuisine{
				AllCuisine:v,
			}

			tempMap[k]=res1

		}
		req := &orderspb.AllStateRequest{StateCuisineCount: tempMap}
		res, err = oc.GetAllStateCusine(c, req)
		c.JSON(200,res.GetAllState())
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
	}

}

func GetTopNumRestaurants(c *gin.Context) {

	user := c.MustGet(gin.AuthUserKey).(string)


	if _, ok := AuthUtil.Secrets[user]; ok {
		num,err := strconv.ParseInt(c.Param("num"),10,64)
		if err != nil {
			log.Fatalf("Enter valid integer for num: %v: ", err)
			return
		}

		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Sorry client cannot talk to server: %v: ", err)
			return
		}

		defer conn.Close();
		oc := orderspb.NewOrdersServiceClient(conn)
		req := &orderspb.TopNumRestaurantRequest{Num:num,RestaurantCount: Restaurant_count}
		res, err := oc.GetTopNumRestaurants(c, req)
		c.JSON(200,res)
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
	}
}

func GetTopNumCuisines(c *gin.Context) {

	user := c.MustGet(gin.AuthUserKey).(string)


	if _, ok := AuthUtil.Secrets[user]; ok {
		num,err := strconv.ParseInt(c.Param("num"),10,64)
		if err != nil {
			log.Fatalf("Enter valid integer for num: %v: ", err)
			return
		}

		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Sorry client cannot talk to server: %v: ", err)
			return
		}

		defer conn.Close();
		oc := orderspb.NewOrdersServiceClient(conn)
		req := &orderspb.TopNumCuisineRequest{Num:num,CuisineCount: Cuisine_count}
		res, err := oc.GetTopNumCuisines(c, req)
		c.JSON(200,res)
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
	}
}

func GetTopNumStateCuisines(c *gin.Context) {

	user := c.MustGet(gin.AuthUserKey).(string)


	if _, ok := AuthUtil.Secrets[user]; ok {
		num,err := strconv.ParseInt(c.Param("num"),10,64)
		state:=c.Param("state")
		if err != nil {
			log.Fatalf("Enter valid integer for num: %v: ", err)
			return
		}

		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Sorry client cannot talk to server: %v: ", err)
			return
		}

		defer conn.Close();
		oc := orderspb.NewOrdersServiceClient(conn)
		//var res *orderspb.AllStateResponse
		var res1 *orderspb.AllCuisine
		var tempMap=make(map[string]*orderspb.AllCuisine)
		for k,v:= range State_cuisine_count{
			res1= &orderspb.AllCuisine{
				AllCuisine:v,
			}

			tempMap[k]=res1

		}
		req := &orderspb.TopNumStatesCuisinesRequest{Num:num, State: state,StateCuisineCount: tempMap}
		res, err := oc.GetTopNumStatesCuisines(c, req)
		c.JSON(200,res)
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
	}
}

func AddOrder(c *gin.Context) {

	user := c.MustGet(gin.AuthUserKey).(string)


	if _, ok := AuthUtil.Secrets[user]; ok {
		body:=c.Request.Body
		content, err:= ioutil.ReadAll(body)
		if err != nil {
			log.Fatalf(err.Error())
			return
		}

		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Sorry client cannot talk to server: %v: ", err)
			return
		}
		var res1 *orderspb.AllCuisine
		var tempMap=make(map[string]*orderspb.AllCuisine)
		for k,v:= range State_cuisine_count{
			res1= &orderspb.AllCuisine{
				AllCuisine:v,
			}

			tempMap[k]=res1

		}
		defer conn.Close();
		oc := orderspb.NewOrdersServiceClient(conn)
		req := &orderspb.AddOrderRequest{Order: string(content),RestaurantCount: Restaurant_count,CuisineCount: Cuisine_count,StateCuisineCount: tempMap,Orders: Orders}
		res, err := oc.AddOrder(c, req)
		c.JSON(200,res)
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
	}
}