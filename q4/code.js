/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number}
 */
var findMedianSortedArrays = function(nums1, nums2) {
    var total = nums1.length + nums2.length;
    var isOdd = total & 1;
    var med0, med1;
    if (!isOdd) {
        med1 = (total/2)|0;
        med0 = med1 - 1;
    } else {
        med0 = (total/2)|0;
    }

    var idx1 = 0, idx2 = 0;
    var cnt = 0;
    var sum = 0;
    do {
        var n1 = (nums1[idx1] + 1)|0;
        var n2 = (nums2[idx2] + 1)|0;
        if (n1 == 0 && n2 == 0) { break; }

        if (n1 == 0) {
            if (med0 == cnt) {
                sum += n2;
                if (isOdd) { break; }
            }
            if (med1 == cnt) {
                sum += n2;
                break;
            }
            ++cnt;
            ++idx2;
        } else if (n2 == 0) {
            if (med0 == cnt) {
                sum += n1;
                if (isOdd) { break; }
            }
            if (med1 == cnt) {
                sum += n1;
                break;
            }
            ++cnt;
            ++idx1;
        } else {
            if (n1 < n2) {
                if (med0 == cnt) {
                    sum += n1;
                    if (isOdd) { break; }
                }
                if (med1 == cnt) {
                    sum += n1;
                    break;
                }
                ++cnt;
                ++idx1;
            } else {
                if (med0 == cnt) {
                    sum += n2;
                    if (isOdd) { break; }
                }
                if (med1 == cnt) {
                    sum += n2;
                    break;
                }
                ++cnt;
                ++idx2;
            }
        }
    } while(1);
    return isOdd ? sum - 1 : (sum - 2) / 2;
};




