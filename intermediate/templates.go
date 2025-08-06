package main

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"strings"
)

func simpleSample() {
	tpl, err := template.New("example").Parse("Welcome, {{.name}}! How ate you?\n")
	if err != nil {
		panic(err)
	}

	// define data:
	data := map[string]interface{}{
		"name": "John", // key must match to template content
	}

	if err := tpl.Execute(os.Stdout, data); err != nil {
		panic(err)
	}
}

func main() {
	simpleSample()

	fmt.Print("Enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	//Define templates:
	templates := map[string]string{
		"welcome":      "Welcome {{.name}}! we are glad you joined.",
		"notification": "{{.name}}, you have a new notification: {{.notification}}",
		"error":        "Oops! An error occurred: {{.errorMessage}}",
	}

	parsedTemplates := make(map[string]*template.Template)
	for tplName, tplStr := range templates {
		parsedTemplates[tplName] = template.Must(template.New(tplName).Parse(tplStr))
	}

	for {
		fmt.Println("\nMenue:")
		fmt.Println("1. Join")
		fmt.Println("2. Get notigfication")
		fmt.Println("3. Get Error")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		var data map[string]interface{}
		var tpl *template.Template

		switch choice {
		case "1":
			tpl = parsedTemplates["welcome"] //use correspondent parsed template
			data = map[string]interface{}{"name": name}
		case "2":
			fmt.Print("Enter you notification: ")
			note, _ := reader.ReadString('\n')
			note = strings.TrimSpace(note)

			tpl = parsedTemplates["notification"] //use correspondent parsed template
			data = map[string]interface{}{"name": name, "notification": note}
		case "3":
			fmt.Print("Enter your error message: ")
			errMsg, _ := reader.ReadString('\n')
			errMsg = strings.TrimSpace(errMsg)
			tpl = parsedTemplates["error"] //use correspondent parsed template
			data = map[string]interface{}{"errorMessage": errMsg}
		case "4":
			fmt.Println("Bye!")
			return //os.Exit(0)
		default:
			fmt.Println("Wrong choice. Please select valid option.")
			continue
		}

		if err := tpl.Execute(os.Stdout, data); err != nil {
			fmt.Println("Failed to execute template:", err)
		}
		fmt.Println()
	}
}
