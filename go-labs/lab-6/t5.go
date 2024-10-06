package main

import (
	"fmt"
	"sync"
)

type Request struct {
	op1       float64
	op2       float64
	operation string
}

type Response struct {
	req    Request
	result float64
	status string
	err    string
}

func calculator(requests chan Request, res chan Response, wg *sync.WaitGroup) {
	defer wg.Done()
	for req := range requests {
		switch req.operation {
		case "+":
			res <- Response{req: req, result: req.op1 + req.op2, status: "Correct", err: ""}
		case "-":
			res <- Response{req: req, result: req.op1 - req.op2, status: "Correct", err: ""}
		case "/":
			res <- div(req.op1, req.op2, req)
		case "*":
			res <- Response{req: req, result: req.op1 * req.op2, status: "Correct", err: ""}
		default:
			res <- Response{req: req, result: 0, status: "Incorrect", err: "invalid operation"}
		}
	}
}

func div(a float64, b float64, req Request) Response {
	if b == 0 {
		return Response{req: req, result: 0, status: "Incorrect", err: "invalid operation: division by zero"}
	} else {
		return Response{req: req, result: a / b, status: "Correct", err: ""}
	}
}

func main() {
	var wg sync.WaitGroup
	reqCh := make(chan Request)
	respCh := make(chan Response)

	operations := []Request{
		{5, 3, "+"},
		{10, 2, "-"},
		{4, 7, "*"},
		{4, 0, "/"},
		{8, 2, "/"},
	}

	wg.Add(1)
	go calculator(reqCh, respCh, &wg)

	go func() {
		for _, op := range operations {
			reqCh <- op
		}
		close(reqCh)
	}()

	go func() {
		for resp := range respCh {
			if resp.status == "Correct" {
				fmt.Printf("Operation %.2f %s %.2f = %.2f\n", resp.req.op1, resp.req.operation, resp.req.op2, resp.result)
			} else {
				fmt.Println("Error: ", resp.err)
			}
		}
	}()

	wg.Wait()
	close(respCh)
}
