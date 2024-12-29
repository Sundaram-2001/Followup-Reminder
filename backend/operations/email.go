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
	ID            int    `json:"id"`
	Email         string `json:"email"`
	Selected_Date string `json:"selectedDate"`
}

func Add_Email(db *sql.DB, res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	res.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	res.Header().Set("Access-Control-Max-Age", "15")

	if req.Method != http.MethodPost {
		http.Error(res, "Method Not Allowed!", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "Error reading request body", http.StatusBadRequest)
		return
	}
	fmt.Println("Raw request body:", string(body))

	var mail Mail
	err = json.Unmarshal(body, &mail)
	if err != nil {
		http.Error(res, "Unable parsing the request body", http.StatusInternalServerError)
		return
	}
	fmt.Printf("Parsed values - Email: %s, Selected_date: %s\n", mail.Email, mail.Selected_Date)

	query := `INSERT INTO emails (email,Selected_Date) values ($1, $2) returning id`
	err = db.QueryRow(query, mail.Email, mail.Selected_Date).Scan(&mail.ID)
	if err != nil {
		fmt.Println("Error inserting data", err)
		http.Error(res, "Error Inserting Data", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"Message":       "Data Inserted Successfully",
		"Email":         mail.Email,
		"Selected Date": mail.Selected_Date,
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
