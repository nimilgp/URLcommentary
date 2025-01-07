sudo -u postgres psql -f createdb.sql

for FILE in `ls setupfiles`
do
    sudo -u postgres psql -d urlc -f ./setupfiles/$FILE
    sleep 2
done