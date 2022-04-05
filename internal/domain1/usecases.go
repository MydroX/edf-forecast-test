package domain1

type UseCase interface {
	ThisIsAUsecase(ctx *context.Context, req Request) (data Data, error)
}