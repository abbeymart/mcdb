# Historical logs

## Versions 0.3.1
- Release date: June 25, 2021
- re-publish refactored types


## Versions 0.3.0
- Release date: June 23, 2021
- refactored types to json-camelCase mapping


## Versions 0.2.6
- Release date: Mar 04, 2021
- added matching json-variables to types


## Versions 0.2.5
- Release date: Jan 26, 2021
- set default sslmode to disable for backward compatibility

## Versions 0.2.4
- Release date: Jan 26, 2021
- Parameterised sslmode and sslCert values

## Versions 0.2.3
- Release date: Dec 14, 2020
- Updated dependencies

## Versions 0.2.2
- Release date: Dec 13, 2020
- Updated dependencies

## Version 0.2.1
- Release date: Dec 11, 2020
- Exposed Db-pgx and DbPool-pgxpool struct dbConn param as DbConn

## Version 0.2.0

- Release date: Dec 08, 2020
- Re-aligned dbConfig type for direct access / updates
- Implemented pgx standard and pool db-connection features
- Tested pg, pgx, pgxPool and mongoDB connection test cases successfully

## Version 0.1.1

- Release date: Dec 06, 2020
- Re-tested with defer CloseDb() for each test cases
- Db Connect module: added MongoDB Connection feature

## Version 0.1.0

- Release date: Dec 05, 2020
- Db Connect module: for RDBMS => PostgresDB, MySQL, SQLite...
