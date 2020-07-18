//+build nobuild

package connection

import (
    //"fmt"

	"github.com/valyala/bytebufferpool"
)

type Storage struct {
    pool *bytebufferpool.Pool
    test *bytebufferpool.ByteBuffer
    length int
}

func (s *Storage) Init() {
    s.test = s.pool.Get()
    s.test.Write([]byte("dsfksdjflkdsjlkfjdslkfjlkdsjf"))
    //s.length = len(s.test)
    s.length = s.test.Len()
}
