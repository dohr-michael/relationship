package neo

import (
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"github.com/dohr-michael/relationship/apis/cfg"
	"io"
)

var driver bolt.Driver

type Conn bolt.Conn

type P map[string]interface{}
type Row []interface{}
type Rows [][]interface{}
type Meta map[string]interface{}

func C(fn func(connection Conn)) {
	conn, err := driver.OpenNeo("bolt://" + cfg.GetNeoUser() + ":" + cfg.GetNeoPassword() + "@" + cfg.GetNeoUrl())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fn(conn.(Conn))
}

func ExecAll(query string, params P, fn func(data Rows, meta Meta)) {
	C(func(c Conn) {
		q, err := c.PrepareNeo(query)
		if err != nil {
			panic(err)
		}
		r, err := q.QueryNeo(params)
		if err != nil {
			panic(err)
		}
		data, meta, err := r.All()
		if err != nil {
			panic(err)
		}
		fn(data, meta)
		q.Close()
	})
}

func ExecForeach(query string, params P, mapper func(row Row, meta Meta)) {
	C(func(c Conn) {
		q, err := c.PrepareNeo(query)
		if err != nil {
			panic(err)
		}
		r, err := q.QueryNeo(params)
		if err != nil {
			panic(err)
		}
		items, meta, err := r.NextNeo()
		for err == nil {
			mapper(items, meta)
			items, meta, err = r.NextNeo()
		}
		if err != io.EOF {
			panic(err)
		}
		q.Close()
	})
}

func init() {
	driver = bolt.NewDriver()
}
