class Solution:
  def canJump(self, nums: List[int]) -> bool:
    goal = len(nums) -1
    for i in range(len(s) - 1, -1, -1):
      if i + nums[i] >= goal:
        goal = i
        return True if goal == 0 else False
# class Solution:
#   def numDecodies(self, s: str) -> int:
#     dp = { len(s): 1}
#     def dfs(i):
#       if i in dp:
#         return dp[i]
#         if s[i] == "0":
#           return 0

#           res = dfs(i + 1)
#           if (i + 1< len(s) and (s[i] == "1" or s[i] == "2" and s[i + 1] in "0123456")):
#             res += dfs(i+2)
#             dp[i] = res
#             return res

#             return dfs(0)