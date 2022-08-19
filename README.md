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
It's available for various operating systems: Linux, MacOS, Windows.

## Usage

```
./sql_commatize --help
```

prints available flags:
```
Usage of sql_commatize:
  -leading_commas
        Use leading commas for separating rows (trailing commas used by default).
  -strings
        Wraps rows with single quotes (for SQL strings).
```
### No arguments provided
When you run the application without any arguments, it uses default arguments:
* only commas are added (input is treated as numbers)
* trailing commas styling is used

For input:
```
1
2
3
```

running

```bash
./sql_commatize_linux_amd64
```

produces:

```
1,
2,
3
```

### `--strings` flag

```bash
./sql_commatize_linux_amd64 --strings
```

`--strings` (or `-strings`) flag wraps each record with single quotes:

```
'1',
'2',
'3'
```

If any records contains a single quotes, program replaces them with double single quotes:

```
What's up?
Foo
Bar
```

would be transformed into:

```sql
'What''s up?',
'Foo',
'Bar'
```

### `--leading_commas` flag
`--leading_commas` (or `-leading_commas`) switches the leading commas styling:

```
1
,2
,3
```