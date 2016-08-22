package main

import "fmt"

type Process interface {
	Start()
	Kill()
	Signal()
	Pid() int
}

type Proc struct {
	foo string
	pid int
}

func (p *Proc) Start()  {}
func (p *Proc) Kill()   {}
func (p *Proc) Signal() {}
func (p *Proc) Pid() int {
	return p.pid
}

func NewProcess() *Proc {
	return &Proc{"foo", 3}
}

type Daemon struct {
	process *Proc
}

func (d *Daemon) Run(p Process) {
	fmt.Println(p.Pid())
}

func main() {

	d := Daemon{
		process: NewProcess(),
	}

	fmt.Printf("d.process.foo = %+v\n", d.process.foo)
	fmt.Printf("d.process.Pid() = %+v\n", d.process.Pid())

	d.Run(d.process)
}
