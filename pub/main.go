package main

import (
	"github.com/asaskevich/govalidator"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type PageData struct {
	Title     string
	FirstName string
}

func init() {

	tpl = template.Must(template.ParseGlob("/Users/seeyagandecha/DMU/Year 4/Final Year Project/Prototype/src/TheTalentedTeapot/pub/templates/*.gohtml"))

}

func main() {
	http.HandleFunc("/", idx)
	http.HandleFunc("/ttt", ttt)
	http.HandleFunc("/gallery", gall)
	http.HandleFunc("/forum", frm)
	http.HandleFunc("/maps", maps)
	http.HandleFunc("/contact", cntct)
	http.HandleFunc("/account", acct)
	http.HandleFunc("/TTTabout", TTTabt)
	http.HandleFunc("/TTTfeedback", TTTfdbk)
	http.HandleFunc("/TTTevents", TTTevents)
	http.HandleFunc("/TTTshop", TTTshop)
	http.HandleFunc("/TTTportfolio", TTTport)
	http.HandleFunc("/TTTservices", TTTserv)
	http.HandleFunc("/TTTsm", TTTsm)

	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img"))))

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))

	govalidator.SetFieldsRequiredByDefault(true)

	http.ListenAndServe(":8080", nil)

}

func idx(w http.ResponseWriter, req *http.Request) {
	pd := PageData{
		Title: "Home",
	}

	err := tpl.ExecuteTemplate(w, "home.gohtml", pd)
	if err != nil {
		log.Println("LOGGED", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
func ttt(w http.ResponseWriter, req *http.Request) {

	pd := PageData{
		Title: "The Talented Teapot",
	}

	err := tpl.ExecuteTemplate(w, "ttt.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func gall(w http.ResponseWriter, req *http.Request) {

	pd := PageData{
		Title: "Gallery",
	}

	err := tpl.ExecuteTemplate(w, "gallery.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func frm(w http.ResponseWriter, req *http.Request) {

	pd := PageData{
		Title: "Forum",
	}

	err := tpl.ExecuteTemplate(w, "forum.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func maps(w http.ResponseWriter, req *http.Request) {

	pd := PageData{
		Title: "Maps",
	}

	err := tpl.ExecuteTemplate(w, "maps.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func cntct(w http.ResponseWriter, req *http.Request) {

	pd := PageData{
		Title: "Contact",
	}

	err := tpl.ExecuteTemplate(w, "contact.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
func acct(w http.ResponseWriter, req *http.Request) {

	pd := PageData{
		Title: "Account",
	}

	var firstname string

	if req.Method == http.MethodPost {

		firstname = req.FormValue("fname")
		pd.FirstName = firstname
	}

	err := tpl.ExecuteTemplate(w, "account.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func TTTabt(w http.ResponseWriter, req *http.Request) {

	pd := PageData{
		Title: "About",
	}

	err := tpl.ExecuteTemplate(w, "TTTabout.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func TTTfdbk(w http.ResponseWriter, req *http.Request) {

	pd := PageData{
		Title: "Contact",
	}

	err := tpl.ExecuteTemplate(w, "TTTfeedback.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
func TTTevents(w http.ResponseWriter, req *http.Request) {

	pd := PageData{
		Title: "Events",
	}

	err := tpl.ExecuteTemplate(w, "TTTevents.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func TTTshop(w http.ResponseWriter, req *http.Request) {

	pd := PageData{
		Title: "Shop",
	}

	err := tpl.ExecuteTemplate(w, "TTTshop.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func TTTport(w http.ResponseWriter, req *http.Request) {

	pd := PageData{
		Title: "Portfolio",
	}

	err := tpl.ExecuteTemplate(w, "TTTportfolio.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func TTTserv(w http.ResponseWriter, req *http.Request) {

	pd := PageData{
		Title: "Services",
	}

	err := tpl.ExecuteTemplate(w, "TTTservices.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
func TTTsm(w http.ResponseWriter, req *http.Request) {

	pd := PageData{
		Title: "Social Media",
	}

	err := tpl.ExecuteTemplate(w, "TTTsm.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

//func login(w http.ResponseWriter, req *http.Request) {
//
//	pd := PageData{
//		Title: "Login",
//	}
//
//	err := tpl.ExecuteTemplate(w, "login.gohtml", pd)
//	if err != nil {
//		log.Println(err)
//		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
//		return
//	}
//}
//
//func signup(w http.ResponseWriter, req *http.Request) {
//
//	pd := PageData{
//		Title: "Sign Up",
//	}
//
//	err := tpl.ExecuteTemplate(w, "signupold.gohtml", pd)
//	if err != nil {
//		log.Println(err)
//		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
//		return
//	}
//}
