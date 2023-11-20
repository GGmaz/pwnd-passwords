# Exercise

## Datapoints
- Time limit: 4 hours
- Programming language: Go


## Task
Your task is to write a go program which reads the `pwnd-passwords.txt` (haveibeenpwnd dump) file (180M unzipped, but we will test your solution with a 5GB file)
and prints out the hash with the highest digit sum (if multiple hashes all have the highest digit sum, print them all) and the calculated digit sum. 
No other output is expected. 

The program should do this task in the most **memory** and **time** efficient way possible. CPU resource utilization is to be neglected.

---

Each line in the file has on the left to the colon a hash which is to be parsed. The colon as well as the number on the right side are to be ignored.

```
7C4A8D09CA3762AF61E59520943DC26494F8941B:37359195
      ^                                     ^
      |                                     |
  hash to be parsed                       ignore
```

The hash is displayed in hexadecimal form. Remove all non-numerical characters before calculating the digit sum.

Example:
```
7C4A8D09CA3762AF61E59520943DC26494F8941B --> 7480937626159520943264948941

--> digit sum = 137
```

Expected output if only one hash has the highest digit sum:
```
7C4A8D09CA3762AF61E59520943DC26494F8941B:137
```


## Result & Questions

Please upload the **resulting application** in this repository and add a **short design document** where you describe your major design decisions.
