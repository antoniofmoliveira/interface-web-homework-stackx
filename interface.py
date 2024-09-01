import requests
import os
from dotenv import load_dotenv
from pymongo import MongoClient

try:
    url = "https://randomuser.me/api/?results=5&inc=name,email,dob"
    response = requests.get(url)
    data = response.json()
    print(data)
except Exception as e:
    print("error:", e)
    exit(1)

try:
    load_dotenv()
    MONGODB_URI = os.getenv('MONGODB_URI')
    print(MONGODB_URI)
except Exception as e:
    print("error:", e)
    exit(1)

try:
    client = MongoClient(MONGODB_URI)
    db = client["stackx"]
except Exception as e:
    print("error:", e)
    exit(1)

try:
    users_collection = db["users"]
    ids = users_collection.insert_many(data['results'])
    print("Inserted multiple documents: ", ids)
except Exception as e:
    print("error:", e)
    exit(1)

client.close()
