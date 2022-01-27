# should output foo and bar on separate lines 
for _ in {1..10}; do
  curl ${LB_IP}:5678
done
