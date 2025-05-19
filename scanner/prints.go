package scanner

import (
	"fmt"
	"todoConsoleApp/todo"

	"github.com/k0kubun/pp"
)

func printHelp() {
	fmt.Println("")
	fmt.Println("----------------------------------")
	fmt.Println("Список команд, которые должны быть доступны в приложении:")
	fmt.Println("")

	fmt.Println("help - эта команда позволяет узнать доступные команды и их формат")
	fmt.Println("")

	fmt.Println("add {заголовок задачи из одного слова} {текст задачи из одного или нескольких слов} ")
	fmt.Println("- эта команда позволяет добавлять новые задачи в список задач")
	fmt.Println("")

	fmt.Println("list - эта команда позволяет получить полный список всех задач")
	fmt.Println("")

	fmt.Println("delete {заголовок существующей задачи} - эта команда позволяет удалить задачу по её заголовку")
	fmt.Println("")

	fmt.Println("done {заголовок существующей задачи} - эта команда позволяет отменить задачу как выполненную")
	fmt.Println("")

	fmt.Println("events - эта команда позволяет получить список всех событий")
	fmt.Println("")

	fmt.Println("exit — эта команда позволяет завершить выполнение программы")
	fmt.Println("----------------------------------")
	fmt.Println("")

}

func printExit() {
	fmt.Println("Завершение программы!")
	fmt.Println("")
}

func printList(tasks *map[string]todo.Task) {
	fmt.Println("Список всех задача:")
	pp.Println(*tasks)
	fmt.Println("")
}

func printAdd(taskName string) {
	fmt.Println("Задача" + "'" + taskName + "'" + "успешно добавлена!")
	fmt.Println("")
}

func printDelete(taskName string) {
	fmt.Println("Задача" + "'" + taskName + "'" + "успешно удалена!")
	fmt.Println("")
}

func printDone(taskName string) {
	fmt.Println("Задача" + "'" + taskName + "'" + "выполнена!")
	fmt.Println("")
}

func printEvents(events []Event) {
	fmt.Println("Список событий: ")
	pp.Println(events)
	fmt.Println("")
}

func printPromt() {
	fmt.Println("Введите команду:")
}

func printError() {
	fmt.Println("Ошибка!")
	fmt.Println("")
}
