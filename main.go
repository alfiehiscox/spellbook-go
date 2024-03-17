package main

import (
	"context"
	"database/sql"
	_ "embed"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/alfiehiscox/spellbook-go/db"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var ddl string
var queries *db.Queries

func main() {

	ctx := context.Background(); 

	// Should probably use a pool
	database, err := sql.Open("sqlite3", ":memory:")
	if err != nil { 
		panic(err) 
	}

	if _, err := database.ExecContext(ctx, ddl); err != nil {
		panic(err)
	}

	queries = db.New(database)

	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/spells", handleGetAllSpells)
	e.POST("/spells", handleCreateSpell)

	e.GET("/spells/:id", handleGetSpell)
	e.PUT("/spells/:id", handleUpdateSpell)
	e.DELETE("/spells/:id", handleDeleteSpell)

	if err := e.Start(":8080"); err != nil {
		fmt.Println(err.Error())	
	}
}

func handleGetAllSpells(c echo.Context) error {
	spells, err := queries.GetSpells(context.Background())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, spells)
}

func handleGetSpell(c echo.Context) error {
	idStr := c.Param("id")

	if idStr == "" {
		return errors.New("No id provided")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return errors.New("Not a correct id")
	}
	
	spell, err := queries.GetSpell(context.Background(), int64(id))
	if err != nil {
		return err
	}
	
	return c.JSON(http.StatusOK, spell)
}

func handleCreateSpell(c echo.Context) error {
	s := new(db.CreateSpellParams)
	if err := c.Bind(s); err != nil {
		return err
	}

	spell, err := queries.CreateSpell(context.Background(), *s)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, spell)
}

func handleUpdateSpell(c echo.Context) error {
	idStr := c.Param("id")

	if idStr == "" {
		return errors.New("No id provided")
	}

	id, err := strconv.Atoi(idStr)

	if err != nil {
		return errors.New("Not a correct id")
	}

	s := new(db.UpdateSpellParams)
	s.ID = int64(id)

	if err := c.Bind(s); err != nil {
		return err
	}

	spell, err := queries.UpdateSpell(context.Background(), *s)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, spell)
}

func handleDeleteSpell(c echo.Context) error {
	idStr := c.Param("id")

	if idStr == "" {
		return errors.New("No id provided")
	}

	id, err := strconv.Atoi(idStr)

	if err != nil {
		return errors.New("Not a correct id")
	}

	if err := queries.DeleteSpell(context.Background(), int64(id)); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
