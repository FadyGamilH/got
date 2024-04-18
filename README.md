# got
got is just an implementation of git using go, trying to implement some of the features git supports

# Coding Session #1:
- Day (1) will be about blobs, how to hash blobs just like git is doing.
### What is the difference btn blobs and files ? 
- Blobs are `Binary Large Objects` which is the content of the file.
- Blob is represented in `SHA-1HASH`, which are 20 bytes that identify this blob in git world.
- Files are the actuall file with it's own metadata, and this is the difference between `files` and `blobs`, if we create a file and move it to another directory, its `creation date` will remain the same, but blobs doesn't keep these info.

- To use the hashing that git is using : 
```bash
echo -n <content> | git hash-object --stdin
```
