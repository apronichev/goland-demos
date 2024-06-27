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
        },
        gotoEdit(id) {
            this.$router.push(`/edit/${id}`)
        },
        addTodo() {
            this.$router.push(`/add`)
        }
    },
    mounted() {
        this.loadTodos()
    },
    template: `
      <h1>Todo List</h1>
      <table>
        <tr>
          <th>ID</th>
          <th>Title</th>
          <th>Due</th>
          <th>Done</th>
          <th>Action</th>
        </tr>
        <tr v-for="todo in todos">
          <td>{{ todo.id }}</td>
          <td>{{ todo.title }}</td>
          <td>{{ todo.due }}</td>
          <td>{{ todo.done }}</td>
          <td>
            <button @click="gotoEdit(todo.id)">Edit</button>
            <button @click="deleteTodo(todo)">Delete</button>
          </td>
        </tr>
      </table>
      <button @click="addTodo">Add Todo</button>
    `
}