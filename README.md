Asynchronous processing using golang


Main: 
1. Generate data on input channel.
2. Content from input channel is extracted and passed to Workers on job Channel.
   Worker allows to reduces no of concurrent jobs.
3. Worker loads data from job Channel gives it to Business logic for processing.
4. Business logic uses input and state data to generate a new output as byte[] on output channel.
   Business Logic also updates a singelton (concept) state map.
   The State keep counter of number of times a cluster Id is processed by business logic.
   Every parallel operation access the state and updates counter against the cluster ID. 
5. Output channel is picked up by OutputWriter which writes the data to console.
  
Output Trace
----------------
```
Generated:  SEQ: 0,  CLUSTER: 1
Generated:  SEQ: 1,  CLUSTER: 2
Generated:  SEQ: 2,  CLUSTER: 2
Generated:  SEQ: 3,  CLUSTER: 4
Generated:  SEQ: 4,  CLUSTER: 1
Generated:  SEQ: 5,  CLUSTER: 3
Generated:  SEQ: 6,  CLUSTER: 0
Generated:  SEQ: 7,  CLUSTER: 0
Generated:  SEQ: 8,  CLUSTER: 1
Generated:  SEQ: 9,  CLUSTER: 0
Generated:  SEQ: 10,  CLUSTER: 4
## WORKER: Created :  1
------------- WORKER [ 1 ] DISPATCHING :  SEQ: 0,  CLUSTER: 1
------------- WORKER [ 1 ] DISPATCHING :  SEQ: 1,  CLUSTER: 2
Generated:  SEQ: 11,  CLUSTER: 1
Generated:  SEQ: 12,  CLUSTER: 2
## WORKER: Created :  3
## WORKER: Created :  2
------------- WORKER [ 1 ] DISPATCHING :  SEQ: 2,  CLUSTER: 2
Generated:  SEQ: 13,  CLUSTER: 4
------------- WORKER [ 3 ] DISPATCHING :  SEQ: 3,  CLUSTER: 4
------------- WORKER [ 3 ] DISPATCHING :  SEQ: 5,  CLUSTER: 3
------------- WORKER [ 2 ] DISPATCHING :  SEQ: 4,  CLUSTER: 1
------------- WORKER [ 3 ] DISPATCHING :  SEQ: 6,  CLUSTER: 0
------------- WORKER [ 2 ] DISPATCHING :  SEQ: 7,  CLUSTER: 0
----------------------- : BUSINESS-LOGIC:  SEQ: 4,  CLUSTER: 1 via Worker: 2  Counter Value:  2
------------- WORKER [ 2 ] DISPATCHING :  SEQ: 8,  CLUSTER: 1
----------------------- : BUSINESS-LOGIC:  SEQ: 3,  CLUSTER: 4 via Worker: 3  Counter Value:  1
Generated:  SEQ: 14,  CLUSTER: 3
----------------------- : BUSINESS-LOGIC:  SEQ: 0,  CLUSTER: 1 via Worker: 1  Counter Value:  1
Generated:  SEQ: 15,  CLUSTER: 4
Generated:  SEQ: 16,  CLUSTER: 1
Generated:  SEQ: 17,  CLUSTER: 0
Generated:  SEQ: 18,  CLUSTER: 2
Generated:  SEQ: 19,  CLUSTER: 1
----------------------------------------OUTPUT: SEQ: 4,  CLUSTER: 1 via Worker: 2
----------------------- : BUSINESS-LOGIC:  SEQ: 7,  CLUSTER: 0 via Worker: 2  Counter Value:  1
------------- WORKER [ 1 ] DISPATCHING :  SEQ: 9,  CLUSTER: 0
----------------------------------------OUTPUT: SEQ: 3,  CLUSTER: 4 via Worker: 3
------------- WORKER [ 3 ] DISPATCHING :  SEQ: 10,  CLUSTER: 4
------------- WORKER [ 2 ] DISPATCHING :  SEQ: 11,  CLUSTER: 1
Generated:  SEQ: 20,  CLUSTER: 0
Generated:  SEQ: 21,  CLUSTER: 1
Generated:  SEQ: 22,  CLUSTER: 3
----------------------- : BUSINESS-LOGIC:  SEQ: 5,  CLUSTER: 3 via Worker: 3  Counter Value:  1
----------------------------------------OUTPUT: SEQ: 0,  CLUSTER: 1 via Worker: 1
----------------------- : BUSINESS-LOGIC:  SEQ: 1,  CLUSTER: 2 via Worker: 1  Counter Value:  1
------------- WORKER [ 3 ] DISPATCHING :  SEQ: 12,  CLUSTER: 2
----------------------- : BUSINESS-LOGIC:  SEQ: 6,  CLUSTER: 0 via Worker: 3  Counter Value:  2
------------- WORKER [ 2 ] DISPATCHING :  SEQ: 13,  CLUSTER: 4
----------------------- : BUSINESS-LOGIC:  SEQ: 2,  CLUSTER: 2 via Worker: 1  Counter Value:  2
------------- WORKER [ 1 ] DISPATCHING :  SEQ: 14,  CLUSTER: 3
----------------------------------------OUTPUT: SEQ: 1,  CLUSTER: 2 via Worker: 1
Generated:  SEQ: 23,  CLUSTER: 3
Generated:  SEQ: 24,  CLUSTER: 2
Generated:  SEQ: 25,  CLUSTER: 2
----------------------- : BUSINESS-LOGIC:  SEQ: 8,  CLUSTER: 1 via Worker: 2  Counter Value:  3
----------------------------------------OUTPUT: SEQ: 7,  CLUSTER: 0 via Worker: 2
----------------------------------------OUTPUT: SEQ: 5,  CLUSTER: 3 via Worker: 3
----------------------------------------OUTPUT: SEQ: 2,  CLUSTER: 2 via Worker: 1
------------- WORKER [ 2 ] DISPATCHING :  SEQ: 15,  CLUSTER: 4
Generated:  SEQ: 26,  CLUSTER: 2
Generated:  SEQ: 27,  CLUSTER: 3
Generated:  SEQ: 28,  CLUSTER: 0
------------- WORKER [ 1 ] DISPATCHING :  SEQ: 17,  CLUSTER: 0
----------------------------------------OUTPUT: SEQ: 6,  CLUSTER: 0 via Worker: 3
----------------------- : BUSINESS-LOGIC:  SEQ: 11,  CLUSTER: 1 via Worker: 2  Counter Value:  4
----------------------- : BUSINESS-LOGIC:  SEQ: 9,  CLUSTER: 0 via Worker: 1  Counter Value:  3
----------------------- : BUSINESS-LOGIC:  SEQ: 10,  CLUSTER: 4 via Worker: 3  Counter Value:  2
------------- WORKER [ 3 ] DISPATCHING :  SEQ: 16,  CLUSTER: 1
----------------------------------------OUTPUT: SEQ: 8,  CLUSTER: 1 via Worker: 2
------------- WORKER [ 1 ] DISPATCHING :  SEQ: 18,  CLUSTER: 2
Generated:  SEQ: 29,  CLUSTER: 0
Generated:  SEQ: 30,  CLUSTER: 1
Generated:  SEQ: 31,  CLUSTER: 3
----------------------------------------OUTPUT: SEQ: 9,  CLUSTER: 0 via Worker: 1
----------------------------------------OUTPUT: SEQ: 11,  CLUSTER: 1 via Worker: 2
------------- WORKER [ 2 ] DISPATCHING :  SEQ: 20,  CLUSTER: 0
----------------------- : BUSINESS-LOGIC:  SEQ: 12,  CLUSTER: 2 via Worker: 3  Counter Value:  3
----------------------- : BUSINESS-LOGIC:  SEQ: 13,  CLUSTER: 4 via Worker: 2  Counter Value:  3
----------------------- : BUSINESS-LOGIC:  SEQ: 14,  CLUSTER: 3 via Worker: 1  Counter Value:  2
----------------------------------------OUTPUT: SEQ: 10,  CLUSTER: 4 via Worker: 3
------------- WORKER [ 3 ] DISPATCHING :  SEQ: 19,  CLUSTER: 1
------------- WORKER [ 1 ] DISPATCHING :  SEQ: 21,  CLUSTER: 1
----------------------------------------OUTPUT: SEQ: 14,  CLUSTER: 3 via Worker: 1
----------------------- : BUSINESS-LOGIC:  SEQ: 16,  CLUSTER: 1 via Worker: 3  Counter Value:  5
----------------------------------------OUTPUT: SEQ: 12,  CLUSTER: 2 via Worker: 3
----------------------- : BUSINESS-LOGIC:  SEQ: 15,  CLUSTER: 4 via Worker: 2  Counter Value:  4
----------------------------------------OUTPUT: SEQ: 13,  CLUSTER: 4 via Worker: 2
------------- WORKER [ 2 ] DISPATCHING :  SEQ: 22,  CLUSTER: 3
----------------------- : BUSINESS-LOGIC:  SEQ: 17,  CLUSTER: 0 via Worker: 1  Counter Value:  4
Generated:  SEQ: 32,  CLUSTER: 2
Generated:  SEQ: 33,  CLUSTER: 1
Generated:  SEQ: 34,  CLUSTER: 4
------------- WORKER [ 3 ] DISPATCHING :  SEQ: 23,  CLUSTER: 3
----------------------------------------OUTPUT: SEQ: 16,  CLUSTER: 1 via Worker: 3
------------- WORKER [ 3 ] DISPATCHING :  SEQ: 24,  CLUSTER: 2
----------------------- : BUSINESS-LOGIC:  SEQ: 19,  CLUSTER: 1 via Worker: 3  Counter Value:  6
Generated:  SEQ: 35,  CLUSTER: 1
Generated:  SEQ: 36,  CLUSTER: 2
Generated:  SEQ: 37,  CLUSTER: 1
----------------------------------------OUTPUT: SEQ: 15,  CLUSTER: 4 via Worker: 2
------------- WORKER [ 1 ] DISPATCHING :  SEQ: 25,  CLUSTER: 2
------------- WORKER [ 2 ] DISPATCHING :  SEQ: 26,  CLUSTER: 2
----------------------- : BUSINESS-LOGIC:  SEQ: 18,  CLUSTER: 2 via Worker: 1  Counter Value:  4
----------------------------------------OUTPUT: SEQ: 17,  CLUSTER: 0 via Worker: 1
----------------------- : BUSINESS-LOGIC:  SEQ: 20,  CLUSTER: 0 via Worker: 2  Counter Value:  5
----------------------- : BUSINESS-LOGIC:  SEQ: 22,  CLUSTER: 3 via Worker: 2  Counter Value:  3
------------- WORKER [ 2 ] DISPATCHING :  SEQ: 27,  CLUSTER: 3
----------------------- : BUSINESS-LOGIC:  SEQ: 23,  CLUSTER: 3 via Worker: 3  Counter Value:  4
------------- WORKER [ 3 ] DISPATCHING :  SEQ: 29,  CLUSTER: 0
Generated:  SEQ: 38,  CLUSTER: 0
Generated:  SEQ: 39,  CLUSTER: 1
Generated:  SEQ: 40,  CLUSTER: 3
------------- WORKER [ 1 ] DISPATCHING :  SEQ: 28,  CLUSTER: 0
----------------------------------------OUTPUT: SEQ: 19,  CLUSTER: 1 via Worker: 3
----------------------------------------OUTPUT: SEQ: 20,  CLUSTER: 0 via Worker: 2
----------------------- : BUSINESS-LOGIC:  SEQ: 21,  CLUSTER: 1 via Worker: 1  Counter Value:  7
----------------------------------------OUTPUT: SEQ: 18,  CLUSTER: 2 via Worker: 1
----------------------- : BUSINESS-LOGIC:  SEQ: 25,  CLUSTER: 2 via Worker: 1  Counter Value:  5
------------- WORKER [ 1 ] DISPATCHING :  SEQ: 30,  CLUSTER: 1
Generated:  SEQ: 41,  CLUSTER: 0
------------- WORKER [ 3 ] DISPATCHING :  SEQ: 31,  CLUSTER: 3
Generated:  SEQ: 42,  CLUSTER: 4
----------------------------------------OUTPUT: SEQ: 23,  CLUSTER: 3 via Worker: 3
------------- WORKER [ 2 ] DISPATCHING :  SEQ: 32,  CLUSTER: 2
Generated:  SEQ: 43,  CLUSTER: 3
----------------------------------------OUTPUT: SEQ: 22,  CLUSTER: 3 via Worker: 2
----------------------- : BUSINESS-LOGIC:  SEQ: 24,  CLUSTER: 2 via Worker: 3  Counter Value:  6
----------------------- : BUSINESS-LOGIC:  SEQ: 26,  CLUSTER: 2 via Worker: 2  Counter Value:  6
----------------------------------------OUTPUT: SEQ: 21,  CLUSTER: 1 via Worker: 1
----------------------- : BUSINESS-LOGIC:  SEQ: 28,  CLUSTER: 0 via Worker: 1  Counter Value:  6
----------------------- : BUSINESS-LOGIC:  SEQ: 29,  CLUSTER: 0 via Worker: 3  Counter Value:  7
----------------------------------------OUTPUT: SEQ: 24,  CLUSTER: 2 via Worker: 3
----------------------- : BUSINESS-LOGIC:  SEQ: 27,  CLUSTER: 3 via Worker: 2  Counter Value:  5
----------------------------------------OUTPUT: SEQ: 25,  CLUSTER: 2 via Worker: 1
----------------------------------------OUTPUT: SEQ: 26,  CLUSTER: 2 via Worker: 2
------------- WORKER [ 2 ] DISPATCHING :  SEQ: 34,  CLUSTER: 4
------------- WORKER [ 1 ] DISPATCHING :  SEQ: 33,  CLUSTER: 1
Generated:  SEQ: 44,  CLUSTER: 3
Generated:  SEQ: 45,  CLUSTER: 2
------------- WORKER [ 3 ] DISPATCHING :  SEQ: 35,  CLUSTER: 1
Generated:  SEQ: 46,  CLUSTER: 3
----------------------- : BUSINESS-LOGIC:  SEQ: 31,  CLUSTER: 3 via Worker: 3  Counter Value:  6
----------------------------------------OUTPUT: SEQ: 29,  CLUSTER: 0 via Worker: 3
------------- WORKER [ 3 ] DISPATCHING :  SEQ: 36,  CLUSTER: 2
Generated:  SEQ: 47,  CLUSTER: 4
Generated:  SEQ: 48,  CLUSTER: 4
------------- WORKER [ 1 ] DISPATCHING :  SEQ: 37,  CLUSTER: 1
----------------------- : BUSINESS-LOGIC:  SEQ: 30,  CLUSTER: 1 via Worker: 1  Counter Value:  8
----------------------------------------OUTPUT: SEQ: 28,  CLUSTER: 0 via Worker: 1
----------------------------------------OUTPUT: SEQ: 27,  CLUSTER: 3 via Worker: 2
------------- WORKER [ 2 ] DISPATCHING :  SEQ: 38,  CLUSTER: 0
Generated:  SEQ: 49,  CLUSTER: 3
----------------------- : BUSINESS-LOGIC:  SEQ: 32,  CLUSTER: 2 via Worker: 2  Counter Value:  7
----------------------- : BUSINESS-LOGIC:  SEQ: 34,  CLUSTER: 4 via Worker: 2  Counter Value:  5
----------------------------------------OUTPUT: SEQ: 32,  CLUSTER: 2 via Worker: 2
------------- WORKER [ 2 ] DISPATCHING :  SEQ: 39,  CLUSTER: 1
Generated:  SEQ: 50,  CLUSTER: 2
----------------------- : BUSINESS-LOGIC:  SEQ: 33,  CLUSTER: 1 via Worker: 1  Counter Value:  9
----------------------- : BUSINESS-LOGIC:  SEQ: 35,  CLUSTER: 1 via Worker: 3  Counter Value:  10
----------------------------------------OUTPUT: SEQ: 30,  CLUSTER: 1 via Worker: 1
------------- WORKER [ 1 ] DISPATCHING :  SEQ: 40,  CLUSTER: 3
Generated:  SEQ: 51,  CLUSTER: 1
------------- WORKER [ 3 ] DISPATCHING :  SEQ: 41,  CLUSTER: 0
Generated:  SEQ: 52,  CLUSTER: 4
----------------------------------------OUTPUT: SEQ: 31,  CLUSTER: 3 via Worker: 3
------------- WORKER [ 1 ] DISPATCHING :  SEQ: 42,  CLUSTER: 4
----------------------------------------OUTPUT: SEQ: 34,  CLUSTER: 4 via Worker: 2
------------- WORKER [ 2 ] DISPATCHING :  SEQ: 44,  CLUSTER: 3
Generated:  SEQ: 53,  CLUSTER: 4
------------- WORKER [ 3 ] DISPATCHING :  SEQ: 43,  CLUSTER: 3
----------------------------------------OUTPUT: SEQ: 35,  CLUSTER: 1 via Worker: 3
Generated:  SEQ: 54,  CLUSTER: 0
Generated:  SEQ: 55,  CLUSTER: 0
----------------------- : BUSINESS-LOGIC:  SEQ: 36,  CLUSTER: 2 via Worker: 3  Counter Value:  8
----------------------- : BUSINESS-LOGIC:  SEQ: 38,  CLUSTER: 0 via Worker: 2  Counter Value:  8
----------------------- : BUSINESS-LOGIC:  SEQ: 37,  CLUSTER: 1 via Worker: 1  Counter Value:  11
----------------------------------------OUTPUT: SEQ: 33,  CLUSTER: 1 via Worker: 1
----------------------- : BUSINESS-LOGIC:  SEQ: 39,  CLUSTER: 1 via Worker: 2  Counter Value:  12
Generated:  SEQ: 56,  CLUSTER: 3
----------------------------------------OUTPUT: SEQ: 38,  CLUSTER: 0 via Worker: 2
Generated:  SEQ: 57,  CLUSTER: 3
------------- WORKER [ 1 ] DISPATCHING :  SEQ: 45,  CLUSTER: 2
----------------------- : BUSINESS-LOGIC:  SEQ: 40,  CLUSTER: 3 via Worker: 1  Counter Value:  7
------------- WORKER [ 2 ] DISPATCHING :  SEQ: 46,  CLUSTER: 3
----------------------------------------OUTPUT: SEQ: 36,  CLUSTER: 2 via Worker: 3
------------- WORKER [ 3 ] DISPATCHING :  SEQ: 47,  CLUSTER: 4
----------------------------------------OUTPUT: SEQ: 37,  CLUSTER: 1 via Worker: 1
----------------------- : BUSINESS-LOGIC:  SEQ: 41,  CLUSTER: 0 via Worker: 3  Counter Value:  9
Generated:  SEQ: 58,  CLUSTER: 3
----------------------------------------OUTPUT: SEQ: 39,  CLUSTER: 1 via Worker: 2
----------------------- : BUSINESS-LOGIC:  SEQ: 44,  CLUSTER: 3 via Worker: 2  Counter Value:  8
----------------------- : BUSINESS-LOGIC:  SEQ: 42,  CLUSTER: 4 via Worker: 1  Counter Value:  6
------------- WORKER [ 2 ] DISPATCHING :  SEQ: 48,  CLUSTER: 4
----------------------------------------OUTPUT: SEQ: 41,  CLUSTER: 0 via Worker: 3
Generated:  SEQ: 59,  CLUSTER: 0
Generated:  SEQ: 60,  CLUSTER: 1
Generated:  SEQ: 61,  CLUSTER: 0
------------- WORKER [ 1 ] DISPATCHING :  SEQ: 50,  CLUSTER: 2
----------------------- : BUSINESS-LOGIC:  SEQ: 43,  CLUSTER: 3 via Worker: 3  Counter Value:  9
----------------------------------------OUTPUT: SEQ: 40,  CLUSTER: 3 via Worker: 1
------------- WORKER [ 3 ] DISPATCHING :  SEQ: 49,  CLUSTER: 3
----------------------- : BUSINESS-LOGIC:  SEQ: 47,  CLUSTER: 4 via Worker: 3  Counter Value:  7
----------------------- : BUSINESS-LOGIC:  SEQ: 46,  CLUSTER: 3 via Worker: 2  Counter Value:  10
----------------------------------------OUTPUT: SEQ: 44,  CLUSTER: 3 via Worker: 2
----------------------- : BUSINESS-LOGIC:  SEQ: 45,  CLUSTER: 2 via Worker: 1  Counter Value:  9
----------------------------------------OUTPUT: SEQ: 43,  CLUSTER: 3 via Worker: 3
Generated:  SEQ: 62,  CLUSTER: 0
Generated:  SEQ: 63,  CLUSTER: 1
Generated:  SEQ: 64,  CLUSTER: 1
------------- WORKER [ 3 ] DISPATCHING :  SEQ: 52,  CLUSTER: 4
------------- WORKER [ 2 ] DISPATCHING :  SEQ: 51,  CLUSTER: 1
------------- WORKER [ 1 ] DISPATCHING :  SEQ: 53,  CLUSTER: 4
----------------------------------------OUTPUT: SEQ: 42,  CLUSTER: 4 via Worker: 1
----------------------- : BUSINESS-LOGIC:  SEQ: 48,  CLUSTER: 4 via Worker: 2  Counter Value:  8
----------------------------------------OUTPUT: SEQ: 46,  CLUSTER: 3 via Worker: 2
------------- WORKER [ 2 ] DISPATCHING :  SEQ: 54,  CLUSTER: 0
Generated:  SEQ: 65,  CLUSTER: 3
----------------------- : BUSINESS-LOGIC:  SEQ: 49,  CLUSTER: 3 via Worker: 3  Counter Value:  11
------------- WORKER [ 3 ] DISPATCHING :  SEQ: 55,  CLUSTER: 0
Generated:  SEQ: 66,  CLUSTER: 1
----------------------------------------OUTPUT: SEQ: 47,  CLUSTER: 4 via Worker: 3
----------------------- : BUSINESS-LOGIC:  SEQ: 50,  CLUSTER: 2 via Worker: 1  Counter Value:  10
----------------------------------------OUTPUT: SEQ: 45,  CLUSTER: 2 via Worker: 1
Generated:  SEQ: 67,  CLUSTER: 2
------------- WORKER [ 1 ] DISPATCHING :  SEQ: 56,  CLUSTER: 3
----------------------- : BUSINESS-LOGIC:  SEQ: 52,  CLUSTER: 4 via Worker: 3  Counter Value:  9
------------- WORKER [ 3 ] DISPATCHING :  SEQ: 57,  CLUSTER: 3
----------------------------------------OUTPUT: SEQ: 49,  CLUSTER: 3 via Worker: 3
Generated:  SEQ: 68,  CLUSTER: 3
----------------------------------------OUTPUT: SEQ: 48,  CLUSTER: 4 via Worker: 2
----------------------- : BUSINESS-LOGIC:  SEQ: 53,  CLUSTER: 4 via Worker: 1  Counter Value:  9
------------- WORKER [ 2 ] DISPATCHING :  SEQ: 59,  CLUSTER: 0
----------------------- : BUSINESS-LOGIC:  SEQ: 51,  CLUSTER: 1 via Worker: 2  Counter Value:  13
----------------------------------------OUTPUT: SEQ: 50,  CLUSTER: 2 via Worker: 1
Generated:  SEQ: 69,  CLUSTER: 1
Generated:  SEQ: 70,  CLUSTER: 3
------------- WORKER [ 1 ] DISPATCHING :  SEQ: 58,  CLUSTER: 3
----------------------- : BUSINESS-LOGIC:  SEQ: 55,  CLUSTER: 0 via Worker: 3  Counter Value:  10
----------------------- : BUSINESS-LOGIC:  SEQ: 56,  CLUSTER: 3 via Worker: 1  Counter Value:  12
----------------------------------------OUTPUT: SEQ: 52,  CLUSTER: 4 via Worker: 3
----------------------------------------OUTPUT: SEQ: 53,  CLUSTER: 4 via Worker: 1
------------- WORKER [ 3 ] DISPATCHING :  SEQ: 60,  CLUSTER: 1
Generated:  SEQ: 71,  CLUSTER: 1
Generated:  SEQ: 72,  CLUSTER: 2
------------- WORKER [ 1 ] DISPATCHING :  SEQ: 61,  CLUSTER: 0
----------------------------------------OUTPUT: SEQ: 51,  CLUSTER: 1 via Worker: 2
----------------------- : BUSINESS-LOGIC:  SEQ: 54,  CLUSTER: 0 via Worker: 2  Counter Value:  11
------------- WORKER [ 2 ] DISPATCHING :  SEQ: 62,  CLUSTER: 0
Generated:  SEQ: 73,  CLUSTER: 3
----------------------- : BUSINESS-LOGIC:  SEQ: 58,  CLUSTER: 3 via Worker: 1  Counter Value:  13
------------- WORKER [ 2 ] DISPATCHING :  SEQ: 63,  CLUSTER: 1
----------------------------------------OUTPUT: SEQ: 54,  CLUSTER: 0 via Worker: 2
----------------------------------------OUTPUT: SEQ: 56,  CLUSTER: 3 via Worker: 1
----------------------- : BUSINESS-LOGIC:  SEQ: 59,  CLUSTER: 0 via Worker: 2  Counter Value:  12
------------- WORKER [ 1 ] DISPATCHING :  SEQ: 64,  CLUSTER: 1
----------------------------------------OUTPUT: SEQ: 55,  CLUSTER: 0 via Worker: 3
----------------------- : BUSINESS-LOGIC:  SEQ: 57,  CLUSTER: 3 via Worker: 3  Counter Value:  14
------------- WORKER [ 3 ] DISPATCHING :  SEQ: 65,  CLUSTER: 3
Generated:  SEQ: 74,  CLUSTER: 2
Generated:  SEQ: 75,  CLUSTER: 4
Generated:  SEQ: 76,  CLUSTER: 2
----------------------- : BUSINESS-LOGIC:  SEQ: 60,  CLUSTER: 1 via Worker: 3  Counter Value:  14
------------- WORKER [ 3 ] DISPATCHING :  SEQ: 66,  CLUSTER: 1
----------------------- : BUSINESS-LOGIC:  SEQ: 62,  CLUSTER: 0 via Worker: 2  Counter Value:  13
Generated:  SEQ: 77,  CLUSTER: 3
----------------------------------------OUTPUT: SEQ: 57,  CLUSTER: 3 via Worker: 3
----------------------------------------OUTPUT: SEQ: 59,  CLUSTER: 0 via Worker: 2
------------- WORKER [ 2 ] DISPATCHING :  SEQ: 67,  CLUSTER: 2
Generated:  SEQ: 78,  CLUSTER: 1
----------------------- : BUSINESS-LOGIC:  SEQ: 61,  CLUSTER: 0 via Worker: 1  Counter Value:  14
----------------------------------------OUTPUT: SEQ: 58,  CLUSTER: 3 via Worker: 1
------------- WORKER [ 1 ] DISPATCHING :  SEQ: 68,  CLUSTER: 3
Generated:  SEQ: 79,  CLUSTER: 0
----------------------- : BUSINESS-LOGIC:  SEQ: 63,  CLUSTER: 1 via Worker: 2  Counter Value:  15
----------------------------------------OUTPUT: SEQ: 61,  CLUSTER: 0 via Worker: 1
----------------------------------------OUTPUT: SEQ: 62,  CLUSTER: 0 via Worker: 2
------------- WORKER [ 1 ] DISPATCHING :  SEQ: 70,  CLUSTER: 3
Generated:  SEQ: 80,  CLUSTER: 3
Generated:  SEQ: 81,  CLUSTER: 3
------------- WORKER [ 2 ] DISPATCHING :  SEQ: 71,  CLUSTER: 1
----------------------- : BUSINESS-LOGIC:  SEQ: 64,  CLUSTER: 1 via Worker: 1  Counter Value:  16
------------- WORKER [ 3 ] DISPATCHING :  SEQ: 69,  CLUSTER: 1
----------------------- : BUSINESS-LOGIC:  SEQ: 65,  CLUSTER: 3 via Worker: 3  Counter Value:  15
Generated:  SEQ: 82,  CLUSTER: 2
----------------------------------------OUTPUT: SEQ: 60,  CLUSTER: 1 via Worker: 3
----------------------- : BUSINESS-LOGIC:  SEQ: 68,  CLUSTER: 3 via Worker: 1  Counter Value:  16
----------------------- : BUSINESS-LOGIC:  SEQ: 66,  CLUSTER: 1 via Worker: 3  Counter Value:  17
Generated:  SEQ: 83,  CLUSTER: 3
----------------------------------------OUTPUT: SEQ: 64,  CLUSTER: 1 via Worker: 1
----------------------------------------OUTPUT: SEQ: 65,  CLUSTER: 3 via Worker: 3
------------- WORKER [ 1 ] DISPATCHING :  SEQ: 73,  CLUSTER: 3
----------------------------------------OUTPUT: SEQ: 63,  CLUSTER: 1 via Worker: 2
----------------------- : BUSINESS-LOGIC:  SEQ: 67,  CLUSTER: 2 via Worker: 2  Counter Value:  11
------------- WORKER [ 2 ] DISPATCHING :  SEQ: 74,  CLUSTER: 2
Generated:  SEQ: 84,  CLUSTER: 1
Generated:  SEQ: 85,  CLUSTER: 4
------------- WORKER [ 3 ] DISPATCHING :  SEQ: 72,  CLUSTER: 2
------------- WORKER [ 3 ] DISPATCHING :  SEQ: 75,  CLUSTER: 4
----------------------------------------OUTPUT: SEQ: 68,  CLUSTER: 3 via Worker: 1
----------------------------------------OUTPUT: SEQ: 66,  CLUSTER: 1 via Worker: 3
------------- WORKER [ 1 ] DISPATCHING :  SEQ: 76,  CLUSTER: 2
----------------------- : BUSINESS-LOGIC:  SEQ: 69,  CLUSTER: 1 via Worker: 3  Counter Value:  18
----------------------------------------OUTPUT: SEQ: 67,  CLUSTER: 2 via Worker: 2
----------------------- : BUSINESS-LOGIC:  SEQ: 70,  CLUSTER: 3 via Worker: 1  Counter Value:  17
Generated:  SEQ: 86,  CLUSTER: 3
Generated:  SEQ: 87,  CLUSTER: 3
Generated:  SEQ: 88,  CLUSTER: 1
------------- WORKER [ 2 ] DISPATCHING :  SEQ: 77,  CLUSTER: 3
----------------------- : BUSINESS-LOGIC:  SEQ: 71,  CLUSTER: 1 via Worker: 2  Counter Value:  18
----------------------------------------OUTPUT: SEQ: 71,  CLUSTER: 1 via Worker: 2
----------------------- : BUSINESS-LOGIC:  SEQ: 74,  CLUSTER: 2 via Worker: 2  Counter Value:  12
------------- WORKER [ 2 ] DISPATCHING :  SEQ: 78,  CLUSTER: 1
Generated:  SEQ: 89,  CLUSTER: 2
Generated:  SEQ: 90,  CLUSTER: 3
Generated:  SEQ: 91,  CLUSTER: 1
------------- WORKER [ 1 ] DISPATCHING :  SEQ: 80,  CLUSTER: 3
----------------------- : BUSINESS-LOGIC:  SEQ: 72,  CLUSTER: 2 via Worker: 3  Counter Value:  13
------------- WORKER [ 3 ] DISPATCHING :  SEQ: 79,  CLUSTER: 0
----------------------- : BUSINESS-LOGIC:  SEQ: 73,  CLUSTER: 3 via Worker: 1  Counter Value:  18
----------------------------------------OUTPUT: SEQ: 69,  CLUSTER: 1 via Worker: 3
----------------------------------------OUTPUT: SEQ: 70,  CLUSTER: 3 via Worker: 1
----------------------- : BUSINESS-LOGIC:  SEQ: 75,  CLUSTER: 4 via Worker: 3  Counter Value:  10
------------- WORKER [ 2 ] DISPATCHING :  SEQ: 82,  CLUSTER: 2
----------------------- : BUSINESS-LOGIC:  SEQ: 77,  CLUSTER: 3 via Worker: 2  Counter Value:  19
----------------------------------------OUTPUT: SEQ: 74,  CLUSTER: 2 via Worker: 2
----------------------------------------OUTPUT: SEQ: 72,  CLUSTER: 2 via Worker: 3
Generated:  SEQ: 92,  CLUSTER: 1
Generated:  SEQ: 93,  CLUSTER: 2
Generated:  SEQ: 94,  CLUSTER: 0
------------- WORKER [ 3 ] DISPATCHING :  SEQ: 81,  CLUSTER: 3
----------------------- : BUSINESS-LOGIC:  SEQ: 76,  CLUSTER: 2 via Worker: 1  Counter Value:  14
----------------------------------------OUTPUT: SEQ: 73,  CLUSTER: 3 via Worker: 1
------------- WORKER [ 1 ] DISPATCHING :  SEQ: 83,  CLUSTER: 3
----------------------- : BUSINESS-LOGIC:  SEQ: 80,  CLUSTER: 3 via Worker: 1  Counter Value:  20
----------------------- : BUSINESS-LOGIC:  SEQ: 78,  CLUSTER: 1 via Worker: 2  Counter Value:  19
----------------------------------------OUTPUT: SEQ: 76,  CLUSTER: 2 via Worker: 1
------------- WORKER [ 1 ] DISPATCHING :  SEQ: 84,  CLUSTER: 1
Generated:  SEQ: 95,  CLUSTER: 3
------------- WORKER [ 2 ] DISPATCHING :  SEQ: 85,  CLUSTER: 4
Generated:  SEQ: 96,  CLUSTER: 2
----------------------- : BUSINESS-LOGIC:  SEQ: 79,  CLUSTER: 0 via Worker: 3  Counter Value:  15
----------------------------------------OUTPUT: SEQ: 77,  CLUSTER: 3 via Worker: 2
------------- WORKER [ 3 ] DISPATCHING :  SEQ: 86,  CLUSTER: 3
----------------------------------------OUTPUT: SEQ: 75,  CLUSTER: 4 via Worker: 3
Generated:  SEQ: 97,  CLUSTER: 3
----------------------------------------OUTPUT: SEQ: 80,  CLUSTER: 3 via Worker: 1
----------------------- : BUSINESS-LOGIC:  SEQ: 82,  CLUSTER: 2 via Worker: 2  Counter Value:  15
------------- WORKER [ 3 ] DISPATCHING :  SEQ: 87,  CLUSTER: 3
------------- WORKER [ 1 ] DISPATCHING :  SEQ: 88,  CLUSTER: 1
Generated:  SEQ: 98,  CLUSTER: 0
Generated:  SEQ: 99,  CLUSTER: 3
----------------------- : BUSINESS-LOGIC:  SEQ: 83,  CLUSTER: 3 via Worker: 1  Counter Value:  21
----------------------------------------OUTPUT: SEQ: 78,  CLUSTER: 1 via Worker: 2
------------- WORKER [ 2 ] DISPATCHING :  SEQ: 89,  CLUSTER: 2
----------------------- : BUSINESS-LOGIC:  SEQ: 81,  CLUSTER: 3 via Worker: 3  Counter Value:  21
----------------------------------------OUTPUT: SEQ: 79,  CLUSTER: 0 via Worker: 3
----------------------- : BUSINESS-LOGIC:  SEQ: 85,  CLUSTER: 4 via Worker: 2  Counter Value:  11
------------- WORKER [ 2 ] DISPATCHING :  SEQ: 90,  CLUSTER: 3
----------------------------------------OUTPUT: SEQ: 82,  CLUSTER: 2 via Worker: 2
----------------------------------------OUTPUT: SEQ: 81,  CLUSTER: 3 via Worker: 3
----------------------- : BUSINESS-LOGIC:  SEQ: 86,  CLUSTER: 3 via Worker: 3  Counter Value:  22
------------- WORKER [ 1 ] DISPATCHING :  SEQ: 91,  CLUSTER: 1
------------- WORKER [ 3 ] DISPATCHING :  SEQ: 92,  CLUSTER: 1
----------------------- : BUSINESS-LOGIC:  SEQ: 84,  CLUSTER: 1 via Worker: 1  Counter Value:  20
----------------------------------------OUTPUT: SEQ: 83,  CLUSTER: 3 via Worker: 1
----------------------- : BUSINESS-LOGIC:  SEQ: 88,  CLUSTER: 1 via Worker: 1  Counter Value:  21
----------------------- : BUSINESS-LOGIC:  SEQ: 87,  CLUSTER: 3 via Worker: 3  Counter Value:  23
----------------------------------------OUTPUT: SEQ: 84,  CLUSTER: 1 via Worker: 1
----------------------------------------OUTPUT: SEQ: 86,  CLUSTER: 3 via Worker: 3
------------- WORKER [ 1 ] DISPATCHING :  SEQ: 93,  CLUSTER: 2
------------- WORKER [ 3 ] DISPATCHING :  SEQ: 94,  CLUSTER: 0
----------------------- : BUSINESS-LOGIC:  SEQ: 89,  CLUSTER: 2 via Worker: 2  Counter Value:  16
----------------------------------------OUTPUT: SEQ: 85,  CLUSTER: 4 via Worker: 2
------------- WORKER [ 2 ] DISPATCHING :  SEQ: 95,  CLUSTER: 3
----------------------- : BUSINESS-LOGIC:  SEQ: 91,  CLUSTER: 1 via Worker: 1  Counter Value:  22
----------------------------------------OUTPUT: SEQ: 88,  CLUSTER: 1 via Worker: 1
----------------------- : BUSINESS-LOGIC:  SEQ: 90,  CLUSTER: 3 via Worker: 2  Counter Value:  24
----------------------- : BUSINESS-LOGIC:  SEQ: 92,  CLUSTER: 1 via Worker: 3  Counter Value:  23
----------------------------------------OUTPUT: SEQ: 89,  CLUSTER: 2 via Worker: 2
------------- WORKER [ 2 ] DISPATCHING :  SEQ: 96,  CLUSTER: 2
----------------------------------------OUTPUT: SEQ: 87,  CLUSTER: 3 via Worker: 3
------------- WORKER [ 3 ] DISPATCHING :  SEQ: 97,  CLUSTER: 3
------------- WORKER [ 1 ] DISPATCHING :  SEQ: 98,  CLUSTER: 0
----------------------- : BUSINESS-LOGIC:  SEQ: 95,  CLUSTER: 3 via Worker: 2  Counter Value:  25
----------------------- : BUSINESS-LOGIC:  SEQ: 94,  CLUSTER: 0 via Worker: 3  Counter Value:  16
----------------------------------------OUTPUT: SEQ: 92,  CLUSTER: 1 via Worker: 3
------------- WORKER [ 3 ] DISPATCHING :  SEQ: 99,  CLUSTER: 3
----------------------------------------OUTPUT: SEQ: 90,  CLUSTER: 3 via Worker: 2
----------------------- : BUSINESS-LOGIC:  SEQ: 93,  CLUSTER: 2 via Worker: 1  Counter Value:  17
----------------------------------------OUTPUT: SEQ: 91,  CLUSTER: 1 via Worker: 1
 ################ WORKER [ 1 ] processed  33 jobs  ###############
 ################ WORKER [ 2 ] processed  33 jobs  ###############
----------------------------------------OUTPUT: SEQ: 95,  CLUSTER: 3 via Worker: 2
 ################ WORKER [ 3 ] processed  34 jobs  ###############
----------------------------------------OUTPUT: SEQ: 94,  CLUSTER: 0 via Worker: 3
----------------------- : BUSINESS-LOGIC:  SEQ: 96,  CLUSTER: 2 via Worker: 2  Counter Value:  18
----------------------- : BUSINESS-LOGIC:  SEQ: 97,  CLUSTER: 3 via Worker: 3  Counter Value:  26
----------------------------------------OUTPUT: SEQ: 93,  CLUSTER: 2 via Worker: 1
----------------------- : BUSINESS-LOGIC:  SEQ: 98,  CLUSTER: 0 via Worker: 1  Counter Value:  17
----------------------------------------OUTPUT: SEQ: 96,  CLUSTER: 2 via Worker: 2
----------------------------------------OUTPUT: SEQ: 98,  CLUSTER: 0 via Worker: 1
----------------------------------------OUTPUT: SEQ: 97,  CLUSTER: 3 via Worker: 3
----------------------- : BUSINESS-LOGIC:  SEQ: 99,  CLUSTER: 3 via Worker: 3  Counter Value:  27
----------------------------------------OUTPUT: SEQ: 99,  CLUSTER: 3 via Worker: 3

```



