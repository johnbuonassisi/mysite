````# Handling not-found errors in Go interfaces

In Go it is common to use interfaces to abstract implementation details. When
different implementations handle errors in different ways it can be challenging to 
handle errors in a consistent way.

I often come across these scenarios when dealing with data stores. Not-found errors
are a common occurrence and calling code often wants to know about them. In this article
we'll explore a pattern for handling not-found errors in Go interfaces.

## The problem with custom errors

The simplest and most obvious approach to communicate a not-found error is with a
common sentinel error. 

```go
var NotFoundErr = errors.New("not found")
```

A data store interface can return this error when data is attempting to be
retrieved that does not exist. 

```go
type UserStore interface {
	Get(id string) (User, error)
}

type PostgresUserStore struct {
	// ...
}

func (s *PostgresUserStore) Get(id string) (User, error) {
	// ...
	// returns NotFoundErr when user not found
}
```

Callers can then check for this error type.

```go
user, err := store.Get(id)
if err != nil {
	return err
}
if errors.Is(err, ErrNotFound) {
	// handle the not-found error
}
```

The problem with this approach is that all implementation needs to agree
to return the same error. This could be challenging to enforce and could
lead to inconsistencies between implementations. If implementations are in
different packages they will need to import the error from some common place.

## The IsNotFound Method

An alternate approach is to add a `IsNotFound` method to the interface, preventing 
all implementation to return the same error.

```go
type UserStore interface {
	Get(id string) (User, error)
	IsNotFound(err error) bool
}
```

Implementations of the interface can return their own error type and check it
within `IsNotFound`.

```go
type PostgresUserStore struct {
	pool *pgxpool.Pool
}
func (s *PostgresUserStore) Get(id string) (User, error) {
	var user User
	err := s.db.QueryRow(context.Background(), "SELECT id, name, age FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		// err == pgx.ErrNoRows when user is not found
		return User{}, err
	}
	return user, err	
}

func (s *PostgresUserStore) IsNotFound(err error) bool {
	return errors.Is(err, pgx.ErrNoRows)
}
```

Callers can simply use the `IsNotFound` method to check for not-found errors.

```go
user, err := store.Get(id)
if err != nil {
	return err
}
if store.IsNotFound(err) {
	// handle the not-found error
}
```

Adding the `IsNotFound` method has solved the problem of relying on a common error
type while also making it very simple for callers to check for not-found errors.

## When to use this pattern?

Here are some guidelines on when to use this pattern:
1. You have multiple implementations of an interface, and those implementations
return different errors types to communicate not-found to a user. 
2. You want to avoid multiple packages importing a common error.
3. You are willing to accept the overhead to implementing the `IsNotFound` method
in each implementation.

## Wrapping-up
Overall, adding a `IsNotFound` error to your interface can be an ergonomic and flexible
way to allow checks for not-found errors. Like any pattern, don't use it blindly, ensure
it brings value to your codebase. Happy coding ;)
````
