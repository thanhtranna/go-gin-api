## example

```cassandraql
1.
query {
  bySex(sex: "male") {
    id
    name
    sex
    mobile
  }
}

2.
mutation {
  updateUserMobile(data: {id: "1", mobile: "13299999999"}) {
    id
    name
    sex
    mobile
  }
}

```