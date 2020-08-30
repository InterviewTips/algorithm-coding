select distinct tmp.Num as ConsecutiveNums
from (
        select a.Id,
            a.Num,
            if (
                (
                    select Num
                    from Logs as b
                    where b.id = a.id -1
                ) = a.Num,
                1,
                0
            ) as b1,
            if (
                (
                    select Num
                    from Logs as b
                    where b.id = a.id + 1
                ) = a.Num,
                1,
                0
            ) as a1,
            if (
                (
                    select Num
                    from Logs as b
                    where b.id = a.id -2
                ) = a.Num,
                1,
                0
            ) as b2,
            if (
                (
                    select Num
                    from Logs as b
                    where b.id = a.id + 2
                ) = a.Num,
                1,
                0
            ) as a2
        from Logs as a
    ) as tmp
where (
        tmp.a1 = 1
        and tmp.b1 = 1
    )
    or (
        tmp.a1 = 1
        and tmp.a2 = 1
    )
    or (
        tmp.b1 = 1
        and tmp.b2 = 1
    )