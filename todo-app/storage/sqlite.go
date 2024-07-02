package storage

import (
	"database/sql"
	"log"
	"todo-app/model"

	_ "modernc.org/sqlite"
)

// todo invoke implement interface
// todo generate constructor
type SQLiteStorage struct {
	dataSourceName string
	db             *sql.DB
}

func (S *SQLiteStorage) Init() {
	db, err := sql.Open("sqlite", S.dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Query(`
create table if not exists todos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title varchar(255) NOT NULL,
    done BOOLEAN,
    due varchar(255) NOT NULL
)`)
	if err != nil {
		log.Fatal(err)
	}
	S.db = db
}

func NewSQLiteStorage(dataSourceName string) *SQLiteStorage {
	return &SQLiteStorage{dataSourceName: dataSourceName}
}

func (S *SQLiteStorage) Get(id int) *model.TodoItem {
	// todo type S.db.Query(" and obsere how FLCC works great here!
	rows, err := S.db.Query("SELECT id, title, done, due FROM todos WHERE id = ?", id)
	if err != nil {
		return nil
	}
	defer rows.Close()
	todo := model.TodoItem{}
	if rows.Next() {
		err = rows.Scan(&todo.Id, &todo.Title, &todo.Done, &todo.Due)
		if err != nil {
			return nil
		}
	}
	return &todo
}

func (S *SQLiteStorage) GetAll() []*model.TodoItem {
	rows, err := S.db.Query(`select id, title, done, due from todos`)
	if err != nil {
		return nil
	}
	defer rows.Close()
	result := make([]*model.TodoItem, 0)
	for rows.Next() {
		todo := model.TodoItem{}
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Done, &todo.Due)
		if err != nil {
			return nil
		}
		result = append(result, &todo)
	}
	return result
}

func (S *SQLiteStorage) Put(item *model.TodoItem) *model.TodoItem {
	if item.Id == NO_ID {
		// todo extract insert() function
		_, err := S.db.Query(` insert into todos(title, done, due) values (?, ?, ?) `,
			item.Title, item.Done, item.Due)
		if err != nil {
			return nil
		}
		return item
	} else {
		// todo extract update() function
		_, err := S.db.Query(`update todos set title = ?, done = ?, due = ? where id = ?`,
			item.Title, item.Done, item.Due, item.Id)
		if err != nil {
			return nil
		}
		return item
	}
}

func (S *SQLiteStorage) Remove(id int) *model.TodoItem {
	_, err := S.db.Query(`delete from todos where id = ?`, id)
	if err != nil {
		return nil
	}
	return nil
}
