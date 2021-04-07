package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
)

type user_entity struct {
	// user_id               string
	// medicalemployee_id    int
	// email                 string
	// email_confirmed       bool
	// email_confirmed_token string
	Username      string `json: email`
	Password_hash string `json: password`
	// salt                  string
	// lockout               bool
	// reset_password_stamp  string
	// reset_password_date   string
}

/* type Cookie struct {
	Name       string
	Value      string
	Path       string
	Domain     string
	Expires    time.Time
	RawExpires string
	//Maxage <= 0 delete(d)
	//Maxage > 0 alive
	MaxAge   int
	Secure   bool
	HttpOnly bool
	Raw      string
	Unparsed []string
}

//Global Session Manager
type SessionManager struct {
	cookieName  string
	lock        sync.Mutex
	provider    Provider
	maxlifetime int64
}

func NewManager(provideName, cookieName string, maxlifetime int64) (*SessionManager, error) {
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errof("Session: unknown provide %q", provideName)
	}
	return &SessionManager{provider: provider, cookieName: cookieName, maxlifetime: maxlifetime}, nil
}

type Provider interface {
	SessionInit(sid string) (Session, error)  //init session, returns session
	SessionRead(sid string) (Session, error)  //returns session represented by the sid, else create new
	SessionDelete(sid string) error           //given sid is deleted from memory
	SessionGC(sid string) (maxLifeTime int64) //deletes expired sessions
}

var provides = make(map[string]Provider)

func Register(name string, provider Provider) {
	//create a session provider for the provided name
	if provider == nil {
		panic("Session: Register provider is nil.")
	}
	if _, dup := provides[name]; dup {
		panic("Session: Reigster called twiced for provider: " + name)
	}
	provides[name] = provider
}

func (manager *SessionManager) sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

//session creation
func (manager *SessionManager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {

}
*/
/* creation of cookies
   expiration := time.Now().Add(365 * 24 * time.Hour)
    cookie := http.Cookie{Name: "username", Value: "astaxie", Expires: expiration}
    http.SetCookie(w, &cookie)
getting a cookie
    cookie, _ := r.Cookie("username")
    fmt.Fprint(w, cookie)
OR
  for _, cookie := range r.Cookies() {
        fmt.Fprint(w, cookie.Name)
    }
*/
var store = sessions.NewCookieStore([]byte("secret-password"))

func Home(w http.ResponseWriter, r *http.Request) {
	//returns a session for the given name after adding it to registry
	session, err := store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	username, found := session.Values["username"]
	// authentication
	if !found || username == "" {
		//if not found or username is blank, redirect them to login page with HTTP 303
		http.Redirect(w, r, "/login.html", http.StatusSeeOther)
		return
	}
	/* 	if r.FormValue("email") != ""{
	   		session.Values["email"] = r.FormValue("email")
	   	}
	   	//save before writing response to client
	   	session.Values["authenticated"] = true
	   	err := session.Save(r,w)
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	} */
}

func SessionLogin(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	//authenication maybe add password?
	if r.FormValue("email") != "" {
		session.Values["email"] = r.FormValue("email")
		// check if its valid email and password from (cookie or DB)????
	}
	//user is enabled authentication
	session.Values["authenticated"] = true
	session.Save(r, w)
}

func SessionLogout(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "cookie-name")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
}

func secret(w http.ResponseWriter, r *http.Request) {

}

type correct struct {
	Correctcredentials bool `json: "correctcredentials"`
}

type incorrect struct {
	Incorrectcredentials bool `json: "incorrectcredentials"`
}

func signin(w http.ResponseWriter, r *http.Request) {
	//allow for CORS
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// t, _ := template.ParseFiles("index.html")
	// t.Execute(w, nil)

	switch r.Method {
	case "POST":
		user := user_entity{}

		correctcred := correct{
			Correctcredentials: true,
		}

		incorrectcred := incorrect{
			Incorrectcredentials: false,
		}

		correctcredJson, err := json.Marshal(correctcred)
		if err != nil {
			fmt.Fprintf(w, "Error: %s", err)
		}

		incorrectcredJson, err := json.Marshal(incorrectcred)
		if err != nil {
			fmt.Fprintf(w, "Error: %s", err)
		}

		jsn, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal("Error reading the body", err)
		}

		//fmt.Printf("ioutil.ReadAll Body: ", string(jsn))

		err = json.Unmarshal(jsn, &user)
		if err != nil {
			log.Fatal("Decoding error: ", err)
		}

		//for testing
		log.Printf("Received: %v\n", user)

		// Get the existing entry present in the database for the given username
		result := db.QueryRow("SELECT password_hash FROM user_entity where username=$1", user.Username)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write(incorrectcredJson)
			return
		}

		// We create another instance of `Credentials` to store the credentials we get from the database
		storedCreds := &user_entity{}
		// Store the obtained password in `storedCreds`
		err = result.Scan(&storedCreds.Password_hash)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write(incorrectcredJson)
			return
		}

		// Compare the stored passwords
		check := storedCreds.Password_hash == user.Password_hash
		if !check {
			w.Header().Set("Content-Type", "application/json")
			w.Write(incorrectcredJson)
			return
		}

		// If we reach this point, that means the users password was correct
		w.Header().Set("Content-Type", "application/json")
		w.Write(correctcredJson)
	}
}

// func newsAggHandler(w http.ResponseWriter, r *http.Request) {
// 	p := NewsAggPage{Title: "Amazing News Aggregator", News: "Some News"}
// 	t, _ := template.ParseFiles("index.html")
// 	t.Execute(w, p)
// }
/*
//used for testing
func server() {
	http.HandleFunc("/", signin)
	// http.ListenAndServe(":8088", nil)
	//http.HandleFunc("/agg", newsAggHandler)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal(err)
	}
}
*/
