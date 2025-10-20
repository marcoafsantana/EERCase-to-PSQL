# EERCase-to-PSQL

A small tool to convert an EER project (exported as JSON from EERCase) into PostgreSQL DDL.

This repository contains a DDL generator that reads `project.json`, prints a human-readable summary of the model, and generates `tbl_create.sql` with steps to create tables, primary keys and handling of weak entities.

## Features
- Generates `CREATE TABLE` statements from entities and attributes described in the JSON.
- Generates `ALTER TABLE ... ADD PRIMARY KEY` for strong entities and super-entities.
- Adds identifying owner columns and composite primary keys for weak entities.
- Prints a readable model report (module `printer`).

## Requirements
- Go 1.20+ (the module is set to go 1.24 in `go.mod`).

## Quick start
1. Replace or edit `project.json` with the JSON exported from EERCase.
2. Run the application: `go run main.go` or `go build && ./EERCase-to-PSQL`.
3. The program prints a project summary to the terminal and writes `tbl_create.sql` to the repository root.

Note: The generator currently produces basic DDL (types, primary keys and weak-entity handling). Foreign keys, indexes and other advanced constraints are not generated automatically yet.

## Project structure
- `main.go` — program entrypoint that unmarshals `project.json` and calls `printer` and `sqlgen`.
- `project.json` — sample EER project used as input.
- `printer/` — code that prints a human-friendly summary of the EER model.
- `sqlgen/` — code that generates SQL (`tbl_create.sql`).
- `models/` — EER model structs used for parsing the JSON.
- `tbl_create.sql` — generated output (usually ignored in git).

## Development notes
- Code is split into small packages for readability and easier maintenance.
- To test different scenarios, modify `project.json` and run the generator to validate behavior for strong, weak and multivalued attributes.

## Contributing
Pull requests are welcome. For larger changes, please open an issue first to discuss the design.

## License
Add your preferred license here (for example, MIT) or remove this section if you prefer.
