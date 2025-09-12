package db

import (
	"database/sql"
	"encoding/json"
	"flipfloppy1/quDnD/src/statblock"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	_ "modernc.org/sqlite"
)

func NewSqliteHandler(logger *log.Logger) (*DbHandler, error) {
	path, err := os.UserConfigDir()
	if err != nil {
		return &DbHandler{}, err
	}
	path = filepath.Join(path, "quDnD", "db", "sqlite.db")
	logger.Println("Path is:", path)
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return &DbHandler{}, err
	}
	err = db.Ping()
	if err != nil {
		logger.Println("Error pinging database:", err.Error())
	}
	_, err = db.Exec(`
		create table if not exists custompages (
			pageid integer primary key,
			document text not null
		) strict;
		`)
	if err != nil {
		logger.Println("Error creating custompages table:", err.Error())
	}
	_, err = db.Exec(`
		create table if not exists cachedpages (
			pageid integer primary key,
			document text not null
		) strict;
		`)
	if err != nil {
		logger.Println("Error creating cachedpages table:", err.Error())
	}

	return &DbHandler{db, logger}, nil
}

type NotFoundErr struct {
	criterion string
}

func (this NotFoundErr) Error() string {
	return fmt.Sprintf("[quDnD] Row not found with criterion '%s'", this.criterion)
}

type DbHandler struct {
	db     *sql.DB
	logger *log.Logger
}

func (this *DbHandler) SetCustomPage(page statblock.PageInfo) error {
	bytes, err := json.Marshal(page)
	if err != nil {
		this.logger.Printf("Error parsing page to db: %d\n", page.PageId)
		return err
	}
	_, err = this.db.Exec(`
		insert into custompages
			(pageid, document)
			values (?, ?)
			on conflict (pageid) do
				update set
				pageid = ?,
				document = ?;`,
		page.PageId,
		string(bytes),
		page.PageId,
		string(bytes),
	)
	if err != nil {
		this.logger.Println("Error updating custom page:", err.Error())
		return err
	}

	return nil
}

func (this *DbHandler) GetCustomPage(pageid int) (statblock.PageInfo, error) {
	row := this.db.QueryRow("select * from custompages where pageid = ?", pageid)
	err := row.Err()
	if err == sql.ErrNoRows {
		this.logger.Println("Custom page requested but not found:", pageid)
		return statblock.PageInfo{}, NotFoundErr{strconv.Itoa(pageid)}
	} else if err != nil {
		this.logger.Println("Error getting custom page:", err.Error())
		return statblock.PageInfo{}, err
	}

	var document string
	row.Scan(&pageid, &document)
	var page statblock.PageInfo
	err = json.Unmarshal([]byte(document), &page)
	if err != nil {
		this.logger.Printf("Invalid page in db: %d\n", pageid)
		return statblock.PageInfo{}, err
	}

	return page, nil
}

func (this *DbHandler) GetCachedPage(pageid int) (statblock.PageInfo, error) {
	row := this.db.QueryRow("select * from cachedpages where pageid = ?", pageid)
	err := row.Err()
	if err != nil {
		if err == sql.ErrNoRows {
			this.logger.Println("Cached page requested but not found:", pageid)
			return statblock.PageInfo{}, NotFoundErr{strconv.Itoa(pageid)}
		} else {
			this.logger.Println("Error getting cached page:", err.Error())
			return statblock.PageInfo{}, err
		}
	}

	var document string
	row.Scan(&pageid, &document)
	var page statblock.PageInfo
	err = json.Unmarshal([]byte(document), &page)
	if err != nil {
		this.logger.Printf("Invalid page in db: %d\n", pageid)
		return statblock.PageInfo{}, err
	}

	return page, nil
}

func (this *DbHandler) SetCachedPage(page statblock.PageInfo) error {
	bytes, err := json.Marshal(page)
	if err != nil {
		this.logger.Printf("Error parsing page to db: %d\n", page.PageId)
		return err
	}
	_, err = this.db.Exec(`
		insert into cachedpages
			(pageid, document)
			values (?, ?)
			on conflict (pageid) do
				update set
				pageid = ?,
				document = ?;`,
		page.PageId,
		string(bytes),
		page.PageId,
		string(bytes),
	)
	if err != nil {
		this.logger.Println("Error updating cached page:", err.Error())
		return err
	}

	return nil
}
