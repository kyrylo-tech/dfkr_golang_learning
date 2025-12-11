package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Address struct {
	City   string `json:"city"`
	Street string `json:"street"`
}

type User struct {
	ID        int      `json:"id"`
	Username  string   `json:"username"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Email     string   `json:"email"`
	IsActive  bool     `json:"isActive"`
	Address   Address  `json:"address"`
	Skills    []string `json:"skills"`
}

func loadUsers(path string) ([]User, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var users []User
	if err := json.Unmarshal(data, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func groupUsersBySkill(users []User) map[string][]User {
	result := make(map[string][]User)

	for _, u := range users {
		for _, skill := range u.Skills {
			skill = strings.TrimSpace(skill)
			if skill == "" {
				continue
			}
			result[skill] = append(result[skill], u)
		}
	}

	return result
}

func main() {
	users, err := loadUsers("users.json")
	if err != nil {
		fmt.Println("Error loading users:", err)
		return
	}

	groups := groupUsersBySkill(users)

	for skill, us := range groups {
		fmt.Printf("=== Skill: %s ===\n", skill)
		for _, u := range us {
			fmt.Printf("- [%d] %s %s (%s, %s)\n",
				u.ID, u.FirstName, u.LastName, u.Username, u.Address.City)
		}
		fmt.Println()
	}
}
