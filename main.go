package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

// create variable tpl of type template
var tpl template.Template

// create variables for for each page's HTML
var (
	homeHTML        []byte
	attractionsHTML []byte
	eventsHTML      []byte
	planHTML        []byte
	buyticketsHTML  []byte
	successHTML     []byte
)

// declare init() and parse the templates in tpl
func init() {
	tpl = *template.Must(template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html"))

	// read templates and store in their respective variable, store error in err and handle the error
	var err error
	homeHTML, err = os.ReadFile("templates/home.html")
	if err != nil {
		log.Fatalln(err)
	}

	attractionsHTML, err = os.ReadFile("templates/attractions.html")
	if err != nil {
		log.Fatalln(err)
	}

	eventsHTML, err = os.ReadFile("templates/events.html")
	if err != nil {
		log.Fatalln(err)
	}

	planHTML, err = os.ReadFile("templates/plan.html")
	if err != nil {
		log.Fatalln(err)
	}

	buyticketsHTML, err = os.ReadFile("templates/tickets.html")
	if err != nil {
		log.Fatalln(err)
	}

	successHTML, err = os.ReadFile("templates/success.html")
	if err != nil {
		log.Fatalln(err)
	}
}

// start function main()
func main() {
	// use handle to serve the image files to the server so that they can be used with HTML
	http.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("images"))))

	// use handlefunc to route "/" to index()
	http.HandleFunc("/", index)

	// use handlefunc to route "/attractions" to attractions()
	http.HandleFunc("/attractions/", attractions)

	// use handlefunc to route "/events" to events()
	http.HandleFunc("/events/", events)

	// use handlefunc to route "/plan" to plan()
	http.HandleFunc("/plan/", plan)

	// use handlefunc to route "/buy-tickets" to tickets()
	http.HandleFunc("/buy-tickets/", tickets)

	// use handlefunc to route "/success" to success()
	http.HandleFunc("/success/", success)

	// start server at 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// declare index()
func index(w http.ResponseWriter, r *http.Request) {

	// execute tpl with index.html, pass homeHTML in as data and use template.HTML() to write it as html code
	err := tpl.ExecuteTemplate(w, "index.html", template.HTML(homeHTML))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

// declare attractions()
func attractions(w http.ResponseWriter, r *http.Request) {
	// execute tpl with attractions.html
	err := tpl.ExecuteTemplate(w, "index.html", template.HTML(attractionsHTML))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
	}
}

// declare events()
func events(w http.ResponseWriter, r *http.Request) {
	// execute tpl with events.html
	err := tpl.ExecuteTemplate(w, "index.html", template.HTML(eventsHTML))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

// declare plan()
func plan(w http.ResponseWriter, r *http.Request) {
	// execute tpl with plan.html
	err := tpl.ExecuteTemplate(w, "index.html", template.HTML(planHTML))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

// declare tickets()
func tickets(w http.ResponseWriter, r *http.Request) {
	// check if the request method is post
	if r.Method == http.MethodPost {
		// if it is, then redirect the link to see other page "success.html"
		http.Redirect(w, r, "/success", http.StatusSeeOther)
	}
	// execute tpl with tickets.html
	err := tpl.ExecuteTemplate(w, "index.html", template.HTML(buyticketsHTML))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

// declare success()
func success(w http.ResponseWriter, r *http.Request) {
	// execute tpl with success.html
	err := tpl.ExecuteTemplate(w, "index.html", template.HTML(successHTML))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
