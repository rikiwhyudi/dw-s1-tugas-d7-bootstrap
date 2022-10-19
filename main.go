package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter()

	// route path folder untuk public
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer((http.Dir("./public")))))

	//routing
	route.HandleFunc("/hello", helloWorld).Methods("GET")
	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")
	route.HandleFunc("/blog", blog).Methods("GET")
	route.HandleFunc("/blog-detail/{id}", blogDetail).Methods("GET")
	route.HandleFunc("/post", addBlog).Methods("POST")
	route.HandleFunc("/process", formAddBlog).Methods("GET")



	fmt.Println("Server running on port 5000");
	http.ListenAndServe("localhost:5000", route)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

	//mengatur header
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	//membuat variabel guna memparsing halaman
	var tmpl, err  = template.ParseFiles("views/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK) 
	tmpl.Execute(w, nil)
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var  tmpl, err = template.ParseFiles("views/form.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message :" +err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}

func blog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/blog.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message :" +err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)

}

func blogDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/blog-detail.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message :" +err.Error()))
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"]) 
	
	data := map[string]interface{} {
	"Title": "Dumbways Web App",
	"Date": "15 Oktober 2022",
	"Duration": "1 Month",
	"React": "React Js",
	"JavaScript": "JavaScript",
	"NodeJs": "Node Js",
	"SocketIO": "Socket IO",
	"Content":`Lorem ipsum dolor sit amet consectetur adipisicing elit. Soluta
	corrupti possimus optio accusamus dignissimos magni modi dolorem
	dolore in, ullam nisi, iusto aspernatur sit cupiditate quasi.
	Numquam dolorum neque praesentium corporis laudantium laboriosam
	alias quae? Nam nesciunt illo iusto! Assumenda vero delectus
	recusandae itaque reiciendis natus quibusdam aspernatur odio porro
	aperiam voluptatibus qui ipsa numquam tempore quas iusto pariatur
	et, dolorum aliquam. Impedit architecto, corrupti beatae id
	laboriosam enim amet ullam totam quaerat repellat illo delectus,
	minima, optio voluptatum neque explicabo iure ipsa cupiditate
	ratione. Aspernatur repudiandae placeat quam iusto neque, doloremque
	accusamus aliquam rerum velit rem id culpa minus earum a suscipit
	blanditiis ducimus molestiae magni laborum beatae esse provident
	vero? Architecto asperiores expedita, eius sint quis debitis fuga?`,
	"Id": 		id,
}
	w.WriteHeader(http.StatusInternalServerError)
	tmpl.Execute(w, data)
}


func formAddBlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message :" +err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)

}

func addBlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Title :" + r.PostForm.Get("inputTitle"))
	fmt.Println("Start :" + r.PostForm.Get("inputStart"))
	fmt.Println("End :" + r.PostForm.Get("inputEnd"))
	fmt.Println("Content :" + r.PostForm.Get("inputContent"))

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}