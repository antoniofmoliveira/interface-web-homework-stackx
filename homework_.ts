import { MongoClient } from 'mongodb';
import { connect, SSL, SSLMode } from 'ts-postgres'
import dotenv from 'dotenv'

dotenv.config()

async function getDataNoSql() {
    const url = "https://randomuser.me/api/?results=5&inc=name,email,dob"
    const response = await fetch(url)
    const data = await response.json()
    if (process.env.MONGODB_URI === undefined) {
        throw new Error('MONGODB_URI not set')
    }
    const client = new MongoClient(process.env.MONGODB_URI);
    await client.connect();
    const usersCollections = client.db("users").collection("users")
    await usersCollections.insertMany(data["results"])
    await client.close()
}

async function getDataSql() {
    const url = "https://randomuser.me/api/?results=5&inc=name,email,dob"
    const response = await fetch(url)
    const data = await response.json()
    if (process.env.COCKROACHDB_URI === undefined) {
        throw new Error('COCKROACHDB_URI not set')
    }
    let pghost = process.env.PG_HOST
    if (pghost === undefined) {
        throw new Error('PG_HOST not set')
    }
    if (process.env.PG_PORT === undefined) {
        throw new Error('PG_PORT not set')
    }
    let pgport = Number.parseInt(process.env.PG_PORT)
    if (process.env.PG_USER === undefined) {
        throw new Error('PG_USER not set')
    }
    let pguser = process.env.PG_USER
    if (process.env.PG_PASSWORD === undefined) {
        throw new Error('PG_PASSWORD not set')
    }
    let pgpassword = process.env.PG_PASSWORD
    if (process.env.PG_DATABASE === undefined) {
        throw new Error('PG_DATABASE not set')
    }
    let pgdatabase = process.env.PG_DATABASE
    if (process.env.PG_SSL === undefined) {
        throw new Error('PG_SSL not set')
    }
    let pgssl: (SSLMode | SSL)
    if (process.env.PG_SSL === "true") {
        pgssl = { mode: SSLMode.Prefer }
    } else {
        pgssl = SSLMode.Disable
    }
    const client = connect({
        "host": pghost,
        "port": pgport,
        "user": pguser,
        "password": pgpassword,
        "database": pgdatabase,
        "ssl": pgssl

    });
    const statement = (await client).prepare("INSERT INTO users (name, email, dob, age) VALUES ($1, $2, $3, $4)")
    for (const user of data["results"]) {
        (await statement).execute([`${user.name.title} ${user.name.first} ${user.name.last}`, `${user.email}`, `${user.dob.date}`, user.dob.age])
    }
    (await statement).close();
    (await client).end()
}

function main() {
    if (process.env.USE_NOSQL === "true") {
        getDataNoSql()
    } else {
        getDataSql()
    }
}

main()