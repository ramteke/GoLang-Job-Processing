Asynchronous processing using golang


Main: 
1. Generate data on input channel
2. Content from input channel is extracted and passed to Workers on job Channel.
   Worker allows to reduces no of concurrent jobs
3. Worker pics data from job Channel gives it to Business logic for processing
4. Business Logic maintains state information about input job
   State is stored as map[id]State using goroutine.
   This allows mutually execulive access to state data
   
   Business logic uses input and state data to generate a new output as byte[] on output channel
5. Output channel is picked up by OutputWriter which writes the data to console
  




