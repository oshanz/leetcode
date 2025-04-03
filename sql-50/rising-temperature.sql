# Write your MySQL query statement below

select w2.id as Id from Weather w1 
inner join Weather w2 on w1.recordDate =  DATE_SUB(w2.recordDate, INTERVAL 1 DAY)
where w2.temperature > w1.temperature
