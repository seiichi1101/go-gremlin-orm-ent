// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/gremlin"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Cars holds the value of the cars edge.
	Cars []*Car `json:"cars,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CarsOrErr returns the Cars value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CarsOrErr() ([]*Car, error) {
	if e.loadedTypes[0] {
		return e.Cars, nil
	}
	return nil, &NotLoadedError{edge: "cars"}
}

// FromResponse scans the gremlin response data into User.
func (u *User) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scanu struct {
		ID   string `json:"id,omitempty"`
		Age  int    `json:"age,omitempty"`
		Name string `json:"name,omitempty"`
	}
	if err := vmap.Decode(&scanu); err != nil {
		return err
	}
	u.ID = scanu.ID
	u.Age = scanu.Age
	u.Name = scanu.Name
	return nil
}

// QueryCars queries the "cars" edge of the User entity.
func (u *User) QueryCars() *CarQuery {
	return (&UserClient{config: u.config}).QueryCars(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", age=")
	builder.WriteString(fmt.Sprintf("%v", u.Age))
	builder.WriteString(", name=")
	builder.WriteString(u.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

// FromResponse scans the gremlin response data into Users.
func (u *Users) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scanu []struct {
		ID   string `json:"id,omitempty"`
		Age  int    `json:"age,omitempty"`
		Name string `json:"name,omitempty"`
	}
	if err := vmap.Decode(&scanu); err != nil {
		return err
	}
	for _, v := range scanu {
		*u = append(*u, &User{
			ID:   v.ID,
			Age:  v.Age,
			Name: v.Name,
		})
	}
	return nil
}

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
