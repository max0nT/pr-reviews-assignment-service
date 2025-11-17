# ERD

```mermaid
 erDiagram

   team ||--o{ user: has
   user ||--o{ pull_request: has
   reviewer o{--|| user: "assigned as"
   reviewer o{--|| pull_request: "assigned to"

 user {
    id bigint
    username varchar(50)
    team_name string
    is_active boolean
 }
 team {
    id bigint
    name string
 }
 pull_request {
    id bigint
    name varchar(50)
    status varchar(50)
    created_by bigint
    created_at timestamp
    merged_at timestamp
 }
 reviewer {
   id bigint
   reviewer_id bigint
   pr_id bigint
 }

```
