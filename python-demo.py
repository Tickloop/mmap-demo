import mmap

# open the file in read/write/binary mode => needed for mmap 
fd = open('./testdata/testfile-1K', 'r+b')

# PROT and MAP flags are sane defaults PROT_READ and MAP_SHARED
mm = mmap.mmap(fd.fileno(), 0)

# read
print(mm[:10])

# properties similar* to list
print(len(mm))

# write
mm[0] = ord('a')

# always close files and memory maps!
mm.close()
fd.close()