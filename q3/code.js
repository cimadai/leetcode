var lengthOfLongestSubstring = function(s) {
    var i = 0,
        l = s.length,
        m = [],
        x = 0,
        target = 0,
        offset = 0,
        maxlen = 0,
        len = 0;

    for(; i < l|0; ++i) {
        target = s[i].codePointAt()|0;
        x = m[target]|0;
        m[target] = i + 1;

        if(x > offset) { offset = x|0; }

        len = i - offset + 1;
        if (maxlen < len) { maxlen = len|0; }
    }

    return maxlen;
};

