// Copyright 2013 The Go Circuit Project
// Use of this source code is governed by the license for
// The Go Circuit Project, found in the LICENSE file.
//
// Authors:
//   2014 Petar Maymounkov <p@gocircuit.org>

package docker

import (
	"io"
	
	xio "github.com/gocircuit/circuit/kit/x/io"
	"github.com/gocircuit/circuit/use/circuit"
	"github.com/gocircuit/circuit/use/errors"
)

func init() {
	circuit.RegisterValue(XContainer{})
}

type XContainer struct {
	Container
}

func (x XContainer) Wait() (*Stat, error) {
	stat, err := x.Container.Wait()
	return stat, errors.Pack(err)
}

func (x XContainer) Signal(sig string) error {
	return errors.Pack(x.Container.Signal(sig))
}

func (x XContainer) Stdin() circuit.X {
	return xio.NewXWriteCloser(x.Container.Stdin())
}

func (x XContainer) Stdout() circuit.X {
	return xio.NewXReadCloser(x.Container.Stdout())
}

func (x XContainer) Stderr() circuit.X {
	return xio.NewXReadCloser(x.Container.Stderr())
}

func (x XContainer) Peek() (*Stat, error) {
	stat, err := x.Container.Peek()
	return stat, errors.Pack(err)
}

type YContainer struct {
	X circuit.X
}

func (y YContainer) Wait() (stat *Stat, err error) {
	r := y.X.Call("Wait")
	stat, _ = r[0].(*Stat)
	return stat, errors.Unpack(r[1])
}

func (y YContainer) Signal(sig string) error {
	r := y.X.Call("Signal", sig)
	return errors.Unpack(r[0])
}

func (y YContainer) Scrub() {
	y.X.Call("Scrub")
}

func (y YContainer) IsDone() bool {
	return y.X.Call("IsDone")[0].(bool)
}

func (y YContainer) Peek() (stat *Stat, err error) {
	r := y.X.Call("Peek")
	stat, _ = r[0].(*Stat)
	return stat, errors.Unpack(r[1])
}

func (y YContainer) Stdin() io.WriteCloser {
	return xio.NewYWriteCloser(y.X.Call("Stdin")[0])
}

func (y YContainer) Stdout() io.ReadCloser {
	return xio.NewYReadCloser(y.X.Call("Stdout")[0])
}

func (y YContainer) Stderr() io.ReadCloser {
	return xio.NewYReadCloser(y.X.Call("Stderr")[0])
}
