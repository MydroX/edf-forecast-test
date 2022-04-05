package domain1

type Repository interface {
	ThisIsARepository(ctx *context.Context, req Request) (data Data, error)
}