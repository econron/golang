package main

import (
	"context"
	"log/slog"
	"os"
	"sync"
	"io"
)

type LogStruct struct {
	Elm1 string
	Elm2 uint64
	Elm3 map[string][]byte
}

func main() {
	// 2024/09/17 20:21:34 INFO Hello World user=hoge test=""
	slog.Info("Hello World", "user", os.Getenv("USER"), "test", os.Getenv("TEST"))

	// we can get slog top level logger explicitly
	logger := slog.Default()
	logger.Info("top level logger")

	// change handler used by logger
	// time=2024-09-17T20:25:31.775+09:00 level=INFO msg="hello world." user=hoge
	logger2 := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger2.Info("hello world.", "user", os.Getenv("USER"))

	// use json handler
	// time=2024-09-17T20:26:43.600+09:00 level=INFO msg="hello world." user=hoge
	// {"time":"2024-09-17T20:26:43.600172+09:00","level":"INFO","msg":"hello world.","user":"hoge"}
	jsonLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	jsonLogger.Info("hello world.", "user", os.Getenv("USER"))

	// use attrs and LogAttrs
	// 2024/09/17 20:29:31 INFO hello, world user=hoge
	slog.LogAttrs(context.Background(), slog.LevelInfo, "hello, world", slog.String("user", os.Getenv("USER")))

	// any argsを投げ込んでいいので、複数のKey Valueを投げ込む
	// {"time":"2024-09-17T20:33:08.212216+09:00","level":"INFO","msg":"tekitou","maps":{"test1":"test1","test2":"test2","test3":"test3"}}
	m := make(map[string]string)
	m["test1"] = "test1"
	m["test2"] = "test2"
	m["test3"] = "test3"
	mulLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	mulLogger.Info("tekitou", "maps", m)

	// 構造体を入れてみる
	// {"time":"2024-09-17T20:38:49.770478+09:00","level":"INFO","msg":"tekitou","struct":{"Elm1":"tekitou","Elm2":12,"Elm3":{"bytekey":"YWJj"}}}
	m2 := make(map[string][]byte)
	m2["bytekey"] = []byte("abc")
	s := LogStruct{
		Elm1: "tekitou",
		Elm2: uint64(12),
		Elm3: m2,
	}
	mulLogger2 := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	mulLogger2.Info("tekitou", "struct", s)

	// log grouping
	gLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	// {"time":"2024-09-17T20:49:29.05917+09:00","level":"INFO","msg":"[1]hello world","G":{"testing":20}}
	gLogger.WithGroup("G").With("testing", os.Getegid()).Info("[1]hello world")
	// {"time":"2024-09-17T20:49:29.059173+09:00","level":"INFO","msg":"[2]hello world","G2":{"testing2":501}}
	gLogger.WithGroup("G2").With("testing2", os.Geteuid()).Info("[2]hello world")
	// {"time":"2024-09-17T20:49:29.059175+09:00","level":"INFO","msg":"[3]hello world","testing3":"okuyamaaron"}
	gLogger.Info("[3]hello world", "testing3", os.Getenv("USER"))

	/*
logger.Info("hello", "key", 23)
	
time: 2023-05-15T16:29:00
level: INFO
message: "hello"
key: 23
---
	*/
}

type IndentHandler struct {
	opts Options
	mu *sync.Mutex
	out io.Writer
}

type Options struct {
	Level slog.Leveler
}

func New(out io.Writer, opts *Options) *IndentHandler {
	h := &IndentHandler{out: out, mu: &sync.Mutex{}}
	if opts != nil {
		h.opts = *opts
	}
	if h.opts.Level == nil {
		h.opts.Level = slog.LevelInfo
	}
	return h
}

// EnableはLoggerの出力メソッド実行前に実行される
func(h *IndentHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.opts.Level.Level()
}

func (h *IndentHandler) Handle(ctx context.Context, r slog.Record) error {
	buf := make([]byte, 0, 1024)
	if !r.Time.IsZero() {
		buf = h.appendAttr(buf, slog.Time(slog.TimeKey, r.Time), 0)
	}
}