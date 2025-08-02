package DAO

import (
	us "CRUD_WebSite_GoLang/Model"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func DataConnection() {
	connStr := "user=postgres dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Ошибка при открытии подключения: %v", err)
	}
	fmt.Println("Подключено к PostgreSQL")
	defer db.Close()

	//Получение всех пользователей
	allUsers, err := db.Query("SELECT * FROM person")
	if err != nil {
		log.Fatalf("Ошибка получения списка пользователей: %v", err)
	}
	fmt.Println("Список получен:")

	for allUsers.Next() {
		var user us.User
		err = allUsers.Scan(&user.ID, &user.Name, &user.Age, &user.Email, &user.Money, &user.AvgGrades, &user.Happiness)
		if err != nil {
			log.Panicf("Ошибка сканирования данных: %v", err)
		}
		fmt.Println(fmt.Sprintf("Users: %s\n with age: %d\n with email: %s", user.Name, user.Age, user.Email))
	}
	defer allUsers.Close()

	//Добавление пользователей
	addUser, err := db.Query("Insert into person(name, age, email, money, avggrades, happiness) " +
		"values ('Tado Hopsky', 33, 'tadohopsky@mail.ru', 1233.2, 4.3, 6.5)")
	if err != nil {
		log.Fatalf("Ошибка при добавлении пользователя: %v", err)
	}
	fmt.Println("Пользователь внесён в базу данных!")
	defer addUser.Close()
}
