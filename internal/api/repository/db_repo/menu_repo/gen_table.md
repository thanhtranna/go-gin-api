#### go_gin_api.menu 
left menu bar table

| Serial number | Name | Description | Types of | Key | Empty | Extra | Default |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | primary key | int unsigned | PRI | NO | auto_increment |  |
| 2 | pid | parent id | int unsigned |  | NO |  | 0 |
| 3 | name | menu name | varchar(32) |  | NO |  |  |
| 4 | link | link address | varchar(100) |  | NO |  |  |
| 5 | icon | icon | varchar(60) |  | NO |  |  |
| 6 | level | menu type 1: first level menu 2: second level menu | tinyint unsigned |  | NO |  | 1 |
| 7 | is_used | enable 1: yes -1: no | tinyint(1) |  | NO |  | 1 |
| 8 | is_deleted | delete 1: yes -1: no | tinyint(1) |  | NO |  | -1 |
| 9 | created_at | created time | timestamp |  | NO | DEFAULT_GENERATED | CURRENT_TIMESTAMP |
| 10 | created_user | founder | varchar(60) |  | NO |  |  |
| 11 | updated_at | update time | timestamp |  | NO | DEFAULT_GENERATED on update CURRENT_TIMESTAMP | CURRENT_TIMESTAMP |
| 12 | updated_user | updater | varchar(60) |  | NO |  |  |
