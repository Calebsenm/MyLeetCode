
#include <iostream>
#include <vector>

class Solution {
public:
    std::vector<int> twoSum(const std::vector<int>& nums, int target) {
        std::vector<int> result; 

        for (size_t i = 0; i < nums.size(); i++) {
            for (size_t j = i + 1; j < nums.size(); j++) { 
                if (nums[i] + nums[j] == target) {
                    result.push_back(i); 
                    result.push_back(j); 
                }
            }
        }

        return result;
    }
};

int main() {
    Solution solution; 
    std::vector<int> numbers = {2, 7, 11, 15}; 
    int target = 9; 

    std::vector<int> resultado = solution.twoSum(numbers, target); 
    std::cout << "Índices encontrados: " << resultado[0] << ", " << resultado[1] << std::endl;
    
    return 0;
}

