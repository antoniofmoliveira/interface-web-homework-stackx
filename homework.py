#!/usr/bin/env python3

import os
import psycopg
import requests
from dotenv import load_dotenv
from pymongo import MongoClient

load_dotenv()
USE_NOSQL = os.getenv('USE_NOSQL')

if USE_NOSQL == 'true':

    try:
        url = "https://randomuser.me/api/?results=5&inc=name,email,dob"
        response = requests.get(url)
        data = response.json()
        print(data)
    except Exception as e:
        print("error:", e)
        exit(1)

    try:
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

else:

    try:
        url = "https://randomuser.me/api/?results=5&inc=name,email,dob"
        response = requests.get(url)
        data = response.json()
        print(data)
    except Exception as e:
        print("error:", e)
        exit(1)

    try:

        COCKROACHDB_URI = os.getenv('COCKROACHDB_URI')
        print(COCKROACHDB_URI)
    except Exception as e:
        print("error:", e)
        exit(1)

    try:
        conn = psycopg.connect(COCKROACHDB_URI,
                               application_name="$ tarefa")
    except Exception as e:
        print("error:", e)
        exit(1)

    cur = conn.cursor()

    for user in data['results']:
        try:
            name = user["name"]["first"] + " " + user["name"]["last"]
            email = user["email"]
            dob = user["dob"]["date"]
            age = user["dob"]["age"]
            sql = "INSERT INTO users (name, email, dob, age) VALUES (%s, %s, %s, %s)"
            cur.execute(sql, (name, email, dob, age))
            conn.commit()
        except Exception as e:
            print("error:", e)
            exit(1)

    conn.close()
