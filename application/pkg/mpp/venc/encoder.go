package venc

type NewDataMsg struct {

}

type Subscriber struct {
    data    chan NewDataMsg
}

type encoder struct {
    vencId      uint32
    subscribers map[*Subscriber]bool
    subscribe   chan *Subscriber //channel for subscription
    unsubscribe chan *Subscriber //channel for unsubscription
    newData     chan *NewDataMsg //channel for new data

    frames      frames
}

var encoders []encoder

func findEncoder(vencId uint32) (*encoder, error) { //
    
    return nil, nil
}

func createEncoder(vencId uint32) (*encoder, error) {
    return nil, nil
}
func (e *encoder) deleteEncoder() {}

func (e *encoder) subscribeEncoder(s *Subscriber) {}
func (e *encoder) unSubscribeEncoder(s *Subscriber) {}

func (e *encoder) routineEncoder() { //goroutine that process data to subscribers
    /*
        receive messages from
            subscribe channel
                on receive check is it new client or existing one (if existing WARN!)
                add client to list
            unsubscribe channel
                on reveive delete from clients list
            new frame channel
                iterate through clients, check input channel for each client if full than increase skip value for client
                for each client send info about new data
            delete itself channel
                on receive force unsubscribe all clients
    */
}
/* Destroy process
    Consider that children don`t have access to parent memory, access is only allowed by pointers
    before we will destroy hub, we should somehow be sure that all children will not try to access hub`s memory

    Option 1
        Parent closes channel, child on read from closed channel will notice it and inform parent by unsubscribe that 
        it is save to delete it from list, consider these parent can detroy itself safe on 0 clients.


*/
