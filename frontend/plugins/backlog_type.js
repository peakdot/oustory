var zeroPad = function (num, places) {
    var zero = places - num.toString().length + 1;
    return Array(+(zero > 0 && zero)).join("0") + num;
}

var backlogNumber = function (backlog) {
    var num = zeroPad(backlog.Number, 3);
    switch (backlog.Type) {
        case "user_story":
            return "U" + num;
        case "technical_task":
            return "T" + num;
        case "bug":
            return "B" + num;
        default:
            return num
    }
}

var backlogEmoji = function (backlogType) {
    switch (backlogType) {
        case "user_story":
            return '🌱';
        case "technical_task":
            return '🔧';
        case "bug":
            return '🐞';
        default:
            return '🔨';
    }
}

export default ({ app }, inject) => {
    inject('backlogNumber', backlogNumber)
    inject('backlogEmoji', backlogEmoji)
}