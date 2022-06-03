SELECT * FROM User_Procedures;
--NAMING CONVENTION:pro_update_topup_in_auto_debit_register

drop procedure pro_update_topup_in_auto_debit_register;
   
BEGIN
   pro_update_topup_in_auto_debit_register(17,'viettel'); 
   dbms_output.put_line('record inserted successfully');    
END;

--26-5
create or replace PROCEDURE pro_update_topup_in_auto_debit_register(
    updateDate NUMBER,
    today NUMBER,
    serviceCode VARCHAR2) IS
BEGIN
    UPDATE AUTO_DEBIT_REGISTER SET PAYMENT_DATE = CASE 
        WHEN payment_option = 3 THEN updateDate      
        WHEN payment_option = 2 AND (scan_time IS null OR scan_time < TO_TIMESTAMP(trunc(last_day(sysdate)-1, 'mm')) ) THEN updateDate
        WHEN payment_option = 2 AND today < payment_date THEN updateDate 
        WHEN payment_option = 2 AND today >= payment_date THEN payment_date END,
    SCAN_TIME = sysdate,
    LAST_UPDATE = sysdate
    WHERE SERVICE_CODE = serviceCode AND updateDate <> payment_date AND is_active IN (0,1) AND payment_option IN (2,3);
END;
--TIME

2/6
create or replace PROCEDURE pro_update_topup_in_auto_debit_register(
    updateDate NUMBER,
    today NUMBER,
    serviceCode VARCHAR2) IS
BEGIN
    UPDATE AUTO_DEBIT_REGISTER SET PAYMENT_DATE = CASE 
        WHEN payment_option = 3 THEN updateDate      
        WHEN payment_option = 2 AND (scan_time IS null OR scan_time < TO_TIMESTAMP(trunc(last_day(sysdate)-1, 'mm')) ) THEN updateDate
        WHEN payment_option = 2 AND today < payment_date THEN updateDate 
        WHEN payment_option = 2 AND today >= payment_date THEN payment_date END,
    SCAN_TIME = sysdate,
    LAST_UPDATE = sysdate
    WHERE ID in 
    (select id from AUTO_DEBIT_REGISTER 
        WHERE SERVICE_CODE = serviceCode
        AND payment_date <> updateDate
        AND is_active IN (0,1) 
        AND payment_option = 3
    UNION
        select id from AUTO_DEBIT_REGISTER 
        WHERE SERVICE_CODE = serviceCode
        AND payment_date <> updateDate
        AND is_active IN (0,1) 
        AND (scan_time IS null OR scan_time < TO_TIMESTAMP(trunc(last_day(sysdate)-1, 'mm')) OR today < payment_date)
        AND payment_option = 2
    );
END;

select sysdate from dual;

select curTime;

select to_timeStamp(sysdate)* (24 * 60 * 60 * 1000) from dual;

--to_timestamp( '23.12.2011 13:01:001', 'DD.MM.YYYY HH24:MI:SSFF3' )

select last_day(sysdate) from dual;

select * from dual;

SELECT
  LAST_DAY(ADD_MONTHS(SYSDATE,-1 )) LAST_DAY_LAST_MONTH,
  LAST_DAY(ADD_MONTHS(SYSDATE,1 )) LAST_DAY_NEXT_MONTH
FROM
  dual;
