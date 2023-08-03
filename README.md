# 1099

A golang program to read hours from a google sheet for a given month, summarize hours into each week (start = Monday) and provide a total number of hours

## Configuration

The config file defaults to /usr/local/etc/1099.toml.  An example file is:

`[google]
credfile = "/usr/local/etc/1099.json"
spreadsheet = "<spreadsheet-id>"
rate=150.0
tabname = "Sheet1"
`

The credfile is a json file obtained when creating a google service account

The spreadsheet id is the ID of the spreadsheet - it must be shared with the service account.

