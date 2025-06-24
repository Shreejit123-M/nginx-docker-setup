from flask import Flask

app = Flask(__name__)

@app.route('/')
def hello():
    return "✅ Hello from Python service on /service2!"

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8002)
