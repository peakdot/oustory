var colors = {
    status: {
        pending: "blue lighten-2",
        todo: "transparent",
        rejected: "red lighten-2",
        in_progress: "orange lighten-2",
        done: "light-green",
    }
};

var statuses = {
    backlog: ["pending", "todo", "in_progress", "done", "rejected"],
    subtask: ["todo", "in_progress", "done"],
}

export default ({ app }, inject) => {
    inject('colors', colors)
    inject('statuses', statuses)
}
