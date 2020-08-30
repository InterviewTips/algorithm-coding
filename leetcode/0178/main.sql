-- # 1、order by 出来，相同的名次需要给上相同的排名
select a.Score as Score,
    (
        select count(distinct Score)
        from Scores as b
        where b.Score >= a.Score
    ) as `Rank`
from Scores as a
order by a.Score desc