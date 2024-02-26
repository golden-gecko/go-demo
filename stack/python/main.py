from flask import Flask, json


api = Flask(__name__)


@api.route('/', methods=['GET'])
def get_companies():
    data = {
        "message": "Hello"
    }

    return json.dumps(data)


if __name__ == "__main__":
    api.run(host="0.0.0.0", port=4000)
