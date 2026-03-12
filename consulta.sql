with relatorio as (
    SELECT
        p.id,
        c.email,
        (pag.data_atualizacao - p.data_criacao) as tempo
    from
        pedidos p
            inner join clientes c on p.cliente_id = c.id
            inner join pagamentos pag on pag.pedido_id = p.id
    where
            pag.status_pagamento like 'processando'
)
SELECT
    *
from
    relatorio
where
        tempo > '00:05:00'