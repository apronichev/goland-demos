export default {
    data() {
        return {
            todos: {}
        }
    },
    methods: {
        async loadTodos() {
            let response = await fetch("http://localhost:8080/todos")
            let todos = await response.json()
            this.todos = {}
            for (let todo of todos) {
                this.todos[todo.id] = todo
            }
        },
        async deleteTodo(todo) {
            let response = await fetch(`http://localhost:8080/todos/${todo.id}`, {"method": "DELETE"})
            if (response.ok) {
                delete this.todos[todo.id]
            }
        }
    },
    mounted() {
        this.loadTodos()
    },
    template: `
      <table>
        <tr>
          <th>Todo ID</th>
          <th>Todo Title</th>
          <th>Due</th>
          <th>Done</th>
          <th>Action</th>
        </tr>
        <tr v-for="todo in todos">
          <td>{{ todo.id }}</td>
          <td>{{ todo.title }}</td>
          <td>{{ todo.due }}</td>
          <td>{{ todo.done }}</td>
          <td><button @click="deleteTodo(todo)">delete</button></td>
        </tr>
      </table>
    `
}