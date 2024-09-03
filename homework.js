import { MongoClient } from 'mongodb';
import dotenv from 'dotenv'
import { connect, SSLMode } from 'ts-postgres'

async function getDataNoSql() {
    let url = "https://randomuser.me/api/?results=5&inc=name,email,dob"
    const response = await fetch(url)
    const data = await response.json()
    dotenv.config()
    const client = new MongoClient(process.env.MONGODB_URI);
    await client.connect();
    let usersCollections = client.db("users").collection("users")
    await usersCollections.insertMany(data["results"])
    await client.close()
}

async function getDataSql() {
    let url = "https://randomuser.me/api/?results=5&inc=name,email,dob"
    const response = await fetch(url)
    const data = await response.json()
    dotenv.config()
    let pghost = process.env.PG_HOST
    let pgport = Number.parseInt(process.env.PG_PORT)
    let pguser = process.env.PG_USER
    let pgpassword = process.env.PG_PASSWORD
    let pgdatabase = process.env.PG_DATABASE
    let pgssl = process.env.PG_SSL
    if (process.env.PG_SSL === "true") {
        pgssl = { mode: SSLMode.Prefer }
    } else {
        pgssl = SSLMode.Disable
    }
    const client = await connect({
        "host": pghost,
        "port": pgport,
        "user": pguser,
        "password": pgpassword,
        "database": pgdatabase,
        "ssl": pgssl
    })
    const statement = await client.prepare("INSERT INTO users (name, email, dob, age) VALUES ($1, $2, $3, $4)")
    for (const user of data["results"]) {
        await statement.execute([`${user.name.title} ${user.name.first} ${user.name.last}`, `${user.email}`, `${user.dob.date}`, user.dob.age])
    }
    await statement.close();
    client.end()

}

function main() {
    if (process.env.USE_NOSQL === "true") {
        getDataNoSql()
    } else {
        getDataSql()
    }
}

main()