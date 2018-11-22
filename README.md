# Archiver for RealMac Software's Clear

**USE AT YOUR OWN RISK**

_**This is untested software.**_

DISCLAIMER: THIS SOFTWARE IS PROVIDED “AS IS” WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED. ALL EXPRESS OR IMPLIED REPRESENTATIONS, CONDITIONS AND WARRANTIES, INCLUDING ANY IMPLIED WARRANTY OF MERCHANTABILITY OR FITNESS FOR A PARTICULAR PURPOSE, ARE DISCLAIMED, EXCEPT TO THE EXTENT THAT SUCH DISCLAIMERS ARE DETERMINED TO BE ILLEGAL.

No affiliation with RealMac Software. This is a bootstrapped Go program that:

1. Reads the Clear database from `$HOME/Library/Containers/com.realmacsoftware.clear.mac/Data/Library/Application Support/com.realmacsoftware.clear.mac/LocalTasks.sqlite`
2. Saves data from the `tasks` and `completed_tasks` to CSV files.

## Requirements

This Go application is built with:

- Go 1.11.2 darwin/amd64
- A laptop running macOS 10.14.2

## Usage

To use, download the binary from https://github.com/zeddee/clear-archiver/releases and run it in the terminal:

```bash
./clear-archiver
```

## Notes

Clear uses an embedded sqlite3 database that's usually saved in:

```
$HOME/Library/Containers/com.realmacsoftware.clear.mac/Data/Library/Application Support/com.realmacsoftware.clear.mac/LocalTasks.sqlite
```

It contains the following tables:

- `tasks`: All current tasks in Clear.
- `completed_tasks`: All tasks marked as "complete" in Clear.
- `lists`: All lists of tasks in Clear.
- `modelmeta_int`: ?
- `task_reminders`: Presumably, information on tasks that have reminders set.
- `version`: ? Presumably metadata related to task versioning.

## TODO

- Allow user to select location of Clear database.
- Allow user to save entries from `lists` database.
- Allow user to set output files for extracted data.
- Allow user to set destination of log output.
