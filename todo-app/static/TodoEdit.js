export default {
    data() {
        return {
            error: "",
            todo: {
                id: 0,
                title: "",
                due: "",
                done: false
            }
        }
    },
    methods: {
        async loadTodo(id) {
            let response = await fetch(`http://localhost:8080/todos/${id}`)
            this.todo = await response.json()
        },
        async updateTodo() {
            this.error = ""
            let response = await fetch(`http://localhost:8080/todos/${this.todo.id}`, {
                "method":"POST",
                "body": JSON.stringify(this.todo)
            })
            if (!response.ok) {
                this.error = await response.text()
                return
            }
            this.$router.push(`/`)
        }
    },
    mounted() {
        if (this.$route.params.id) {
            this.loadTodo(this.$route.params.id)
        }
    },
    template: `
      <h1>{{ todo.id ? "Edit Todo" : "Create New Todo" }}</h1>
      <form>
        <article v-if="error">
          <aside>{{ error }}</aside>
        </article>
        <label for="title">Title</label><input id="title" type="text" v-model="todo.title">
        <label for="due">Due</label><input id="due" type="text" v-model="todo.due">
        <label for="done">Done</label><input id="done" type="checkbox" v-model="todo.done">
        <br/>
        <button @click.prevent="updateTodo">Save</button>
        &nbsp;
        <button @click.prevent="this.$router.push('/')">Cancel</button>
      </form>
    `
}