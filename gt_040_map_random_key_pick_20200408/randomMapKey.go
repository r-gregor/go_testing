package main

import (
    "fmt"
    "math/rand"
)

type stringIndex struct {
    s string
    n int
}

// M is logically a map[string]string that also supports efficiently choosing a
// random key.
type M struct {
    m map[string]stringIndex
    s []string
}

func NewM() *M {
    return &M{m: make(map[string]stringIndex)}
}

func (m *M) Get(k string) (string, bool) {
    si, ok := m.m[k]
    if !ok {
        return "", false
    }
    return si.s, true
}

func (m *M) Put(k, v string) {
    if _, ok := m.m[k]; ok {
        return
    }
    m.m[k] = stringIndex{v, len(m.s)}
    m.s = append(m.s, k)
}

func (m *M) Delete(k string) {
    si, ok := m.m[k]
    if !ok {
        return
    }
    // Remove element m.s[si.n] from the map by swapping in the last element
    // in its place.
    last := m.s[len(m.s)-1]
    m.s[si.n] = last
    m.s = m.s[:len(m.s)-1]
    m.m[last] = stringIndex{last, si.n}
    delete(m.m, k)
}

func (m *M) RandomKey() string {
    if len(m.s) == 0 {
        panic("RandomKey called on empty M")
    }
    return m.s[rand.Intn(len(m.s))]
}

func main() {
    m := NewM()
    m.Put("Mexico", "Mexico City")
    m.Put("Yemen", "San'a")
    m.Put("Ecuador", "Quito")
    m.Put("France", "Paris")
    m.Put("Nepal", "Kathmandu")
    m.Delete("Ecuador")
    m.Delete("Mexico")
    fmt.Println(m.RandomKey())
    fmt.Println(m.RandomKey())
    fmt.Println(m.RandomKey())
    fmt.Println(m.RandomKey())
    fmt.Println(m.RandomKey())
    fmt.Println(m.RandomKey())
}
