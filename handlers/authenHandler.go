package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/goweb4/models"
	"github.com/goweb4/utils"
	"github.com/stretchr/objx"
)

type HomePageVars struct {
	Name string
}

func Index(w http.ResponseWriter, r *http.Request) {
	var HomeVars HomePageVars
	if authCookie, err := r.Cookie("auth"); err == nil {
		var cookieData interface{}
		cookieData = objx.MustFromBase64(authCookie.Value)
		name := cookieData.(objx.Map)["name"].(string)
		HomeVars.Name = name
	}
	utils.GenerateTemplate(w, HomeVars, "index")
}

func Login(w http.ResponseWriter, r *http.Request) {
	utils.GenerateTemplate(w, "", "login_register", "login", "register")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var info models.User
	err := utils.MapFormValues(&info, r)
	if err != nil {
		fmt.Println("cannot decode login info: ", err)
	} else {
		if checkCredential(info) {
			setSession(info.UserName, w)
			http.Redirect(w, r, "/index", http.StatusSeeOther)
		}
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	clearSession(w)
	http.Redirect(w, r, "/login", 302)
}

func checkCredential(info models.User) bool {
	user := models.GetUser(info)
	if user == (models.User{}) {
		return false
	}
	return true
}

func setSession(userName string, response http.ResponseWriter) {
	authCookieValue := objx.New(map[string]interface{}{
		"name": userName,
	}).MustBase64()
	expiration := time.Now().Add(365 * 24 * time.Hour)
	http.SetCookie(response, &http.Cookie{
		Name:    "auth",
		Expires: expiration,
		Value:   authCookieValue})
}

func clearSession(response http.ResponseWriter) {
	http.SetCookie(response, &http.Cookie{
		Name:    "auth",
		Value:   "",
		MaxAge:  -1,
		Expires: time.Unix(1, 0),
	})
}