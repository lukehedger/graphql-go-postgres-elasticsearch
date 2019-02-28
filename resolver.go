package api

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

func (r *Resolver) Hello(args struct{ ID string }) (string, error) {
	var name string
	rows, err := r.DB.Query("SELECT name FROM people WHERE id = $1 LIMIT $2", args.ID, 1)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&name)
		if err != nil {
			return "", err
		}
	}
	err = rows.Err()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Hello, %v", name), nil
}

func (r *Resolver) Search(args struct{ Query string }) (*PersonResolver, error) {
	res, err := r.ES.Search(
		r.ES.Search.WithContext(context.Background()),
		r.ES.Search.WithIndex("people"),
		r.ES.Search.WithBody(strings.NewReader(
			`{"query": { "match": { "name": "`+args.Query+`" } }}`,
		)),
		r.ES.Search.WithPretty(),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var body map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		return nil, err
	}
	hit := body["hits"].(map[string]interface{})["hits"].([]interface{})[0]
	source := hit.(map[string]interface{})["_source"].(map[string]interface{})
	person := Person{
		ID:   source["id"].(string),
		Name: source["name"].(string),
	}

	return &PersonResolver{person}, nil
}

func (r *PersonResolver) ID() string {
	return r.Person.ID
}

func (r *PersonResolver) Name() string {
	return r.Person.Name
}
