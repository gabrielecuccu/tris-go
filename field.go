package main

type field [3][3]cell

func newEmptyField() *field {
    return &field {
        {emptyCell, emptyCell, emptyCell},
        {emptyCell, emptyCell, emptyCell},
        {emptyCell, emptyCell, emptyCell},
    }
}