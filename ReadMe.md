## Go Concurrency

Concurrency Management: The sync.WaitGroup is a synchronization primitive that helps you wait for a collection of goroutines to finish executing. You use Add(1) to indicate that you're launching a new goroutine, and then Done() is called to signal that the goroutine has completed its work.

Ensures Completion: By using defer wg.Done(), you ensure that wg.Done() will be called when the fetchWather function exits, whether it exits normally or due to an error. This is important because if the function were to panic or return prematurely, you would still want to decrement the wait group counter to reflect that the goroutine has finished.

Prevents Deadlocks: If you forget to call wg.Done(), the main goroutine would wait indefinitely because the counter would never reach zero. Using defer guarantees that Done() will always be executed, preventing potential deadlocks.

Why Close the Channel?
Signaling Completion: Closing a channel signals to any goroutines reading from it that no more values will be sent on that channel. This is particularly useful in your example, where multiple goroutines are sending weather data to the channel. By closing the channel, you inform the receiver (in this case, the main goroutine) that it can stop reading values.

Preventing Deadlocks: If the channel is not closed, the goroutine reading from the channel (for result := range ch) may wait indefinitely for new values, leading to potential deadlocks. Closing the channel ensures that this goroutine can finish its execution cleanly.

Understanding the Code Snippet

go func() {
    wg.Wait()
    close(ch)
}()

Anonymous Goroutine: This snippet starts an anonymous goroutine that waits for all the fetching goroutines to finish. The wg.Wait() call blocks this goroutine until all the goroutines have called wg.Done().

Channel Closure: Once all the goroutines have completed, close(ch) is executed. This closes the channel ch, allowing the main goroutine to exit its loop that reads from the channel without waiting for more values.

Flow of Execution
The main goroutine launches several fetching goroutines and adds to the wait group.
The anonymous goroutine is started to wait for all the fetching goroutines to finish.

Once all fetching goroutines are done (when wg.Wait() returns), the anonymous goroutine closes the channel.
The main goroutine, which is reading from the channel in a for range loop, will detect that the channel is closed and terminate its loop, allowing the program to exit gracefully.

In summary, closing the channel after all goroutines are done is essential for proper resource management and ensuring that the reading side can finish its operation without waiting indefinitely.
