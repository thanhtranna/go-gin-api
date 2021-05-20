#### go_gin_api.admin 
administrator table

| Serial number | Name | Description | Types of | Key | Empty | Extra | Default |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | primary key | int unsigned | PRI | NO | auto_increment |  |
| 2 | username | username | varchar(32) | UNI | NO |  |  |
| 3 | password | password | varchar(100) |  | NO |  |  |
| 4 | nickname | nickname | varchar(60) |  | NO |  |  |
| 5 | mobile | phone number | varchar(20) |  | NO |  |  |
| 6 | is_used | is used 1: yes -1: no | tinyint(1) |  | NO |  | 1 |
| 7 | is_deleted | delete 1: yes -1: no | tinyint(1) |  | NO |  | -1 |
| 8 | created_at | created time | timestamp |  | NO | DEFAULT_GENERATED | CURRENT_TIMESTAMP |
| 9 | created_user | founder | varchar(60) |  | NO |  |  |
| 10 | updated_at | update time | timestamp |  | NO | DEFAULT_GENERATED on update CURRENT_TIMESTAMP | CURRENT_TIMESTAMP |
| 11 | updated_user | updater | varchar(60) |  | NO |  |  |
