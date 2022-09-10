package routes

import (
	"encoding/json"
	"net/http"

	"github.com/claudemuller/fishwife/internal/pkg"
)

type Reminder struct {
	ID          int         `json:"id"`
	User        string      `json:"user"`
	Description string      `json:"description"`
	Done        bool        `json:"done"`
	CreatedAt   string      `json:"created_at"`
	UpdatedAt   interface{} `json:"updated_at"`
	DeletedAt   interface{} `json:"deleted_at"`
}

func GetReminders(w http.ResponseWriter, _ *http.Request, app pkg.AppState) {
	//rows, err := app.DB.Query("SELECT * from sales")
	//if err != nil {
	//	log.Fatal("db query error: %w", err)
	//}
	//defer rows.Close()

	// var sale Sale
	// var sales []Sale

	//for rows.Next() {
	//	if err = rows.Scan(&sale.ID, &sale.Name, &sale.Description, &sale.CreatedAt, &sale.UpdatedAt, &sale.DeletedAt); err != nil {
	//		log.Fatal("db query error: %w", err)
	//	}
	//	sales = append(sales, sale)
	//}

	//if err = rows.Err(); err != nil {
	//	log.Fatal("db query error: %w", err)
	//}

	reminder := Reminder{
		ID:          1,
		User:        "claude",
		Description: "test reminder",
	}

	res := pkg.Response{
		Data: reminder,
	}

	resJSON, err := json.Marshal(res)
	if err != nil {
		pkg.SendInternalServerError(w)

		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(resJSON)
}
