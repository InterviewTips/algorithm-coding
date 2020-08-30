-- # Write your MySQL query statement below

-- # 思路
-- # 确保某条记录大于 1jk0
-- # 1) 遍历每一天
-- # 2) 是否高峰日
-- # 3) (当天 > 100 and ((前一天 >=100 and 后一天 >= 100) or (前两天 >=100) or (后两天 >= 100)))

select tmp.id, s.visit_date, s.people from
(select id,
-- # 前两天
if (
    (select people from stadium as b where b.id = a.id-2) >= 100, -- 边界问题 1-2 = -1, 问题不大，此时 if 结果为 0
    1,
    0
) as b2,
-- # 前一天
if (
    (select people from stadium as b where b.id = a.id-1) >= 100,
    1,
    0
) as b1,
-- # 后一天
if (
    (select people from stadium as b where b.id = a.id+1) >= 100,
    1,
    0
) as a1,
-- # 后两天
if (
    (select people from stadium as b where b.id = a.id+2) >= 100,
    1,
    0
) as a2 from stadium as a where a.people >= 100) as tmp 
left join 
stadium as s on tmp.id = s.id 
where (tmp.b1 = 1 and tmp.b2 = 1) or (tmp.a1 = 1 and tmp.a2 = 1) or (tmp.b1 = 1 and tmp.a1 = 1);

