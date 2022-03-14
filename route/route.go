package route

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"strings"
	//"sync"
	"time"
	"bytes"

	//"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	dbop "codefood-rahman/database"
	//"codefood-rahman/middleware"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my app : %s", "user app")
}

//----------------------------------------------------------------------------------------
type RetGetMessageCategory struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Data []dbop.RecipeCategory `json:"data,omitempty"`
}

func returnAllRecipeCategories(w http.ResponseWriter, r *http.Request) {
	//ignore if its not http get call

	var recipeCategory []dbop.RecipeCategory = dbop.GetAllRecipeCategories()

	temp := new(RetGetMessageCategory)
	temp.Success = true
	temp.Message = "Success"
	temp.Data = recipeCategory

	json.NewEncoder(w).Encode(temp)
}

type RetInsertMessageCategory struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Data dbop.RecipeCategory `json:"data,omitempty"`
}

type RetSuccessInsertUpdateMessageCategory struct{
	Success bool `json:"success"`
	Message string `json:"message"`
	Data dbop.RecipeCategory `json:"data,omitempty"`
}

type RetFailedInsertUpdateMessageCategory struct{
	Success bool `json:"success"`
	Message string `json:"message"`
}

func createNewRecipeCategories(w http.ResponseWriter, r *http.Request) {
	//ignore if its not http get call
	reqBody, _ := ioutil.ReadAll(r.Body)

	//validation (?)
	mappedReqBody := make(map[string]string)
	err := json.Unmarshal(reqBody, &mappedReqBody)

	if(reqBody == nil || mappedReqBody["name"] == "" || err != nil){
		w.WriteHeader(http.StatusBadRequest)

		temp := new(RetFailedInsertUpdateMessageCategory)
		temp.Success = false
		temp.Message = "Name is required"
		//temp.Data = nil

		json.NewEncoder(w).Encode(temp)
	}else{
		//add time
		type EnhanceReqBody struct{
			Name string `json:"name"`
			CreatedAt string `json:"createdAt"`
			UpdatedAt string `json:"updatedAt"`
		}

		enhanceReqBody := new(EnhanceReqBody)
		enhanceReqBody.Name = mappedReqBody["name"]
		enhanceReqBody.CreatedAt = string(time.Now().Format(time.RFC3339))
		enhanceReqBody.UpdatedAt = string(time.Now().Format(time.RFC3339))

		// fmt.Print(string(reqBody))
		// fmt.Print(createdAt)
		// fmt.Print(mappedReqBody["name"])
		//marshalledNewReqBody = json.Marshaler(enhanceReqBody)

		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(enhanceReqBody)

		recipeCategory := dbop.AddRecipeCategory(reqBodyBytes.Bytes())

		//it means its failed
		if (dbop.RecipeCategory{}) == recipeCategory {
			w.WriteHeader(http.StatusBadRequest)

			temp := new(RetFailedInsertUpdateMessageCategory)
			temp.Success = false
			temp.Message = "Error on inserting data"

			json.NewEncoder(w).Encode(temp)
		} else {//it means it is success

			w.WriteHeader(http.StatusOK)

			temp := new(RetSuccessInsertUpdateMessageCategory)
			temp.Success = true
			temp.Message = "Success"
			temp.Data = recipeCategory

			json.NewEncoder(w).Encode(temp)
		}
	}
}

func updateRecipeCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
  key := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)

	//validation
	mappedReqBody := make(map[string]string)
	err := json.Unmarshal(reqBody, &mappedReqBody)

	if(reqBody == nil || mappedReqBody["name"] == "" || err != nil){
		w.WriteHeader(http.StatusBadRequest)

		temp := new(RetFailedInsertUpdateMessageCategory)
		temp.Success = false
		temp.Message = "Name is required"
		//temp.Data = nil

		json.NewEncoder(w).Encode(temp)
	}else{
		//check if the id is exist
		idCheckRes := dbop.FindRecipeCategoryId(key)
		if(idCheckRes == 0){
			//error
			w.WriteHeader(http.StatusBadRequest)

			temp := new(RetFailedInsertUpdateMessageCategory)
			temp.Success = false
			temp.Message = "Recipe Category with id "+key+" not found"
			//temp.Data = nil

			json.NewEncoder(w).Encode(temp)
		}else{
			//add time
			type EnhanceReqBody struct{
				Name string `json:"name"`
				CreatedAt string `json:"createdAt"`
				UpdatedAt string `json:"updatedAt"`
			}

			enhanceReqBody := new(EnhanceReqBody)
			enhanceReqBody.Name = mappedReqBody["name"]
			enhanceReqBody.UpdatedAt = string(time.Now().Format(time.RFC3339))

			reqBodyBytes := new(bytes.Buffer)
			json.NewEncoder(reqBodyBytes).Encode(enhanceReqBody)

			updateRes := dbop.UpdateRecipeCategory(reqBodyBytes.Bytes(), idCheckRes)

			w.WriteHeader(http.StatusOK)
			temp := new(RetSuccessInsertUpdateMessageCategory)
			temp.Success = true
			temp.Message = "Success"
			temp.Data = updateRes

			json.NewEncoder(w).Encode(temp)
		}
	}
}

func deleteRecipeCategory(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
  key := vars["id"]

	//validation
	//check if the id is exist
	idCheckRes := dbop.FindRecipeCategoryId(key)
	if(idCheckRes == 0){
		//error
		w.WriteHeader(http.StatusBadRequest)

		temp := new(RetFailedInsertUpdateMessageCategory)
		temp.Success = false
		temp.Message = "Recipe Category with id "+key+" not found"

		json.NewEncoder(w).Encode(temp)
	}else{
		deleteRes := dbop.DeleteRecipeCategory(idCheckRes)

		if(deleteRes == true){
			w.WriteHeader(http.StatusOK)
			temp := new(RetFailedInsertUpdateMessageCategory)
			temp.Success = true
			temp.Message = "Success"

			json.NewEncoder(w).Encode(temp)
		}
	}
}

func HandleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	//	myRouter.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	myRouter.Handle("/", http.FileServer(http.Dir("./static")))

	myRouter.HandleFunc("/recipe-categories", createNewRecipeCategories).Methods("POST")
	myRouter.HandleFunc("/recipe-categories/{id}", updateRecipeCategory).Methods("PUT")
	myRouter.HandleFunc("/recipe-categories/{id}", deleteRecipeCategory).Methods("DELETE")
	myRouter.HandleFunc("/recipe-categories", returnAllRecipeCategories)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
