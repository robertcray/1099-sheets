# 1099

A golang program to read hours from a google sheet for a given month, summarize hours into each week (start = Monday) and provide a total number of hours

## Configuration

The config file defaults to /usr/local/etc/1099.toml.  An example file is:

```[google]
credfile = "/usr/local/etc/1099.json"
spreadsheet = "<spreadsheet-id>"
rate=150.0
tabname = "Sheet1"
```

The credfile is a json file obtained when creating a google service account

The spreadsheet id is the ID of the spreadsheet - it must be shared with the service account.

## Output

When run the program assumes last month and the current year.

The tab name is configured in the config file but the program assumes the date is in column "A" and the time is in column "C".  Times are recorded as whole integers ending in either "m" or "h" (minutes or hours).  So for example 1.5 hours would be recorded as 90m, two hours as 2h or 2hr.

Example output:

```Week 1 - 08/01/2022(Mon) - 08/07/2022(Sun)     6 hours, 0 minutes
Week 2 - 08/08/2022(Mon) - 08/14/2022(Sun)     8 hours, 0 minutes
Week 3 - 08/15/2022(Mon) - 08/21/2022(Sun)     8 hours, 0 minutes
Week 4 - 08/22/2022(Mon) - 08/28/2022(Sun)     7 hours, 0 minutes
Week 5 - 08/29/2022(Mon) - 08/31/2022(Wed)     5 hours, 0 minutes

Total: 34 hours, 0 minutes
Total rate: $4080.00 (at $120.00/hr)```

The hourly rate is taken from the config file but can be overriden on the command line.

Recorded time is summarized for each week.

