package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"

	model "sevgifidan.com/webAPI/fourthapp/models"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	//todo: page nesnesi
	page := model.Page{ID: 3, Name: "Kullanıcılar", Description: "Kullanıcı Listesi", URI: "/users"}
	//todo: data loaders
	users := loadUsers()
	interests := loadInterests()
	interestsMappings := loadInterestMappings()
	//todo: işlem
	var newUsers []model.User
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
	viewModel := model.UserViewModel{Page: page, Users: newUsers}

	t, _ := template.ParseFiles("template/page.html")
	t.Execute(w, viewModel)
}

func loadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func loadUsers() []model.User {
	bytes, _ := ioutil.ReadFile("json/users.json")
	var data []model.User
	json.Unmarshal([]byte(bytes), &data)
	return data
}
func loadInterests() []model.Interest {
	bytes, _ := ioutil.ReadFile("json/interests.json")
	var data []model.Interest
	json.Unmarshal([]byte(bytes), &data)
	return data
}
func loadInterestMappings() []model.InterestMapping {
	bytes, _ := ioutil.ReadFile("json/userInterestMappings.json")
	var data []model.InterestMapping
	json.Unmarshal([]byte(bytes), &data)
	return data
}
