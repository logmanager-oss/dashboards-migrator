# dashboards-migrator

## Description

Dashboards migrator is a CLI tool for migrating dashboards from Logmanager 3 to Logmanager 4.

## How it works

Dashboards migrator takes a dashboard exported from LM3 as input, iterates over its contents and migrates it to version compliant with LM4.

You need to provide an input path to exported LM3 dashboard and output path to which dashboards migrator will write LM4 dashboard. Output file generated from dashboard migrator is ready to be imported to LM4.

If you want LM4 dashboard to be targetting specific index-pattern you can provide it by using `-ip` flag but it's not mandatory - index-pattern can also be assigned during dashboard import.

## Disclaimer

**While Dashboards migrator should be able to deal with most common use cases effectively, there is a possibility that more complicated dashboards might not be migrated fully or with errors! Always make sure to import and test your migrated dashboard to make sure it works as expected.**

**Please make sure to report any errors you encounter so we can improve dashboards migrator functionality.**

## Usage

```
Usage of ./dashboards-migrator:
  -h
        Print help
  -i string
        Path to input file containing LM3 dashboard (Mandatory)
  -ip string
        Index-pattern string (Default: empty)
  -o string
        Path to output file containing LM4 dashboard (Mandatory)
```

**Examples:**

1. Read LM3 dashboard from `log-overview.json` file, set index-pattern to `lm-*` and output migration results to `log-overview.ndjson`

`./dashboards-migrator -i log-overview.json -o log-overview.ndjson -ip lm-*`
