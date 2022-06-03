UPDATE AUTO_DEBIT_REGISTER SET PAYMENT_DATE = 1, SCAN_TIME = '01-MAY-22 04.11.47.000000000 PM',
    LAST_UPDATE = '01-MAY-22 04.11.47.000000000 PM'
WHERE ID in 
    (select id from AUTO_DEBIT_REGISTER 
        WHERE SERVICE_CODE = 'topup_Viettel'
        AND is_active IN (0,1) 
        AND payment_option = 3
    UNION
        select id from AUTO_DEBIT_REGISTER 
        WHERE SERVICE_CODE = 'topup_Viettel'
        AND is_active IN (0,1) 
        AND (scan_time IS null OR scan_time < TO_TIMESTAMP(trunc(last_day(sysdate)-1, 'mm')) OR 0 < payment_date)
        AND payment_option = 2
    );
    
    
    call pro_update_topup_in_auto_debit_register(2,2,'topup_Viettel');
    
     
    (select * from AUTO_DEBIT_REGISTER 
        WHERE SERVICE_CODE = 'topup_Viettel'
        AND is_active IN (0,1) 
        AND payment_option = 3
    UNION
        select * from AUTO_DEBIT_REGISTER 
        WHERE SERVICE_CODE = 'topup_Viettel'
        AND is_active IN (0,1) 
        AND (scan_time IS null OR scan_time < TO_TIMESTAMP(trunc(last_day(sysdate)-1, 'mm')) OR 0 < payment_date)
        AND payment_option = 2
    );


select * from auto_debit_register where reference_1 in (select bill_ID from REMINDER_TRACKING_DATA where AMOUNT>0 AND ERROR_CODE =0 and user_ID = 0946265079
and service_code like 'evn_hcm') and is_active in (0,1);

select * from auto_debit_register and service_code in('topup_hcm','topup_mobifone') and is_active in (0,1) 
select * from auto_debit_register where service_code in('topup_Viettel','topup_Mobifone') and is_active in (0,1) ;
select * from auto_debit_register where service_code in('topup_Viettel','topup_Mobifone') 
and is_active in (0,1) and last_update > '03-JUN-22 03.10.00.000000000 PM' and payment_option in (2,3);



SELECT * FROM V_PRIORITY_BILL_REMIND WHERE SERVICECODE = 'evn_hcm';

select * from REMINDER_TRACKING_DATA where AMOUNT>0 AND ERROR_CODE =0 and service_code = 'evn_hcm' and remind_time > 1654011128529
    and remind_type = 'VTTI_EVN_HCM';

select * from REMINDER_TRACKING_DATA where AMOUNT>0 AND ERROR_CODE =0 and service_code = 'evn_hcm';

SELECT * FROM AUTO_DEBIT_REGISTER WHERE user_ID='01639383882';

SELECT * FROM AUTO_DEBIT_REGISTER WHERE service_code = 'evn_hcm' and payment_date in(1,2,3);

SELECT * FROM AUTO_DEBIT_REGISTER WHERE reference_1='20160615-636110-0699';

SELECT * FROM AUTO_DEBIT_REGISTER order by id desc;


call generate_table_priority_bill_remind_v2();

SELECT * FROM REMINDER_TRACKING_DATA WHERE SERVICE_CODE = 'evn_hcm' 
and last_update >1654136102737 and error_code = 0 and amount>0;

SELECT * FROM AUTO_DEBIT_REGISTER WHERE reference_1='PE03000001579';
select * from bill where billid = 'PE03000001578';
select * from bill_available where bill_id = 1280761;


select * from priority_bill_remind where billid = 'PE03000001579';
SELECT * FROM AUTO_DEBIT_REGISTER WHERE reference_1='PE03000001579';


SELECT * FROM AUTO_DEBIT_REGISTER 
WHERE service_code='topup_Viettel' and is_active= 1
and payment_option = 3;

SELECT * FROM AUTO_DEBIT_REGISTER;

SELECT * FROM AUTO_DEBIT_REGISTER 
WHERE service_code='topup_Mobifone' and is_active= 1
and payment_option = 3;

call pro_update_topup_in_auto_debit_register(2,2,'topup_Mobifone');


select id from AUTO_DEBIT_REGISTER 
WHERE SERVICE_CODE = 'topup_Viettel'
AND payment_date <> 3
AND is_active IN (0,1) 
AND payment_option = 3
UNION
select id from AUTO_DEBIT_REGISTER 
WHERE SERVICE_CODE = 'topup_Viettel'
AND payment_date <> 3
AND is_active IN (0,1) 
AND (scan_time IS null OR scan_time < TO_TIMESTAMP(trunc(last_day(sysdate)-1, 'mm')) OR 1 < payment_date)
AND payment_option = 2;