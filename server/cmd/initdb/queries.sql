-- get shares with just 1 record
select s.symbol, s.name, s.exchange, count(r.*) 
from share s
left join record r on r.share_id = s.id
group by s.symbol, s.name, s.exchange
having count(r.*) = 1

-- get shares %
with yest_results as (
	select share_id, last_sale from record
	where recorded_at='2021-01-12'
), today_results as (
	select share_id, last_sale from record
	where recorded_at='2021-01-13'
) 
select s.symbol, s.name, s.exchange, 
	yr.last_sale "yest_price", tr.last_sale "today_price", 
	tr.last_sale - yr.last_sale "change",
	(tr.last_sale - yr.last_sale)*100 / yr.last_sale "percent"
from share s
left join yest_results yr on yr.share_id = s.id
left join today_results tr on tr.share_id = s.id
where tr.last_sale is not null
and yr.last_sale is not null
and yr.last_sale <> 0
order by "percent" desc


-- get shares that existed yesterday but not today
with yest_results as (
	select share_id from record
	where recorded_at='2021-01-12'
), today_results as (
	select share_id from record
	where recorded_at='2021-01-13'
) 
select s.symbol, s.name, s.exchange
from share s
join today_results tr on tr.share_id = s.id
where tr.share_id not in (select yr.share_id from yest_results yr)
