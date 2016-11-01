/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) {
    var m = {}, i = 0, l = nums.length;
    while(i < l) {
        var e = nums[i]|0, x = m[e];
        if (x !== undefined) {
            return [x, i];
        }
        m[target - e] = i|0;
        i=(i+1)|0;
    }
    return [];
};
