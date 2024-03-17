// noinspection JSUnusedLocalSymbols

document.addEventListener('DOMContentLoaded', function() {
    const form = document.getElementById('todo-form');
    const todoList = document.getElementById('todos');

    form.addEventListener('submit', async function(e) {
        e.preventDefault();
        const title = document.getElementById('title').value;
        const description = document.getElementById('description').value;

        // Add_todo_via_API_call
        await addTodo({
            title,
            description,
            completed: false // assuming new todos are not completed by default
        });
        fetchTodos(); // Refresh_the_todo_list
    });

    fetchTodos();
});

async function addTodo(todo) {
    try {
        await fetch('http://localhost:8080/todos', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                ...todo,
                id: Date.now().toString(), // simple ID generation
            }),
        });
    } catch (error) {
        console.error('Error adding todo:', error);
    }
}

async function deleteTodo(todoId) {
    try {
        await fetch(`http://localhost:8080/todos/${todoId}`, {
            method: 'DELETE',
        });
    } catch (error) {
        console.error('Error deleting todo:', error);
    }
}

async function fetchTodos() {
    try {
        const response = await fetch('http://localhost:8080/todos');
        const todos = await response.json();

        const todoList = document.getElementById('todos');
        todoList.innerHTML = ''; // Clear current todos

        todos.forEach(todo => {
            const todoItem = document.createElement('li');
            todoItem.textContent = `${todo.title}: ${todo.description}`;

            const deleteBtn = document.createElement('button');
            deleteBtn.textContent = 'Delete';
            deleteBtn.onclick = function() {
                deleteTodo(todo.id).then(() => fetchTodos()); // Delete and refresh list
            };
            todoItem.appendChild(deleteBtn);

            todoList.appendChild(todoItem);
        });
    } catch (error) {
        console.error('Error fetching todos:', error);
    }
}
