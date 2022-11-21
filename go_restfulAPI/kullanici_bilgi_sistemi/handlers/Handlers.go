package handlers

import (
	"encoding/json"
	"net/http"

	. "sevgifidan.com/kullaniciSistemi/dataloaders"
	. "sevgifidan.com/kullaniciSistemi/models"
)

func Run() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	//todo: page nesnesi
	page := Page{ID: 7, Name: "Kullanıcılar", Description: "Kullanıcı Listesi", URI: "/users"}
	//todo: data loaders
	users := LoadUsers()
	interests := LoadInterests()
	interestsMappings := LoadInterestMappings()
	//todo: işlem
	var newUsers []User
	for _, user := range users {
		for _, interestMapping := range interestsMappings {
			if user.ID == interestMapping.UserID {
				for _, interest := range interests {
					if interestMapping.InterestID == interest.ID {
						user.Interests = append(user.Interests, interest)
					}
				}
			}
		}
		newUsers = append(newUsers, user)
		//log.Println(user.FirstName)
	}
	//todo: view-model nesnesi
	viewModel := UserViewModel{Page: page, Users: newUsers}
	data, _ := json.Marshal(viewModel)
	w.Write([]byte(data))
}
