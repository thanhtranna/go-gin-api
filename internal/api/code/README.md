## Error code rules

-The error code needs to be defined in the `code` package.

#### The error code is 5 digits

| 1 | 01 | 01 |
| :------ | :------ | :------ |
| Service level error code | Module level error code | Specific error code |

-Service-level error code: <br>
1 digit to indicate, for example, 1 is a system-level error; 2 is a normal error, usually caused by an illegal operation by the user.
<br>
-Module-level error code: <br>
2 digits, for example, 01 is the user module; 02 is the order module.
<br>
-Specific error code:<br>
2 digits for display. For example, 01 means the mobile phone number is illegal; 02 means the verification code is entered incorrectly.