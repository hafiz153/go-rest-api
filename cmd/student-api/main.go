package main
import(
	"net/http"
	"log"
	"github.com/hafizul16103123/student-api/internal/config"
)
func main(){
	// Load config
	cfg:=config.MustLoad()
	// Database setup
	// setup router
	router:=http.NewServeMux()
	router.HandleFunc("GET /",func(w http.ResponseWriter ,r *http.Request){
		w.Write([]byte("Welcome to student API"))
	})
		// setup server
	server:=http.Server{
		Addr:cfg.Addr,
		Handler:router,
	}

	err:=server.ListenAndServe()
	if err!= nil{
		log.Fatal("Failed to start server")
	}

}
