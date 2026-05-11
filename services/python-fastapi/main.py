from fastapi import FastAPI

app = FastAPI()

@app.get("/ping")
def ping():
    return {"message": "pong", "service": "python-fastapi"}

@app.get("/compute")
def compute():
    total = 0
    for i in range(1000000):
        total += i
    return {"result": total, "service": "python-fastapi"}