package watch

import (

)

// broker is all about events
type broker struct {
		streamCount int
		in chan<-
		reqChan <-chan
}


readStream := func(
		done <-chan interface{},
		chanStream <-chan <-chan interface{},
) <-chan interface {
		valStream := make(chan interface{})
		go func() {
				defer close(valStream)
				for {
						var stream <-chan interface{}
						select {
						case maybeStream, ok := <-chanStream:
								if !ok {
										return
								}
								stream = maybeStream
						case <-done:
								return
						}
						for val := range orDone(done, stream) {
								select {
								case valStream <-val:
								case <-done:
								}
						}
				}
		}()
}	

func main() {
		go readStream()
}
