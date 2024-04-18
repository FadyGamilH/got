# got
got is just an implementation of git using go, trying to implement some of the features git supports

# Part (1):
- This part is all about blobs, how to hash blobs just like git is doing.
### What is the difference btn blobs and files ? 
- Blobs are `Binary Large Objects` which is the content of the file.
- Blob is represented in `SHA-1HASH`, which are 20 bytes that identify this blob in git world.
- Files are the actuall file with it's own metadata, and this is the difference between `files` and `blobs`, if we create a file and move it to another directory, its `creation date` will remain the same, but blobs doesn't keep these info.

- To use the hashing that git is using : 
```bash
echo -n <content> | git hash-object --stdin
```

- To get the Sha1-Hash of a content : 
```bash
echo -n <content> | shasum -a 1
```

- So git uses the shasum (sha1-hash) to give you the final hash of the content, but if we used the 2 commands above on the same content string we will get different hashes ? 
    - Because `git hash-object` command uses `shasum` command with other data : 
        ```bash
        shasum "blob <content-size>\o<conent-data>"
        ```
    - Example : 
    ```bash
        > echo -n 'blob 11\0git with go' | shasum -a 1
        fd790ad99ff75c6c383962e2f0bc1ffeabc22142  -
                    
        > echo -n 'git with go' | shasum -a 1         
        1cd141fb363aca49017c83d35f3609d1819eb3b0  -
                    
        > echo -n 'git with go' | git hash-object --stdin                                                 
        fd790ad99ff75c6c383962e2f0bc1ffeabc22142
    ```