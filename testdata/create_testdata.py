import random
import time

# create N random characters and print countB
KB = 1024
MB = 1024 * KB
GB = 1024 * MB

def character_generator():
    while True:
        offset = random.randint(0, 25)
        yield chr(ord('a') + offset)

def generate_test_data(n: int) -> str:
    testdata = ""
    for next_chr in character_generator():
        testdata += next_chr
        n -= 1
        if n == 0:
            break
    return testdata

def write_test_data(testdata: str, file_ext: str):
    with open(f'./testfile-{file_ext}', 'w') as f:
        f.write(testdata)
    
def main():
    filesizes = {
        '1K': KB,
        '4K': 4 * KB,
        '64K': 64 * KB,
        '1M': MB,
        '4M': 4 * MB,
        '64M': 64 * MB,
        # '1G': GB,
        # '4G': 4 * GB
    }
    for file_ext, n in filesizes.items():
        start = time.time()
        print(f'Creating file - {file_ext}')
        write_test_data(generate_test_data(n), file_ext)
        print(f'Done...took {time.time() - start}')

if __name__ == '__main__':
    main()