package todo

type List struct {
	Tasks map[string]Task
}

func NewList() *List {
	return &List{
		Tasks: make(map[string]Task),
	}
}

func (l *List) AddTask(task *Task) string {
	l.Tasks[task.Name] = *task

	return ""
}

func (l *List) DeleteTask(name string) string {
	_, ok := l.Tasks[name]
	if !ok {
		return unknownTask
	}
	delete(l.Tasks, name)

	return ""
}

func (l *List) DoneTask(name string) string {
	task, ok := l.Tasks[name]
	if !ok {
		return unknownTask
	}

	task.Done()
	l.Tasks[name] = task

	return ""
}

func (l *List) ListTasks() *map[string]Task {
	return &l.Tasks
}
