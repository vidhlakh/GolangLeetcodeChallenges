# Definition for a binary tree node.
from collections import deque
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None
    
class Solution:
    def getTargetCopy(self, original: TreeNode, cloned: TreeNode, target: TreeNode) -> TreeNode:
        if not original:
            return None
        
        stack =deque([(original,cloned)]) #og,clone 
        
        while stack:
            og, clone = stack.popleft()
            if og == target:
                return clone
            if og.left:
                stack.append((og.left, clone.left))
            if og.right:
                stack.append((og.right, clone.right))
        return None
s = Solution()
original = TreeNode(4)

cloned = original
s.getTargetCopy(original, cloned, 4)