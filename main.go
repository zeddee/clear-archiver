package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"path"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

// Application is used to share data
type Application struct {
	db      *sql.DB
	timeNow string
}

func main() {
	app := Application{}
	dbPath := path.Join(os.Getenv("HOME"), "Library/Containers/com.realmacsoftware.clear.mac/Data/Library/Application Support/com.realmacsoftware.clear.mac/LocalTasks.sqlite")
	var err error
	app.db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	db := app.db
	defer db.Close()

	app.timeNow = time.Now().Format("2006Jan02T150405")

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		//	log.Println("Successfully pinged db")
	}

	// Configure log output. We'll stick to the best practice default of sending everything to stdout for now.
	log.SetOutput(os.Stdout)
	/*
		log.SetLevel(log.WarnLevel)
		logFile, err := os.OpenFile("errors.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			log.Warningf("Could not write to errors.log: %s", err)
		}
		log.SetOutput(logFile)
	*/

	// Get Task records from tasks and completed_tasks tables
	if err := app.GetTasks("tasks"); err != nil {
		log.Error(err)
	}

	if err := app.GetTasks("completed_tasks"); err != nil {
		log.Error(err)
	}

	// Get List records from the lists table
	// === DOESN'T WORK. UNABLE TO WRITE OUT TO FILE.
	//	if err := app.GetLists(); err != nil {
	//		log.Error(err)
	//	}
}

// Task is the structure of a task as stored in Clear
type Task struct {
	ID             int    `json:"id"`
	Identifier     string `json:"identifier"`
	ListIdentifier string `json:"list_identifier"`
	Title          string `json:"title"`
	PrevIdentifier string `json:"prev_identifier"`
	NextIdentifier string `json:"next_identifier"`
}

// GetTasks extracts all tasks from a given table in a Clear database
func (app *Application) GetTasks(tableName string) error {
	db := app.db

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s", tableName))
	if err != nil {
		return err
	}

	outputFile := app.timeNow + "_" + tableName + ".csv"
	f, err := os.OpenFile(outputFile, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("Failed to open and write to %s: %s", outputFile, err)
	}
	defer f.Close()
	writer := csv.NewWriter(f)

	writeColumnHeaders(writer, db, tableName)

	for rows.Next() {
		t := Task{}
		if err := rows.Scan(&t.ID, &t.Identifier, &t.ListIdentifier, &t.Title, &t.PrevIdentifier, &t.NextIdentifier); err != nil {
			// Don't kill operation; just log the error and continue extracting records
			log.Warning(t.ID, ":", err)
		}

		// Convert data types of all members to string, to write as CSV record
		record := []string{
			strconv.Itoa(t.ID),
			t.Identifier,
			t.Title,
			t.PrevIdentifier,
			t.NextIdentifier,
		}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("Failed to write record %d to destination csv: %s", t.ID, err)
		}
	}
	return nil
}

// List is the structure of a list as stored in Clear
type List struct {
	ID            int     `json:"id"`
	Identifier    string  `json:"identifier"`
	Title         string  `json:"title"`
	Scroll        float64 `json:"scroll"`
	PrevIdentfier string  `json:"prev_identifier"`
}

// GetLists gets all lists from the Clear sqlite3 database
// (No idea why this is not working)
func (app *Application) GetLists() error {
	db := app.db
	rows, err := db.Query(`SELECT * FROM lists`)
	if err != nil {
		return err
	}

	outputFile := app.timeNow + "_" + "lists.csv"
	f2, err := os.OpenFile(outputFile, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("Failed to open and write to %s: %s", outputFile, err)
	}
	defer f2.Close()
	writer := csv.NewWriter(f2)

	writeColumnHeaders(writer, db, "lists")

	for rows.Next() {
		r := List{}
		if err := rows.Scan(&r.ID, &r.Identifier, &r.Title, &r.Scroll, &r.PrevIdentfier); err != nil {
			// Don't kill operation; just log the error and continue extracting records
			log.Warning(r.ID, ":", err)
		}

		record := []string{
			strconv.Itoa(r.ID),
			r.Identifier,
			r.Title,
			strconv.Itoa(int(r.Scroll)),
			r.PrevIdentfier,
		}
		log.Println(record)
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("Failed to write record %d to destination csv: %s", r.ID, err)
		}
	}
	return nil
}

func writeColumnHeaders(writer *csv.Writer, db *sql.DB, tableName string) error {
	row, err := db.Query(`SELECT * FROM $1 VALUES($1) LIMIT 1`, tableName)
	if err != nil {
		return fmt.Errorf("Failed to access table %s to get column names: %s", tableName, err)
	}
	columnNames, err := row.Columns()
	if err != nil {
		return fmt.Errorf("Failed to get column names for tasks table: %s", err)
	}
	if err := writer.Write(columnNames); err != nil {
		return fmt.Errorf("Failed to write task table headers: %s", err)
	}
	return nil
}
