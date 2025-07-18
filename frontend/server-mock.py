url = {
    "v1/user/auth/github" : {
        "accesstoken": "asd"
    },
    "v1/user/auth/github" : {
        "accesstoken": "asd"
    },
    "v1/integrations/all": {
        "integrations": [
            {
                "name": "github",
                "status": "inactive",
            }
        ]
    },
    "v1/user/details" : {
        "orgstats":{"repocount":2344},
        "default_pattern_count":"81",
        "custom_pattern_count":"21"
    },
    "v1/alerts/all?state=open" : {
        "alerts" : [
            {"alert_name":"Atlassian API Token", "state":"active", "source":"GITHUB","source_repo":"reponame", "date":"2 month old", "path":"reponame/path/to/.../view.py#L25"},
            {"alert_name":"Atlassian API Token", "state":"active", "source":"GITHUB","source_repo":"reponame", "date":"2 month old", "path":"reponame/path/to/.../view.py#L25"},
            {"alert_name":"Atlassian API Token", "state":"active", "source":"GITHUB","source_repo":"reponame", "date":"2 month old", "path":"reponame/path/to/.../view.py#L25"}
        ]
    } ,
    "v1/alerts/all?state=open" : {
        "alerts" : [
            {"alert_name":"Atlassian API Token", "state":"active", "source":"GITHUB","source_repo":"reponame", "date":"2 month old", "path":"reponame/path/to/.../view.py#L25"},
            {"alert_name":"Atlassian API Token", "state":"active", "source":"GITHUB","source_repo":"reponame", "date":"2 month old", "path":"reponame/path/to/.../view.py#L25"},
            {"alert_name":"Atlassian API Token", "state":"active", "source":"GITHUB","source_repo":"reponame", "date":"2 month old", "path":"reponame/path/to/.../view.py#L25"}
        ]
    } ,
    "v1/v1/secrets/alert_over_time": {
        "total_alerts":"18,964",
        "compared_to_last_month": {
            "movement":"negative",
            "movement_value": "2.8%"
        },
        "alerts_over_time_count":[
            {"date":"2024-12-01", "value_high":181, "value_medium":50, "value_low":105},
            {"date":"2024-11-01", "value_high":178, "value_medium":70, "value_low":200},
            {"date":"2024-10-01", "value_high":177, "value_medium":90, "value_low":200},
            {"date":"2024-09-01", "value_high":171, "value_medium":180, "value_low":200},
            {"date":"2024-08-01", "value_high":175, "value_medium":130, "value_low":200},
            {"date":"2024-07-01", "value_high":175, "value_medium":99, "value_low":200},
            {"date":"2024-06-01", "value_high":170, "value_medium":100, "value_low":200}
        ]
    },
    "api/v1/secrets/alerts/custom": {
        "open_count":21,
        "closed_count":1,
        "open_alerts" : [
            {"alert_name":"Atlassian API Token", "state":"active", "source":"GITHUB","source_name":"reponame", "date":"2 month old", "path":"reponame/path/to/.../view.py#L25"},
            {"alert_name":"Atlassian API Token", "state":"active", "source":"GITHUB","source_name":"reponame", "date":"2 month old", "path":"reponame/path/to/.../view.py#L25"},
            {"alert_name":"Atlassian API Token", "state":"active", "source":"GITHUB","source_name":"reponame", "date":"2 month old", "path":"reponame/path/to/.../view.py#L25"}
        ],
        "closed_alerts" : [
            {"alert_name":"Atlassian API Token", "state":"active", "source":"GITHUB","source_name":"reponame", "date":"2 month old", "path":"reponame/path/to/.../view.py#L25"},
        ]
    },
    "v1/v1/secrets/coverage": {
        "repositories" : {
            "percent_enabled":"94",
            "enabled":"1.4k",
            "not_enabled":"103"
        },
        "secret_scannin":{
            "percent_protected":"95",
            "active_users":"200",
            "inactive_users":"23"
        }
    },
    "v1/v1/secrets/top_open_alerts_by_repository": [
        {"repository_name":"temp1", "repository_link":"", "open_alerts":511, "critical":48, "high":237, "medium":173, "low":53},
        {"repository_name":"temp2", "repository_link":"", "open_alerts":231, "critical":44, "high":200, "medium":190, "low":130},
        {"repository_name":"temp3", "repository_link":"", "open_alerts":200, "critical":100, "high":200, "medium":150, "low":100},
        {"repository_name":"temp3", "repository_link":"", "open_alerts":200, "critical":100, "high":200, "medium":150, "low":100},
        {"repository_name":"temp3", "repository_link":"", "open_alerts":200, "critical":100, "high":200, "medium":150, "low":100},
        {"repository_name":"temp3", "repository_link":"", "open_alerts":200, "critical":100, "high":200, "medium":150, "low":100},
        {"repository_name":"temp3", "repository_link":"", "open_alerts":200, "critical":100, "high":200, "medium":150, "low":100},
        {"repository_name":"temp3", "repository_link":"", "open_alerts":200, "critical":100, "high":200, "medium":150, "low":100},
        {"repository_name":"temp3", "repository_link":"", "open_alerts":200, "critical":100, "high":200, "medium":150, "low":100},
        {"repository_name":"temp3", "repository_link":"", "open_alerts":200, "critical":100, "high":200, "medium":150, "low":100}
    ],
    "v1/user/create": {
        "auth-token": "TOKEN-123"
    },
    "v1/user/login": {
        "accesstoken": "TOKEN-123"
    },

    "api/integrations/all" : {
    "integrations": [
        {
            "name": "github",
            "status": "inactive",
            "action_url": "https://github.com/login/oauth/authorize?client_id=Iv23liw3m0dtDHQ9OzCH\u0026redirect_uri=http%3A%2F%2Flocalhost%3A8080%2Fgithub%2Fcallback\u0026response_type=code\u0026scope=read%3Aorg+read%3Auser\u0026state=abdcd"
        }
    ]
}
}
from flask import Flask, jsonify, request
from flask_cors import CORS
from time import sleep


app = Flask(__name__)
CORS(app)

@app.route('/<path:subpath>', methods=['GET','POST'])
def mock_api(subpath):
    # sleep(1)
    raw_query = request.query_string.decode("utf-8")
    route = f"{subpath}"
    if len(raw_query) > 1 :
        route = f"{subpath}?{raw_query}"

    print(route)
    if route in url:
        return jsonify(url[route])
    return jsonify({"error": "Route not found"}), 404


@app.route('/', methods=['GET'])
def root():
    return jsonify({"error": "Specify a valid route"}), 400


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000, debug=True)