package main

import (
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"html/template"
	"httprouter"
	"io/ioutil"
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

type User struct {
	Email    string
	UserName string
	Password string
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
	http.HandleFunc("/login", login)
	http.HandleFunc("/signup", signup)

	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img"))))

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))

	http.ListenAndServe(":8080", nil)

	r := httprouter.New()
	http.Handle("/", r)
	r.POST("/api/checkusername", checkUserName)
	r.POST("/api/createuser", createUser)

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

func shop(w http.ResponseWriter, req *http.Request) {

	pd := PageData{
		Title: "Shop",
	}

	err := tpl.ExecuteTemplate(w, "shop.gohtml", pd)
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
func login(w http.ResponseWriter, req *http.Request) {

	pd := PageData{
		Title: "Login",
	}

	err := tpl.ExecuteTemplate(w, "login.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func signup(w http.ResponseWriter, req *http.Request) {

	pd := PageData{
		Title: "Sign Up",
	}

	err := tpl.ExecuteTemplate(w, "signup.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func checkUserName(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)
	bs, err := ioutil.ReadAll(req.Body)
	sbs := string(bs)
	stdlog.Println("REQUEST BODY: ", sbs)
	q, err := datastore.NewQuery("Users").Filter("UserName=", sbs).Count(ctx)
	stdlog.Println("ERR: ", err)
	stdlog.Println("QUANTITY: ", q)
	if err != nil {
		fmt.Fprint(res, "false")
		return
	}
	if q >= 1 {
		fmt.Fprint(res, "true")
	} else {
		fmt.Fprint(res, "false")
	}
}

func createUser(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)
	NewUser := User{
		Email:    req.FormValue("Email"),
		UserName: req.FormValue("Username"),
		Password: req.FormValue("Password"),
	}
	key := datastore.NewIncompleteKey(ctx, "Users", nil)
	key, _ = datastore.Put(ctx, key, &NewUser)
	http.Redirect(res, req, "/", 302)
}
