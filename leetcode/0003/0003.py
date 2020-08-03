# coding=utf-8


"""
public class Solution {
    public int lengthOfLongestSubstring(String s) {
        int n = s.length();
        Set<Character> set = new HashSet<>();
        int ans = 0, i = 0, j = 0;
        while (i < n && j < n) {
            // try to extend the range [i, j]
            if (!set.contains(s.charAt(j))){
                set.add(s.charAt(j++));
                ans = Math.max(ans, j - i);
            }
            else {
                set.remove(s.charAt(i++));
            }
        }
        return ans;
    }
}
"""


class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        n = len(s)
        ans = i = j = 0
        r = set()
        # 遍历所有 s[i:] 的情况
        while i < n and j < n:
            if s[j] not in r:
                # 加入边界
                r.add(s[j])
                j += 1
                ans = max(ans, j - i)
                pass
            else:
                r.remove(s[i])
                i += 1

        return ans


if __name__ == '__main__':
    print(Solution().lengthOfLongestSubstring("hello"))
