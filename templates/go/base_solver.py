
import bisect

def main():
    # Read size of the array
    n = int(input())

    # Read n elements in one line, split by space
    arr = list(map(int, input().split()))

    el = int(input())

    res = bisect.bisect_left(arr, el)

    print(res)

if __name__ == "__main__":
    main()