# IQueue
_IQueue_ is a web-app which allows you create, manage and track some user queues.
## Installing
1. Install golang.
1. Customize database:
    1. Create PostgreSQL database.
    1. Configure database in `models/db.yml`.
    1. Run `models/init.sql` code for your database.
1. Run app:
    ```bash
      $ go build -o iqueue
      $ ./iqueue
    ```
