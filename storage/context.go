package storage

import "context"

type downloadFilenameKey struct{}

func NewDownloadFilenameContext(ctx context.Context, filename string) context.Context {
	return context.WithValue(ctx, downloadFilenameKey{}, filename)
}

func FromDownloadFilenameContext(ctx context.Context) (filename string, ok bool) {
	filename, ok = ctx.Value(downloadFilenameKey{}).(string)
	return
}

type downloadBeforeKey struct{}

type DownloadBeforeFunc func(info DownloadFileInfoer)

func NewDownloadBeforeContext(ctx context.Context, f DownloadBeforeFunc) context.Context {
	return context.WithValue(ctx, downloadBeforeKey{}, f)
}

func FromDownloadBeforeContext(ctx context.Context) (f DownloadBeforeFunc, ok bool) {
	f, ok = ctx.Value(downloadBeforeKey{}).(DownloadBeforeFunc)
	return
}

type renameKey struct{}

type RenameFunc func(filename string) string

func NewRenameContext(ctx context.Context, f RenameFunc) context.Context {
	return context.WithValue(ctx, renameKey{}, f)
}

func FromRenameContext(ctx context.Context) (f RenameFunc, ok bool) {
	f, ok = ctx.Value(renameKey{}).(RenameFunc)
	return
}
