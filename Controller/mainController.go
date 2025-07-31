package Controller

import (
	modelStruct "CRUD_WebSite_GoLang/Model"
	"html/template"
	"net/http"
)

func HandleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/contacts", viewContacts)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func viewContacts(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/contacts.html")
	_ = tmpl.Execute(w, nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	user1 := modelStruct.User{
		Name:      "John Doe",
		Age:       17,
		Email:     "johndoe1375@mail.ru",
		Money:     1123.50,
		AvgGrades: 4.3,
		Happiness: 6.2,
		Hobbies:   []string{"Programming", "Basketball", "Sex"}}

	tmpl, _ := template.ParseFiles("templates/home.html")
	_ = tmpl.Execute(w, user1)
}
