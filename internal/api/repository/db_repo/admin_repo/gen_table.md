#### go_gin_api.admin
Administrator table

| Serial number | name | description | type | key | empty | extra | default value |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | primary key | int(11) unsigned | PRI | NO | auto_increment | |
| 2 | username | username | varchar(32) | UNI | NO | | |
| 3 | password | password | varchar(32) | | NO | | |
| 4 | nickname | nickname | varchar(60) | | NO | | |
| 5 | mobile | Mobile number | varchar(20) | | NO | | |
| 6 | is_used | whether to enable 1: Yes -1: No | tinyint(1) | | NO | | 1 |
| 7 | is_deleted | whether to delete 1: Yes -1: No | tinyint(1) | | NO | | -1 |
| 8 | created_at | created time | timestamp | | NO | | CURRENT_TIMESTAMP |
| 9 | created_user | created by | varchar(60) | | NO | | |
| 10 | updated_at | update time | timestamp | | NO | on update CURRENT_TIMESTAMP | CURRENT_TIMESTAMP |
| 11 | updated_user | updated by | varchar(60) | | NO | | |