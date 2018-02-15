var obj = require('./data');
var request = require('request');

for (var i = 2; i < obj.data.users.length; i++) {
    var current = obj.data.users[i];
    if (current.user.question) {
        request({
            url: 'BASE_SPONSORSHIP_PORTAL_URL/addParticipant',
            method: 'POST',
            json: {
                "name": current.user.name,
                "email": current.user.email,
                "resumeId": current.user.question.file.path.replace("uploads/", ""),
                "token": "INSERT TOKEN HERE"
                }},
            function (error, response, body) {
                if (error) console.log(error);
                if (!error && response.statusCode == 200) {
                    console.log(body)
                }
            }
        );
    }
}
