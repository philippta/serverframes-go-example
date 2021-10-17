package main

import (
	"embed"
	"html/template"
	"net/http"
	"strconv"
)

//go:embed *.html
var FS embed.FS

type Product struct {
	ID    int
	Name  string
	Price int
}

type Cart struct {
	Products []Product
}

type PageData struct {
	Products []Product
	Cart     Cart
}

func main() {
	products := []Product{
		{1, "Laptop", 1500},
		{2, "Mouse", 50},
		{3, "Keyboard", 100},
		{4, "Monitor", 500},
		{5, "CPU", 1000},
		{6, "RAM", 250},
		{7, "HDD", 500},
		{8, "GPU", 250},
		{9, "Speaker", 100},
		{10, "Headphones", 50},
		{11, "Camera", 250},
	}
	cart := Cart{
		Products: []Product{
			{1, "Laptop", 1500},
		},
	}

	tmpl := shoppingTemplate()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, PageData{
			Products: products,
			Cart:     cart,
		})
	})

	http.HandleFunc("/add/", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(r.URL.Path[len("/add/"):])
		for _, p := range products {
			if p.ID == id {
				cart.Products = append(cart.Products, p)
				break
			}
		}

		tmpl.ExecuteTemplate(w, "navbarcart", cart.Products)
		tmpl.ExecuteTemplate(w, "productlist", PageData{products, cart})
		tmpl.ExecuteTemplate(w, "cart", cart)
	})

	http.HandleFunc("/remove/", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(r.URL.Path[len("/remove/"):])
		for i, p := range cart.Products {
			if p.ID == id {
				cart.Products = append(cart.Products[:i], cart.Products[i+1:]...)
				break
			}
		}

		tmpl.ExecuteTemplate(w, "navbarcart", cart.Products)
		tmpl.ExecuteTemplate(w, "productlist", PageData{products, cart})
		tmpl.ExecuteTemplate(w, "cart", cart)
	})

	http.ListenAndServe(":8080", nil)
}

func shoppingTemplate() *template.Template {
	funcs := template.FuncMap{
		"incart": func(p Product, c Cart) bool {
			for _, v := range c.Products {
				if v.ID == p.ID {
					return true
				}
			}
			return false
		},
		"total": func(c Cart) int {
			total := 0
			for _, v := range c.Products {
				total += v.Price
			}
			return total
		},
	}
	return template.Must(template.New("shoppingcart.html").Funcs(funcs).ParseFS(FS,
		"shoppingcart.html",
		"productlist.html",
		"cart.html",
	))
}
