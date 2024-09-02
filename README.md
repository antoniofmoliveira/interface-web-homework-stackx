# Tarefa StackX - interface web

## Objetivo: 

O Objetivo desta tarefa é criar uma interface web que consuma os dados da API http://randomuser.meLinks to an external site.

## Dicas:

📌Conforme mostrado na aula 3 do módulo. Você deve consultar a API e tratar os dados para pegar somente o nome, email, data de nascimento e idade.

📌 Você deve tratar esses dados e enviar para um banco de dados próprio. Divirtam-se.


## The homework

the url ```https://randomuser.me/api/?results=5&inc=name,email,dob```

returns:

```json
{
    "results": [
        {
            "name": {
                "title": "Mr",
                "first": "Nils",
                "last": "Charles"
            },
            "email": "nils.charles@example.com",
            "dob": {
                "date": "1971-02-14T07:24:27.451Z",
                "age": 53
            }
        },

        ...

    ],
    "info": {
        "seed": "2e9ff79bec090711",
        "results": 1,
        "page": 1,
        "version": "1.4"
    }
}
```

Documents in a **NoSQL Collection**:

```json
[
  {
    "_id": "66d210043b6ac374ae2d7d7b",
    "name": {
      "title": "Mrs",
      "first": "Silke",
      "last": "Mortensen"
    },
    "email": "silke.mortensen@example.com",
    "dob": {
      "date": "1970-10-28T04:12:55.979Z",
      "age": 53
    }
  },
  ...
  {
    "_id": "66d210043b6ac374ae2d7d7f",
    "name": {
      "title": "Mr",
      "first": "Christian",
      "last": "Omahony"
    },
    "email": "christian.omahony@example.com",
    "dob": {
      "date": "1996-12-01T15:21:23.944Z",
      "age": 27
    }
  }
]
```
or a json representation of a row in a table in a **SQL database**:

```json
{
  "id": "1e8ab1d4-efbc-467d-a904-abc85f1a0fb5",
  "name": "Svyatovid Gotra",
  "email": "svyatovid.gotra@example.com",
  "dob": "1989-01-25T19:47:08.583Z",
  "age": "35"
}
```

also the **SQL database** needs preparation:

```sql
CREATE DATABASE stackx;
CREATE USER *******;
USE stackx;
CREATE TABLE users (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name STRING NOT NULL,
    email STRING NOT NULL,
    dob STRING NOT NULL,
    age INT NOT NULL
)
GRANT ALL ON stackx.* TO *******;
```

the `.env` should be:

```bash
MONGODB_URI="mongodb://user:password@0.0.0.0/?retryWrites=true&w=majority"
COCKROACHDB_URI="postgresql://user:password@0.0.0.0:0000/stackx?sslmode=disable"
USE_NOSQL="true"
```
A *Golang 1.23.0* ~/.local instalation.

A venv with *Python 3.12.3*.

**Docker** *containers* were used for the databases *MongoDB 7.0.12* and *CockroachDB 20.1*.

## Running

```bash
$ go run .
```

or 

```bash
$ go run homework.go
```

or compile and install in your gobin folder. 

GOOS = "aix", "android", "darwin", "dragonfly", "freebsd", "illumos", "ios", "js", "linux", "netbsd", "openbsd", "plan9", "solaris", "windows".

GOARCH = "386", "amd64", "arm", "arm64", "mips", "mips64", "mips64le", "mipsle", "ppc64", "ppc64le", "riscv64", "s390x", "wasm".

```bash
$ GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build  -ldflags="-w -s" -o /homework
$ go install
$ homework
```

For Python:

```bash
$ python homework.py
```

or if python interpreter is in your `PATH` variable

```
$ ./homework.py
```
