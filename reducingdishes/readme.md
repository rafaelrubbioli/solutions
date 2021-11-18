# Reducing Dishes

https://leetcode.com/problems/reducing-dishes/

### Example 1:

Input: [-1,-8,0,5,-9]

Output: 14

### Example 2:

Input: [2 3 4]

Output: 20

### Example 3:

Input: [-5 -4 -1]

Output: 0

### Example 4:

Input: [-3 -2 -1 0 3 5]

Output: 35

## Solution
The max like-time coefficient will be the sum of the positive tails.
example: 

input = [-1, -8, 0, 5, -9]

sorted = [-9, -8, -1, 0, 5]

tail = (-9) + (-8) + (-1) + 0 + 5 = -13 (0)

tail = (-8) + (-1) + 0 + 5 = -4 (0)

tail = (-1) + 0 + 5 = 4 (4)

tail = 0 + 5 = 5 (9)

tail = 5 = 5 (14)


