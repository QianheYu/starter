package main

import (
    "bufio"
    "context"
    "errors"
    "flag"
    "fmt"
    "golang.org/x/sync/errgroup"
    "io"
    "os"
    "os/exec"
    "os/signal"
    "strings"
	"log"
)


var file = flag.String("file", "commands.txt", "Designation of programme documents")
type Process struct {
    c *exec.Cmd
	Stdout io.ReadCloser
	Stderr io.ReadCloser
}

func main() {
	flag.Parse()
	
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()
	eg, ctx := errgroup.WithContext(ctx)
	
	eg.Go(func() error {
        <-ctx.Done()
		log.Println("[main] canceled")
		return nil
    })
	
	cstr := make(chan string)
	defer close(cstr)
	
	eg.Go(func() error {
        for {
            select {
            case buf := <-cstr:
                log.Print(buf)
            case <-ctx.Done():
                return nil
            }
		}
    })
	
	fd, err := os.Open(*file)
	if err != nil {
		log.Fatalf("[main] stderr: %v", err)
	}
	defer fd.Close()
	scanner := bufio.NewScanner(fd)
	
	p := make(map[string]*Process)
	c :=make(chan struct{})
	for scanner.Scan() {
		command := scanner.Text()
		if len(command) <= 0 {
			continue
		}
		var ps *Process
		// 启动进程
		
		eg.Go(func() error {
			log.Printf("[main] start %s\n", command)
			commandParam :=strings.SplitN(command, ",", 0)
			if len(commandParam) > 0 {
				ps = &Process{
					c: exec.CommandContext(ctx, commandParam[0], commandParam[1:]...),
				}
				if ps.Stderr, err = ps.c.StderrPipe(); err != nil {
					return err
				}
				if ps.Stdout, err = ps.c.StdoutPipe(); err != nil {
					return err
				}
				p[commandParam[0]] = ps 
			} else {
				ps = &Process{
					c: exec.CommandContext(ctx, command),
				}
				if ps.Stderr, err = ps.c.StderrPipe(); err != nil {
					return err
				}
				if ps.Stdout, err = ps.c.StdoutPipe(); err != nil {
					return err
				}
				p[command] = ps
			}
			
			c <- struct{}{}
			return ps.c.Run()
        })
		<-c
		eg.Go(func() error {
//			if ps.Stdout, err = ps.c.StdoutPipe(); err != nil {
//				return err
//			}
            stdout := bufio.NewReader(ps.Stdout)
			for {
                select {
                case <-ctx.Done():
					return nil
                default:
					str, err := stdout.ReadString('\n')
					if len(str) <= 0 {
						continue
					}
					if err != nil && !errors.Is(err, io.EOF) {
						cstr <- fmt.Sprintf("[%s] stdout: %v", command, err)
					} else {
						cstr <- fmt.Sprintf("[%s] strout: %s", command, str)
					}
                }
			}
        })
		
		eg.Go(func() error {
//			if ps.Stderr, err = ps.c.StderrPipe(); err != nil {
//				return err
//			}
            stderr := bufio.NewReader(ps.Stderr)
			for {
                select {
                case <-ctx.Done():
					return nil
                default:
					str, err :=stderr.ReadString('\n')
					if len(str) <=0 {
						continue
					}
					if err != nil && !errors.Is(err, io.EOF) {
						cstr <- fmt.Sprintf("[%s] stderr: %v", command, err)
					} else {
						cstr <- fmt.Sprintf("[%s] stderr: %s", command, str)
					}
                }
			}
        })
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("[main] stderr: %v", fmt.Errorf("error scanning file %s: %w", *file, err))
	}
	
	if err := eg.Wait(); err != nil {
		log.Fatalf("[main] stderr: %v", fmt.Errorf("wait goroutine exit err: %w", err))
	}
	log.Println("[main] exit")
}