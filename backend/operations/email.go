package operations

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	_ "github.com/lib/pq"
)

type Mail struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

func Add_Email(db *sql.DB, res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "Method Not Allowed!", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "Error reading request body", http.StatusBadRequest)
		return
	}
	//unmarshal the request
	var mail Mail
	err = json.Unmarshal(body, &mail)
	if err != nil {
		http.Error(res, "Unable parsing the request body", http.StatusInternalServerError)
		return
	}
	query := `INSERT INTO emails  (email) values ($1) returning id`
	err = db.QueryRow(query, mail.Email).Scan(&mail.ID)
	if err != nil {
		fmt.Println("Error inserting data", err)
		http.Error(res, "Error Inserting Data", http.StatusInternalServerError)
		return
	}
	response := map[string]interface{}{
		"Message": "Data Inserted Successfully",
		"Email":   mail.ID,
		"ID":      mail.Email,
	}
	res.Header().Set("Content-Type", "application/json")
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	res.Write(jsonResponse)
}
