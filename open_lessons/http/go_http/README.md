Требования

Сделать веб сервер для голосования с 2 ручками:
- Для сохранения результатов голосования (vote), принимает номер паспорта и id кандидата
- Для получения результатов (как по конкретному кандидату так и по всем кандидатам)

`bombardier -c 100 -d 60s -r 100000 -l -m POST -b '{"passport":"pass", "candidate_id": 1}' localhost:8000/vote`
`bombardier -c 125 -n 1000000 -l -m POST -b '{"passport":"pass", "candidate_id": 1}' localhost:8000/vote`
`bombardier -c 125 -n 1000000 -l -m GET localhost:8000/stats`