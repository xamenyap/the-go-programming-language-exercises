package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
	mux.HandleFunc("/create", db.create)
	mux.HandleFunc("/update", db.update)
	mux.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

var mut sync.RWMutex

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

const listTpl = `
<html>
<h1>Items</h1>
<table>
	<tr>
		<th style="border:1px solid #dddddd; text-align: left; padding: 8px;">Item</th>
		<th style="border:1px solid #dddddd; text-align: left; padding: 8px;">Price</th>
	</tr>
	{{range $item, $price := .}} 
		<tr>
			<td style="border:1px solid #dddddd; text-align: left; padding: 8px;">{{$item}}</td>
			<td style="border:1px solid #dddddd; text-align: left; padding: 8px;">{{$price}}</td>
		</tr>
	{{else}}
		<tr>
			<td style="border:1px solid #dddddd; text-align: left; padding: 8px;" colspan="2">No items</td>
		</tr>
	{{end}}
</table>
</html>
`

func (db database) list(w http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.New("list").Parse(listTpl))

	if err := tpl.Execute(w, db); err != nil {
		w.WriteHeader(http.StatusNotFound) // 404
		_, _ = fmt.Fprintf(w, "error: %v\n", err)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	mut.RLock()
	defer mut.RUnlock()

	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		_, _ = fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	_, _ = fmt.Fprintf(w, "%s\n", price)
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	mut.Lock()
	defer mut.Unlock()

	if exist, ok := db[item]; ok {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "already exists: %q\n", exist)
		return
	}

	parsedPrice, _ := strconv.ParseFloat(price, 32)
	db[item] = dollars(parsedPrice)

	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "item created: %q\n", item)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	mut.Lock()
	defer mut.Unlock()

	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	parsedPrice, _ := strconv.ParseFloat(price, 32)
	db[item] = dollars(parsedPrice)

	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "item updated: %q\n", item)
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	mut.Lock()
	defer mut.Unlock()

	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	delete(db, item)
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "item deleted: %q\n", item)
}
