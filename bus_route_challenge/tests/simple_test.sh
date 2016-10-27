#!/bin/bash
sleep 5
{
    curl -s "localhost:8088/api/direct?dep_sid=3&arr_sid=4" | grep -E 'true|false' > /dev/null &&
    curl -s "localhost:8088/api/direct?dep_sid=0&arr_sid=1" | grep -E 'true|false' > /dev/null &&
    echo TEST PASSED!
} || {
    curl -s "localhost:8080/api/direct?dep_sid=3&arr_sid=4" | grep -E 'true|false' > /dev/null &&
    curl -s "localhost:8080/api/direct?dep_sid=0&arr_sid=1" | grep -E 'true|false' > /dev/null &&
    echo WRONG PORT!
} || {
    echo TEST FAILED!
}
