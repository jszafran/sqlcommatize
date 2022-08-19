# SQL Commatizer

## Why this project was created?

If you work a lot with SQL then you probably often have a situation that you're 
copying some column values (i.e. from SQL query result or Excel sheet) and then paste them into SQL IN clause.

For example you copy 3 rows from Excel that contains values 1, 2, 3.
```sql
SELECT * 
FROM foo
WHERE 1=1
AND ID IN (
    -- <the list of IDs that you paste>
    1
    2
    3
)
```

This is not a valid SQL syntax though.

In order for your query to work, you'll need to delimit your records with commas.
```sql
SELECT * 
FROM foo
WHERE 1=1
  AND ID IN (
    1,
    2,
    3
)
```

If `ID` is a string type, then you'll also need to wrap values with single quotes:
```sql
SELECT * FROM foo
WHERE ID IN (
    '1',
    '2',
    '3'
)
```

Handling large number of records manually becomes very cumbersome quickly.

`SQL Commatizer` does the tedious work for you.

## How it works?
`SQL Commatizer` is a command-line tool that:
* Reads the content of your clipboard
* Adds commas (and optionally single quotes) to records
* Pastes processed records back to clipboard

It was written in Golang and compiles to a single binary file.
It's available for various OSes: Linux, MacOS, Windows.


