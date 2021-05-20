#### go_gin_api.authorized_api 
authorized interface address table

| Serial number | Name | Description | Types of | Key | Empty | Extra | Default |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | primary key | int unsigned | PRI | NO | auto_increment |  |
| 2 | business_key | caller key | varchar(30) |  | NO |  |  |
| 3 | method | request method | varchar(30) |  | NO |  |  |
| 4 | api | request address | varchar(100) |  | NO |  |  |
| 5 | is_deleted | delete 1: yes -1: no | tinyint(1) |  | NO |  | -1 |
| 6 | created_at | created time | timestamp |  | NO | DEFAULT_GENERATED | CURRENT_TIMESTAMP |
| 7 | created_user | founder | varchar(60) |  | NO |  |  |
| 8 | updated_at | update time | timestamp |  | NO | DEFAULT_GENERATED on update CURRENT_TIMESTAMP | CURRENT_TIMESTAMP |
| 9 | updated_user | updater | varchar(60) |  | NO |  |  |
