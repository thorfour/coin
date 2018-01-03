var DEBUG = false
var executable = "./coin";
var cp = require("child_process");
var request = require('request')
var headers = {
    'User-Agent': 'Super Agent/0.0.1',
    'Content-Type': 'application/json',
}

exports.handler = function(req, res) {

    // spawn routine
    var proc = cp.spawnSync(executable, [], {stdio: 'pipe', encoding: "utf8"});
    var quote = proc.stdout;

    var respType = "in_channel";
    // Check for no response, means there was an error
    if (quote === "") {
        quote = proc.stderr
        respType = "ephemeral";

        if (DEBUG) {
            console.log(proc.stderr);
        }
    }

    // Parse quote into json for slack
    var resp = '{ "response_type" : "' + respType + '", "text" : "' + quote + '" }';
    if (DEBUG) {
        console.log(resp)
    }

    var options = {
        url: req.body.response_url,
        method: 'POST',
        headers: headers,
        form: resp
    }

    if (DEBUG) {
        console.log(req.body);
    }

    // Return json
    request(options);

    res.status(200).end();
};
