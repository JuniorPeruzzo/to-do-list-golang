package main

import (
	"bufio"
	"fmt"
	"os"
)

var tasks []string

func main() {
	for {
		fmt.Println("1 - Adicionar Tarefa")
		fmt.Println("2 - Listar Tarefas")
		fmt.Println("3 - Marcar Tarefa como Concluída")
		fmt.Println("4 - Sair")
		fmt.Print("Escolha uma opção: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addTask()
		case 2:
			listTasks()
		case 3:
			completeTask()
		case 4:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println("Escolha inválida.")
		}
	}
}

func addTask() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite a tarefa: ")
	task, _ := reader.ReadString('\n')
	tasks = append(tasks, task)
	fmt.Println("Tarefa adicionada!")
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("Nenhuma tarefa para listar.")
		return
	}
	fmt.Println("Tarefas:")
	for i, task := range tasks {
		fmt.Printf("%d. %s", i+1, task)
	}
}

func completeTask() {
	listTasks()
	fmt.Print("Escolha o número da tarefa concluída: ")
	var taskNumber int
	fmt.Scan(&taskNumber)

	if taskNumber > 0 && taskNumber <= len(tasks) {
		tasks = append(tasks[:taskNumber-1], tasks[taskNumber:]...)
		fmt.Println("Tarefa marcada como concluída!")
	} else {
		fmt.Println("Número de tarefa inválido.")
	}
}