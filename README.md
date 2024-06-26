# Go Persistence Extension V0 (Work In Progress)

A Go package that provides a plugin-like features to reduce the boilerplate code 
when implementing the repository layers. Now it'll support for the Gorm ORM first.

Example: 
```go
type UserRepository struct {
	DB *gorm.DB
	contract.Basic[*entity.User, uint] // composite with basic operation interface
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
		Basic: gorme.NewBasicRepository[*entity.User, uint](db), // use constructor
	}
}
```
You'll get basic data operation methods immediately! Save lots of time to write boilerplate code of repository layer!
```go
UserRepository.GetBy(ctx, &entity.User{Username: "johndoe"})
UserRepository.GetById(ctx, userId)
UserRepository.DeleteById(ctx, userId)
UserRepository.FindAll(ctx, limit)
// ...
// .....
```
