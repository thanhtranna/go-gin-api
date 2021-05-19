## DDM

Dynamic Data Masking (DDM for short) can prevent sensitive data from being exposed to unauthorized users.

| Type | Requirements | Example | Description
| ---- | ---- | ---- | ----
| Mobile phone number | First 3 and last 4 | 132****7986 | Fixed length 11 digits
| Email Address | First 1 Next 1 | l**w@gmail.com | Only the mailbox name before @ is masked
| Name | Incognito | *Hong Zhang | Hide the last name
| Password | No output | ****** |
| Bank card number | First 6 and last 4 | 622888******5676 | Bank card number up to 19 digits
| ID number | First 1 After 1 | 1******7 | Fixed length 18 digits

#### Code example

```
// return value
type message struct {
	Email ddm.Email `json:"email"`
}

msg := new(message)
msg.Email = ddm.Email("xinliangnote@163.com")
...

```