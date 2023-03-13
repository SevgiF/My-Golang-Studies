package main

import (
	"fmt"
	"net/http"

	helper "sevgifidan.com/registerLogin/helpers"
)

func main() {

	uName, email, pwd, pwdConfirm := "", "", "", ""

	mux := http.NewServeMux()
	//todo: SIGNUP
	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		uName = r.FormValue("username")
		email = r.FormValue("email")
		pwd = r.FormValue("password")
		pwdConfirm = r.FormValue("confirm")

		uNameCheck := helper.IsEmpty(uName)
		emailCheck := helper.IsEmpty(email)
		pwdCheck := helper.IsEmpty(pwd)
		pwdConfirmCheck := helper.IsEmpty(pwdConfirm)

		if uNameCheck || emailCheck || pwdCheck || pwdConfirmCheck {
			fmt.Fprintf(w, "There is empty data.")
			return
		}

		if pwd == pwdConfirm {
			//todo: Veritabanına kullanıcıyı kaydet.
			fmt.Fprintf(w, "Registration successful.")
		} else {
			fmt.Fprintf(w, "Password information must be the same.")
		}

	})
	//todo: LOGIN
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		email = r.FormValue("email")
		pwd = r.FormValue("password")

		emailCheck := helper.IsEmpty(email)
		pwdCheck := helper.IsEmpty(pwd)

		if emailCheck || pwdCheck {
			fmt.Fprintf(w, "There is empty data.")
			return
		}

		//Veritabanında bir veriymiş gibi değişken tanımladık
		dbPwd := "1234!*."
		dbEmail := "cihan.ozhan@hotmail.com"

		if email == dbEmail && pwd == dbPwd {
			fmt.Fprintf(w, "Login successful.")
		} else {
			fmt.Fprintf(w, "Login failed.")
		}
	})

	http.ListenAndServe(":8080", mux)

}

/*
	for key, value := range r.Form {
		fmt.Printf("%s = %s", key, value)
	}*/
