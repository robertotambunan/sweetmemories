# Sweet Memories

### How to install
go get -v github.com/robertotambunan/sweetmemories

### Usage
Sweet memories is data cache. If you want to save data for a spesific time without any caching tools, you can simply use sweet memories.

### Example Case
If you want to not get data from database everytime, you can simply use sweet memories. Use one sweet memories Object to save any kind of data! One object, one responsibility.


### Example

```
var (
	rememberData sweetmemories.Memory
)


	data := []Data{}

	if rememberdata.IsForget() { // if cache is nil
		rememberData.Remember(getData(), 600*time.Second) //Fill your cache memory
	}

	w, err := rememberData.GetMemory() //Get your memory
	if err != nil && len(w) == 0 {
		log.Println("Data Cache is Failing", err)
		data = getData() // If fail, we can't do anything except get data from your source
	} else {
		e := json.Unmarshal(w, &data)
		if e != nil {
			log.Println(e)
			data = getData() // If fail, we can't do anything except get data from your source
		}
	}
```

```

func getData() []Data {
	result := []Data{}
	results, err := db.Query("SELECT product_id, product_name, product_cat FROM product")
	if err != nil {
		return result
	}

	for results.Next() {
		var id int
		var cat int
		var name string

		err = results.Scan(&id, &name, &cat)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, Data{
			ProductID:   id,
			ProductCat:  cat,
			ProductName: name,
		})
	}
	return result
}

type Data struct {
	ProductID   int    `json:"product_id"`
	ProductCat  int    `json:"product_cat"`
	ProductName string `json:"product_name"`
}


```