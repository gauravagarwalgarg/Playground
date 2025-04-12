package main

import "fmt"

func printTasks(taskItems []string, itemLimit int) {
  fmt.Println("List of my Todos")
  for index, task := range taskItems {
	  fmt.Printf("%d: %s\n", index+1, task)
  }
}

func addTask(taskItems []string, newTask string) []string {
		var updatedTaskItems = append(taskItems, newTask)
		return updatedTaskItems

}

func main() {
		var taskOne = "Check the Go Course"
		//no := 10
		fmt.Println("Welcome to our TODO-list App!")

		var taskItems = []string {"Check the Go Course", "Sleep", "Eat!"}

		printTasks(taskItems)

		for _, tasks := range taskItems {
				fmt.Println(tasks);
		}

		fmt.Println("List of my TODO's:")
		fmt.Println(taskOne)
		fmt.Println("Sleep")
		fmt.Println("Eat!")
		fmt.Println("Repeat!")

		fmt.Println()

		fmt.Println("Tutorials")
		fmt.Println(taskOne)
		fmt.Println("Create a TODO")
		
		fmt.Println()

		fmt.Println("Reward")
		fmt.Println(taskOne)
		fmt.Println("Eat")
}
