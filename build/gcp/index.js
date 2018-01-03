var DEBUG = true;
var executable = "./coin";
var cp = require("child_process");

exports.handler = function(req, res) {

    var queryStr = JSON.stringify(req.body);

    if (DEBUG) {
        console.log(queryStr);
    }

    var proc = cp.spawnSync(executable, [queryStr], {stdio: 'pipe', encoding: "utf8"});
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

    res.status(200).send(resp)
};
