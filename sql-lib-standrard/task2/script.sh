#!/bin/bash

REPLACED=""
REPLACE=""
# while there are more parameters passed (the reason why this is changing is because of shift
while test $# -gt 0; do
    case $1 in
        -h|--help)
            echo "options:"
            echo "-h --help                         show brief help"
            echo "1                                 auto create new file and replace without variable"
            echo "2                                 auto create new file and replace using variable"
            echo "3 [replaced_var] [time_replace]   ex: ./script.sh 3 \\\${START_TIME} '2021-12-10 00:00:00' 
                                                    this will create a new file"
            echo "4                                 replace in exist file"
            exit 0
                ;;
        1)
            cat product_filter.json | sed -e 's/"${START_TIME}"/"2021-10-20 00:00:00"/ ; s/"${END_TIME}"/"2021-10-21 12:12:12"/' > new_product_filter_1.json
            exit 1
            ;;
        2)
            replaced_start='${START_TIME}'
            replace_start=`date +"%Y-%m-%d %T"`
            replaced_end='${END_TIME}'
            replace_end=`date -d "+12 hour" +"%Y-%m-%d %T"`
            cat product_filter.json | sed -e "s@$replaced_start@$replace_start@ ; s@$replaced_end@$replace_end@" > new_product_filter_2.json
            exit 2
            ;;
        3)
            shift
            if test $# -gt 0; then
                REPLACED=$1
            else
                echo "no REPLACED specified"
                exit 3
            fi

            shift
            if test $# -gt 0; then
                REPLACE=$1
            else
                echo "no REPLACE specified"
                exit 3
            fi
            cat product_filter.json | sed -e "s@$REPLACED@$REPLACE@" > new_product_filter_3.json
            exit 3
            ;;
        4)
            sed -i 's/"${START_TIME}"/"2021-10-20 00:00:00"/ ; s/"${END_TIME}"/"2021-10-20 12:12:12"/' product_filter.json
            exit 4
            ;;
    esac
done

echo "./scritpt.sh [--help|-h]      to get more information"