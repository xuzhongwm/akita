package datarecording

import (
	"database/sql"
	"fmt"
	"os"
	"reflect"
	"strings"

	// Need to use SQLite connections.
	"github.com/fatih/structs"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/xid"
	"github.com/tebeka/atexit"
)

// DataRecorder is a backend that can record and store data
type DataRecorder interface {
	//Init establishes a connection to the database
	Init()

	//CreateTable creates a new table with given filename
	CreateTable(table string, sampleEntry any)

	//DataInsert writes a same-type task into table that already exists
	DataInsert(table string, entry any)

	//ListTable returns a slice containing names of all tables
	ListTables() []string

	//Flush flushes all the baffered task into database
	Flush()
}

// SQLiteWriter is the writer that writes data into SQLite database
type SQLiteWriter struct {
	*sql.DB
	statement *sql.Stmt

	dbName     string
	tables     map[string][]any
	batchSize  int
	tableCount int
	taskCount  int
}

// NewSQLiteWriter creates a new SQLiteWriter.
func NewSQLiteWriter(path string) *SQLiteWriter {
	w := &SQLiteWriter{
		dbName:    path,
		batchSize: 100000,
		tables:    make(map[string][]any),
	}

	atexit.Register(func() { w.Flush() })

	return w
}

// Init establishes a connection to the databse
func (t *SQLiteWriter) Init() {
	if t.dbName == "" {
		t.dbName = "akita_data_recording_" + xid.New().String()
	}

	filename := t.dbName + ".sqlite3"
	_, err := os.Stat(filename)
	if err == nil {
		panic(fmt.Errorf("file %s already exists", filename))
	}

	fmt.Fprintf(os.Stderr, "Database created for recording: %s\n", filename)

	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		panic(err)
	}

	t.DB = db
}

func (t *SQLiteWriter) CreateTable(table string, sampleEntry any) {
	t.tableCount++
	n := structs.Names(sampleEntry)
	fields := strings.Join(n, ", \n\t")
	tableName := table
	createTableSQL := `CREATE TABLE ` + tableName + ` (` + "\n\t" + fields + "\n" + `);`

	t.mustExecute(createTableSQL)
	fmt.Printf("Table %s created successfully\n", tableName)

	storedTasks := []any{sampleEntry}
	t.tables[tableName] = storedTasks
	t.taskCount++
	if t.taskCount >= t.batchSize {
		t.Flush()
	}

}

func (t *SQLiteWriter) DataInsert(table string, entry any) {
	storedTasks, exists := t.tables[table]
	if !exists {
		panic(fmt.Errorf("table %s does not exist", table))
	}

	stdTask := storedTasks[0]
	if reflect.TypeOf(stdTask) != reflect.TypeOf(entry) {
		panic(fmt.Errorf("task %s can't be written into table %s", entry, table))
	}

	storedTasks = append(storedTasks, entry)
	t.tables[table] = storedTasks
	t.taskCount += 1
	if t.taskCount >= t.batchSize {
		t.Flush()
	}
}

func (t *SQLiteWriter) Flush() {
	if t.taskCount == 0 {
		return
	}

	t.mustExecute("BEGIN TRANSACTION")
	defer t.mustExecute("COMMIT TRANSACTION")

	for tableName, storedEntries := range t.tables {
		sampleEntry := storedEntries[0]
		t.prepareStatement(tableName, sampleEntry)
		for _, task := range storedEntries {
			v := structs.Values(task)
			_, err := t.statement.Exec(v...)
			if err != nil {
				panic(err)
			}
		}
	}

	t.taskCount = 0
}

func (t *SQLiteWriter) mustExecute(query string) sql.Result {
	res, err := t.Exec(query)
	if err != nil {
		fmt.Printf("Failed to execute: %s\n", query)
		panic(err)
	}
	return res
}

func (t *SQLiteWriter) prepareStatement(table string, task any) {
	n := structs.Names(task)
	for i := 0; i < len(n); i++ {
		n[i] = "?"
	}
	entryToFill := "(" + strings.Join(n, ", ") + ")"
	sqlStr := "INSERT INTO " + table + " VALUES " + entryToFill

	stmt, err := t.Prepare(sqlStr)
	if err != nil {
		panic(err)
	}

	t.statement = stmt
}

// SQLiteTraceReader is a reader that reads trace data from a SQLite database.
type SQLiteReader struct {
	*sql.DB

	filename string
}

// NewSQLiteTraceReader creates a new SQLiteTraceReader.
func NewSQLiteReader(filename string) *SQLiteReader {
	r := &SQLiteReader{
		filename: filename + ".sqlite3",
	}

	return r
}

// Init establishes a connection to the database.
func (r *SQLiteReader) Init() {
	db, err := sql.Open("sqlite3", r.filename)
	if err != nil {
		panic(err)
	}

	r.DB = db
}

// ListTable returns a slice containing names of all tables
func (r *SQLiteReader) ListTables() []string {
	tableNames := make([]string, 0)
	query := `SELECT name FROM sqlite_master WHERE type='table';`

	rows, err := r.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var tableName string
		err := rows.Scan(&tableName)
		if err != nil {
			panic(err)
		}
		fmt.Println(tableName)
		tableNames = append(tableNames, tableName)
	}

	return tableNames

}
