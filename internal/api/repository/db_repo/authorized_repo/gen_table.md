#### go_gin_api.authorized
Authorized caller table

| Serial number | name | description | type | key | empty | extra | default value |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | primary key | int(11) unsigned | PRI | NO | auto_increment | |
| 2 | business_key | caller key | varchar(32) | | NO | | |
| 3 | business_secret | caller secret | varchar(60) | | NO | | |
| 4 | business_developer | caller developer | varchar(60) | | NO | | |
| 5 | remark | remarks | varchar(255) | | NO | | |
| 6 | is_used | whether to enable 1: yes -1: no | tinyint(1) | | NO | | -1 |
| 7 | is_deleted | whether to delete 1: yes -1: no | tinyint(1) | | NO | | -1 |
| 8 | created_at | created time | timestamp | | NO | | CURRENT_TIMESTAMP |
| 9 | created_user | created by | varchar(60) | | NO | | |
| 10 | updated_at | update time | timestamp | | NO | on update CURRENT_TIMESTAMP | CURRENT_TIMESTAMP |
| 11 | updated_user | updated by | varchar(60) | | NO | | |