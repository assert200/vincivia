
-- Get shares with one record
select s.symbol, s.name, s.exchange, count(r.*) 
from share s
left join record r on r.share_id = s.id
group by s.symbol, s.name, s.exchange
having count(*) = 1