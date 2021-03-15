while true
do
  curl -s -o /dev/null -I -w "Status Code : %{http_code}\n" http://localhost:8000/
done
