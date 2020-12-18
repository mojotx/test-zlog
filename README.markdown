# Sample code for demonstrating bug with GoLand IDE

I noticed that when using [Zerolog](https://github.com/rs/zerolog) and [GORM](https://v1.gorm.io/docs/),
the GoLand 2020.3 IDE does not recognize that `log.Fatal().Msg()` calls `os.Exit(1)`, exiting the program.

The warning given is something like:

'db' may have 'nil' or other unexpected value as its corresponding error variable may be not 'nil'

```go
	db, err = gorm.Open("mysql", sqlString)
	if err != nil {
		log.Fatal().Msgf("Unable to connect to '%s'", sqlString)
	}
	// Next line is flagged as using 'db' with 'nil' value
	db = db.BlockGlobalUpdate(true)
	defer func() {
		_ = db.Close()
	}()
```
