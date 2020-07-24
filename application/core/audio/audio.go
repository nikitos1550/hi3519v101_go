package audio

import (
    "sync"
)

type audio struct {
    sync.RWMutex

    Id      int
    name    string

    Codec   int
}
