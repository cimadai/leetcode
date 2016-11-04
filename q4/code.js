/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number}
 */
var findMedianSortedArrays = function(nums1, nums2) {
    var total = nums1.length + nums2.length;
    var isOdd = total & 1;
    var med0, med1;
    if (isOdd) {
        med0 = (total/2)|0;
    } else {
        med1 = (total/2)|0;
        med0 = med1 - 1;
    }

    var idx1 = 0, idx2 = 0, sum = 0;
    for (var i = 0, l = isOdd ? med0 : med1; i < l + 1; ++i) {
        var n1 = (nums1[idx1] + 1)|0,
            n2 = (nums2[idx2] + 1)|0,
            n  = n1 == 0 ? (++idx2, n2) : n2 == 0 ? (++idx1, n1) : n1 < n2 ? (++idx1, n1) : (++idx2, n2);

        if (med0 == i || med1 == i) {
            sum += n;
        }
    }

    return isOdd ? sum - 1 : (sum - 2) / 2;
};

