URL="http://localhost:8000/b"
BATCHSIZE=200

PROCESSED_REQUEST=0
RATE_LIMITED_REQUESTS=0

START_TIME=$(date +%s)
for(( i = 0; i < $BATCHSIZE; i++ ))
do
  code=$(curl -s $URL --write-out '%{http_code}' -o /dev/null)
  echo "Status Code :" $code
  if [[ $code != 200 ]]; then
    ((RATE_LIMITED_REQUESTS++))
  else
    ((PROCESSED_REQUEST++))
  fi
done
END_TIME=$(date +%s)

DIFF=$(($END_TIME - $START_TIME))
echo "$PROCESSED_REQUEST requests processed in $DIFF seconds @ $(($PROCESSED_REQUEST/$DIFF)) requests/s"