package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"golangdockerex/models"
)

type HomeHellow struct {
	Message string
}

var (
	client   *mongo.Client
	err      error
	collName string
	dbNAME   string
)

func main() {

	dbNAME = os.Getenv("DBNAME")
	collName = os.Getenv("COLLNAME")
	dbConnectionString := os.Getenv("dbConnectionString")
	client, err = mongo.NewClient(options.Client().ApplyURI(dbConnectionString))

	if err != nil {
		log.Printf("Failed to create client: %v", err)
	}
	//context.Backgound() should be used in main, init,and tests, TODO when you don't know what context to use
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Printf("Failed to connect to cluster: %v", err)
	}

	// Force a connection to verify our connection string
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Printf("Failed to ping cluster: %v", err)
	}

	r := chi.NewRouter()

	r.Get("/", HomePage)
	r.Post("/addAlert", AddCryptoAlert)

	http.ListenAndServe(":3000", r)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	res := HomeHellow{
		Message: "WHALEOCME HERE",
	}
	json.NewEncoder(w).Encode(res)
}

//Adds a new crypto Alert
//TODO -> Possibly some middle ware for Auth, for payload validation
func AddCryptoAlert(w http.ResponseWriter, req *http.Request) {
	var reqData models.CryptoNotification
	/*
		if err := render.Bind(req, &contact); err != nil {
				//handle error
			}
	*/
	err := json.NewDecoder(req.Body).Decode((&reqData))
	//TODO: return 403 bad request on empty body request
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Body malformed"))
		//panic(err)
		//log.Fatal(err)
		return
	}
	//if the request is good, pass the data to the manager to do work,
	//in this layer we just return 200 or 400 or 505 errors on data bad request

	//TODO use middle ware pattern styles

	resp := AddCryptoAlertREPO(reqData)

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(resp)
}

func AddCryptoAlertREPO(cryptoNotification models.CryptoNotification) *models.CryptoNotification {
	if err := client.Ping(context.TODO() /*repo.DBHandler.Ctx*/, readpref.Primary()); err != nil {
		// Can't connect to Mongo server
		log.Fatal(err)
	}

	collection := client.Database(dbNAME).Collection(collName)

	insertResult, err := collection.InsertOne(context.TODO() /*repo.DBHandler.Ctx*/, cryptoNotification)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(insertResult)

	return &cryptoNotification
}
