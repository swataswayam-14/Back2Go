package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type User struct {
	Id              string           `json:"id"`
	PersonalDetails *PersonalDetails `json:"personaldetails"`
	CompanyDetails  *CompanyDetails  `json:"companydetails"`
	Subscribed      bool             `json:"subscribed"`
}
type PersonalDetails struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int32  `json:"age"`
	Gender    string `json:"gender"`
	Married   bool   `json:"married"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
type CompanyDetails struct {
	Name       string  `json:"name"`
	Role       string  `json:"role"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

var users = []User{
	{"1", &PersonalDetails{"John", "Doe", 30, "Male", false, "john@example.com", "password"}, &CompanyDetails{"Google", "Engineer", "Tech", 120000}, true},
	{"2", &PersonalDetails{"Jane", "Smith", 28, "Female", true, "jane@example.com", "password"}, &CompanyDetails{"Microsoft", "Manager", "HR", 90000}, true},
	{"3", &PersonalDetails{"Alice", "Johnson", 35, "Female", true, "alice@example.com", "password"}, &CompanyDetails{"Amazon", "Architect", "Cloud", 150000}, false},
	{"4", &PersonalDetails{"Bob", "Brown", 40, "Male", false, "bob@example.com", "password"}, &CompanyDetails{"Meta", "Designer", "UI/UX", 85000}, true},
	{"5", &PersonalDetails{"Charlie", "Davis", 29, "Male", true, "charlie@example.com", "password"}, &CompanyDetails{"Tesla", "Analyst", "Finance", 100000}, false},
	{"6", &PersonalDetails{"Emma", "Wilson", 32, "Female", false, "emma@example.com", "password"}, &CompanyDetails{"Netflix", "Director", "Media", 180000}, true},
	{"7", &PersonalDetails{"David", "White", 45, "Male", true, "david@example.com", "password"}, &CompanyDetails{"Apple", "Engineer", "Hardware", 140000}, true},
	{"8", &PersonalDetails{"Olivia", "Martinez", 27, "Female", false, "olivia@example.com", "password"}, &CompanyDetails{"Uber", "Marketing", "Sales", 95000}, true},
	{"9", &PersonalDetails{"Ethan", "Anderson", 31, "Male", true, "ethan@example.com", "password"}, &CompanyDetails{"Google", "Security", "Cyber", 130000}, false},
	{"10", &PersonalDetails{"Sophia", "Thomas", 36, "Female", true, "sophia@example.com", "password"}, &CompanyDetails{"Airbnb", "Executive", "Strategy", 110000}, false},
	{"11", &PersonalDetails{"Michael", "Moore", 50, "Male", true, "michael@example.com", "password"}, &CompanyDetails{"IBM", "Lead", "IT", 125000}, true},
	{"12", &PersonalDetails{"Ava", "Lee", 24, "Female", false, "ava@example.com", "password"}, &CompanyDetails{"Spotify", "Content", "Music", 89000}, false},
	{"13", &PersonalDetails{"James", "Garcia", 38, "Male", true, "james@example.com", "password"}, &CompanyDetails{"LinkedIn", "Sales", "B2B", 99000}, true},
	{"14", &PersonalDetails{"Isabella", "Clark", 41, "Female", true, "isabella@example.com", "password"}, &CompanyDetails{"YouTube", "Creator", "Media", 78000}, true},
	{"15", &PersonalDetails{"Benjamin", "Rodriguez", 34, "Male", false, "benjamin@example.com", "password"}, &CompanyDetails{"Twitter", "Moderator", "Content", 72000}, true},
	{"16", &PersonalDetails{"Mia", "Lewis", 30, "Female", true, "mia@example.com", "password"}, &CompanyDetails{"Pinterest", "Designer", "Creative", 97000}, false},
	{"17", &PersonalDetails{"William", "Walker", 39, "Male", false, "william@example.com", "password"}, &CompanyDetails{"Snapchat", "Developer", "Mobile", 108000}, true},
	{"18", &PersonalDetails{"Charlotte", "Hall", 26, "Female", true, "charlotte@example.com", "password"}, &CompanyDetails{"Reddit", "Moderator", "Community", 71000}, false},
	{"19", &PersonalDetails{"Henry", "Young", 29, "Male", true, "henry@example.com", "password"}, &CompanyDetails{"Twitch", "Streamer", "Gaming", 95000}, true},
	{"20", &PersonalDetails{"Amelia", "King", 33, "Female", false, "amelia@example.com", "password"}, &CompanyDetails{"Discord", "Manager", "Community", 102000}, true},
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getSingleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, user := range users {
		if user.Id == params["id"] {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	user.Id = strconv.Itoa(rand.Intn(100000))
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}
func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	var updatedUser User
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	for i, user := range users {
		if user.Id == id {
			if updatedUser.PersonalDetails != nil {
				if updatedUser.PersonalDetails.Firstname != "" {
					users[i].PersonalDetails.Firstname = updatedUser.PersonalDetails.Firstname
				}
				if updatedUser.PersonalDetails.Lastname != "" {
					users[i].PersonalDetails.Lastname = updatedUser.PersonalDetails.Lastname
				}
				if updatedUser.PersonalDetails.Age != 0 {
					users[i].PersonalDetails.Age = updatedUser.PersonalDetails.Age
				}
				if updatedUser.PersonalDetails.Gender != "" {
					users[i].PersonalDetails.Gender = updatedUser.PersonalDetails.Gender
				}
				users[i].PersonalDetails.Married = updatedUser.PersonalDetails.Married
				if updatedUser.PersonalDetails.Email != "" {
					users[i].PersonalDetails.Email = updatedUser.PersonalDetails.Email
				}
				if updatedUser.PersonalDetails.Password != "" {
					users[i].PersonalDetails.Password = updatedUser.PersonalDetails.Password
				}
			}

			if updatedUser.CompanyDetails != nil {
				if updatedUser.CompanyDetails.Name != "" {
					users[i].CompanyDetails.Name = updatedUser.CompanyDetails.Name
				}
				if updatedUser.CompanyDetails.Role != "" {
					users[i].CompanyDetails.Role = updatedUser.CompanyDetails.Role
				}
				if updatedUser.CompanyDetails.Department != "" {
					users[i].CompanyDetails.Department = updatedUser.CompanyDetails.Department
				}
				if updatedUser.CompanyDetails.Salary != 0 {
					users[i].CompanyDetails.Salary = updatedUser.CompanyDetails.Salary
				}
			}

			users[i].Subscribed = updatedUser.Subscribed
			json.NewEncoder(w).Encode(users[i])
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	for index, user := range users {
		if user.Id == id {
			users = append(users[:index], users[index+1:]...)
			json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

func filterUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	queryParams := r.URL.Query()

	var filteredUsers []User

	for _, user := range users {
		matched := true

		if ageStr, exists := queryParams["age"]; exists {
			age, _ := strconv.Atoi(ageStr[0])
			if user.PersonalDetails.Age != int32(age) {
				matched = false
			}
		}
		if gender, exists := queryParams["gender"]; exists {
			if user.PersonalDetails.Gender != gender[0] {
				matched = false
			}
		}
		if department, exists := queryParams["department"]; exists {
			if user.CompanyDetails.Department != department[0] {
				matched = false
			}
		}

		if matched {
			filteredUsers = append(filteredUsers, user)
		}
	}
	json.NewEncoder(w).Encode(filteredUsers)
}

func sortUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	sortBy := r.URL.Query().Get("by")
	order := r.URL.Query().Get("order")

	sort.Slice(users, func(i, j int) bool {
		if sortBy == "age" {
			if order == "desc" {
				return users[i].PersonalDetails.Age > users[j].PersonalDetails.Age
			}
			return users[i].PersonalDetails.Age < users[j].PersonalDetails.Age
		} else if sortBy == "salary" {
			if order == "desc" {
				return users[i].CompanyDetails.Salary < users[j].CompanyDetails.Salary
			}
		}
		return false
	})
	json.NewEncoder(w).Encode(users)
}

func searchUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query().Get("q")

	var searchedResults []User

	for _, user := range users {
		if strings.Contains(strings.ToLower(user.PersonalDetails.Firstname), strings.ToLower(query)) || strings.Contains(strings.ToLower(user.PersonalDetails.Lastname), strings.ToLower(query)) || strings.Contains(strings.ToLower(user.CompanyDetails.Department), strings.ToLower(query)) || strings.Contains(strings.ToLower(user.CompanyDetails.Name), strings.ToLower(query)) || strings.Contains(strings.ToLower(user.CompanyDetails.Role), strings.ToLower(query)) {
			searchedResults = append(searchedResults, user)
		}
	}
	json.NewEncoder(w).Encode(searchedResults)
}

func getTopUsers(w http.ResponseWriter, r *http.Request) {
	topCount := 5
	query := r.URL.Query().Get("top")

	if query != "" {
		count, err := strconv.Atoi(query)
		if err == nil {
			topCount = count
		}
	}
	sort.Slice(users, func(i, j int) bool {
		return users[i].CompanyDetails.Salary > users[j].CompanyDetails.Salary
	})
	if topCount > len(users) {
		topCount = len(users)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users[:topCount])
}

func toggleSubscription(w http.ResponseWriter, r *http.Request) {
	var subscribedUsers []User
	for index, user := range users {
		if users[index].Subscribed {
			subscribedUsers = append(subscribedUsers, user)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subscribedUsers)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", getAllUsers).Methods("GET")
	r.HandleFunc("/users/{id}", getSingleUser).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
	r.HandleFunc("/users/filter", filterUsers).Methods("GET")
	r.HandleFunc("/users/sort", sortUsers).Methods("GET")
	r.HandleFunc("/users/search", searchUsers).Methods("GET")

	r.HandleFunc("/users/top", getTopUsers).Methods("GET")
	r.HandleFunc("/users/subscribers", toggleSubscription).Methods("GET")

	fmt.Printf("Starting server at port 3000\n")
	log.Fatal(http.ListenAndServe(":3000", r))
}
