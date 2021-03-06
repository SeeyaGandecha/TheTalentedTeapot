package main

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/securecookie"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var tpl *template.Template

type PageData struct {
	Title     string
	FirstName string
}

func init() {

	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))

}

func main() {
	http.HandleFunc("/", idx)
	http.HandleFunc("/ttt", ttt)
	http.HandleFunc("/gallery", gall)
	http.HandleFunc("/forum", frm)
	http.HandleFunc("/maps", maps)
	http.HandleFunc("/contact", cntct)
	//http.HandleFunc("/account", acct)
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
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("files"))))
	http.HandleFunc("/account", accountpage)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/welcomeuser", userPage)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/upload", upload)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))


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
//func acct(w http.ResponseWriter, req *http.Request) {
//
//	pd := PageData{
//		Title: "Account",
//	}
//
//	var firstname string
//
//	if req.Method == http.MethodPost {
//
//		firstname = req.FormValue("fname")
//		pd.FirstName = firstname
//	}
//
//	err := tpl.ExecuteTemplate(w, "account.gohtml", pd)
//	if err != nil {
//		log.Println(err)
//		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
//		return
//	}
//}

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

//func account(w http.ResponseWriter, r *http.Request) {
//
//	pd := PageData{
//				Title: "Account",
//			}
//			err := tpl.ExecuteTemplate(w, "account.gohtml", pd)
//			if err != nil {
//				log.Println(err)
//				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
//				return
//			}
//}

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func accountpage(w http.ResponseWriter, r *http.Request) {

	uuid := getUuid(r)
	if uuid != "" {
		http.Redirect(w, r, "/welcomeuser", 302)
		return
	}
	msg := getMsg(w, r, "message")
	var u = &User{}
	u.Errors = make(map[string]string)
	if msg != "" {
		u.Errors["message"] = msg
		render(w, "signin", u)
	} else {
		u := &User{}
		render(w, "signin", u)
	}

}

func login(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("uname")
	pass := r.FormValue("password")
	u := &User{Username: name, Password: pass}
	redirect := "/account"
	if name != "" && pass != "" {
		if b, uuid := userExists(u); b == true {
			setSession(&User{Uuid: uuid}, w)
			redirect = "/welcomeuser"
		} else {
			setMsg(w, "message", "please signup or enter a valid username and password!")
		}
	} else {
		setMsg(w, "message", "Username or Password field are empty!")
	}
	http.Redirect(w, r, redirect, 302)
}

func logout(w http.ResponseWriter, r *http.Request) {
	clearSession(w, "session")
	http.Redirect(w, r, "/account", 302)
}

func userPage(w http.ResponseWriter, r *http.Request) {
	uuid := getUuid(r)
	u := getUserFromUuid(uuid)
	if uuid != "" {
		render(w, "internal", u)
	} else {
		setMsg(w, "message", "Please login first!")
		http.Redirect(w, r, "/account", 302)
	}
}

func signup(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		u := &User{}
		u.Errors = make(map[string]string)
		u.Errors["lname"] = getMsg(w, r, "lname")
		u.Errors["fname"] = getMsg(w, r, "fname")
		u.Errors["email"] = getMsg(w, r, "email")
		u.Errors["username"] = getMsg(w, r, "username")
		u.Errors["password"] = getMsg(w, r, "password")
		render(w, "signup", u)
	case "POST":
		if n := checkUser(r.FormValue("userName")); n == true {
			setMsg(w, "username", "User already exists. Please enter a unique username!")
			http.Redirect(w, r, "/signup", 302)
			return
		}
		u := &User{
			Uuid:     Uuid(),
			Fname:    r.FormValue("fName"),
			Lname:    r.FormValue("lName"),
			Email:    r.FormValue("email"),
			Username: r.FormValue("userName"),
			Password: r.FormValue("password"),
		}
		result, err := govalidator.ValidateStruct(u)
		if err != nil {
			e := err.Error()
			if re := strings.Contains(e, "Lname"); re == true {
				setMsg(w, "lname", "Please enter a valid Last Name")
			}
			if re := strings.Contains(e, "Email"); re == true {
				setMsg(w, "email", "Please enter a valid Email Address!")
			}
			if re := strings.Contains(e, "Fname"); re == true {
				setMsg(w, "fname", "Please enter a valid First Name")
			}
			if re := strings.Contains(e, "Username"); re == true {
				setMsg(w, "username", "Please enter a valid Username!")
			}
			if re := strings.Contains(e, "Password"); re == true {
				setMsg(w, "password", "Please enter a Password!")
			}

		}
		if r.FormValue("password") != r.FormValue("cpassword") {
			setMsg(w, "password", "The passwords you entered do not Match!")
			http.Redirect(w, r, "/signup", 302)
			return
		}

		if result == true {
			u.Password = enyptPass(u.Password)
			saveData(u)
			http.Redirect(w, r, "/account", 302)
			return
		}
		http.Redirect(w, r, "/signup", 302)

	}
}

func render(w http.ResponseWriter, name string, data interface{}) {
	tmpl, err := template.ParseGlob("templates/*.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	tmpl.ExecuteTemplate(w, name, data)
}


func upload(w http.ResponseWriter, r *http.Request) {

	uuid := getUuid(r)
	u := getUserFromUuid(uuid)
	if uuid != "" {
		render(w, "/upload", u)
	} else {
		setMsg(w, "message", "Please login first!")
		http.Redirect(w, r, "/upload", 302)
	}
	fmt.Println("Method", r.Method)



	fn := ""
	if r.Method == http.MethodPost {
		// open
		f, h, err := r.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		fn = h.Filename

		df, err := os.Create("./assets/"+fn)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer df.Close()

		_, err = io.Copy(df, f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}

	tpl.ExecuteTemplate(w, "upload.gohtml", fn)
}