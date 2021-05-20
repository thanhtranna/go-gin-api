#### go_gin_api.authorized 
caller table

| Serial number | Name | Description | Types of | Key | Empty | Extra | Default |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | primary key | int unsigned | PRI | NO | auto_increment |  |
| 2 | business_key | caller key | varchar(32) | UNI | NO |  |  |
| 3 | business_secret | caller secret | varchar(60) |  | NO |  |  |
| 4 | business_developer | caller developer | varchar(60) |  | NO |  |  |
| 5 | remark | remarks | varchar(255) |  | NO |  |  |
| 6 | is_used | enable 1: yes -1: no | tinyint(1) |  | NO |  | 1 |
| 7 | is_deleted | delete 1: yes -1: no | tinyint(1) |  | NO |  | -1 |
| 8 | created_at | created time | timestamp |  | NO | DEFAULT_GENERATED | CURRENT_TIMESTAMP |
| 9 | created_user | founder | varchar(60) |  | NO |  |  |
| 10 | updated_at | update time | timestamp |  | NO | DEFAULT_GENERATED on update CURRENT_TIMESTAMP | CURRENT_TIMESTAMP |
| 11 | updated_user | updater | varchar(60) |  | NO |  |  |
