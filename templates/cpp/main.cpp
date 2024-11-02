#include <iostream>
#include <vector>

size_t binary_search(const std::vector<int>& array, int element) {
    size_t start = 0;
    size_t end = array.size();

    while (start < end) {
        size_t pivot = (start + end) / 2; // assume there is no overflow
        if (element < array[pivot]) {
            end = pivot;
        } else {
            start = pivot + 1;
        }
    }

    return start;
}

int main() {
    std::ios_base::sync_with_stdio(false);
    std::cin.tie(NULL);

    size_t arr_size;
    std::cin >> arr_size;
    std::vector<int> arr(arr_size);

    for (unsigned int i = 0; i < arr.size(); i++) {
        std::cin >> arr[i];
    }
    int inserted;

    std::cin >> inserted;

    std::cout << binary_search(arr, inserted);
    return 0;
}