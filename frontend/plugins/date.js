var coolDate = function (app, date) {
    var seconds = Math.floor((new Date() - date) / 1000);
    if (seconds < 86400 * 30) {
        var interval = Math.floor(seconds / 86400);
        if (interval >= 1) {
            if (interval == 1) {
                return app.i18n.i("_time.yesterday")[0];
            }
            return interval + ' ' + app.i18n.i("_time.days_ago");
        }
        interval = Math.floor(seconds / 3600);
        if (interval >= 1) {
            if (interval == 1) {
                return app.i18n.i("_time.hour_ago")[0];
            }
            return interval + ' ' + app.i18n.i("_time.hours_ago");
        }
        interval = Math.floor(seconds / 60);
        if (interval >= 1) {
            if (interval == 1) {
                return app.i18n.i("_time.minute_ago")[0];
            }
            return interval + ' ' + app.i18n.i("_time.minutes_ago");
        }
        return app.i18n.i("_time.just_now")[0];
    } else {
        return fullDate(date);
    }
}

var fullDate = function (date) {
    return date.toLocaleString("mn-MN");
}

export default ({ app }, inject) => {
    // Inject $hello(msg) in Vue, context and store.
    inject('fullDate', dateStr => {
        var date = new Date(dateStr);
        return fullDate(date)
    })
    inject('coolDate', dateStr => {
        var date = new Date(dateStr);
        return coolDate(app, date)
    })
}