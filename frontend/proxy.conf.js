const PROXY_CONFIG = {
    "/webapi/v1/*": {
        "target": "http://localhost:9090",
        "secure": false
    }
};

module.exports = PROXY_CONFIG;
