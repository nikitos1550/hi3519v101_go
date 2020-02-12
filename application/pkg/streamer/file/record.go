//+build streamerFile

package file

/*
    /opt/storage/[hash] 
    /opt/storage/[hash]/info
    /opt/storage/[hash]/chunk[id].h264
    /opt/storage/[hash]/chunk[id].info 

*/

type chunk struct {

}

type record struct {
    hash    string
    chunks  []chunk
}

func (r *record) Load(hash string) {}
func (r *record) Read() {}

