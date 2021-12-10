package xreflect

import (
    "reflect"
    "testing"

    "github.com/stretchr/testify/assert"
)

/**
 * @author Rancho
 * @date 2021/12/11
 */

type Test struct {
    ID   int
    Name string
}

func TestIsBasedOn(t *testing.T) {
    var b *bool
    assert.True(t, IsBasedOn(reflect.ValueOf(b).Type(), reflect.Bool))

    var i *int
    assert.True(t, IsBasedOn(reflect.ValueOf(i).Type(), reflect.Int))

    var a [3]int
    assert.True(t, IsBasedOn(reflect.ValueOf(a).Type(), reflect.Array))

    var c chan int
    assert.True(t, IsBasedOn(reflect.ValueOf(c).Type(), reflect.Chan))

    var f func()
    assert.True(t, IsBasedOn(reflect.ValueOf(f).Type(), reflect.Func))

    // FIXME call of reflect.Value.Type on zero Value
    // var itf interface{}
    // assert.True(t, IsBasedOn(reflect.ValueOf(itf).Type(), reflect.Interface))

    var m map[int]int
    assert.True(t, IsBasedOn(reflect.ValueOf(m).Type(), reflect.Map))

    var slice []int
    assert.True(t, IsBasedOn(reflect.ValueOf(slice).Type(), reflect.Slice))

    var s *string
    assert.True(t, IsBasedOn(reflect.ValueOf(s).Type(), reflect.String))

    var test *Test
    assert.True(t, IsBasedOn(reflect.ValueOf(test).Type(), reflect.Struct))

    var tests []*Test
    assert.True(t, IsBasedOn(reflect.ValueOf(tests).Type(), reflect.Slice))
}
