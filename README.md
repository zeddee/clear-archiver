# Archiver for RealMac Software's Clear

No affiliation with RealMac Software. This is a bootstrapped Go program that:

1. Reads the Clear database from `$HOME/Library/Containers/com.realmacsoftware.clear.mac/Data/Library/Application Support/com.realmacsoftware.clear.mac/LocalTasks.sqlite`
2. Saves data from the `tasks` and `completed_tasks` to CSV files.

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
