package dbop

import (
	//"context"
	"database/sql"
	"encoding/json"
	//"errors"
	"fmt"
	"log"
	//"time"

	_ "github.com/go-sql-driver/mysql"
	//"codefood-rahman/middleware"
)

type RecipeCategory struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    CreatedAt string `json:"createdAt"`
    UpdatedAt string `json:"updatedAt"`
}

const (
	username = "root"
	password = ""
	hostname = "127.0.0.1:3306"
	dbname   = "codefood-db"
)

var db *sql.DB

func init() {
	db = prepareDb(dbname)
	defer db.Close()

}

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func prepareDb(dbname string) *sql.DB {
	db, err := sql.Open("mysql", dsn(dbname))
	if err != nil {
		log.Printf("error %s during the open db\n", err)
	}
	return db
}

func FindRecipeCategoryId(input_id string) int{
	db := prepareDb(dbname)
	defer db.Close()
	results, err := db.Query("SELECT id from tb_recipe_category WHERE id = ?", input_id)
	if err != nil {
		log.Fatal("An error occured during the query db to get id ", err)
		return -1
	}
	var id int
	for results.Next() {
		err = results.Scan(&id)
		if err != nil {
			log.Fatal("an error occured during the scan db to get id ", err)
			return -1
		}
	}
	fmt.Println("got id --> ", id)
	return id
}

func DeleteRecipeCategory(input_id int) bool {
	db := prepareDb(dbname)
	defer db.Close()

	return deleteRecipeCategory(input_id)
}

func deleteRecipeCategory(input_id int) bool {
	db := prepareDb(dbname)
	defer db.Close()
	_, err := db.Query("DELETE FROM tb_recipe_category WHERE id = ?", input_id)
	if err != nil {
		log.Fatal("An error occured during the query db to get sigle Recipe ", err)
		return false
	}
	return true
}

func UpdateRecipeCategory(reqBody []byte, input_id int) RecipeCategory {
	db := prepareDb(dbname)
	defer db.Close()

	var newRecipe RecipeCategory
	json.Unmarshal(reqBody, &newRecipe)

	_, err := updateRecipeCategory(newRecipe, int64(input_id))
	if err != nil {
		log.Fatal("Failed to update Recipe Category ", err)
		return RecipeCategory{}
	}

	return getSingleRecipeCategory(input_id)
}

func updateRecipeCategory(newRecipe RecipeCategory, recipeId int64) (int64, error) {
	db := prepareDb(dbname)
	defer db.Close()
	stmt, err := db.Prepare("UPDATE tb_recipe_category SET name = ?, updated_at = ? WHERE id = ?")
	if err != nil {
		log.Fatal("An error occured during the update user %w", err)
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(newRecipe.Name, newRecipe.UpdatedAt, recipeId)

	if err != nil {
		log.Fatal("an error occured during the exec db to update : %w", err)
		return 0, err
	}
	return res.RowsAffected()
}

func getSingleRecipeCategory(input_id int) RecipeCategory{
	db := prepareDb(dbname)
	defer db.Close()
	results, err := db.Query("SELECT * from tb_recipe_category WHERE id = ?", input_id)
	if err != nil {
		log.Fatal("An error occured during the query db to get sigle Recipe ", err)
	}
	var temp RecipeCategory
	for results.Next() {
		err = results.Scan(&temp.ID, &temp.Name, &temp.CreatedAt, &temp.UpdatedAt)
		if err != nil {
			log.Fatal("an error occured during the scan db to get single Recipe ", err)
		}
	}
	return temp
}

func GetAllRecipeCategories() []RecipeCategory{
	db := prepareDb(dbname)
	defer db.Close()
	results, err := db.Query("SELECT * from tb_recipe_category")
	if err != nil {
		panic(err.Error())
	}
	var recipeCategoryArr []RecipeCategory
	for results.Next() {
		var temp RecipeCategory
		err = results.Scan(&temp.ID, &temp.Name, &temp.CreatedAt, &temp.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
		recipeCategoryArr = append(recipeCategoryArr, temp)
	}
	return recipeCategoryArr
}

func AddRecipeCategory(reqBody []byte) RecipeCategory {
	var recipeCategory RecipeCategory
	db := prepareDb(dbname)
	defer db.Close()
	json.Unmarshal(reqBody, &recipeCategory)
	id, err := insertRecipeCategory(db, recipeCategory)
	if err != nil {
		log.Println("Failed to insert into db ", err)
		//return null if it is success
		recipeCategory = RecipeCategory{}
	}

	log.Printf("Inserted row with ID of %d\n", id)
	return getSingleRecipeCategory(int(id))
}

func insertRecipeCategory(db *sql.DB, recipe RecipeCategory) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO tb_recipe_category VALUES (?,?,?,?)")
	if err != nil {
		return -1, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(recipe.ID, recipe.Name, recipe.CreatedAt, recipe.UpdatedAt)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}
