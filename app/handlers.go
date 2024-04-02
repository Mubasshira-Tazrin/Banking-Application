package app
import(
	"encoding/xml"
	"encoding/json"
	"fmt"
	"net/http"
)
type Customer struct {
	Name    string `json:"full name"`
	City    string `json:"city"`
	Zipcode string `json:"Zip_Code"`
}
// Handler function for the "/greet" route
func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Greetings!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "saima", City: "dhaka", Zipcode: "1200"},
		{Name: "saa", City: "ctg", Zipcode: "1000"},
	}
	if r.Header.Get("Content-Type")=="application/xml"{
	w.Header().Add("Content-Type", "application/xml")
	xml.NewEncoder(w).Encode(customers)
	} else{
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
    }
}