# Instructions

**You need a PostgreSQL server**

 Run SQL script [Create table script](../contrib/sql/db.postgres.sql)



## Run service

1. Download dependencies 

    ```
    go mod download
    ```
2. Create config.yml in ./contrib

3. Build

    ```
    go build -o main .
    ```
4. Run
    ```
    ./main 
    ```




