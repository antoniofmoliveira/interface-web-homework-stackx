# Tarefa StackX - interface web

## Objetivo: 

O Obetivo desta tarefa é criar uma interface web que consuma os dados da API http://randomuser.meLinks to an external site.

## Dicas:

📌Conforme mostrado na aula 3 do módulo. Você deve consultar a API e tratar os dados para pegar somente o nome, email, data de nascimento e idade.

📌 Você deve tratar esses dados e enviar para um banco de dados próprio. Divirtam-se.


## The homework


the url ```https://randomuser.me/api/?results=5&inc=name,email,dob```

returns 

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

Documents in Collection stackx.user

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
