// This file was generated by counterfeiter
package dbfakes

import (
	"database/sql"
	"database/sql/driver"
	"sync"

	"github.com/concourse/atc/db"
)

type FakeConn struct {
	BeginStub        func() (db.Tx, error)
	beginMutex       sync.RWMutex
	beginArgsForCall []struct{}
	beginReturns     struct {
		result1 db.Tx
		result2 error
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	closeReturns     struct {
		result1 error
	}
	DriverStub        func() driver.Driver
	driverMutex       sync.RWMutex
	driverArgsForCall []struct{}
	driverReturns     struct {
		result1 driver.Driver
	}
	ExecStub        func(query string, args ...interface{}) (sql.Result, error)
	execMutex       sync.RWMutex
	execArgsForCall []struct {
		query string
		args  []interface{}
	}
	execReturns struct {
		result1 sql.Result
		result2 error
	}
	PingStub        func() error
	pingMutex       sync.RWMutex
	pingArgsForCall []struct{}
	pingReturns     struct {
		result1 error
	}
	PrepareStub        func(query string) (*sql.Stmt, error)
	prepareMutex       sync.RWMutex
	prepareArgsForCall []struct {
		query string
	}
	prepareReturns struct {
		result1 *sql.Stmt
		result2 error
	}
	QueryStub        func(query string, args ...interface{}) (*sql.Rows, error)
	queryMutex       sync.RWMutex
	queryArgsForCall []struct {
		query string
		args  []interface{}
	}
	queryReturns struct {
		result1 *sql.Rows
		result2 error
	}
	QueryRowStub        func(query string, args ...interface{}) *sql.Row
	queryRowMutex       sync.RWMutex
	queryRowArgsForCall []struct {
		query string
		args  []interface{}
	}
	queryRowReturns struct {
		result1 *sql.Row
	}
	SetMaxIdleConnsStub        func(n int)
	setMaxIdleConnsMutex       sync.RWMutex
	setMaxIdleConnsArgsForCall []struct {
		n int
	}
	SetMaxOpenConnsStub        func(n int)
	setMaxOpenConnsMutex       sync.RWMutex
	setMaxOpenConnsArgsForCall []struct {
		n int
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConn) Begin() (db.Tx, error) {
	fake.beginMutex.Lock()
	fake.beginArgsForCall = append(fake.beginArgsForCall, struct{}{})
	fake.recordInvocation("Begin", []interface{}{})
	fake.beginMutex.Unlock()
	if fake.BeginStub != nil {
		return fake.BeginStub()
	} else {
		return fake.beginReturns.result1, fake.beginReturns.result2
	}
}

func (fake *FakeConn) BeginCallCount() int {
	fake.beginMutex.RLock()
	defer fake.beginMutex.RUnlock()
	return len(fake.beginArgsForCall)
}

func (fake *FakeConn) BeginReturns(result1 db.Tx, result2 error) {
	fake.BeginStub = nil
	fake.beginReturns = struct {
		result1 db.Tx
		result2 error
	}{result1, result2}
}

func (fake *FakeConn) Close() error {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	} else {
		return fake.closeReturns.result1
	}
}

func (fake *FakeConn) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeConn) CloseReturns(result1 error) {
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConn) Driver() driver.Driver {
	fake.driverMutex.Lock()
	fake.driverArgsForCall = append(fake.driverArgsForCall, struct{}{})
	fake.recordInvocation("Driver", []interface{}{})
	fake.driverMutex.Unlock()
	if fake.DriverStub != nil {
		return fake.DriverStub()
	} else {
		return fake.driverReturns.result1
	}
}

func (fake *FakeConn) DriverCallCount() int {
	fake.driverMutex.RLock()
	defer fake.driverMutex.RUnlock()
	return len(fake.driverArgsForCall)
}

func (fake *FakeConn) DriverReturns(result1 driver.Driver) {
	fake.DriverStub = nil
	fake.driverReturns = struct {
		result1 driver.Driver
	}{result1}
}

func (fake *FakeConn) Exec(query string, args ...interface{}) (sql.Result, error) {
	fake.execMutex.Lock()
	fake.execArgsForCall = append(fake.execArgsForCall, struct {
		query string
		args  []interface{}
	}{query, args})
	fake.recordInvocation("Exec", []interface{}{query, args})
	fake.execMutex.Unlock()
	if fake.ExecStub != nil {
		return fake.ExecStub(query, args...)
	} else {
		return fake.execReturns.result1, fake.execReturns.result2
	}
}

func (fake *FakeConn) ExecCallCount() int {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	return len(fake.execArgsForCall)
}

func (fake *FakeConn) ExecArgsForCall(i int) (string, []interface{}) {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	return fake.execArgsForCall[i].query, fake.execArgsForCall[i].args
}

func (fake *FakeConn) ExecReturns(result1 sql.Result, result2 error) {
	fake.ExecStub = nil
	fake.execReturns = struct {
		result1 sql.Result
		result2 error
	}{result1, result2}
}

func (fake *FakeConn) Ping() error {
	fake.pingMutex.Lock()
	fake.pingArgsForCall = append(fake.pingArgsForCall, struct{}{})
	fake.recordInvocation("Ping", []interface{}{})
	fake.pingMutex.Unlock()
	if fake.PingStub != nil {
		return fake.PingStub()
	} else {
		return fake.pingReturns.result1
	}
}

func (fake *FakeConn) PingCallCount() int {
	fake.pingMutex.RLock()
	defer fake.pingMutex.RUnlock()
	return len(fake.pingArgsForCall)
}

func (fake *FakeConn) PingReturns(result1 error) {
	fake.PingStub = nil
	fake.pingReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConn) Prepare(query string) (*sql.Stmt, error) {
	fake.prepareMutex.Lock()
	fake.prepareArgsForCall = append(fake.prepareArgsForCall, struct {
		query string
	}{query})
	fake.recordInvocation("Prepare", []interface{}{query})
	fake.prepareMutex.Unlock()
	if fake.PrepareStub != nil {
		return fake.PrepareStub(query)
	} else {
		return fake.prepareReturns.result1, fake.prepareReturns.result2
	}
}

func (fake *FakeConn) PrepareCallCount() int {
	fake.prepareMutex.RLock()
	defer fake.prepareMutex.RUnlock()
	return len(fake.prepareArgsForCall)
}

func (fake *FakeConn) PrepareArgsForCall(i int) string {
	fake.prepareMutex.RLock()
	defer fake.prepareMutex.RUnlock()
	return fake.prepareArgsForCall[i].query
}

func (fake *FakeConn) PrepareReturns(result1 *sql.Stmt, result2 error) {
	fake.PrepareStub = nil
	fake.prepareReturns = struct {
		result1 *sql.Stmt
		result2 error
	}{result1, result2}
}

func (fake *FakeConn) Query(query string, args ...interface{}) (*sql.Rows, error) {
	fake.queryMutex.Lock()
	fake.queryArgsForCall = append(fake.queryArgsForCall, struct {
		query string
		args  []interface{}
	}{query, args})
	fake.recordInvocation("Query", []interface{}{query, args})
	fake.queryMutex.Unlock()
	if fake.QueryStub != nil {
		return fake.QueryStub(query, args...)
	} else {
		return fake.queryReturns.result1, fake.queryReturns.result2
	}
}

func (fake *FakeConn) QueryCallCount() int {
	fake.queryMutex.RLock()
	defer fake.queryMutex.RUnlock()
	return len(fake.queryArgsForCall)
}

func (fake *FakeConn) QueryArgsForCall(i int) (string, []interface{}) {
	fake.queryMutex.RLock()
	defer fake.queryMutex.RUnlock()
	return fake.queryArgsForCall[i].query, fake.queryArgsForCall[i].args
}

func (fake *FakeConn) QueryReturns(result1 *sql.Rows, result2 error) {
	fake.QueryStub = nil
	fake.queryReturns = struct {
		result1 *sql.Rows
		result2 error
	}{result1, result2}
}

func (fake *FakeConn) QueryRow(query string, args ...interface{}) *sql.Row {
	fake.queryRowMutex.Lock()
	fake.queryRowArgsForCall = append(fake.queryRowArgsForCall, struct {
		query string
		args  []interface{}
	}{query, args})
	fake.recordInvocation("QueryRow", []interface{}{query, args})
	fake.queryRowMutex.Unlock()
	if fake.QueryRowStub != nil {
		return fake.QueryRowStub(query, args...)
	} else {
		return fake.queryRowReturns.result1
	}
}

func (fake *FakeConn) QueryRowCallCount() int {
	fake.queryRowMutex.RLock()
	defer fake.queryRowMutex.RUnlock()
	return len(fake.queryRowArgsForCall)
}

func (fake *FakeConn) QueryRowArgsForCall(i int) (string, []interface{}) {
	fake.queryRowMutex.RLock()
	defer fake.queryRowMutex.RUnlock()
	return fake.queryRowArgsForCall[i].query, fake.queryRowArgsForCall[i].args
}

func (fake *FakeConn) QueryRowReturns(result1 *sql.Row) {
	fake.QueryRowStub = nil
	fake.queryRowReturns = struct {
		result1 *sql.Row
	}{result1}
}

func (fake *FakeConn) SetMaxIdleConns(n int) {
	fake.setMaxIdleConnsMutex.Lock()
	fake.setMaxIdleConnsArgsForCall = append(fake.setMaxIdleConnsArgsForCall, struct {
		n int
	}{n})
	fake.recordInvocation("SetMaxIdleConns", []interface{}{n})
	fake.setMaxIdleConnsMutex.Unlock()
	if fake.SetMaxIdleConnsStub != nil {
		fake.SetMaxIdleConnsStub(n)
	}
}

func (fake *FakeConn) SetMaxIdleConnsCallCount() int {
	fake.setMaxIdleConnsMutex.RLock()
	defer fake.setMaxIdleConnsMutex.RUnlock()
	return len(fake.setMaxIdleConnsArgsForCall)
}

func (fake *FakeConn) SetMaxIdleConnsArgsForCall(i int) int {
	fake.setMaxIdleConnsMutex.RLock()
	defer fake.setMaxIdleConnsMutex.RUnlock()
	return fake.setMaxIdleConnsArgsForCall[i].n
}

func (fake *FakeConn) SetMaxOpenConns(n int) {
	fake.setMaxOpenConnsMutex.Lock()
	fake.setMaxOpenConnsArgsForCall = append(fake.setMaxOpenConnsArgsForCall, struct {
		n int
	}{n})
	fake.recordInvocation("SetMaxOpenConns", []interface{}{n})
	fake.setMaxOpenConnsMutex.Unlock()
	if fake.SetMaxOpenConnsStub != nil {
		fake.SetMaxOpenConnsStub(n)
	}
}

func (fake *FakeConn) SetMaxOpenConnsCallCount() int {
	fake.setMaxOpenConnsMutex.RLock()
	defer fake.setMaxOpenConnsMutex.RUnlock()
	return len(fake.setMaxOpenConnsArgsForCall)
}

func (fake *FakeConn) SetMaxOpenConnsArgsForCall(i int) int {
	fake.setMaxOpenConnsMutex.RLock()
	defer fake.setMaxOpenConnsMutex.RUnlock()
	return fake.setMaxOpenConnsArgsForCall[i].n
}

func (fake *FakeConn) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.beginMutex.RLock()
	defer fake.beginMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.driverMutex.RLock()
	defer fake.driverMutex.RUnlock()
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	fake.pingMutex.RLock()
	defer fake.pingMutex.RUnlock()
	fake.prepareMutex.RLock()
	defer fake.prepareMutex.RUnlock()
	fake.queryMutex.RLock()
	defer fake.queryMutex.RUnlock()
	fake.queryRowMutex.RLock()
	defer fake.queryRowMutex.RUnlock()
	fake.setMaxIdleConnsMutex.RLock()
	defer fake.setMaxIdleConnsMutex.RUnlock()
	fake.setMaxOpenConnsMutex.RLock()
	defer fake.setMaxOpenConnsMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeConn) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ db.Conn = new(FakeConn)
