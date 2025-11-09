package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Task struct {
	ID                 int
	Title, Description string
}

var tasks []Task
var nextID int = 1

func addTask(title, description string) {
	task := Task{
		ID:          nextID,
		Title:       title,
		Description: description,
	}
	tasks = append(tasks, task)
	nextID++
}
func showAllTasks() {
	if len(tasks) == 0 {
		fmt.Print("No tasks yet")
		return
	}
	fmt.Print("All Tasks\n")
	for _, task := range tasks {
		fmt.Print(task.Title, " | ", task.Description, " | ", task.ID)
	}
}
func deleteTask(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Task was deleted")
			return
		}
	}
	fmt.Println("Task not found")
}

func main() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN environment variable is not set")
	}
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Print("Bot started: %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		text := update.Message.Text
		chatID := update.Message.Chat.ID

		// msg := tgbotapi.NewMessage(chatID, "You write: "+text)
		// bot.Send(msg)
		if strings.HasPrefix(text, "/list") {
			var message string
			if len(tasks) == 0 {
				message = "No tasks found"
			} else {
				for _, task := range tasks {
					message += fmt.Sprintf("%d. %s | %s\n", task.ID, task.Title, task.Description)
				}
			}
			msg := tgbotapi.NewMessage(chatID, message)
			bot.Send(msg)
		}
		if strings.HasPrefix(text, "/add") {
			parts := strings.Split(text, " ")
			if len(parts) < 3 {
				msg := tgbotapi.NewMessage(chatID, "Invalid format")
				bot.Send(msg)
				continue
			}
			title := parts[1]
			description := strings.Join(parts[2:], " ")
			addTask(title, description)
			msg := tgbotapi.NewMessage(chatID, "Task added")
			bot.Send(msg)
		}
		if strings.HasPrefix(text, "/delete") {
			parts := strings.Fields(text)

			if len(parts) < 2 {
				msg := tgbotapi.NewMessage(chatID, "Write id")
				bot.Send(msg)
				continue
			}

			id, err := strconv.Atoi(parts[1])
			if err != nil {
				msg := tgbotapi.NewMessage(chatID, "Invalid ID")
				bot.Send(msg)
				continue
			}
			deleteTask(id)
			msg := tgbotapi.NewMessage(chatID, "Tasks deleted")
			bot.Send(msg)
		}
	}

	// for {
	// 	fmt.Println("\n=== Tasks Manager===")
	// 	fmt.Println("1. Add Task")
	// 	fmt.Println("2. Show All Tasks")
	// 	fmt.Println("3. Delete Task")
	// 	fmt.Println("4. Exit")
	// 	fmt.Println("Choise action: ")

	// 	var choise int
	// 	fmt.Scan(&choise)

	// 	if choise == 1 {
	// 		var title, description string
	// 		fmt.Println("Enter Title: ")
	// 		fmt.Scan(&title)
	// 		fmt.Println("Enter Description: ")
	// 		fmt.Scan(&description)
	// 		addTask(title, description)
	// 	} else if choise == 2 {
	// 		showAllTasks()
	// 	} else if choise == 3 {
	// 		var id int
	// 		fmt.Println("Enter ID task for delete: ")
	// 		fmt.Scan(&id)
	// 		deleteTask(id)
	// 	} else if choise == 4 {
	// 		fmt.Println("Goodbye!")
	// 		break
	// 	} else {
	// 		fmt.Println("Invalid choise")
	// 	}
	// }
}

// главная функция, с неё всё начинается
// func main() {
// 	height, kg := getInput()

// 	bmi := calculateBMI(height, kg)
// 	outputResult(bmi)
// }

// func calculateBMI(height, weight float64) float64 {
// 	const double = 2
// 	bmi := weight / math.Pow(height/100, double)
// 	return bmi
// }
// func outputResult(bmi float64) {
// 	fmt.Printf("Body Mass Index: %.0f\n", bmi)
// }
// func getInput() (float64, float64) {
// 	var height float64
// 	fmt.Println("Write youre height: ")
// 	fmt.Scan(&height)
// 	var kg float64
// 	fmt.Println("Write your weight: ")
// 	fmt.Scan(&kg)
// 	return height, kg
// }

// type Student struct {
// 	name  string
// 	Grade int
// }

// func (s Student) IsPassed() bool {
// 	return s.Grade >= 50
// }

// func main() {
// 	students := []Student{
// 		{name: "Alice", Grade: 80},
// 		{name: "Bob", Grade: 40},
// 		{name: "Charlie", Grade: 90},
// 	}

// 	for _, student := range students {
// 		fmt.Printf("%s passed: %t\n", student.name, student.IsPassed())
// 	}
// }
