--
-- @lc app=leetcode id=182 lang=mysql
--
-- [182] Duplicate Emails
--

-- @lc code=start
# Write your MySQL query statement below

select email, count(*) as num
from Person
group by email
having count(*) > 1



