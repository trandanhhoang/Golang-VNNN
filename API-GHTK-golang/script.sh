# curl -X POST -H "Token: 617Fd51A9f1481164C2b9b429E8fC3E73365A82e" \
#  -H "Content-Type: application/json;" \
#  -d '{"name":"hoang shop test","first_address":"41 ywuw","province":"Gia Lai","district":"Kbang","tel":"0969141679","email":"danhhoang5920@gmail.com"}' \
#  "https://services.giaohangtietkiem.vn/services/shops/add"

# curl -X POST -H "Token: 617Fd51A9f1481164C2b9b429E8fC3E73365A82e" \
#  -H "Content-Type: application/json" \
#  -d '{"products":[{"name":"bút","weight":0.1,"quantity":1,"product_code": 1241},{"name":"tẩy","weight":0.2,"quantity":1,"product_code": 1254}],"order":{"id":"a4","pick_name":"HCM-nội thành","pick_address":"590 CMT8 P.11","pick_province":"TP. Hồ Chí Minh","pick_district":"Quận 3","pick_ward":"Phường 1","pick_tel":"0911222333","tel":"0911222333","name":"GHTK - HCM - Noi Thanh","address":"123 nguyễn chí thanh","province":"TP. Hồ Chí Minh","district":"Quận 1","ward":"Phường Bến Nghé","hamlet":"Khác","is_freeship":"1","pick_date":"2016-09-30","pick_money":47000,"note":"Khối lượng tính cước tối đa: 1.00 kg","value":3000000,"transport":"fly","pick_option":"cod","deliver_option" : "xteam","pick_session" : 2}}' "https://services.giaohangtietkiem.vn/services/shipment/order"

curl -H "Token: 617Fd51A9f1481164C2b9b429E8fC3E73365A82e" \
  "https://services.giaohangtietkiem.vn/authentication-request-sample"