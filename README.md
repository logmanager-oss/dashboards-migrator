# dashboards-migrator

## Description

Dashboards migrator is a tool for migrating dashboards from Logmanager 3 to Logmanager 4.

## How it works

Dashboards migrator takes a dashboard exported from LM3 as input, iterates over its contents and migrates it to version compliant with LM4.

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
