package controllers

import (
	"context"
	"encoding/json"

	"net/http"
	"strings"

	"github.com/Funmi4194/funmimod/configs"
	"github.com/Funmi4194/funmimod/models"
	"github.com/Funmi4194/funmimod/responses"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetConnection("user")
var valid = validator.New()

// var user models.User

func CreateAccount() http.HandlerFunc {
	var user models.User
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "applacation/json")

		//json.decoder  to convert json into go
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			user.Name = strings.ToLower(user.Name)
			// log.Fatal(err) insteadof log we can just validate the request body
			//w.writeheader functionis used for setting the API status
			w.WriteHeader(http.StatusBadRequest)
			response := responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			}
			json.NewEncoder(w).Encode(response)
			return
		}
		//validate the required field

		if validationErr := valid.Struct(&user); validationErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data: map[string]interface{}{
					"data": validationErr.Error(),
				},
			}

			json.NewEncoder(w).Encode(response)
			return
		}

		newUser := models.User{
			//ID generates another primtive objectID
			//ID:         primitive.NewObjectID(),
			Name:       strings.ToLower(user.Name),
			Age:        user.Age,
			Profession: user.Profession,
			Location:   user.Location,
		}

		//newUser.Name = strings.ToLower(newUser.Name)
		result, err := userCollection.InsertOne(context.Background(), newUser)

		if err != nil {
			//validates
			w.WriteHeader(http.StatusInternalServerError)
			response := responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data: map[string]interface{}{
					"data": result,
				},
			}
			json.NewEncoder(w).Encode(response)
			return
		}
		w.WriteHeader(http.StatusCreated)
		response := responses.UserResponse{
			Status:  http.StatusCreated,
			Message: "success",
			Data: map[string]interface{}{
				"data": result,
			},
		}
		json.NewEncoder(w).Encode(response)
	}
}

func GetUserDetails() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var user models.User
		params := mux.Vars(r)
		username := strings.ToLower(params["Name"])

		filter := bson.D{primitive.E{Key: "Name", Value: username}}
		err := userCollection.FindOne(context.Background(), filter).Decode(&user)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			}
			json.NewEncoder(w).Encode(response)
			return
		}
		w.WriteHeader(http.StatusOK)
		response := responses.UserResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data: map[string]interface{}{
				"data": map[string]interface{}{
					"data": user,
				},
			},
		}
		json.NewEncoder(w).Encode(response)

	}

}
