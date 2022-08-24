create or replace function utilize_credit_code_v1_1(credit_code character varying, reference integer, ms_client_id smallint) returns integer
    language plpgsql
as
$$
declare
__version int;
__max_count int;
__code_id int;
__current_count int;
__affected_rows int := 0;
__retry_count int := 0;
__log_id int;
begin

        while __retry_count < 3 loop

            if not exists(
                select 1
                from microservice_clients
                where id = ms_client_id
            ) then
                raise EXCEPTION 'undefined client id: "%"', ms_client_id using DETAIL = '{"code":6}';
end if;

select id, _version, max_usable_count, current_used_count
into __code_id, __version, __max_count, __current_count
from codes where "text" = credit_code limit 1;

if __version is null then
                raise EXCEPTION 'undefined credit code: "%"', credit_code using DETAIL = '{"code":1}';
end if;

            if __current_count >= __max_count then
                raise EXCEPTION 'credit limitation reached: "%"', __max_count using DETAIL = '{"code":2}';
end if;

update codes
set current_used_count = current_used_count + 1,
    _version = _version + 1
where codes.text = credit_code
  and codes._version = __version
;

GET DIAGNOSTICS __affected_rows := ROW_COUNT;

if __affected_rows != 0 then

begin
insert into code_usage_logs (code_id, reference_id, client_id)
values (__code_id, reference, ms_client_id)
    returning id into __log_id;
return __log_id;
exception
                    when SQLSTATE '23505' THEN
                        raise EXCEPTION 'already utilized for this reference id' using DETAIL = '{"code":3}';
when others then
                        raise EXCEPTION 'unknown error to insert into usage logs' using DETAIL = '{"code":4}';
end;


else
                __retry_count := __retry_count + 1;
end if;
end loop;

        raise EXCEPTION 'unknown error to process this code' using DETAIL = '{"code":5}';

end;
$$;