package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/conf"
	"log"
	"net/http"
	"petclinic/internal/handler"
	"petclinic/internal/support"
	"strconv"
)

var configFile = flag.String("f", "../../configs/petclinic-local.yaml", "the configs file")

/*func main() {
	flag.Parse()

	var c model.Config
	conf.MustLoad(*configFile, &c)

	ctx := service.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}*/

func main() {
	// load configuration file
	flag.Parse()
	var c support.Config
	conf.MustLoad(*configFile, &c)
	// static resources
	http.Handle("/resources/css/", http.StripPrefix("/resources/css/", http.FileServer(http.Dir("../../web/static/resources/css/"))))
	http.Handle("/resources/images/", http.StripPrefix("/resources/images/", http.FileServer(http.Dir("../../web/static/resources/images/"))))
	http.Handle("/static/resources/images/", http.StripPrefix("/static/resources/images/", http.FileServer(http.Dir("../../web/static/resources/images/"))))
	http.Handle("/resources/fonts/", http.StripPrefix("/resources/fonts/", http.FileServer(http.Dir("../../web/static/resources/fonts/"))))
	// welcome and error
	http.HandleFunc("/", handler.ParseTemplateFile("../../web/templates/welcome.html", map[string]string{"welcome": "Welcome to My Pet Clinic by Golang!"}))
	http.HandleFunc("/error", handler.ParseTemplateFile("../../web/templates/error.html", map[string]string{"message": "Demo internal server error"}))
	// vets
	http.HandleFunc("/vets", handler.GetVets(nil))
	http.HandleFunc("/vets/new", handler.NewVetForm(nil))
	// owners
	http.HandleFunc("/owners/find", handler.FindOwner(nil))
	http.HandleFunc("/owners", handler.GetOwners(nil))
	http.HandleFunc("/owners/new", handler.NewOwnerForm(nil))
	http.HandleFunc("/owners/detail", handler.GetOwnerDetail(nil))
	// pet
	http.HandleFunc("/pet/visit/new", handler.NewVisitForm(nil))
	http.HandleFunc("/pet/new", handler.NewPetForm(nil))
	// start server
	fmt.Printf("Starting HTTP server at port %v\n", c.Port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(c.Port), nil))
}
