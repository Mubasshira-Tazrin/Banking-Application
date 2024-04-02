package app
import(
	"fmt"
     "log"
	"net/http"
)
func Start(){
	fmt.Println("Server is listening on port 8082...")

	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	// Register the handler function for the root URL
	http.HandleFunc("/", rootHandler)

	// Start the HTTP server on port 8081
	log.Fatal(http.ListenAndServe("localhost:8082", nil))
}