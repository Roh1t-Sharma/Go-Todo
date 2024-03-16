document.addEventListener('DOMContentLoaded', fetchTodos);
document.getElementById('todo-form').addEventListener('submit', createTodo);

function fetchTodos() {
    fetch('http://localhost:8080/todos')
        .then(response => response.json())
        .then(todos => {
            const listElement = document.getElementById('todos');
            listElement.innerHTML = ''; // Clear current todos
            todos.forEach(todo => {
                const todoElement = document.createElement('li');
                todoElement.textContent = `${todo.title}: ${todo.description} - Completed: ${todo.completed ? 'Yes' : 'No'}`;
                listElement.appendChild(todoElement);
            });
        });
}

function createTodo(event) {
    event.preventDefault();
    const title = document.getElementById('title').value;
    const description = document.getElementById('description').value;
    fetch('http://localhost:8080/todos', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ id: generateID(), title, description, completed: false }),
    })
        .then(response => response.json())
        .then(() => {
            fetchTodos(); // Refresh_the_todo_list
            document.getElementById('title').value = '';
            document.getElementById('description').value = '';
        });
}

function generateID() {
    return '_' + Math.random().toString(36).substr(2, 9);
}
