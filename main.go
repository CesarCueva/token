package main

import (
	"log"
	"net/http"
	"os"

	"tokens/controller"
	"tokens/router"
	cService "tokens/service/cache"
	gService "tokens/service/globe"
	"tokens/usecase"
	"tokens/variables"
)

// Globeintegration - Handlers interface
type Globeintegration interface {
	IndexRoute(w http.ResponseWriter, r *http.Request)
	GetToken(w http.ResponseWriter, r *http.Request)
}

func main() {
	variables.Init()

	// Creates services (redis, globe)
	//redisClient := rService.NewRedisClient(os.Getenv("redisHost"))
	globeClient := gService.NewGlobeClient(os.Getenv("globeHost"), os.Getenv("timeoutGlobe"))
	cacheClient := cService.NewCacheClient()

	// Creates usecase
	u := usecase.NewUsecase(cacheClient, globeClient)

	// Creates controller
	c := controller.NewController(u)

	// Creates router
	r := router.NewRouter(c)

	//router.HandleFunc(os.Getenv("accessTokenEndPoint"), globeintegration.GetToken).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:3030", r))

	// newRequest := GetToken()
	// fmt.Println(newRequest)

}
