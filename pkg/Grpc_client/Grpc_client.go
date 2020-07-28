package Grpc_client

import (
	"github.com/gin-gonic/gin"
	"github.com/shubham491/order-analysis/pkg/AuthUtil"
	"github.com/shubham491/order-analysis/pkg/services/orders/orderspb"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"strconv"
)

func GetAllRestaurants(c *gin.Context) {

	user := c.MustGet(gin.AuthUserKey).(string)


	if _, ok := AuthUtil.Secrets[user]; ok {
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Sorry client cannot talk to server: %v: ", err)
		}

		defer conn.Close();
		oc := orderspb.NewOrdersServiceClient(conn)
		req := &orderspb.AllRestaurantRequest{}
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
		req := &orderspb.AllCuisineRequest{}
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
		req := &orderspb.AllStateRequest{}
		res, err := oc.GetAllState(c, req)
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
		req := &orderspb.TopNumRestaurantRequest{Num:num}
		res, err := oc.GetTopNumRestaurants(c, req)
		c.JSON(200,res)
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
	}
}