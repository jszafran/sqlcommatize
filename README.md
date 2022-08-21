# sqlcommatize

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

`sqlcommatize` does the tedious work for you.

## How does it work?
`sqlcommatize` is a command-line tool that:
* Reads the content of your clipboard
* Adds commas (and optionally single quotes) to records
* Pastes processed records back to clipboard

It is written in Golang and compiles to a single binary file.
Multiplatform support (Linux, MacOS, Windows) is achieved thanks to `golang.design/x/clipboard` module.



## Usage

Command
```
./sqlcommatize --help
```

prints available flags:
```
Usage of sqlcommatize:
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
./sqlcommatize
```

produces:

```
1,
2,
3
```

### `-strings` flag

```bash
./sqlcommatize -strings
```

`-strings` flag wraps each record with single quotes:

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

### `-leading_commas` flag
`-leading_commas` changes the comma styling from trailing to leading:

```
1
,2
,3
```

### Keyboard shortcuts
For usage convenience, you can create 2 keyboard shortcuts:
* one for numeric values `/path/to/sqlcommatize` (or optionally `/path/to/sqlcommatize -leading_commas` for leading commas)
* one for string values `/path/to/sqlcommatize -strings` (`/path/to/sqlcommatize -leading_commas -strings`)

and use them quickly whenever you need to paste some rows into SQL query.
